package stringvalidators

import (
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

// GroupConfigurationTemplateID prevents configurations that the API rejects
// when a group is attached to a configuration template.
func GroupConfigurationTemplateID() validator.String {
	return stringvalidator.ConflictsWith(
		path.MatchRoot("request_configurations"),
		path.MatchRoot("admin_owner_id"),
		path.MatchRoot("require_mfa_to_approve"),
		path.MatchRoot("custom_request_notification"),
	)
}

// ResourceConfigurationTemplateID prevents configurations that the API rejects
// when a resource is attached to a configuration template.
func ResourceConfigurationTemplateID() validator.String {
	return stringvalidator.ConflictsWith(
		path.MatchRoot("request_configurations"),
		path.MatchRoot("admin_owner_id"),
		path.MatchRoot("require_mfa_to_approve"),
		path.MatchRoot("require_mfa_to_connect"),
		path.MatchRoot("ticket_propagation"),
		path.MatchRoot("custom_request_notification"),
	)
}
