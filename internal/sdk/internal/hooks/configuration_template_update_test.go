package hooks

import (
	"encoding/json"
	"io"
	"net/http"
	"reflect"
	"strings"
	"testing"
)

// Minimal assertions so the hook tests need no extra module dependency.
type requireT struct{}

var require requireT

func (requireT) NoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func (requireT) Equal(t *testing.T, want, got any, msg ...string) {
	t.Helper()
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("want %v, got %v %s", want, got, strings.Join(msg, " "))
	}
}

func (requireT) False(t *testing.T, v bool) {
	t.Helper()
	if v {
		t.Fatal("expected false")
	}
}

func (requireT) Len(t *testing.T, v []string, n int, msg ...string) {
	t.Helper()
	if len(v) != n {
		t.Fatalf("want len %d, got %d (%v) %s", n, len(v), v, strings.Join(msg, " "))
	}
}

func (requireT) JSONEq(t *testing.T, want, got string) {
	t.Helper()
	var w, g any
	if err := json.Unmarshal([]byte(want), &w); err != nil {
		t.Fatalf("bad want JSON %q: %v", want, err)
	}
	if err := json.Unmarshal([]byte(got), &g); err != nil {
		t.Fatalf("bad got JSON %q: %v", got, err)
	}
	if !reflect.DeepEqual(w, g) {
		t.Fatalf("JSON mismatch\nwant %s\n got %s", want, got)
	}
}

type roundTripFunc func(*http.Request) (*http.Response, error)

func (f roundTripFunc) Do(req *http.Request) (*http.Response, error) {
	return f(req)
}

type templateHookHarness struct {
	t        *testing.T
	hook     *configurationTemplateUpdateHook
	client   HTTPClient
	requests []string
}

func newTemplateHookHarness(t *testing.T) *templateHookHarness {
	t.Helper()
	h := &templateHookHarness{t: t, hook: &configurationTemplateUpdateHook{}}
	upstream := roundTripFunc(func(req *http.Request) (*http.Response, error) {
		h.requests = append(h.requests, req.Method+" "+req.URL.Path)
		return &http.Response{
			StatusCode: http.StatusOK,
			Status:     "200 OK",
			Header:     make(http.Header),
			Body:       http.NoBody,
			Request:    req,
		}, nil
	})
	_, h.client = h.hook.SDKInit("https://example.com/v1", upstream)
	return h
}

func (h *templateHookHarness) put(path, body string) *http.Response {
	h.t.Helper()
	req, err := http.NewRequest(
		http.MethodPut,
		"https://example.com/v1"+path,
		strings.NewReader(body),
	)
	require.NoError(h.t, err)
	res, err := h.client.Do(req)
	require.NoError(h.t, err)
	return res
}

// parentBody runs the BeforeRequest hook for a parent update and returns the body.
func (h *templateHookHarness) parentBody(operationID, body string) string {
	h.t.Helper()
	req, err := http.NewRequest(http.MethodPut, "https://example.com/v1/x", strings.NewReader(body))
	require.NoError(h.t, err)
	req, err = h.hook.BeforeRequest(
		BeforeRequestContext{HookContext: HookContext{OperationID: operationID}},
		req,
	)
	require.NoError(h.t, err)
	got, err := io.ReadAll(req.Body)
	require.NoError(h.t, err)
	return string(got)
}

const (
	groupParentBody    = `{"groups":[{"group_id":"g1","admin_owner_id":"owner","custom_request_notification":"hi","request_configurations":[],"require_mfa_to_approve":false,"requires_forfeit_reason":false,"name":"Engineering"}]}`
	groupStrippedBody  = `{"groups":[{"group_id":"g1","name":"Engineering"}]}`
	resourceParentBody = `{"resources":[{"resource_id":"r1","admin_owner_id":"owner","description":"Prod","match_remote_name":false,"request_configurations":[],"require_mfa_to_approve":false,"require_mfa_to_connect":false,"ticket_propagation":{"enabled_on_grant":true},"name":"Production"}]}`
	resourceStripped   = `{"resources":[{"resource_id":"r1","description":"Prod","match_remote_name":false,"name":"Production"}]}`
)

func TestConfigurationTemplateHookStripsParentPutWhileLinked(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name, operationID, linkPath, body, want string
	}{
		{
			"group",
			"updateGroups",
			"/groups/g1/configuration-template",
			groupParentBody,
			groupStrippedBody,
		},
		{
			"resource",
			"updateResources",
			"/resources/r1/configuration-template",
			resourceParentBody,
			resourceStripped,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			h := newTemplateHookHarness(t)

			h.put(test.linkPath, `{"configuration_template_id":"template-1"}`)

			require.JSONEq(t, test.want, h.parentBody(test.operationID, test.body))
		})
	}
}

func TestConfigurationTemplateHookLeavesUnlinkedParentPutAlone(t *testing.T) {
	t.Parallel()

	t.Run("never linked", func(t *testing.T) {
		t.Parallel()
		h := newTemplateHookHarness(t)
		require.JSONEq(t, groupParentBody, h.parentBody("updateGroups", groupParentBody))
	})

	t.Run("null template PUT first", func(t *testing.T) {
		t.Parallel()
		h := newTemplateHookHarness(t)
		h.put("/groups/g1/configuration-template", `{"configuration_template_id":null}`)
		require.JSONEq(t, groupParentBody, h.parentBody("updateGroups", groupParentBody))
	})

	t.Run("unlink after link", func(t *testing.T) {
		t.Parallel()
		h := newTemplateHookHarness(t)
		h.put("/groups/g1/configuration-template", `{"configuration_template_id":"template-1"}`)
		h.put("/groups/g1/configuration-template", `{"configuration_template_id":null}`)
		require.JSONEq(t, groupParentBody, h.parentBody("updateGroups", groupParentBody))
	})

	t.Run("other entity is not affected", func(t *testing.T) {
		t.Parallel()
		h := newTemplateHookHarness(t)
		h.put("/groups/other/configuration-template", `{"configuration_template_id":"template-1"}`)
		require.JSONEq(t, groupParentBody, h.parentBody("updateGroups", groupParentBody))
	})

	t.Run("unrelated operation", func(t *testing.T) {
		t.Parallel()
		h := newTemplateHookHarness(t)
		h.put("/groups/g1/configuration-template", `{"configuration_template_id":"template-1"}`)
		body := `{"group_id":"g1","admin_owner_id":"owner"}`
		require.JSONEq(t, body, h.parentBody("createGroup", body))
	})
}

func TestConfigurationTemplateHookIgnoresFailedLink(t *testing.T) {
	t.Parallel()

	hook := &configurationTemplateUpdateHook{}
	failing := roundTripFunc(func(req *http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: http.StatusBadRequest,
			Status:     "400",
			Header:     make(http.Header),
			Body:       http.NoBody,
			Request:    req,
		}, nil
	})
	_, client := hook.SDKInit("https://example.com/v1", failing)

	req, err := http.NewRequest(
		http.MethodPut,
		"https://example.com/v1/groups/g1/configuration-template",
		strings.NewReader(`{"configuration_template_id":"template-1"}`),
	)
	require.NoError(t, err)
	_, err = client.Do(req)
	require.NoError(t, err)

	require.False(t, hook.isLinked("groups", "g1"))
}

func TestConfigurationTemplateHookSkipsSiblingsWhileLinked(t *testing.T) {
	t.Parallel()

	h := newTemplateHookHarness(t)
	h.put("/groups/g1/configuration-template", `{"configuration_template_id":"template-1"}`)
	h.put("/resources/r1/configuration-template", `{"configuration_template_id":"template-1"}`)
	require.Len(t, h.requests, 2)

	for _, operation := range []string{"message-channels", "on-call-schedules", "visibility"} {
		res := h.put("/groups/g1/"+operation, `{"visibility":"GLOBAL"}`)
		require.Equal(t, http.StatusOK, res.StatusCode)
	}
	res := h.put("/resources/r1/visibility", `{"visibility":"GLOBAL"}`)
	require.Equal(t, http.StatusOK, res.StatusCode)
	require.Len(t, h.requests, 2, "linked sibling PUTs must not reach the API")

	// Resources have no template-owned message-channel/on-call sibling in this hook.
	h.put("/resources/r1/message-channels", `{"message_channel_ids":[]}`)
	require.Len(t, h.requests, 3)

	// An unlinked entity's siblings go through.
	h.put("/groups/other/visibility", `{"visibility":"GLOBAL"}`)
	h.put("/groups/other/message-channels", `{"message_channel_ids":[]}`)
	require.Len(t, h.requests, 5)
}

func TestConfigurationTemplateHookSkipsEmptyVisibilityBeforeLink(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		body     string
		wantSent bool
	}{
		{"missing visibility", `{"visibility_group_ids":[]}`, false},
		{"null visibility", `{"visibility":null}`, false},
		{"empty visibility", `{"visibility":""}`, false},
		{"set visibility", `{"visibility":"GLOBAL"}`, true},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			h := newTemplateHookHarness(t)
			h.put("/groups/g1/visibility", test.body)
			require.Equal(t, test.wantSent, len(h.requests) == 1)
		})
	}

	t.Run("empty message-channels still sent while unlinked", func(t *testing.T) {
		t.Parallel()
		h := newTemplateHookHarness(t)
		h.put("/groups/g1/message-channels", `{"message_channel_ids":[]}`)
		require.Len(t, h.requests, 1)
	})
}

func TestConfigurationTemplateHookForwardsLinkRequestBody(t *testing.T) {
	t.Parallel()

	var got string
	hook := &configurationTemplateUpdateHook{}
	upstream := roundTripFunc(func(req *http.Request) (*http.Response, error) {
		b, _ := io.ReadAll(req.Body)
		got = string(b)
		return &http.Response{
			StatusCode: http.StatusOK,
			Status:     "200",
			Header:     make(http.Header),
			Body:       http.NoBody,
			Request:    req,
		}, nil
	})
	_, client := hook.SDKInit("https://example.com/v1", upstream)

	body := `{"configuration_template_id":"template-1"}`
	req, err := http.NewRequest(
		http.MethodPut,
		"https://example.com/v1/groups/g1/configuration-template",
		strings.NewReader(body),
	)
	require.NoError(t, err)
	_, err = client.Do(req)
	require.NoError(t, err)
	require.Equal(t, body, got)
}
