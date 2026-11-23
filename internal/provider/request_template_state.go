package provider

import (
	tfTypes "github.com/opalsecurity/terraform-provider-opal/v3/internal/provider/types"
)

// restoreAPIComputedCustomFieldAttrs copies computed custom-field attrs from the
// API refresh snapshot back onto data after refreshPlan.
//
// Speakeasy's refreshPlan overlays the Terraform plan onto the refreshed model.
// For Optional+Computed fields omitted in HCL (callout name, callout required),
// the plan value is null, so refreshPlan would wipe the API values. SuppressDiff
// used to make the plan value unknown (skipped by Into); without it we restore.
func restoreAPIComputedCustomFieldAttrs(
	fromAPI []tfTypes.RequestTemplateCustomField,
	data *RequestTemplateResourceModel,
) {
	if data == nil || len(fromAPI) == 0 || len(data.CustomFields) == 0 {
		return
	}
	for i := range data.CustomFields {
		if i >= len(fromAPI) {
			break
		}
		api := fromAPI[i]
		cur := &data.CustomFields[i]

		if !api.Name.IsNull() && !api.Name.IsUnknown() && api.Name.ValueString() != "" {
			if cur.Name.IsNull() || cur.Name.IsUnknown() || cur.Name.ValueString() == "" {
				cur.Name = api.Name
			}
		}
		if !api.Required.IsNull() && !api.Required.IsUnknown() {
			if cur.Required.IsNull() || cur.Required.IsUnknown() {
				cur.Required = api.Required
			}
		}
	}
}
