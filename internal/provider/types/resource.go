// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type Resource struct {
	AdminOwnerID          types.String           `tfsdk:"admin_owner_id"`
	AppID                 types.String           `tfsdk:"app_id"`
	AutoApproval          types.Bool             `tfsdk:"auto_approval"`
	Description           types.String           `tfsdk:"description"`
	ID                    types.String           `tfsdk:"id"`
	IsRequestable         types.Bool             `tfsdk:"is_requestable"`
	MaxDuration           types.Int64            `tfsdk:"max_duration"`
	Name                  types.String           `tfsdk:"name"`
	ParentResourceID      types.String           `tfsdk:"parent_resource_id"`
	RecommendedDuration   types.Int64            `tfsdk:"recommended_duration"`
	RemoteInfo            *ResourceRemoteInfo    `tfsdk:"remote_info"`
	RemoteResourceID      types.String           `tfsdk:"remote_resource_id"`
	RemoteResourceName    types.String           `tfsdk:"remote_resource_name"`
	RequestConfigurations []RequestConfiguration `tfsdk:"request_configurations"`
	RequestTemplateID     types.String           `tfsdk:"request_template_id"`
	RequireMfaToApprove   types.Bool             `tfsdk:"require_mfa_to_approve"`
	RequireMfaToConnect   types.Bool             `tfsdk:"require_mfa_to_connect"`
	RequireMfaToRequest   types.Bool             `tfsdk:"require_mfa_to_request"`
	RequireSupportTicket  types.Bool             `tfsdk:"require_support_ticket"`
	ResourceType          types.String           `tfsdk:"resource_type"`
}