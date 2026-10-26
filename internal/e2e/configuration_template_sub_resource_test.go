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

// TestE2EResourceConfigurationTemplate exercises the GET and PUT sub-resource
// endpoints added in opal#32208 (EPRD-4793) through the generated SDK against
// a live backend.
//
// The sub-tests mirror the 14 cases in the backend's own integration test
// (`TestAPIResourceConfigurationTemplate` in opal), minus two cases that the
// typed SDK request cannot model — a non-UUID string value for the ID and a
// completely omitted field. Those are exercised through a raw HTTP helper
// (`rawPutResourceLink`) in a dedicated sub-test block at the bottom so they still
// run under the same `make e2e` entry point.
//
// Each sub-test creates its own resource and templates via t.Cleanup, so the
// suite can be re-run as many times as needed without leaving fixtures
// behind.
func TestE2EResourceConfigurationTemplate(t *testing.T) {
	e := requireEnv(t)
	client := newClient(e)
	ctx := context.Background()

	// -------------------- GET -----------------------------------------------

	t.Run("get_returns_linked_template", func(t *testing.T) {
		t.Parallel()
		// Template first so its cleanup (LIFO) runs after the resource's;
		// otherwise the delete fails with "template is in use by 1 resource".
		templateID := createTemplate(t, ctx, client, e, "get-linked")
		resourceID := createResource(t, ctx, client, e, "get-linked")

		// Seed the link via the sub-resource PUT. The point of the GET test
		// is read-after-write via the sub-resource, not the attach itself.
		putResp := putResourceLink(t, ctx, client, resourceID, strPtr(templateID))
		if putResp.StatusCode != http.StatusOK {
			t.Fatalf("seed attach: status=%d", putResp.StatusCode)
		}

		resp := getResourceLink(t, ctx, client, resourceID)
		if resp.StatusCode != http.StatusOK {
			t.Fatalf("GET status=%d", resp.StatusCode)
		}
		if resp.ResourceConfigurationTemplate == nil ||
			resp.ResourceConfigurationTemplate.ConfigurationTemplateID == nil {
			t.Fatalf(
				"GET body: expected linked template, got %+v",
				resp.ResourceConfigurationTemplate,
			)
		}
		if got := *resp.ResourceConfigurationTemplate.ConfigurationTemplateID; got != templateID {
			t.Fatalf("GET body: template id = %s, want %s", got, templateID)
		}
	})

	t.Run("get_returns_null_when_unlinked", func(t *testing.T) {
		t.Parallel()
		resourceID := createResource(t, ctx, client, e, "get-unlinked")

		resp := getResourceLink(t, ctx, client, resourceID)
		if resp.StatusCode != http.StatusOK {
			t.Fatalf("GET status=%d", resp.StatusCode)
		}
		if resp.ResourceConfigurationTemplate == nil {
			t.Fatalf("GET body: expected present body with null id, got nil")
		}
		if resp.ResourceConfigurationTemplate.ConfigurationTemplateID != nil {
			t.Fatalf(
				"GET body: expected null template id, got %q",
				*resp.ResourceConfigurationTemplate.ConfigurationTemplateID,
			)
		}
	})

	t.Run("get_404_for_nonexistent_resource", func(t *testing.T) {
		t.Parallel()
		resp := getResourceLink(t, ctx, client, uuid.NewString())
		if resp.StatusCode != http.StatusNotFound {
			t.Fatalf("GET status=%d, want 404", resp.StatusCode)
		}
	})

	// -------------------- PUT: state machine --------------------------------

	t.Run("put_attach_from_unlinked", func(t *testing.T) {
		t.Parallel()
		templateID := createTemplate(t, ctx, client, e, "put-attach")
		resourceID := createResource(t, ctx, client, e, "put-attach")

		resp := putResourceLink(t, ctx, client, resourceID, strPtr(templateID))
		if resp.StatusCode != http.StatusOK {
			t.Fatalf("PUT status=%d", resp.StatusCode)
		}
		if resp.ResourceConfigurationTemplate == nil ||
			resp.ResourceConfigurationTemplate.ConfigurationTemplateID == nil {
			t.Fatalf(
				"PUT body: expected linked template id, got %+v",
				resp.ResourceConfigurationTemplate,
			)
		}
		if got := *resp.ResourceConfigurationTemplate.ConfigurationTemplateID; got != templateID {
			t.Fatalf("PUT body: template id = %s, want %s", got, templateID)
		}

		// Read-after-write via the sibling GET must match.
		got := getResourceLink(t, ctx, client, resourceID)
		if got.ResourceConfigurationTemplate == nil ||
			got.ResourceConfigurationTemplate.ConfigurationTemplateID == nil ||
			*got.ResourceConfigurationTemplate.ConfigurationTemplateID != templateID {
			t.Fatalf(
				"GET after attach: body=%+v, want id=%s",
				got.ResourceConfigurationTemplate,
				templateID,
			)
		}
	})

	t.Run("put_unlink_forks_template_and_preserves_request_configurations", func(t *testing.T) {
		t.Parallel()
		templateID := createTemplate(t, ctx, client, e, "put-unlink")
		resourceID := createResource(t, ctx, client, e, "put-unlink")

		putResp := putResourceLink(t, ctx, client, resourceID, strPtr(templateID))
		if putResp.StatusCode != http.StatusOK {
			t.Fatalf("seed attach: status=%d", putResp.StatusCode)
		}

		resp := putResourceLink(t, ctx, client, resourceID, nil) // unlink
		if resp.StatusCode != http.StatusOK {
			t.Fatalf("unlink status=%d", resp.StatusCode)
		}
		if resp.ResourceConfigurationTemplate == nil {
			t.Fatalf("unlink body: expected present body with null id, got nil")
		}
		if resp.ResourceConfigurationTemplate.ConfigurationTemplateID != nil {
			t.Fatalf(
				"unlink body: template id = %q, want null",
				*resp.ResourceConfigurationTemplate.ConfigurationTemplateID,
			)
		}

		// Read via main GET /resources/{id} to confirm the forked private
		// configuration still carries the request configurations the template
		// had (the whole point of fork-on-unlink: behaviour is preserved).
		full, err := client.Resources.GetID(
			ctx,
			operations.GetResourceIDRequest{ID: resourceID},
		)
		if err != nil {
			t.Fatalf("GET resource after unlink: %v", err)
		}
		if full.StatusCode != http.StatusOK || full.Resource == nil {
			t.Fatalf(
				"GET resource after unlink: status=%d resource=%v",
				full.StatusCode,
				full.Resource,
			)
		}
		if full.Resource.ConfigurationTemplateID != nil {
			t.Fatalf(
				"Resource.ConfigurationTemplateID after unlink: want nil, got %q",
				*full.Resource.ConfigurationTemplateID,
			)
		}
		if len(full.Resource.RequestConfigurations) == 0 {
			t.Fatalf(
				"Resource.RequestConfigurations after unlink: want >= 1 (fork should preserve the template's rows), got 0",
			)
		}
	})

	t.Run("put_repoint_template_a_to_template_b", func(t *testing.T) {
		t.Parallel()
		templateA := createTemplate(t, ctx, client, e, "put-repoint-a")
		templateB := createTemplate(t, ctx, client, e, "put-repoint-b")
		resourceID := createResource(t, ctx, client, e, "put-repoint")

		if r := putResourceLink(
			t,
			ctx,
			client,
			resourceID,
			strPtr(templateA),
		); r.StatusCode != http.StatusOK {
			t.Fatalf("seed attach to A: status=%d", r.StatusCode)
		}
		resp := putResourceLink(t, ctx, client, resourceID, strPtr(templateB))
		if resp.StatusCode != http.StatusOK {
			t.Fatalf("repoint status=%d", resp.StatusCode)
		}
		if resp.ResourceConfigurationTemplate == nil ||
			resp.ResourceConfigurationTemplate.ConfigurationTemplateID == nil ||
			*resp.ResourceConfigurationTemplate.ConfigurationTemplateID != templateB {
			t.Fatalf(
				"repoint body: got=%+v, want id=%s",
				resp.ResourceConfigurationTemplate,
				templateB,
			)
		}
	})

	t.Run("put_noop_when_null_to_null", func(t *testing.T) {
		t.Parallel()
		resourceID := createResource(t, ctx, client, e, "put-noop-null")

		// Create starts unlinked; PUT null must be accepted as a no-op.
		resp := putResourceLink(t, ctx, client, resourceID, nil)
		if resp.StatusCode != http.StatusOK {
			t.Fatalf("null->null status=%d", resp.StatusCode)
		}
		if resp.ResourceConfigurationTemplate == nil ||
			resp.ResourceConfigurationTemplate.ConfigurationTemplateID != nil {
			t.Fatalf("null->null body: want null id, got %+v", resp.ResourceConfigurationTemplate)
		}
	})

	t.Run("put_noop_when_already_linked_to_same_template", func(t *testing.T) {
		t.Parallel()
		templateID := createTemplate(t, ctx, client, e, "put-noop-same")
		resourceID := createResource(t, ctx, client, e, "put-noop-same")

		if r := putResourceLink(
			t,
			ctx,
			client,
			resourceID,
			strPtr(templateID),
		); r.StatusCode != http.StatusOK {
			t.Fatalf("seed attach: status=%d", r.StatusCode)
		}
		resp := putResourceLink(t, ctx, client, resourceID, strPtr(templateID))
		if resp.StatusCode != http.StatusOK {
			t.Fatalf("same-template status=%d", resp.StatusCode)
		}
		if resp.ResourceConfigurationTemplate == nil ||
			resp.ResourceConfigurationTemplate.ConfigurationTemplateID == nil ||
			*resp.ResourceConfigurationTemplate.ConfigurationTemplateID != templateID {
			t.Fatalf(
				"same-template body: got=%+v, want id=%s",
				resp.ResourceConfigurationTemplate,
				templateID,
			)
		}
	})

	// -------------------- PUT: validation / error cases ---------------------

	t.Run("put_400_when_target_does_not_exist", func(t *testing.T) {
		t.Parallel()
		resourceID := createResource(t, ctx, client, e, "put-400-missing-tpl")

		resp := putResourceLink(t, ctx, client, resourceID, strPtr(uuid.NewString()))
		if resp.StatusCode != http.StatusBadRequest {
			t.Fatalf("missing-template PUT status=%d, want 400", resp.StatusCode)
		}
	})

	t.Run("put_404_for_nonexistent_resource", func(t *testing.T) {
		t.Parallel()
		resp := putResourceLink(t, ctx, client, uuid.NewString(), nil)
		if resp.StatusCode != http.StatusNotFound {
			t.Fatalf("missing-resource PUT status=%d, want 404", resp.StatusCode)
		}
	})

	// Wire-level validation that the generated SDK can't model: a non-UUID
	// id value and an entirely omitted `configuration_template_id` field.
	// These matter for anything hand-written against the REST API, so cover
	// them via a raw HTTP client.
	t.Run("put_400_when_body_is_malformed", func(t *testing.T) {
		t.Parallel()
		resourceID := createResource(t, ctx, client, e, "put-400-malformed")

		resp := rawPutResourceLink(t, e, resourceID, `{"configuration_template_id":"not-a-uuid"}`)
		if resp.StatusCode != http.StatusBadRequest {
			t.Fatalf("malformed body PUT status=%d, want 400", resp.StatusCode)
		}
		_ = readBody(t, resp)
	})

	t.Run("put_400_when_configuration_template_id_field_is_omitted", func(t *testing.T) {
		t.Parallel()
		resourceID := createResource(t, ctx, client, e, "put-400-omitted")

		resp := rawPutResourceLink(t, e, resourceID, `{}`)
		if resp.StatusCode != http.StatusBadRequest {
			t.Fatalf("omitted-field PUT status=%d, want 400", resp.StatusCode)
		}
		body := readBody(t, resp)
		// The backend commits to this exact phrasing (see opal
		// TestAPIResourceConfigurationTemplate/put_400_when_configuration_template_id_field_is_omitted),
		// so pin it here too — the message is part of the public contract.
		const want = "configuration_template_id is required"
		if !strings.Contains(body, want) {
			t.Fatalf("omitted-field error body = %q, want it to contain %q", body, want)
		}
	})
}
