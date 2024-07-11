// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// # UpdateConfigurationTemplateInfo Object
// ### Description
// The `ConfigurationTemplate` object is used to represent an update to a configuration template.
//
// ### Usage Example
// Use in the `PUT Configuration Templates` endpoint.
type UpdateConfigurationTemplateInfo struct {
	// The ID of the owner of the configuration template.
	AdminOwnerID *string `json:"admin_owner_id,omitempty"`
	// The IDs of the break glass users linked to the configuration template.
	BreakGlassUserIds []string `json:"break_glass_user_ids,omitempty"`
	// The ID of the configuration template.
	ConfigurationTemplateID string `json:"configuration_template_id"`
	// The IDs of the audit message channels linked to the configuration template.
	LinkedAuditMessageChannelIds []string `json:"linked_audit_message_channel_ids,omitempty"`
	// The IDs of the on-call schedules linked to the configuration template.
	MemberOncallScheduleIds []string `json:"member_oncall_schedule_ids,omitempty"`
	// The name of the configuration template.
	Name *string `json:"name,omitempty"`
	// The request configuration list linked to the configuration template.
	RequestConfigurations []RequestConfiguration `json:"request_configurations,omitempty"`
	// A bool representing whether or not to require MFA for reviewers to approve requests for this configuration template.
	RequireMfaToApprove *bool `json:"require_mfa_to_approve,omitempty"`
	// A bool representing whether or not to require MFA to connect to resources associated with this configuration template.
	RequireMfaToConnect *bool `json:"require_mfa_to_connect,omitempty"`
	// Configuration for ticket propagation, when enabled, a ticket will be created for access changes related to the users in this resource.
	TicketPropagation *TicketPropagationConfiguration `json:"ticket_propagation,omitempty"`
	// Visibility infomation of an entity.
	Visibility *VisibilityInfo `json:"visibility,omitempty"`
}

func (o *UpdateConfigurationTemplateInfo) GetAdminOwnerID() *string {
	if o == nil {
		return nil
	}
	return o.AdminOwnerID
}

func (o *UpdateConfigurationTemplateInfo) GetBreakGlassUserIds() []string {
	if o == nil {
		return nil
	}
	return o.BreakGlassUserIds
}

func (o *UpdateConfigurationTemplateInfo) GetConfigurationTemplateID() string {
	if o == nil {
		return ""
	}
	return o.ConfigurationTemplateID
}

func (o *UpdateConfigurationTemplateInfo) GetLinkedAuditMessageChannelIds() []string {
	if o == nil {
		return nil
	}
	return o.LinkedAuditMessageChannelIds
}

func (o *UpdateConfigurationTemplateInfo) GetMemberOncallScheduleIds() []string {
	if o == nil {
		return nil
	}
	return o.MemberOncallScheduleIds
}

func (o *UpdateConfigurationTemplateInfo) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *UpdateConfigurationTemplateInfo) GetRequestConfigurations() []RequestConfiguration {
	if o == nil {
		return nil
	}
	return o.RequestConfigurations
}

func (o *UpdateConfigurationTemplateInfo) GetRequireMfaToApprove() *bool {
	if o == nil {
		return nil
	}
	return o.RequireMfaToApprove
}

func (o *UpdateConfigurationTemplateInfo) GetRequireMfaToConnect() *bool {
	if o == nil {
		return nil
	}
	return o.RequireMfaToConnect
}

func (o *UpdateConfigurationTemplateInfo) GetTicketPropagation() *TicketPropagationConfiguration {
	if o == nil {
		return nil
	}
	return o.TicketPropagation
}

func (o *UpdateConfigurationTemplateInfo) GetVisibility() *VisibilityInfo {
	if o == nil {
		return nil
	}
	return o.Visibility
}
