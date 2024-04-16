package provider

import (
    "context"
    "github.com/hashicorp/terraform-plugin-framework/attr"
    "github.com/hashicorp/terraform-plugin-framework/diag"
    "github.com/hashicorp/terraform-plugin-framework/types"
)

// resourceOpalOwnerStateUpgradeV0 defines the state upgrade method for the opal_owner resource from v2.0.2 to v3.0.0.
func resourceOpalOwnerStateUpgradeV0(ctx context.Context, rawState map[string]attr.Value, meta interface{}) (map[string]attr.Value, diag.Diagnostics) {
    var diags diag.Diagnostics

    // Initialize the upgraded state map
    upgradedState := make(map[string]attr.Value)

    // Check if the old state has the 'user' block list and convert it to 'user_ids' list of strings
    if users, exists := rawState["user"]; exists {
        // Convert users to a list of user IDs
        var userIds []attr.Value
        usersList, ok := users.(types.List)
        if !ok {
            diags.Append(diag.NewErrorDiagnostic("Incorrect type for 'user'", "Expected 'user' to be a list of strings"))
            return nil, diags
        }
        for _, user := range usersList.Elements() {
            userStr, ok := user.(types.String)
            if !ok {
                diags.Append(diag.NewErrorDiagnostic("Incorrect type for 'user' item", "Expected 'user' item to be a string"))
                return nil, diags
            }
            userIds = append(userIds, types.StringValue(userStr.ValueString()))
        }

        // Create a new list value for 'user_ids'
        userIDsValue, userIDsDiags := types.ListValue(types.StringType, userIds)
        if userIDsDiags.HasError() {
            return nil, userIDsDiags
        }

        // Add the 'user_ids' to the upgraded state
        upgradedState["user_ids"] = userIDsValue
    } else {
        // If 'user' block list does not exist, add a diagnostic error as 'user_ids' is now required
        diags.Append(diag.NewErrorDiagnostic("Missing required attribute", "The 'user' attribute is required for the opal_owner resource."))
    }

    return upgradedState, diags
}
