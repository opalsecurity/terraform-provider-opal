//go:build e2e

package e2e

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/opalsecurity/terraform-provider-opal/v3/internal/sdk"
	"github.com/opalsecurity/terraform-provider-opal/v3/internal/sdk/models/operations"
	"github.com/opalsecurity/terraform-provider-opal/v3/internal/sdk/models/shared"
)

// env holds the configuration used by every e2e test. Resolved once per test
// run inside requireEnv to keep the per-sub-test setup noise down.
type env struct {
	BaseURL       string
	Token         string
	ResourceAppID string // CUSTOM-type app that owns the test resources
	GroupAppID    string // OPAL-type app that owns the test groups
	AdminOwnerID  string
}

// requireEnv returns the e2e configuration or t.Skip()s the test. The default
// base URL / apps / owner match the local mprocs dev DB (see
// internal/e2e/README.md); only the bearer token is strictly required.
func requireEnv(t *testing.T) env {
	t.Helper()

	token := os.Getenv("OPAL_E2E_TOKEN")
	if token == "" {
		t.Skip("OPAL_E2E_TOKEN not set; skipping e2e test (see internal/e2e/README.md)")
	}

	return env{
		Token:   token,
		BaseURL: getenvDefault("OPAL_E2E_BASE_URL", "http://localhost:3000/v1"),
		ResourceAppID: getenvDefault(
			"OPAL_E2E_RESOURCE_APP_ID",
			"849aa50a-5869-4988-8396-e8aee2f4a538",
		),
		GroupAppID: getenvDefault(
			"OPAL_E2E_GROUP_APP_ID",
			"4d16279d-252e-4f5b-92fd-2819fab06ab0",
		),
		AdminOwnerID: getenvDefault(
			"OPAL_E2E_ADMIN_OWNER_ID",
			"ee664aad-b2f5-4011-9e19-32776613825a",
		),
	}
}

func getenvDefault(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

// newClient builds the generated SDK against the configured base URL. Tests
// use the SDK rather than hand-rolled HTTP so a regression in the generator
// (schema → SDK) surfaces as a failing test rather than a silent client-side
// drift.
func newClient(e env) *sdk.OpalAPI {
	return sdk.New(
		sdk.WithServerURL(e.BaseURL),
		sdk.WithSecurity(shared.Security{BearerAuth: e.Token}),
	)
}

// uniqueName returns a per-run name prefix that keeps parallel sub-tests from
// colliding on the "name must be unique within the org" constraint.
func uniqueName(prefix string) string {
	return fmt.Sprintf("eprd3921-e2e-%s-%d", prefix, time.Now().UnixNano())
}

// createResource creates a CUSTOM resource in the configured app and registers
// a t.Cleanup that deletes it. Returns the resource ID.
func createResource(
	t *testing.T,
	ctx context.Context,
	client *sdk.OpalAPI,
	e env,
	suffix string,
) string {
	t.Helper()

	resp, err := client.Resources.Create(ctx, shared.CreateResourceInfo{
		AppID:        e.ResourceAppID,
		Name:         uniqueName("res-" + suffix),
		ResourceType: shared.ResourceTypeEnumCustom,
	})
	if err != nil {
		t.Fatalf("create resource: %v", err)
	}
	if resp.StatusCode/100 != 2 || resp.Resource == nil {
		t.Fatalf("create resource: status=%d body=%v", resp.StatusCode, resp.Resource)
	}
	id := resp.Resource.ID

	t.Cleanup(func() {
		del, err := client.Resources.Delete(
			context.Background(),
			operations.DeleteResourceRequest{ID: id},
		)
		if err != nil {
			t.Logf("cleanup: delete resource %s: %v", id, err)
			return
		}
		// 200 is the documented success; 404 is tolerated in case an earlier
		// step in the test (e.g. the unlink-fork path) rewrote the resource's
		// identity in a way we aren't modeling here.
		if del.StatusCode != http.StatusOK && del.StatusCode != http.StatusNotFound {
			t.Logf("cleanup: delete resource %s returned status=%d", id, del.StatusCode)
		}
	})

	return id
}

// createGroup creates an OPAL_GROUP in the configured OPAL app and registers
// a t.Cleanup that deletes it. Returns the group ID.
func createGroup(
	t *testing.T,
	ctx context.Context,
	client *sdk.OpalAPI,
	e env,
	suffix string,
) string {
	t.Helper()

	resp, err := client.Groups.Create(ctx, shared.CreateGroupInfo{
		AppID:     e.GroupAppID,
		Name:      uniqueName("grp-" + suffix),
		GroupType: shared.GroupTypeEnumOpalGroup,
	})
	if err != nil {
		t.Fatalf("create group: %v", err)
	}
	if resp.StatusCode/100 != 2 || resp.Group == nil {
		t.Fatalf("create group: status=%d body=%v", resp.StatusCode, resp.Group)
	}
	id := resp.Group.ID

	t.Cleanup(func() {
		del, err := client.Groups.Delete(
			context.Background(),
			operations.DeleteGroupRequest{ID: id},
		)
		if err != nil {
			t.Logf("cleanup: delete group %s: %v", id, err)
			return
		}
		// 200 is the documented success; 404 is tolerated in case an earlier
		// step in the test rewrote the group's identity in a way we aren't
		// modeling here.
		if del.StatusCode != http.StatusOK && del.StatusCode != http.StatusNotFound {
			t.Logf("cleanup: delete group %s returned status=%d", id, del.StatusCode)
		}
	})

	return id
}

// createTemplate creates a configuration template with one minimal
// request-configuration row (so the unlink-fork test can observe that the
// forked private config inherits request configurations from the template)
// and registers a t.Cleanup to delete it.
func createTemplate(
	t *testing.T,
	ctx context.Context,
	client *sdk.OpalAPI,
	e env,
	suffix string,
) string {
	t.Helper()

	resp, err := client.ConfigurationTemplates.Create(ctx, shared.CreateConfigurationTemplateInfo{
		Name:                    uniqueName("tpl-" + suffix),
		AdminOwnerID:            e.AdminOwnerID,
		RequireMfaToConnect:     false,
		MemberOnCallScheduleIds: []string{},
		Visibility:              shared.VisibilityInfo{Visibility: shared.VisibilityTypeEnumGlobal},
		RequestConfigurations: []shared.RequestConfiguration{
			{
				AllowRequests:        true,
				AutoApproval:         true,
				Priority:             0,
				RequireMfaToRequest:  false,
				RequireSupportTicket: false,
			},
		},
	})
	if err != nil {
		t.Fatalf("create template: %v", err)
	}
	if resp.StatusCode/100 != 2 || resp.ConfigurationTemplate == nil ||
		resp.ConfigurationTemplate.ConfigurationTemplateID == nil {
		t.Fatalf("create template: status=%d body=%+v", resp.StatusCode, resp.ConfigurationTemplate)
	}
	id := *resp.ConfigurationTemplate.ConfigurationTemplateID

	t.Cleanup(func() {
		del, err := client.ConfigurationTemplates.Delete(
			context.Background(),
			operations.DeleteConfigurationTemplateRequest{ConfigurationTemplateID: id},
		)
		if err != nil {
			t.Logf("cleanup: delete template %s: %v", id, err)
			return
		}
		if del.StatusCode != http.StatusOK && del.StatusCode != http.StatusNotFound {
			t.Logf("cleanup: delete template %s returned status=%d", id, del.StatusCode)
		}
	})

	return id
}

// strPtr is the one-liner the generated SDK's nullable fields want.
func strPtr(s string) *string { return &s }

// putResourceLink is a thin wrapper around UpdateResourceConfigurationTemplate
// that returns the resulting link (nil templateID means unlink).
func putResourceLink(
	t *testing.T,
	ctx context.Context,
	client *sdk.OpalAPI,
	resourceID string,
	templateID *string,
) *operations.UpdateResourceConfigurationTemplateResponse {
	t.Helper()
	resp, err := client.Resources.UpdateResourceConfigurationTemplate(
		ctx,
		operations.UpdateResourceConfigurationTemplateRequest{
			ResourceID: resourceID,
			ResourceConfigurationTemplate: shared.ResourceConfigurationTemplate{
				ConfigurationTemplateID: templateID,
			},
		},
	)
	if err != nil {
		t.Fatalf("PUT resources/{id}/configuration-template: %v", err)
	}
	return resp
}

// getResourceLink is a thin wrapper around GetResourceConfigurationTemplate.
func getResourceLink(
	t *testing.T,
	ctx context.Context,
	client *sdk.OpalAPI,
	resourceID string,
) *operations.GetResourceConfigurationTemplateResponse {
	t.Helper()
	resp, err := client.Resources.GetResourceConfigurationTemplate(
		ctx,
		operations.GetResourceConfigurationTemplateRequest{
			ResourceID: resourceID,
		},
	)
	if err != nil {
		t.Fatalf("GET resources/{id}/configuration-template: %v", err)
	}
	return resp
}

// rawPutResourceLink bypasses the SDK and sends an arbitrary request body.
// The SDK's typed request can't model two of the failure modes the backend
// protects against — a malformed UUID string or an entirely omitted
// `configuration_template_id` field — so those cases round-trip through this
// helper instead.
func rawPutResourceLink(t *testing.T, e env, resourceID, body string) *http.Response {
	t.Helper()
	return rawPutLink(t, e, "resources/"+resourceID+"/configuration-template", body)
}

// putGroupLink is a thin wrapper around UpdateGroupConfigurationTemplate that
// returns the resulting link (nil templateID means unlink).
func putGroupLink(
	t *testing.T,
	ctx context.Context,
	client *sdk.OpalAPI,
	groupID string,
	templateID *string,
) *operations.UpdateGroupConfigurationTemplateResponse {
	t.Helper()
	resp, err := client.Groups.UpdateGroupConfigurationTemplate(
		ctx,
		operations.UpdateGroupConfigurationTemplateRequest{
			GroupID: groupID,
			GroupConfigurationTemplate: shared.GroupConfigurationTemplate{
				ConfigurationTemplateID: templateID,
			},
		},
	)
	if err != nil {
		t.Fatalf("PUT groups/{id}/configuration-template: %v", err)
	}
	return resp
}

// getGroupLink is a thin wrapper around GetGroupConfigurationTemplate.
func getGroupLink(
	t *testing.T,
	ctx context.Context,
	client *sdk.OpalAPI,
	groupID string,
) *operations.GetGroupConfigurationTemplateResponse {
	t.Helper()
	resp, err := client.Groups.GetGroupConfigurationTemplate(
		ctx,
		operations.GetGroupConfigurationTemplateRequest{
			GroupID: groupID,
		},
	)
	if err != nil {
		t.Fatalf("GET groups/{id}/configuration-template: %v", err)
	}
	return resp
}

// rawPutGroupLink is the groups-side equivalent of rawPutResourceLink.
func rawPutGroupLink(t *testing.T, e env, groupID, body string) *http.Response {
	t.Helper()
	return rawPutLink(t, e, "groups/"+groupID+"/configuration-template", body)
}

// rawPutLink is the shared transport used by the two typed-SDK bypass helpers
// above. The caller supplies the resource-/groups-rooted path (without a
// leading slash); this helper only cares about Authorization + Content-Type.
func rawPutLink(t *testing.T, e env, path, body string) *http.Response {
	t.Helper()
	req, err := http.NewRequest(
		http.MethodPut,
		e.BaseURL+"/"+path,
		bytes.NewReader([]byte(body)),
	)
	if err != nil {
		t.Fatalf("build raw request: %v", err)
	}
	req.Header.Set("Authorization", "Bearer "+e.Token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("raw PUT: %v", err)
	}
	return resp
}

// readBody reads and closes an http.Response body. Used for raw-request
// assertions that need to inspect the error payload.
func readBody(t *testing.T, resp *http.Response) string {
	t.Helper()
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("read body: %v", err)
	}
	return string(b)
}
