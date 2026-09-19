package hooks

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"testing"
)

func jsonResponse(body string) *http.Response {
	header := make(http.Header)
	header.Set("Content-Type", "application/json")
	return &http.Response{
		StatusCode: http.StatusOK,
		Header:     header,
		Body:       io.NopCloser(strings.NewReader(body)),
	}
}

func runHook(t *testing.T, res *http.Response) string {
	t.Helper()
	out, err := (&maxDurationIndefiniteHook{}).AfterSuccess(AfterSuccessContext{}, res)
	if err != nil {
		t.Fatalf("AfterSuccess returned error: %v", err)
	}
	body, err := io.ReadAll(out.Body)
	if err != nil {
		t.Fatalf("reading rewritten body: %v", err)
	}
	return string(body)
}

// maxDurations pulls every request configuration's max_duration_minutes out of
// a response, so assertions do not depend on JSON key ordering.
func maxDurations(t *testing.T, body string) []any {
	t.Helper()
	var payload any
	decoder := json.NewDecoder(strings.NewReader(body))
	decoder.UseNumber()
	if err := decoder.Decode(&payload); err != nil {
		t.Fatalf("decoding body %q: %v", body, err)
	}

	var found []any
	var walk func(node any)
	walk = func(node any) {
		switch value := node.(type) {
		case map[string]any:
			for key, child := range value {
				if key == requestConfigurationsKey {
					if configurations, ok := child.([]any); ok {
						for _, configuration := range configurations {
							object, ok := configuration.(map[string]any)
							if !ok {
								continue
							}
							found = append(found, object[maxDurationKey])
						}
					}
				}
				walk(child)
			}
		case []any:
			for _, child := range value {
				walk(child)
			}
		}
	}
	walk(payload)
	return found
}

func TestMaxDurationIndefiniteHook(t *testing.T) {
	t.Parallel()

	for _, testCase := range []struct {
		name string
		body string
		want []any
	}{
		{
			name: "absent max duration becomes the sentinel",
			body: `{"group_id":"g1","request_configurations":[{"priority":0}]}`,
			want: []any{json.Number("-1")},
		},
		{
			name: "explicit null becomes the sentinel",
			body: `{"group_id":"g1","request_configurations":[{"priority":0,"max_duration_minutes":null}]}`,
			want: []any{json.Number("-1")},
		},
		{
			name: "a real duration is preserved",
			body: `{"group_id":"g1","request_configurations":[{"priority":0,"max_duration_minutes":120}]}`,
			want: []any{json.Number("120")},
		},
		{
			name: "already indefinite is left alone",
			body: `{"group_id":"g1","request_configurations":[{"priority":0,"max_duration_minutes":-1}]}`,
			want: []any{json.Number("-1")},
		},
		{
			name: "zero is preserved",
			body: `{"group_id":"g1","request_configurations":[{"priority":0,"max_duration_minutes":0}]}`,
			want: []any{json.Number("0")},
		},
		{
			name: "every configuration in a list is normalized independently",
			body: `{"request_configurations":[{"priority":0},{"priority":1,"max_duration_minutes":525600}]}`,
			want: []any{json.Number("-1"), json.Number("525600")},
		},
		{
			name: "paginated responses are reached",
			body: `{"results":[{"group_id":"g1","request_configurations":[{"priority":0}]},` +
				`{"group_id":"g2","request_configurations":[{"priority":0,"max_duration_minutes":60}]}]}`,
			want: []any{json.Number("-1"), json.Number("60")},
		},
		{
			name: "empty configuration lists are untouched",
			body: `{"group_id":"g1","request_configurations":[]}`,
			want: nil,
		},
	} {
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			got := maxDurations(t, runHook(t, jsonResponse(testCase.body)))
			if len(got) != len(testCase.want) {
				t.Fatalf("got %v durations, want %v", got, testCase.want)
			}
			for i := range got {
				if got[i] != testCase.want[i] {
					t.Errorf("configuration %d: got %v, want %v", i, got[i], testCase.want[i])
				}
			}
		})
	}
}

func TestMaxDurationIndefiniteHookIsIdempotent(t *testing.T) {
	t.Parallel()

	body := `{"group_id":"g1","request_configurations":[{"priority":0}]}`
	first := runHook(t, jsonResponse(body))
	second := runHook(t, jsonResponse(first))
	if first != second {
		t.Errorf("second pass changed the body:\n first: %s\nsecond: %s", first, second)
	}
}

func TestMaxDurationIndefiniteHookLeavesOtherResponsesIntact(t *testing.T) {
	t.Parallel()

	for _, testCase := range []struct {
		name        string
		contentType string
		body        string
	}{
		{
			name:        "no request configurations",
			contentType: "application/json",
			body:        `{"group_id":"g1","name":"engineering"}`,
		},
		{
			name:        "recommended duration is not touched",
			contentType: "application/json",
			body:        `{"request_configurations":[{"max_duration_minutes":120}],"recommended_duration_minutes":null}`,
		},
		{
			name:        "non json bodies",
			contentType: "text/plain",
			body:        `request_configurations`,
		},
		{
			name:        "malformed json",
			contentType: "application/json",
			body:        `{"request_configurations":[`,
		},
	} {
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			header := make(http.Header)
			header.Set("Content-Type", testCase.contentType)
			res := &http.Response{
				StatusCode: http.StatusOK,
				Header:     header,
				Body:       io.NopCloser(strings.NewReader(testCase.body)),
			}
			if got := runHook(t, res); got != testCase.body {
				t.Errorf("body was rewritten:\n got: %s\nwant: %s", got, testCase.body)
			}
		})
	}
}

// Decoding through float64 would render large integers in scientific notation
// and silently corrupt them.
func TestMaxDurationIndefiniteHookPreservesLargeIntegers(t *testing.T) {
	t.Parallel()

	body := `{"cursor":9007199254740993,"request_configurations":[{"priority":0}]}`
	got := runHook(t, jsonResponse(body))
	if !strings.Contains(got, "9007199254740993") {
		t.Errorf("large integer was reformatted: %s", got)
	}
}

func TestMaxDurationIndefiniteHookHandlesEmptyBody(t *testing.T) {
	t.Parallel()

	res := &http.Response{StatusCode: http.StatusNoContent, Header: make(http.Header), Body: http.NoBody}
	if _, err := (&maxDurationIndefiniteHook{}).AfterSuccess(AfterSuccessContext{}, res); err != nil {
		t.Fatalf("AfterSuccess returned error: %v", err)
	}
	if _, err := (&maxDurationIndefiniteHook{}).AfterSuccess(AfterSuccessContext{}, nil); err != nil {
		t.Fatalf("AfterSuccess on nil response returned error: %v", err)
	}
}
