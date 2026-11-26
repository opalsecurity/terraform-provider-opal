package hooks

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"testing"
)

func TestPrepareAccessRuleWriteBody(t *testing.T) {
	t.Parallel()

	// GET→PUT shape: empty ruleClauses next to a TF query string must become
	// a query-only body (omit clears tags under PUT replace).
	in := `{"name":"x","query":"{\"nodeFilters\":{\"entityTag\":{\"key\":\"team\"}}}","ruleClauses":{"when":{"clauses":[]}},"status":"ACTIVE"}`
	out, changed, err := prepareAccessRuleWriteBody([]byte(in))
	if err != nil {
		t.Fatalf("prepare: %v", err)
	}
	if !changed {
		t.Fatal("expected rewrite")
	}

	var payload map[string]json.RawMessage
	if err := json.Unmarshal(out, &payload); err != nil {
		t.Fatalf("decode rewritten: %v", err)
	}
	if payload["query"][0] != '{' {
		t.Fatalf("query should be an object, got %s", payload["query"])
	}
	if _, ok := payload["ruleClauses"]; ok {
		t.Fatalf(
			"empty ruleClauses should be stripped when query is set, got %s",
			payload["ruleClauses"],
		)
	}

	// Object query with no ruleClauses — no-op.
	_, changed, err = prepareAccessRuleWriteBody([]byte(`{"query":{"nodeFilters":{}}}`))
	if err != nil {
		t.Fatalf("prepare object: %v", err)
	}
	if changed {
		t.Fatal("object query without ruleClauses should not be rewritten")
	}
}

func TestPrepareAccessRuleWriteBodyTagOnlyOmitsNullQuery(t *testing.T) {
	t.Parallel()

	// Tag-only PUT: drop a null query so the side is omitted (clears OQ).
	out, changed, err := prepareAccessRuleWriteBody([]byte(
		`{"ruleClauses":{"when":{"clauses":[{"selectors":[{"key":"team","value":"eng","connection_id":"00000000-0000-4000-8000-000000000001"}]}]}},"query":null}`,
	))
	if err != nil {
		t.Fatalf("prepare tag-only: %v", err)
	}
	if !changed {
		t.Fatal("expected null query to be dropped for tag-only write")
	}
	var payload map[string]json.RawMessage
	if err := json.Unmarshal(out, &payload); err != nil {
		t.Fatalf("decode tag-only: %v", err)
	}
	if _, ok := payload["query"]; ok {
		t.Fatalf("query should be omitted on tag-only write, got %s", payload["query"])
	}
	if _, ok := payload["ruleClauses"]; !ok {
		t.Fatal("ruleClauses should remain on tag-only write")
	}
}

func TestPrepareAccessRuleWriteBodyHybridKeepsBoth(t *testing.T) {
	t.Parallel()

	// REST AND: TF jsonencode query + real tag selectors must both stay.
	in := `{"ruleClauses":{"when":{"clauses":[{"selectors":[{"key":"access-rule","value":"include","connection_id":"00000000-0000-4000-8000-000000000001"}]}]}},"query":"{\"nodeFilters\":{\"entityTag\":{\"key\":\"dept\",\"value\":\"eng\"}}}","status":"PAUSED"}`
	out, changed, err := prepareAccessRuleWriteBody([]byte(in))
	if err != nil {
		t.Fatalf("prepare: %v", err)
	}
	if !changed {
		t.Fatal("expected query string expansion")
	}
	var payload map[string]json.RawMessage
	if err := json.Unmarshal(out, &payload); err != nil {
		t.Fatalf("decode: %v", err)
	}
	if _, ok := payload["ruleClauses"]; !ok {
		t.Fatal("ruleClauses with selectors should be kept for hybrid AND")
	}
	if payload["query"][0] != '{' {
		t.Fatalf("query should be expanded object, got %s", payload["query"])
	}
}

func TestPrepareAccessRuleWriteBodyHybridKeepsUnlessSelectors(t *testing.T) {
	t.Parallel()

	// unless-only tags next to a TF query string are still a hybrid AND.
	in := `{"ruleClauses":{"when":{"clauses":[]},"unless":{"clauses":[{"selectors":[{"key":"access-rule","value":"exclude","connection_id":"00000000-0000-4000-8000-000000000001"}]}]}},"query":"{\"nodeFilters\":{\"entityTag\":{\"key\":\"dept\",\"value\":\"eng\"}}}","status":"PAUSED"}`
	out, changed, err := prepareAccessRuleWriteBody([]byte(in))
	if err != nil {
		t.Fatalf("prepare: %v", err)
	}
	if !changed {
		t.Fatal("expected query string expansion")
	}
	var payload map[string]json.RawMessage
	if err := json.Unmarshal(out, &payload); err != nil {
		t.Fatalf("decode: %v", err)
	}
	if _, ok := payload["ruleClauses"]; !ok {
		t.Fatal("unless selectors must keep ruleClauses (hybrid AND)")
	}
	if payload["query"][0] != '{' {
		t.Fatalf("query should be expanded object, got %s", payload["query"])
	}
}

func TestCollapseAccessRuleQueryObjects(t *testing.T) {
	t.Parallel()

	in := `{"name":"x","query":{"nodeFilters":{"entityTag":{"key":"team"}}},"uses_opal_query":true}`
	out, changed, err := collapseAccessRuleQueryObjects([]byte(in))
	if err != nil {
		t.Fatalf("collapse: %v", err)
	}
	if !changed {
		t.Fatal("expected rewrite")
	}

	var payload map[string]any
	if err := json.Unmarshal(out, &payload); err != nil {
		t.Fatalf("decode rewritten: %v", err)
	}
	query, ok := payload["query"].(string)
	if !ok {
		t.Fatalf("query should be a string, got %T (%v)", payload["query"], payload["query"])
	}
	if !strings.Contains(query, `"nodeFilters"`) {
		t.Fatalf("unexpected query string: %s", query)
	}
}

func TestPrepareAccessRuleWriteBodyStripsEmptyClausesOnNestedQuery(t *testing.T) {
	t.Parallel()

	// CW / UI nested query as a TF string + empty GET ruleClauses must become
	// query-only on the wire (PUT omit clears tags).
	in := `{"name":"cw","query":"{\"nodeFilters\":{\"allOf\":[{\"entityTypes\":[\"USER\"]},{\"allOf\":[{\"entityTag\":{\"key\":\"division\",\"value\":\"Finance\"}},{\"anyOf\":[{\"entityTag\":{\"key\":\"jobLevel\",\"value\":\"M3\"}}]}]}]}}","ruleClauses":{"when":{"clauses":[]}},"status":"PAUSED"}`
	out, changed, err := prepareAccessRuleWriteBody([]byte(in))
	if err != nil {
		t.Fatalf("prepare: %v", err)
	}
	if !changed {
		t.Fatal("expected rewrite")
	}
	var payload map[string]json.RawMessage
	if err := json.Unmarshal(out, &payload); err != nil {
		t.Fatalf("decode: %v", err)
	}
	if payload["query"][0] != '{' {
		t.Fatalf("query should be object, got %s", payload["query"])
	}
	if _, ok := payload["ruleClauses"]; ok {
		t.Fatalf("empty ruleClauses should be stripped, got %s", payload["ruleClauses"])
	}
	var query map[string]any
	if err := json.Unmarshal(payload["query"], &query); err != nil {
		t.Fatalf("decode query object: %v", err)
	}
	if _, ok := query["nodeFilters"]; !ok {
		t.Fatalf("expected nodeFilters in expanded query, got %v", query)
	}
}

func TestAccessRuleQueryHookRoundTrip(t *testing.T) {
	t.Parallel()

	hook := &accessRuleQueryHook{}

	reqBody := `{"admin_owner_id":"11111111-1111-1111-1111-111111111111","description":"d","name":"n","query":"{\"nodeFilters\":{\"entityTag\":{\"key\":\"access-rule\",\"value\":\"include\",\"connectionId\":\"22222222-2222-2222-2222-222222222222\"}}}","status":"ACTIVE"}`
	req, err := http.NewRequest(
		http.MethodPost,
		"http://127.0.0.1/v1/access-rules",
		strings.NewReader(reqBody),
	)
	if err != nil {
		t.Fatalf("new request: %v", err)
	}
	outReq, err := hook.BeforeRequest(
		BeforeRequestContext{HookContext: HookContext{OperationID: "create_access_rule"}},
		req,
	)
	if err != nil {
		t.Fatalf("BeforeRequest: %v", err)
	}
	got, err := io.ReadAll(outReq.Body)
	if err != nil {
		t.Fatalf("read body: %v", err)
	}
	var payload map[string]json.RawMessage
	if err := json.Unmarshal(got, &payload); err != nil {
		t.Fatalf("decode: %v", err)
	}
	if payload["query"][0] != '{' {
		t.Fatalf("expected object query on the wire, got %s", payload["query"])
	}

	res := jsonResponse(
		`{"access_rule_id":"33333333-3333-3333-3333-333333333333","admin_owner_id":"11111111-1111-1111-1111-111111111111","description":"d","name":"n","query":{"nodeFilters":{"entityTag":{"key":"access-rule"}}},"ruleClauses":{"when":{"clauses":[]}},"status":"ACTIVE","uses_opal_query":true}`,
	)
	outRes, err := hook.AfterSuccess(
		AfterSuccessContext{HookContext: HookContext{OperationID: "create_access_rule"}},
		res,
	)
	if err != nil {
		t.Fatalf("AfterSuccess: %v", err)
	}
	body, err := io.ReadAll(outRes.Body)
	if err != nil {
		t.Fatalf("read response: %v", err)
	}
	var resp map[string]any
	if err := json.Unmarshal(body, &resp); err != nil {
		t.Fatalf("decode response: %v", err)
	}
	if _, ok := resp["query"].(string); !ok {
		t.Fatalf("response query should be string for SDK *string model, got %T", resp["query"])
	}
}

func TestNormalizeAccessRuleReadBodyStripsEmptyRuleClauses(t *testing.T) {
	t.Parallel()

	// OQ GET returns empty when.clauses; optional-only TF must not retain them.
	in := `{"name":"x","query":{"nodeFilters":{"entityTag":{"key":"team"}}},"ruleClauses":{"when":{"clauses":[]}},"uses_opal_query":true}`
	out, changed, err := normalizeAccessRuleReadBody([]byte(in))
	if err != nil {
		t.Fatalf("normalize: %v", err)
	}
	if !changed {
		t.Fatal("expected rewrite")
	}
	var payload map[string]any
	if err := json.Unmarshal(out, &payload); err != nil {
		t.Fatalf("decode: %v", err)
	}
	if _, ok := payload["ruleClauses"]; ok {
		t.Fatalf("empty ruleClauses should be stripped, got %v", payload["ruleClauses"])
	}
	if _, ok := payload["query"].(string); !ok {
		t.Fatalf("query should still collapse to string, got %T", payload["query"])
	}
}

func TestNormalizeAccessRuleReadBodyKeepsHybridSelectors(t *testing.T) {
	t.Parallel()

	in := `{"query":{"nodeFilters":{}},"ruleClauses":{"when":{"clauses":[{"selectors":[{"key":"team","value":"eng","connection_id":"00000000-0000-4000-8000-000000000001"}]}]}}}`
	out, changed, err := normalizeAccessRuleReadBody([]byte(in))
	if err != nil {
		t.Fatalf("normalize: %v", err)
	}
	if !changed {
		t.Fatal("expected query collapse")
	}
	var payload map[string]any
	if err := json.Unmarshal(out, &payload); err != nil {
		t.Fatalf("decode: %v", err)
	}
	if _, ok := payload["ruleClauses"]; !ok {
		t.Fatal("hybrid selectors must keep ruleClauses")
	}
}
