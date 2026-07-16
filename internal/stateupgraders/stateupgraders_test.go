package stateupgraders_test

import (
	"context"
	"testing"

	fwresource "github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"github.com/opalsecurity/terraform-provider-opal/v3/internal/provider"
	"github.com/opalsecurity/terraform-provider-opal/v3/internal/stateupgraders"
)

// groupV0RawStateJSON is the raw state of an opal_group resource as written by
// the v2.x (schema version 0) provider, e.g. v2.0.1.
const groupV0RawStateJSON = `{
	"id": "f9df6f2b-9970-4d38-a13e-1b5a0b1f0f0f",
	"name": "Test Group",
	"description": "A test group",
	"group_type": "OPAL_GROUP",
	"app_id": "9a8f66f6-59ac-4e39-9e33-04ab6c4b0f0f",
	"admin_owner_id": "7c86c85d-0651-43e2-a748-d69d658b0f0f",
	"require_mfa_to_approve": false,
	"manage_resources": true,
	"remote_info": null,
	"visibility": "LIMITED",
	"visibility_group": [{"id": "1b978423-db0a-4037-a4cf-f79c60cb0f0f"}],
	"audit_message_channel": [{"id": "cd453f4c-737d-4b74-b4d0-71c1b1180f0f"}],
	"on_call_schedule": [{"id": "9546209c-42c2-4801-96d7-9ec42df00f0f"}],
	"resource": null,
	"request_configuration": [{
		"group_ids": null,
		"is_requestable": true,
		"auto_approval": false,
		"require_mfa_to_request": false,
		"require_support_ticket": false,
		"max_duration": 120,
		"recommended_duration": 60,
		"request_template_id": null,
		"priority": 0,
		"reviewer_stage": [{
			"operator": "AND",
			"require_manager_approval": false,
			"reviewer": [{"id": "3d34b1f4-9d3f-4b30-b9c3-9e5ec2960f0f"}]
		}]
	}]
}`

// resourceV0RawStateJSON is the raw state of an opal_resource resource as
// written by the v2.x (schema version 0) provider, e.g. v2.0.1.
const resourceV0RawStateJSON = `{
	"id": "f9df6f2b-9970-4d38-a13e-1b5a0b1f0f0f",
	"name": "Test Resource",
	"description": "A test resource",
	"resource_type": "CUSTOM",
	"app_id": "9a8f66f6-59ac-4e39-9e33-04ab6c4b0f0f",
	"admin_owner_id": "7c86c85d-0651-43e2-a748-d69d658b0f0f",
	"require_mfa_to_approve": false,
	"require_mfa_to_connect": false,
	"remote_info": null,
	"visibility": "LIMITED",
	"visibility_group": [{"id": "1b978423-db0a-4037-a4cf-f79c60cb0f0f"}],
	"audit_message_channel": [{"id": "cd453f4c-737d-4b74-b4d0-71c1b1180f0f"}],
	"on_call_schedule": null,
	"resource": null,
	"request_configuration": [{
		"group_ids": null,
		"is_requestable": true,
		"auto_approval": false,
		"require_mfa_to_request": false,
		"require_support_ticket": false,
		"max_duration": 120,
		"recommended_duration": 60,
		"request_template_id": null,
		"priority": 0,
		"reviewer_stage": [{
			"operator": "AND",
			"require_manager_approval": false,
			"reviewer": [{"id": "3d34b1f4-9d3f-4b30-b9c3-9e5ec2960f0f"}]
		}]
	}]
}`

// currentSchemaType returns the tftypes type of a resource's current schema,
// which is the type terraform-plugin-framework decodes upgraded state against.
func currentSchemaType(t *testing.T, ctx context.Context, r fwresource.Resource) tftypes.Type {
	t.Helper()
	schemaResp := fwresource.SchemaResponse{}
	r.Schema(ctx, fwresource.SchemaRequest{}, &schemaResp)
	if schemaResp.Diagnostics.HasError() {
		t.Fatalf("unexpected schema diagnostics: %v", schemaResp.Diagnostics)
	}
	return schemaResp.Schema.Type().TerraformType(ctx)
}

// runUpgrader runs a v0 state upgrader over raw v0 state JSON and returns the
// upgraded dynamic value.
func runUpgrader(
	t *testing.T,
	ctx context.Context,
	upgrader func(context.Context, fwresource.UpgradeStateRequest, *fwresource.UpgradeStateResponse),
	rawStateJSON string,
) *tfprotov6.DynamicValue {
	t.Helper()
	req := fwresource.UpgradeStateRequest{
		RawState: &tfprotov6.RawState{JSON: []byte(rawStateJSON)},
	}
	resp := &fwresource.UpgradeStateResponse{}
	upgrader(ctx, req, resp)
	if resp.Diagnostics.HasError() {
		t.Fatalf("unexpected upgrade diagnostics: %v", resp.Diagnostics)
	}
	if resp.DynamicValue == nil {
		t.Fatal("expected upgrader to set a DynamicValue")
	}
	return resp.DynamicValue
}

// TestGroupStateUpgraderV0MatchesCurrentSchema upgrades v2.x opal_group state
// and decodes the result against the current opal_group schema, exactly as
// terraform-plugin-framework does after calling the upgrader. If the
// upgrader's output object type is missing attributes that were since added
// to the schema, this fails with an error like "error decoding object;
// expected N attributes, got M".
func TestGroupStateUpgraderV0MatchesCurrentSchema(t *testing.T) {
	ctx := context.Background()
	dynamicValue := runUpgrader(t, ctx, stateupgraders.GroupStateUpgraderV0, groupV0RawStateJSON)
	schemaType := currentSchemaType(t, ctx, provider.NewGroupResource())
	upgradedState, err := dynamicValue.Unmarshal(schemaType)
	if err != nil {
		t.Fatalf("upgraded state is not compatible with the current schema: %v", err)
	}
	if !upgradedState.Type().Equal(schemaType) {
		t.Fatalf("upgraded state type %v does not match current schema type %v", upgradedState.Type(), schemaType)
	}
}

// TestResourceStateUpgraderV0MatchesCurrentSchema is the opal_resource
// equivalent of TestGroupStateUpgraderV0MatchesCurrentSchema.
func TestResourceStateUpgraderV0MatchesCurrentSchema(t *testing.T) {
	ctx := context.Background()
	dynamicValue := runUpgrader(t, ctx, stateupgraders.ResourceStateUpgraderV0, resourceV0RawStateJSON)
	schemaType := currentSchemaType(t, ctx, provider.NewResourceResource())
	upgradedState, err := dynamicValue.Unmarshal(schemaType)
	if err != nil {
		t.Fatalf("upgraded state is not compatible with the current schema: %v", err)
	}
	if !upgradedState.Type().Equal(schemaType) {
		t.Fatalf("upgraded state type %v does not match current schema type %v", upgradedState.Type(), schemaType)
	}
}
