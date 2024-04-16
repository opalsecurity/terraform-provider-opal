package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// resourceOpalResourceStateUpgradeV0 defines the v0 to v1 upgrade path for the opal_resource resource.
func resourceOpalResourceStateUpgradeV0(ctx context.Context, state map[string]attr.Value, provider interface{}) (map[string]attr.Value, diag.Diagnostics) {
	var diags diag.Diagnostics

	// Upgrade from v2.0.2 to v3.0.0 schema changes
	// Example: Convert visibility_group to visibility_group_ids and ensure it is a list of string IDs
	if visibilityGroup, exists := state["visibility_group"]; exists {
		// Convert to list of string IDs
		var visibilityGroupIds []attr.Value
		if vGroup, ok := visibilityGroup.(types.String); ok {
			visibilityGroupIds = append(visibilityGroupIds, vGroup)
		} else {
			diags.AddError("Error converting visibility_group", "Failed to convert visibility_group to types.String")
			return state, diags
		}
		listValue, listDiags := types.ListValue(types.StringType, visibilityGroupIds)
		diags = append(diags, listDiags...)
		state["visibility_group_ids"] = listValue
		delete(state, "visibility_group")
	}

	// Example: Convert request_configuration to request_configurations and ensure it is a required list
	if requestConfiguration, exists := state["request_configuration"]; exists {
		var requestConfigurations []attr.Value
		if reqConfig, ok := requestConfiguration.(types.String); ok {
			requestConfigurations = append(requestConfigurations, reqConfig)
		} else {
			diags.AddError("Error converting request_configuration", "Failed to convert request_configuration to types.String")
			return state, diags
		}
		listValue, listDiags := types.ListValue(types.StringType, requestConfigurations)
		diags = append(diags, listDiags...)
		state["request_configurations"] = listValue
		delete(state, "request_configuration")
	}

	// Example: Ensure visibility is set and has a valid value
	if visibility, exists := state["visibility"]; !exists || visibility.IsNull() {
		state["visibility"] = types.StringValue("LIMITED") // Default to LIMITED if not set
	}

	// Example: Remove manage_resources as it is no longer present in v3.0.0
	delete(state, "manage_resources")

	// Return the new state and any diagnostics
	return state, diags
}
