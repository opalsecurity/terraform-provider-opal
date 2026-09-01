package stringvalidators

import (
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

// GroupConfigurationTemplateID prevents configurations that the API rejects
// when a group is attached to a configuration template.
//
// Visibility stays required for untemplated groups (the previous schema marked
// it required). ExactlyOneOf with visibility restores that without forcing a
// value that the template would ignore.
func GroupConfigurationTemplateID() validator.String {
	return stringvalidator.All(
		stringvalidator.ExactlyOneOf(
			path.MatchRoot("configuration_template_id"),
			path.MatchRoot("request_configurations"),
		),
		stringvalidator.ExactlyOneOf(
			path.MatchRoot("configuration_template_id"),
			path.MatchRoot("visibility"),
		),
		stringvalidator.ConflictsWith(
			path.MatchRoot("admin_owner_id"),
			path.MatchRoot("require_mfa_to_approve"),
			path.MatchRoot("custom_request_notification"),
			path.MatchRoot("extensions_duration_in_minutes"),
			path.MatchRoot("visibility_group_ids"),
		),
	)
}
