package provider

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/go-test/deep"
)

func TestResourceOpalResourceStateUpgradeV0(t *testing.T) {
	// Define a sample state as it would have been in v2.0.2
	stateV202 := map[string]attr.Value{
		"visibility_group": types.StringValue("example_visibility_group"),
		"request_configuration": types.StringValue("example_request_configuration"),
		"manage_resources": types.BoolValue(true),
		"visibility": types.StringValue(""),
	}

	// Define what we expect the state to look like after the upgrade to v3.0.0
	expectedStateV300 := map[string]attr.Value{
		"visibility_group_ids": func() attr.Value {
			v, _ := types.ListValue(types.StringType, []attr.Value{types.StringValue("example_visibility_group")})
			return v
		}(),
		"request_configurations": func() attr.Value {
			v, _ := types.ListValue(types.StringType, []attr.Value{types.StringValue("example_request_configuration")})
			return v
		}(),
		"visibility": types.StringValue("LIMITED"),
	}

	// Perform the upgrade
	upgradedState, _ := resourceOpalResourceStateUpgradeV0(context.Background(), stateV202, nil)

	// Check for differences
	if diff := deep.Equal(upgradedState, expectedStateV300); diff != nil {
		t.Error(diff)
	}
}
