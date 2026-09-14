package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

var (
	_ resource.ResourceWithValidateConfig = &GroupResource{}
	_ resource.ResourceWithValidateConfig = &ResourceResource{}
)

// Attributes that must not appear in HCL alongside configuration_template_id.
// Includes ExactlyOneOf siblings (visibility, request_configurations) so a
// same-apply unknown template ID still conflicts at plan/validate time.
// Framework ConflictsWith/ExactlyOneOf skip when any involved value is unknown.
var groupConfigurationTemplateConflictAttributes = []string{
	"admin_owner_id",
	"require_mfa_to_approve",
	"custom_request_notification",
	"extensions_duration_in_minutes",
	"visibility_group_ids",
	"visibility",
	"request_configurations",
}

var resourceConfigurationTemplateConflictAttributes = []string{
	"admin_owner_id",
	"require_mfa_to_approve",
	"require_mfa_to_connect",
	"ticket_propagation",
	"custom_request_notification",
	"extensions_duration_in_minutes",
	"visibility_group_ids",
	"visibility",
	"request_configurations",
}

func (r *GroupResource) ValidateConfig(_ context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	validateConfigurationTemplateConflicts(req.Config.Raw, groupConfigurationTemplateConflictAttributes, &resp.Diagnostics)
}

func (r *ResourceResource) ValidateConfig(_ context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	validateConfigurationTemplateConflicts(req.Config.Raw, resourceConfigurationTemplateConflictAttributes, &resp.Diagnostics)
}

// validateConfigurationTemplateConflicts rejects HCL that sets
// configuration_template_id together with template-owned attributes, including
// when either side is still unknown (e.g. configuration_template_id =
// opal_configuration_template.t.id in the same apply).
func validateConfigurationTemplateConflicts(
	configRaw tftypes.Value,
	conflictAttributes []string,
	diagnostics *diag.Diagnostics,
) {
	if configRaw.IsNull() || !configRaw.IsKnown() {
		return
	}

	config, err := terraformObjectValues(configRaw)
	if err != nil {
		diagnostics.AddError("Unable to inspect configuration", err.Error())
		return
	}

	if !configurationTemplateIDConfiguredInConfig(config) {
		return
	}

	for _, attribute := range conflictAttributes {
		if attributeOmitted(config, attribute) {
			continue
		}
		diagnostics.AddAttributeError(
			path.Root(attribute),
			"Invalid Attribute Combination",
			fmt.Sprintf(
				`Attribute %q cannot be specified when "configuration_template_id" is specified`,
				attribute,
			),
		)
	}
}

// configurationTemplateIDConfiguredInConfig is true when the attribute is
// present in config. Unknown values count; null and known empty string do not.
func configurationTemplateIDConfiguredInConfig(config map[string]tftypes.Value) bool {
	configured, ok := config["configuration_template_id"]
	if !ok || configured.IsNull() {
		return false
	}
	if !configured.IsKnown() {
		return true
	}
	var id string
	if err := configured.As(&id); err != nil || id == "" {
		return false
	}
	return true
}
