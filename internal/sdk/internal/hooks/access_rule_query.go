package hooks

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

// accessRuleQueryHook bridges the Terraform string attribute and the REST object.
//
// Terraform exposes `query` as a JSON-encoded string (jsonencode(...)) so the
// recursive OpalNodeQueryBody schema does not explode Speakeasy's terraform
// generator. The public API still expects/returns a JSON object. This hook
// expands the string on write and collapses the object on read so the SDK's
// *string model round-trips cleanly.
//
// Write semantics match REST: at least one of ruleClauses / query; both AND.
// REST PUT omits a side to clear it; query / rule_clauses are optional-only in
// TF so omit plans null and matches that. This hook expands encoding, strips
// empty GET ruleClauses on write and read (no when/unless selectors) so they
// are not re-sent or retained in state beside query, and drops an explicit null
// query. Real tag selectors next to a TF query string are kept (hybrid AND).
type accessRuleQueryHook struct{}

var (
	_ beforeRequestHook = (*accessRuleQueryHook)(nil)
	_ afterSuccessHook  = (*accessRuleQueryHook)(nil)
)

var accessRuleQueryWriteOps = map[string]struct{}{
	"create_access_rule": {},
	"update_access_rule": {},
}

var accessRuleQueryReadOps = map[string]struct{}{
	"create_access_rule": {},
	"update_access_rule": {},
	"get_access_rule":    {},
	"get_access_rules":   {},
}

func (h *accessRuleQueryHook) BeforeRequest(
	hookCtx BeforeRequestContext,
	req *http.Request,
) (*http.Request, error) {
	if _, ok := accessRuleQueryWriteOps[hookCtx.OperationID]; !ok || req.Body == nil {
		return req, nil
	}

	body, err := io.ReadAll(req.Body)
	if err != nil {
		return req, fmt.Errorf("read access-rule request: %w", err)
	}
	if err := req.Body.Close(); err != nil {
		return req, fmt.Errorf("close access-rule request: %w", err)
	}

	rewritten, changed, err := prepareAccessRuleWriteBody(body)
	if err != nil {
		return req, err
	}
	if !changed {
		req.Body = io.NopCloser(bytes.NewReader(body))
		return req, nil
	}

	req.Body = io.NopCloser(bytes.NewReader(rewritten))
	req.ContentLength = int64(len(rewritten))
	req.GetBody = func() (io.ReadCloser, error) {
		return io.NopCloser(bytes.NewReader(rewritten)), nil
	}
	return req, nil
}

func (h *accessRuleQueryHook) AfterSuccess(
	hookCtx AfterSuccessContext,
	res *http.Response,
) (*http.Response, error) {
	if _, ok := accessRuleQueryReadOps[hookCtx.OperationID]; !ok {
		return res, nil
	}
	if res == nil || res.Body == nil || res.Body == http.NoBody {
		return res, nil
	}
	if !strings.Contains(res.Header.Get("Content-Type"), "json") {
		return res, nil
	}

	body, err := io.ReadAll(res.Body)
	_ = res.Body.Close()
	res.Body = io.NopCloser(bytes.NewReader(body))
	if err != nil {
		return res, fmt.Errorf("read access-rule response: %w", err)
	}
	if !bytes.Contains(body, []byte(`"query"`)) && !bytes.Contains(body, []byte(`"ruleClauses"`)) {
		return res, nil
	}

	rewritten, changed, err := normalizeAccessRuleReadBody(body)
	if err != nil || !changed {
		return res, nil
	}

	res.Body = io.NopCloser(bytes.NewReader(rewritten))
	res.ContentLength = int64(len(rewritten))
	if res.Header.Get("Content-Length") != "" {
		res.Header.Set("Content-Length", strconv.Itoa(len(rewritten)))
	}
	return res, nil
}

// ruleClausesHasSelectors reports whether ruleClauses carries any tag/attribute
// selectors under when or unless (a real tag-based condition). Empty
// when.clauses from an OQ GET must not count — those are stripped so they do
// not ride along on PUT.
func ruleClausesHasSelectors(raw json.RawMessage) bool {
	type clause struct {
		Selectors          []json.RawMessage `json:"selectors"`
		AttributeSelectors []json.RawMessage `json:"attribute_selectors"`
	}
	var clauses struct {
		When struct {
			Clauses []clause `json:"clauses"`
		} `json:"when"`
		Unless struct {
			Clauses []clause `json:"clauses"`
		} `json:"unless"`
	}
	if err := json.Unmarshal(raw, &clauses); err != nil {
		return false
	}
	for _, c := range append(clauses.When.Clauses, clauses.Unless.Clauses...) {
		if len(c.Selectors) > 0 || len(c.AttributeSelectors) > 0 {
			return true
		}
	}
	return false
}

// prepareAccessRuleWriteBody expands a TF query string into a REST object and
// aligns the body with REST at-least-one / AND / PUT-omit semantics.
func prepareAccessRuleWriteBody(body []byte) ([]byte, bool, error) {
	var payload map[string]json.RawMessage
	if err := json.Unmarshal(body, &payload); err != nil {
		return nil, false, fmt.Errorf("decode access-rule request: %w", err)
	}

	changed := false

	raw, hasQuery := payload["query"]
	querySet := hasQuery && len(raw) > 0 && !bytes.Equal(raw, []byte("null"))
	queryIsString := querySet && raw[0] == '"'
	clauses, hasClauses := payload["ruleClauses"]
	hasSelectors := hasClauses && ruleClausesHasSelectors(clauses)

	// Expand TF jsonencode(...) string into a REST object.
	if queryIsString {
		var encoded string
		if err := json.Unmarshal(raw, &encoded); err != nil {
			return nil, false, fmt.Errorf("decode access-rule query string: %w", err)
		}
		if !json.Valid([]byte(encoded)) {
			return nil, false, fmt.Errorf("access-rule query is not valid JSON")
		}
		payload["query"] = json.RawMessage(encoded)
		changed = true
	}

	// Optional-only TF may still serialize an explicit null; omit it so PUT
	// clears OQ. Stale Computed objects no longer reach the wire.
	if hasQuery && bytes.Equal(raw, []byte("null")) {
		delete(payload, "query")
		changed = true
	}

	if querySet || queryIsString {
		// Empty GET clauses must not ride along (omit = clear tags under PUT).
		// Real selectors stay — hybrid AND with query.
		if hasClauses && !hasSelectors {
			delete(payload, "ruleClauses")
			changed = true
		}
	}

	if !changed {
		return body, false, nil
	}
	out, err := json.Marshal(payload)
	if err != nil {
		return nil, false, fmt.Errorf("encode access-rule request: %w", err)
	}
	return out, true, nil
}

// normalizeAccessRuleReadBody collapses query objects to strings and drops empty
// GET ruleClauses so optional-only TF state does not retain a phantom block.
func normalizeAccessRuleReadBody(body []byte) ([]byte, bool, error) {
	decoder := json.NewDecoder(bytes.NewReader(body))
	decoder.UseNumber()

	var payload any
	if err := decoder.Decode(&payload); err != nil {
		return nil, false, nil
	}

	changed := collapseQueryObjects(payload)
	if stripEmptyRuleClauses(payload) {
		changed = true
	}
	if !changed {
		return body, false, nil
	}

	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	enc.SetEscapeHTML(false)
	if err := enc.Encode(payload); err != nil {
		return nil, false, nil
	}
	return bytes.TrimSuffix(buf.Bytes(), []byte("\n")), true, nil
}

// collapseAccessRuleQueryObjects is kept for unit tests that only exercise
// query object→string collapsing.
func collapseAccessRuleQueryObjects(body []byte) ([]byte, bool, error) {
	return normalizeAccessRuleReadBody(body)
}

// stripEmptyRuleClauses removes ruleClauses with no when/unless selectors so
// OQ-only GETs do not leave an empty block in optional-only Terraform state.
func stripEmptyRuleClauses(node any) bool {
	changed := false
	switch value := node.(type) {
	case map[string]any:
		if raw, ok := value["ruleClauses"]; ok && raw != nil {
			encoded, err := json.Marshal(raw)
			if err == nil && !ruleClausesHasSelectors(encoded) {
				delete(value, "ruleClauses")
				changed = true
			}
		}
		for _, child := range value {
			if stripEmptyRuleClauses(child) {
				changed = true
			}
		}
	case []any:
		for _, child := range value {
			if stripEmptyRuleClauses(child) {
				changed = true
			}
		}
	}
	return changed
}

func collapseQueryObjects(node any) bool {
	changed := false
	switch value := node.(type) {
	case map[string]any:
		if raw, ok := value["query"]; ok {
			if obj, isObj := raw.(map[string]any); isObj {
				encoded, err := json.Marshal(obj)
				if err == nil {
					value["query"] = string(encoded)
					changed = true
				}
			}
		}
		for _, child := range value {
			if collapseQueryObjects(child) {
				changed = true
			}
		}
	case []any:
		for _, child := range value {
			if collapseQueryObjects(child) {
				changed = true
			}
		}
	}
	return changed
}
