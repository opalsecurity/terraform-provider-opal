package stringvalidators

import (
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

// ResourceConfigurationTemplateID prevents configurations that the API rejects
// when a resource is attached to a configuration template.
func ResourceConfigurationTemplateID() validator.String {
	return stringvalidator.All(
		stringvalidator.ExactlyOneOf(
			path.MatchRoot("configuration_template_id"),
			path.MatchRoot("request_configurations"),
		),
		stringvalidator.ConflictsWith(
			path.MatchRoot("admin_owner_id"),
			path.MatchRoot("require_mfa_to_approve"),
			path.MatchRoot("require_mfa_to_connect"),
			path.MatchRoot("ticket_propagation"),
			path.MatchRoot("custom_request_notification"),
		),
	)
}
