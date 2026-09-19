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

const (
	requestConfigurationsKey = "request_configurations"
	maxDurationKey           = "max_duration_minutes"

	// indefiniteDuration is the sentinel the public API documents for "no
	// maximum duration". It matches common.NullableDurationIndefinite in the
	// Opal backend.
	indefiniteDuration = -1
)

// maxDurationIndefiniteHook restores the -1 indefinite-duration sentinel on
// responses that omit it.
//
// The API accepts max_duration_minutes = -1 to mean "no maximum", but strips
// the sentinel on write (requestconfigurationutils/api.go, update.go) and
// returns the stored NULL on read, so the field comes back absent. Terraform
// cannot reconcile that on its own: core requires a planned value to equal a
// non-null config value, so a config of -1 can never be planned as null, and
// refreshed state of null then diffs against it on every plan (PLAT-654).
//
// Normalizing the response means state always holds -1 for an indefinite
// duration, which matches config and leaves the plan untouched. The v2 provider
// did the same thing in opal/request_configuration.go; the v3 regeneration
// dropped it, which is what reintroduced the drift.
//
// This is a no-op once the API returns -1 itself, so it can be deleted then.
type maxDurationIndefiniteHook struct{}

var _ afterSuccessHook = (*maxDurationIndefiniteHook)(nil)

func (h *maxDurationIndefiniteHook) AfterSuccess(_ AfterSuccessContext, res *http.Response) (*http.Response, error) {
	if res == nil || res.Body == nil || res.Body == http.NoBody {
		return res, nil
	}
	if !strings.Contains(res.Header.Get("Content-Type"), "json") {
		return res, nil
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return res, fmt.Errorf("read response for max duration normalization: %w", err)
	}
	if err := res.Body.Close(); err != nil {
		return res, fmt.Errorf("close response for max duration normalization: %w", err)
	}
	// Restore the body on every path below, including the ones that bail out.
	res.Body = io.NopCloser(bytes.NewReader(body))

	// Cheap gate: most responses carry no request configurations at all.
	if !bytes.Contains(body, []byte(`"`+requestConfigurationsKey+`"`)) {
		return res, nil
	}

	// UseNumber keeps numeric literals byte-exact; decoding into float64 would
	// reformat large integers such as IDs and durations.
	decoder := json.NewDecoder(bytes.NewReader(body))
	decoder.UseNumber()

	var payload any
	if err := decoder.Decode(&payload); err != nil {
		// Not something we can rewrite - leave the original body for the SDK to
		// report on.
		return res, nil
	}

	if !normalizeIndefiniteMaxDuration(payload) {
		return res, nil
	}

	rewritten, err := json.Marshal(payload)
	if err != nil {
		return res, nil
	}

	res.Body = io.NopCloser(bytes.NewReader(rewritten))
	res.ContentLength = int64(len(rewritten))
	if res.Header.Get("Content-Length") != "" {
		res.Header.Set("Content-Length", strconv.Itoa(len(rewritten)))
	}
	return res, nil
}

// normalizeIndefiniteMaxDuration walks the decoded response and fills in the
// sentinel wherever a request configuration omits max_duration_minutes. It
// reports whether anything changed.
//
// The walk is shape-agnostic on purpose: request configurations are nested
// differently across single, paginated, and bulk responses, and new endpoints
// should not need to be registered here.
func normalizeIndefiniteMaxDuration(node any) bool {
	changed := false
	switch value := node.(type) {
	case map[string]any:
		for key, child := range value {
			if key == requestConfigurationsKey {
				changed = applyIndefiniteMaxDuration(child) || changed
			}
			changed = normalizeIndefiniteMaxDuration(child) || changed
		}
	case []any:
		for _, child := range value {
			changed = normalizeIndefiniteMaxDuration(child) || changed
		}
	}
	return changed
}

func applyIndefiniteMaxDuration(node any) bool {
	configurations, ok := node.([]any)
	if !ok {
		return false
	}

	changed := false
	for _, configuration := range configurations {
		object, ok := configuration.(map[string]any)
		if !ok {
			continue
		}
		// An explicit value, including an explicit null meaning "no maximum",
		// is only rewritten when it carries no duration.
		if existing, present := object[maxDurationKey]; present && existing != nil {
			continue
		}
		object[maxDurationKey] = json.Number(strconv.Itoa(indefiniteDuration))
		changed = true
	}
	return changed
}
