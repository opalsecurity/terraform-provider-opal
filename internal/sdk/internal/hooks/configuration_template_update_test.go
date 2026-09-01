package hooks

import (
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestConfigurationTemplateUpdateHook(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		operationID string
		body        string
		wantBody    string
	}{
		{
			name:        "group attach sends REST-compatible minimal update",
			operationID: "updateGroups",
			body:        `{"groups":[{"group_id":"group-id","configuration_template_id":"template-id","admin_owner_id":"owner-id","custom_request_notification":"message","request_configurations":[],"require_mfa_to_approve":false,"name":"Engineering"}]}`,
			wantBody:    `{"groups":[{"configuration_template_id":"template-id","group_id":"group-id"}]}`,
		},
		{
			name:        "resource attach sends REST-compatible minimal update",
			operationID: "updateResources",
			body:        `{"resources":[{"resource_id":"resource-id","configuration_template_id":"template-id","admin_owner_id":"owner-id","custom_request_notification":"message","description":"Production repository","match_remote_name":false,"request_configurations":[],"require_mfa_to_approve":false,"require_mfa_to_connect":false,"ticket_propagation":{"enabled_on_grant":true},"name":"Production"}]}`,
			wantBody:    `{"resources":[{"configuration_template_id":"template-id","resource_id":"resource-id"}]}`,
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

func TestConfigurationTemplateUpdateHookSuppressesBlockedFollowUp(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		operationID string
		statusCode  int
		body        string
		wantStatus  int
	}{
		{
			name:        "linked resource visibility update",
			operationID: "updateResourceVisibility",
			statusCode:  http.StatusBadRequest,
			body:        `{"detail":"Cannot use API to update resources linked to configuration templates"}`,
			wantStatus:  http.StatusOK,
		},
		{
			name:        "linked group message channel update",
			operationID: "updateGroupMessageChannels",
			statusCode:  http.StatusBadRequest,
			body:        `{"detail":"Cannot use API to update groups linked to configuration templates"}`,
			wantStatus:  http.StatusOK,
		},
		{
			name:        "unrelated bad request",
			operationID: "updateResourceVisibility",
			statusCode:  http.StatusBadRequest,
			body:        `{"detail":"invalid visibility"}`,
			wantStatus:  http.StatusBadRequest,
		},
		{
			name:        "empty visibility level on group visibility update",
			operationID: "updateGroupVisibility",
			statusCode:  http.StatusBadRequest,
			body:        `{"Status":"Bad Request","Message":"Unrecognized visibility level: "}`,
			wantStatus:  http.StatusOK,
		},
		{
			name:        "empty visibility level on resource visibility update",
			operationID: "updateResourceVisibility",
			statusCode:  http.StatusBadRequest,
			body:        `{"Status":"Bad Request","Message":"Unrecognized visibility level: "}`,
			wantStatus:  http.StatusOK,
		},
		{
			name:        "non-empty invalid visibility level is a genuine error",
			operationID: "updateGroupVisibility",
			statusCode:  http.StatusBadRequest,
			body:        `{"Status":"Bad Request","Message":"Unrecognized visibility level: FOO"}`,
			wantStatus:  http.StatusBadRequest,
		},
		{
			name:        "visibility message on non-visibility operation is a genuine error",
			operationID: "updateGroupMessageChannels",
			statusCode:  http.StatusBadRequest,
			body:        `{"Status":"Bad Request","Message":"Unrecognized visibility level: "}`,
			wantStatus:  http.StatusBadRequest,
		},
		{
			name:        "unrelated operation",
			operationID: "updateResources",
			statusCode:  http.StatusBadRequest,
			body:        `{"detail":"Cannot use API to update resources linked to configuration templates"}`,
			wantStatus:  http.StatusBadRequest,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			hook := &configurationTemplateUpdateHook{}
			res := &http.Response{
				StatusCode: test.statusCode,
				Status:     http.StatusText(test.statusCode),
				Header:     make(http.Header),
				Body:       io.NopCloser(strings.NewReader(test.body)),
			}

			got, err := hook.AfterError(AfterErrorContext{
				HookContext: HookContext{OperationID: test.operationID},
			}, res, nil)
			require.NoError(t, err)
			require.Equal(t, test.wantStatus, got.StatusCode)

			if test.wantStatus == http.StatusOK {
				require.Equal(t, http.NoBody, got.Body)
			} else {
				gotBody, readErr := io.ReadAll(got.Body)
				require.NoError(t, readErr)
				require.JSONEq(t, test.body, string(gotBody))
			}
		})
	}
}
