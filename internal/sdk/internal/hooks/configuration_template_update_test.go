package hooks

import (
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

type roundTripFunc func(*http.Request) (*http.Response, error)

func (f roundTripFunc) Do(req *http.Request) (*http.Response, error) {
	return f(req)
}

func TestConfigurationTemplateUpdateHook(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		operationID string
		body        string
		wantBody    string
	}{
		{
			name:        "group attach removes template-owned fields",
			operationID: "updateGroups",
			body:        `{"groups":[{"group_id":"group-id","configuration_template_id":"template-id","admin_owner_id":"owner-id","custom_request_notification":"message","request_configurations":[],"require_mfa_to_approve":false,"name":"Engineering"}]}`,
			wantBody:    `{"groups":[{"configuration_template_id":"template-id","group_id":"group-id","name":"Engineering"}]}`,
		},
		{
			name:        "resource attach removes template-owned fields",
			operationID: "updateResources",
			body:        `{"resources":[{"resource_id":"resource-id","configuration_template_id":"template-id","admin_owner_id":"owner-id","custom_request_notification":"message","description":"Production repository","match_remote_name":false,"request_configurations":[],"require_mfa_to_approve":false,"require_mfa_to_connect":false,"ticket_propagation":{"enabled_on_grant":true},"name":"Production"}]}`,
			wantBody:    `{"resources":[{"configuration_template_id":"template-id","description":"Production repository","match_remote_name":false,"name":"Production","resource_id":"resource-id"}]}`,
		},
		{
			name:        "ordinary update is unchanged",
			operationID: "updateGroups",
			body:        `{"groups":[{"group_id":"group-id","request_configurations":[]}]}`,
			wantBody:    `{"groups":[{"group_id":"group-id","request_configurations":[]}]}`,
		},
		{
			name:        "null template ID is unchanged",
			operationID: "updateResources",
			body:        `{"resources":[{"resource_id":"resource-id","configuration_template_id":null,"name":"Production"}]}`,
			wantBody:    `{"resources":[{"resource_id":"resource-id","configuration_template_id":null,"name":"Production"}]}`,
		},
		{
			name:        "empty template ID is unchanged",
			operationID: "updateGroups",
			body:        `{"groups":[{"group_id":"group-id","configuration_template_id":"","name":"Engineering"}]}`,
			wantBody:    `{"groups":[{"group_id":"group-id","configuration_template_id":"","name":"Engineering"}]}`,
		},
		{
			name:        "unrelated operation is unchanged",
			operationID: "createGroup",
			body:        `{"configuration_template_id":"template-id","request_configurations":[]}`,
			wantBody:    `{"configuration_template_id":"template-id","request_configurations":[]}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			req, err := http.NewRequest(http.MethodPut, "https://example.com", strings.NewReader(test.body))
			require.NoError(t, err)

			hook := &configurationTemplateUpdateHook{}
			req, err = hook.BeforeRequest(BeforeRequestContext{
				HookContext: HookContext{OperationID: test.operationID},
			}, req)
			require.NoError(t, err)

			gotBody, err := io.ReadAll(req.Body)
			require.NoError(t, err)
			require.JSONEq(t, test.wantBody, string(gotBody))
		})
	}
}

func TestConfigurationTemplateUpdateHookSkipsFollowUpsAfterSuccessfulAttach(t *testing.T) {
	t.Parallel()

	requests := 0
	client := roundTripFunc(func(req *http.Request) (*http.Response, error) {
		requests++
		return &http.Response{
			StatusCode: http.StatusOK,
			Status:     "200 OK",
			Header:     make(http.Header),
			Body:       http.NoBody,
			Request:    req,
		}, nil
	})

	hook := &configurationTemplateUpdateHook{}
	_, wrapped := hook.SDKInit("https://example.com/v1", client)

	attach, err := http.NewRequest(
		http.MethodPut,
		"https://example.com/v1/groups",
		strings.NewReader(`{"groups":[{"group_id":"group-id","configuration_template_id":"template-id","admin_owner_id":"owner-id"}]}`),
	)
	require.NoError(t, err)
	attach, err = hook.BeforeRequest(BeforeRequestContext{
		HookContext: HookContext{OperationID: "updateGroups"},
	}, attach)
	require.NoError(t, err)

	_, err = wrapped.Do(attach)
	require.NoError(t, err)
	require.Equal(t, 1, requests)

	for _, operation := range []string{"message-channels", "on-call-schedules", "visibility"} {
		followUp, requestErr := http.NewRequest(
			http.MethodPut,
			"https://example.com/v1/groups/group-id/"+operation,
			http.NoBody,
		)
		require.NoError(t, requestErr)

		res, requestErr := wrapped.Do(followUp)
		require.NoError(t, requestErr)
		require.Equal(t, http.StatusOK, res.StatusCode)
		require.Equal(t, 1, requests, operation+" should not reach the API")
	}

	unrelated, err := http.NewRequest(
		http.MethodPut,
		"https://example.com/v1/groups/other-group/visibility",
		http.NoBody,
	)
	require.NoError(t, err)
	_, err = wrapped.Do(unrelated)
	require.NoError(t, err)
	require.Equal(t, 2, requests)
}
