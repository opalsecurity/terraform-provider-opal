package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/opalsecurity/terraform-provider-opal/v3/internal/provider/types"
)

var _ resource.ResourceWithModifyPlan = &RequestTemplateResource{}

// ModifyPlan preserves generated callout names across updates without the
// index-based SuppressDiff that reused unrelated field names when the list
// reshuffled.
//
// When config omits custom_fields[].name for a CALLOUT, copy the name from a
// prior-state callout with the same callout text. If none matches, leave name
// null so the API generates callout-N.
//
// Config and state are read as typed models (fully known). The plan is only
// touched via GetAttribute/SetAttribute on custom_fields[i].name — a full
// Plan.Get into the generated model fails when nested metadata is unknown
// ("known after apply"), because *RequestTemplateCustomFieldMetadata cannot
// represent unknowns.
func (r *RequestTemplateResource) ModifyPlan(
	ctx context.Context,
	req resource.ModifyPlanRequest,
	resp *resource.ModifyPlanResponse,
) {
	if req.Plan.Raw.IsNull() || req.State.Raw.IsNull() {
		return
	}

	var config, state RequestTemplateResourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	if config.CustomFields == nil {
		return
	}

	stateCalloutNameByText := map[string]string{}
	for _, field := range state.CustomFields {
		if field.Type.IsNull() || field.Type.IsUnknown() || field.Type.ValueString() != "CALLOUT" {
			continue
		}
		if field.Name.IsNull() || field.Name.IsUnknown() || field.Name.ValueString() == "" {
			continue
		}
		text := calloutText(field)
		if text == "" {
			continue
		}
		stateCalloutNameByText[text] = field.Name.ValueString()
	}

	for i, configField := range config.CustomFields {
		if configField.Type.IsNull() || configField.Type.IsUnknown() ||
			configField.Type.ValueString() != "CALLOUT" {
			continue
		}
		// User explicitly set a name — leave the plan alone.
		if !configField.Name.IsNull() && !configField.Name.IsUnknown() {
			continue
		}

		namePath := path.Root("custom_fields").AtListIndex(i).AtName("name")
		text := calloutText(configField)
		priorName, ok := stateCalloutNameByText[text]
		if !ok || text == "" {
			// New/changed callout copy: ensure name stays null for API generation.
			var planName types.String
			resp.Diagnostics.Append(req.Plan.GetAttribute(ctx, namePath, &planName)...)
			if resp.Diagnostics.HasError() {
				return
			}
			if !planName.IsNull() && !planName.IsUnknown() {
				resp.Diagnostics.Append(resp.Plan.SetAttribute(ctx, namePath, types.StringNull())...)
				if resp.Diagnostics.HasError() {
					return
				}
			}
			continue
		}

		var planName types.String
		resp.Diagnostics.Append(req.Plan.GetAttribute(ctx, namePath, &planName)...)
		if resp.Diagnostics.HasError() {
			return
		}
		if planName.IsNull() || planName.IsUnknown() || planName.ValueString() != priorName {
			resp.Diagnostics.Append(resp.Plan.SetAttribute(ctx, namePath, types.StringValue(priorName))...)
			if resp.Diagnostics.HasError() {
				return
			}
		}
	}
}

func calloutText(field tfTypes.RequestTemplateCustomField) string {
	if field.Metadata == nil || field.Metadata.CalloutData == nil {
		return ""
	}
	if field.Metadata.CalloutData.Text.IsNull() || field.Metadata.CalloutData.Text.IsUnknown() {
		return ""
	}
	return field.Metadata.CalloutData.Text.ValueString()
}
