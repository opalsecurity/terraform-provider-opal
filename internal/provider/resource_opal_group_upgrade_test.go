package provider

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/go-test/deep"
)

func TestResourceOpalGroupStateUpgradeV0(t *testing.T) {
	// Define a sample state as it would have been in v2.0.2
	stateV202 := map[string]attr.Value{
		"audit_message_channel": types.StringValue("example_audit_message_channel"),
		"on_call_schedule":      types.StringValue("example_on_call_schedule"),
		"visibility_group":      types.StringValue("example_visibility_group"),
		"manage_resources":      types.BoolValue(true),
		"resource":              types.StringValue("example_resource"),
		"request_configuration": types.StringValue("example_request_configuration"),
	}

	// Define what we expect the state to look like after the upgrade to v3.0.0
	expectedStateV300 := map[string]attr.Value{
		"message_channel_ids": func() attr.Value {
			v, _ := types.ListValue(types.StringType, []attr.Value{types.StringValue("example_audit_message_channel")})
			return v
		}(),
		"on_call_schedule_ids": func() attr.Value {
			v, _ := types.ListValue(types.StringType, []attr.Value{types.StringValue("example_on_call_schedule")})
			return v
		}(),
		"visibility_group_ids": func() attr.Value {
			v, _ := types.ListValue(types.StringType, []attr.Value{types.StringValue("example_visibility_group")})
			return v
		}(),
		"request_configurations": func() attr.Value {
			v, _ := types.ListValue(types.StringType, []attr.Value{types.StringValue("example_request_configuration")})
			return v
		}(),
		"visibility": types.StringValue("private"),
	}

	// Perform the upgrade
	upgradedState, err := resourceOpalGroupStateUpgradeV0(context.Background(), stateV202, nil)
	if err != nil {
		t.Fatalf("error upgrading state: %s", err)
	}

	// Check for differences
	if diff := deep.Equal(upgradedState, expectedStateV300); diff != nil {
		t.Error(diff)
	}
}
