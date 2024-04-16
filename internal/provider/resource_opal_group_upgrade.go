package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// resourceOpalGroupStateUpgradeV0 upgrades the state from version 0 to version 1.
func resourceOpalGroupStateUpgradeV0(ctx context.Context, rawState map[string]attr.Value, meta interface{}) (map[string]attr.Value, diag.Diagnostics) {
	var diags diag.Diagnostics

	// Initialize the upgraded state map
	upgradedState := make(map[string]attr.Value)

	// Upgrade logic for message_channel_ids
	if v, ok := rawState["audit_message_channel"]; ok {
		// Assuming audit_message_channel was a single string ID, now it needs to be a list of string IDs
		var messageChannelIDs types.List
		messageChannelIDs, diags = types.ListValue(types.StringType, []attr.Value{v})
		if diags.HasError() {
			return nil, diags
		}
		upgradedState["message_channel_ids"] = messageChannelIDs
	}

	// Upgrade logic for on_call_schedule_ids
	if v, ok := rawState["on_call_schedule"]; ok {
		// Assuming on_call_schedule was a single string ID, now it needs to be a list of string IDs
		var onCallScheduleIDs types.List
		onCallScheduleIDs, diags = types.ListValue(types.StringType, []attr.Value{v})
		if diags.HasError() {
			return nil, diags
		}
		upgradedState["on_call_schedule_ids"] = onCallScheduleIDs
	}

	// Upgrade logic for visibility_group_ids
	if v, ok := rawState["visibility_group"]; ok {
		// Assuming visibility_group was a single string ID, now it needs to be a list of string IDs
		var visibilityGroupIDs types.List
		visibilityGroupIDs, diags = types.ListValue(types.StringType, []attr.Value{v})
		if diags.HasError() {
			return nil, diags
		}
		upgradedState["visibility_group_ids"] = visibilityGroupIDs
	}

	// Upgrade logic for request_configurations
	if v, ok := rawState["request_configuration"]; ok {
		// Assuming request_configuration was a single configuration, now it needs to be a list of configurations
		var requestConfigurations types.List
		requestConfigurations, diags = types.ListValue(types.StringType, []attr.Value{v})
		if diags.HasError() {
			return nil, diags
		}
		upgradedState["request_configurations"] = requestConfigurations
	}

	// Set visibility to a default value if it was not previously set
	if _, ok := rawState["visibility"]; !ok {
		upgradedState["visibility"] = types.StringValue("private") // Set a default visibility if it was not set before
	}

	// Remove manage_resources as it is no longer used
	delete(rawState, "manage_resources")

	// Remove resource as it has been moved to a separate resource
	delete(rawState, "resource")

	// Return the upgraded state
	return upgradedState, diags
}
