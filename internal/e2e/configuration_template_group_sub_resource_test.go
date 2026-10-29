//go:build e2e

package e2e

import (
	"context"
	"net/http"
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/opalsecurity/terraform-provider-opal/v3/internal/sdk/models/operations"
)

// TestE2EGroupConfigurationTemplate exercises the GET and PUT sub-resource
// endpoints added in the EPRD-4793 groups stack through the generated SDK
// against a live backend.
//
// The sub-tests mirror the 14 cases in the backend's own integration test
// (`TestAPIGroupConfigurationTemplate` in opal), minus two cases that the
// typed SDK request cannot model — a non-UUID string value for the ID and a
// completely omitted field. Those are exercised through a raw HTTP helper
// (`rawPutGroupLink`) in a dedicated sub-test block at the bottom so they
// still run under the same `make e2e` entry point.
//
// Each sub-test creates its own group and templates via t.Cleanup, so the
// suite can be re-run as many times as needed without leaving fixtures
// behind. Template creation is registered *before* group creation so the
// LIFO t.Cleanup order releases the group's link to the template before the
// template itself is deleted (otherwise the delete fails with "template is
// in use").
func TestE2EGroupConfigurationTemplate(t *testing.T) {
	e := requireEnv(t)
	client := newClient(e)
	ctx := context.Background()

	// -------------------- GET -----------------------------------------------

	t.Run("get_returns_linked_template", func(t *testing.T) {
		t.Parallel()
		templateID := createTemplate(t, ctx, client, e, "get-linked")
		groupID := createGroup(t, ctx, client, e, "get-linked")

		// Seed the link via the sub-resource PUT. The point of the GET test
		// is read-after-write via the sub-resource, not the attach itself.
		putResp := putGroupLink(t, ctx, client, groupID, strPtr(templateID))
		if putResp.StatusCode != http.StatusOK {
			t.Fatalf("seed attach: status=%d", putResp.StatusCode)
		}

		resp := getGroupLink(t, ctx, client, groupID)
		if resp.StatusCode != http.StatusOK {
			t.Fatalf("GET status=%d", resp.StatusCode)
		}
		if resp.GroupConfigurationTemplate == nil ||
			resp.GroupConfigurationTemplate.ConfigurationTemplateID == nil {
			t.Fatalf(
				"GET body: expected linked template, got %+v",
				resp.GroupConfigurationTemplate,
			)
		}
		if got := *resp.GroupConfigurationTemplate.ConfigurationTemplateID; got != templateID {
			t.Fatalf("GET body: template id = %s, want %s", got, templateID)
		}
	})

	t.Run("get_returns_null_when_unlinked", func(t *testing.T) {
		t.Parallel()
		groupID := createGroup(t, ctx, client, e, "get-unlinked")

		resp := getGroupLink(t, ctx, client, groupID)
		if resp.StatusCode != http.StatusOK {
			t.Fatalf("GET status=%d", resp.StatusCode)
		}
		if resp.GroupConfigurationTemplate == nil {
			t.Fatalf("GET body: expected present body with null id, got nil")
		}
		if resp.GroupConfigurationTemplate.ConfigurationTemplateID != nil {
			t.Fatalf(
				"GET body: expected null template id, got %q",
				*resp.GroupConfigurationTemplate.ConfigurationTemplateID,
			)
		}
	})

	t.Run("get_404_for_nonexistent_group", func(t *testing.T) {
		t.Parallel()
		resp := getGroupLink(t, ctx, client, uuid.NewString())
		if resp.StatusCode != http.StatusNotFound {
			t.Fatalf("GET status=%d, want 404", resp.StatusCode)
		}
	})

	// -------------------- PUT: state machine --------------------------------

	t.Run("put_attach_from_unlinked", func(t *testing.T) {
		t.Parallel()
		templateID := createTemplate(t, ctx, client, e, "put-attach")
		groupID := createGroup(t, ctx, client, e, "put-attach")

		resp := putGroupLink(t, ctx, client, groupID, strPtr(templateID))
		if resp.StatusCode != http.StatusOK {
			t.Fatalf("PUT status=%d", resp.StatusCode)
		}
		// PUT no longer returns a body (x-speakeasy-ignore'd); verify via GET.
		assertGroupLinkedTo(t, ctx, client, groupID, strPtr(templateID), "attach")
	})

	t.Run("put_unlink_forks_template_and_preserves_request_configurations", func(t *testing.T) {
		t.Parallel()
		templateID := createTemplate(t, ctx, client, e, "put-unlink")
		groupID := createGroup(t, ctx, client, e, "put-unlink")

		putResp := putGroupLink(t, ctx, client, groupID, strPtr(templateID))
		if putResp.StatusCode != http.StatusOK {
			t.Fatalf("seed attach: status=%d", putResp.StatusCode)
		}

		resp := putGroupLink(t, ctx, client, groupID, nil) // unlink
		if resp.StatusCode != http.StatusOK {
			t.Fatalf("unlink status=%d", resp.StatusCode)
		}
		// PUT no longer returns a body (x-speakeasy-ignore'd); verify via GET.
		assertGroupLinkedTo(t, ctx, client, groupID, nil, "unlink")

		// Read via main GET /groups/{id} to confirm the forked private
		// configuration still carries the request configurations the template
		// had (the whole point of fork-on-unlink: behaviour is preserved).
		full, err := client.Groups.GetGroup(
			ctx,
			operations.GetGroupRequest{ID: groupID},
		)
		if err != nil {
			t.Fatalf("GET group after unlink: %v", err)
		}
		if full.StatusCode != http.StatusOK || full.Group == nil {
			t.Fatalf(
				"GET group after unlink: status=%d group=%v",
				full.StatusCode,
				full.Group,
			)
		}
		if len(full.Group.RequestConfigurations) == 0 {
			t.Fatalf(
				"Group.RequestConfigurations after unlink: want >= 1 (fork should preserve the template's rows), got 0",
			)
		}
	})

	t.Run("put_repoint_template_a_to_template_b", func(t *testing.T) {
		t.Parallel()
		templateA := createTemplate(t, ctx, client, e, "put-repoint-a")
		templateB := createTemplate(t, ctx, client, e, "put-repoint-b")
		groupID := createGroup(t, ctx, client, e, "put-repoint")

		if r := putGroupLink(
			t,
			ctx,
			client,
			groupID,
			strPtr(templateA),
		); r.StatusCode != http.StatusOK {
			t.Fatalf("seed attach to A: status=%d", r.StatusCode)
		}
		resp := putGroupLink(t, ctx, client, groupID, strPtr(templateB))
		if resp.StatusCode != http.StatusOK {
			t.Fatalf("repoint status=%d", resp.StatusCode)
		}
		// PUT no longer returns a body (x-speakeasy-ignore'd); verify via GET.
		assertGroupLinkedTo(t, ctx, client, groupID, strPtr(templateB), "repoint")
	})

	t.Run("put_noop_when_null_to_null", func(t *testing.T) {
		t.Parallel()
		groupID := createGroup(t, ctx, client, e, "put-noop-null")

		// Create starts unlinked; PUT null must be accepted as a no-op.
		resp := putGroupLink(t, ctx, client, groupID, nil)
		if resp.StatusCode != http.StatusOK {
			t.Fatalf("null->null status=%d", resp.StatusCode)
		}
		assertGroupLinkedTo(t, ctx, client, groupID, nil, "null->null")
	})

	t.Run("put_noop_when_already_linked_to_same_template", func(t *testing.T) {
		t.Parallel()
		templateID := createTemplate(t, ctx, client, e, "put-noop-same")
		groupID := createGroup(t, ctx, client, e, "put-noop-same")

		if r := putGroupLink(
			t,
			ctx,
			client,
			groupID,
			strPtr(templateID),
		); r.StatusCode != http.StatusOK {
			t.Fatalf("seed attach: status=%d", r.StatusCode)
		}
		resp := putGroupLink(t, ctx, client, groupID, strPtr(templateID))
		if resp.StatusCode != http.StatusOK {
			t.Fatalf("same-template status=%d", resp.StatusCode)
		}
		assertGroupLinkedTo(t, ctx, client, groupID, strPtr(templateID), "same-template")
	})

	// -------------------- PUT: validation / error cases ---------------------

	t.Run("put_400_when_target_does_not_exist", func(t *testing.T) {
		t.Parallel()
		groupID := createGroup(t, ctx, client, e, "put-400-missing-tpl")

		resp := putGroupLink(t, ctx, client, groupID, strPtr(uuid.NewString()))
		if resp.StatusCode != http.StatusBadRequest {
			t.Fatalf("missing-template PUT status=%d, want 400", resp.StatusCode)
		}
	})

	t.Run("put_404_for_nonexistent_group", func(t *testing.T) {
		t.Parallel()
		resp := putGroupLink(t, ctx, client, uuid.NewString(), nil)
		if resp.StatusCode != http.StatusNotFound {
			t.Fatalf("missing-group PUT status=%d, want 404", resp.StatusCode)
		}
	})

	// Wire-level validation that the generated SDK can't model: a non-UUID
	// id value and an entirely omitted `configuration_template_id` field.
	// These matter for anything hand-written against the REST API, so cover
	// them via a raw HTTP client.
	t.Run("put_400_when_body_is_malformed", func(t *testing.T) {
		t.Parallel()
		groupID := createGroup(t, ctx, client, e, "put-400-malformed")

		resp := rawPutGroupLink(t, e, groupID, `{"configuration_template_id":"not-a-uuid"}`)
		if resp.StatusCode != http.StatusBadRequest {
			t.Fatalf("malformed body PUT status=%d, want 400", resp.StatusCode)
		}
		_ = readBody(t, resp)
	})

	t.Run("put_400_when_configuration_template_id_field_is_omitted", func(t *testing.T) {
		t.Parallel()
		groupID := createGroup(t, ctx, client, e, "put-400-omitted")

		resp := rawPutGroupLink(t, e, groupID, `{}`)
		if resp.StatusCode != http.StatusBadRequest {
			t.Fatalf("omitted-field PUT status=%d, want 400", resp.StatusCode)
		}
		body := readBody(t, resp)
		// The backend commits to this exact phrasing (see opal
		// TestAPIGroupConfigurationTemplate/put_400_when_configuration_template_id_field_is_omitted),
		// so pin it here too — the message is part of the public contract.
		const want = "configuration_template_id is required"
		if !strings.Contains(body, want) {
			t.Fatalf("omitted-field error body = %q, want it to contain %q", body, want)
		}
	})
}
