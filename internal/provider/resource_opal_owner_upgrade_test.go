package provider

import (
    "context"
    "testing"

    "github.com/hashicorp/terraform-plugin-framework/attr"
    "github.com/hashicorp/terraform-plugin-framework/types"
    "github.com/go-test/deep"
)

func TestResourceOpalOwnerStateUpgradeV0(t *testing.T) {
    // Define a sample state as it would have been in v2.0.2
    stateV202 := map[string]attr.Value{
        "user": func() attr.Value {
            v, _ := types.ListValue(types.StringType, []attr.Value{
                types.StringValue("user1"),
                types.StringValue("user2"),
            })
            return v
        }(),
    }

    // Define what we expect the state to look like after the upgrade to v3.0.0
    expectedStateV300 := map[string]attr.Value{
        "user_ids": func() attr.Value {
            v, _ := types.ListValue(types.StringType, []attr.Value{
                types.StringValue("user1"),
                types.StringValue("user2"),
            })
            return v
        }(),
    }

    // Perform the upgrade
    upgradedState, diags := resourceOpalOwnerStateUpgradeV0(context.Background(), stateV202, nil)

    // Check for errors
    if diags.HasError() {
        t.Fatalf("unexpected diagnostics: %s", diags)
    }

    // Check for differences
    if diff := deep.Equal(upgradedState, expectedStateV300); diff != nil {
        t.Error(diff)
    }
}
