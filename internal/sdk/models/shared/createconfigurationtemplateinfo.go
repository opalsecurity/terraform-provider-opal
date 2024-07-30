// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// # CreateConfigurationTemplateInfo Object
// ### Description
// The `CreateConfigurationTemplateInfo` object is used to store creation info for a configuration template.
//
// ### Usage Example
// Use in the `POST Configuration Templates` endpoint.
type CreateConfigurationTemplateInfo struct {
	// The ID of the owner of the configuration template.
	AdminOwnerID string `json:"admin_owner_id"`
	// The IDs of the break glass users linked to the configuration template.
	BreakGlassUserIds []string `json:"break_glass_user_ids,omitempty"`
	// The IDs of the audit message channels linked to the configuration template.
	LinkedAuditMessageChannelIds []string `json:"linked_audit_message_channel_ids,omitempty"`
	// The IDs of the on-call schedules linked to the configuration template.
	MemberOncallScheduleIds []string `json:"member_oncall_schedule_ids,omitempty"`
	// The name of the configuration template.
	Name string `json:"name"`
	// The request configuration list of the configuration template. If not provided, the default request configuration will be used.
	RequestConfigurations []RequestConfiguration `json:"request_configurations,omitempty"`
	// A bool representing whether or not to require MFA for reviewers to approve requests for this configuration template.
	RequireMfaToApprove bool `json:"require_mfa_to_approve"`
	// A bool representing whether or not to require MFA to connect to resources associated with this configuration template.
	RequireMfaToConnect bool `json:"require_mfa_to_connect"`
	// Configuration for ticket propagation, when enabled, a ticket will be created for access changes related to the users in this resource.
	TicketPropagation *TicketPropagationConfiguration `json:"ticket_propagation,omitempty"`
	// Visibility infomation of an entity.
	Visibility VisibilityInfo `json:"visibility"`
}

func (o *CreateConfigurationTemplateInfo) GetAdminOwnerID() string {
	if o == nil {
		return ""
	}
	return o.AdminOwnerID
}

func (o *CreateConfigurationTemplateInfo) GetBreakGlassUserIds() []string {
	if o == nil {
		return nil
	}
	return o.BreakGlassUserIds
}

func (o *CreateConfigurationTemplateInfo) GetLinkedAuditMessageChannelIds() []string {
	if o == nil {
		return nil
	}
	return o.LinkedAuditMessageChannelIds
}

func (o *CreateConfigurationTemplateInfo) GetMemberOncallScheduleIds() []string {
	if o == nil {
		return nil
	}
	return o.MemberOncallScheduleIds
}

func (o *CreateConfigurationTemplateInfo) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *CreateConfigurationTemplateInfo) GetRequestConfigurations() []RequestConfiguration {
	if o == nil {
		return nil
	}
	return o.RequestConfigurations
}

func (o *CreateConfigurationTemplateInfo) GetRequireMfaToApprove() bool {
	if o == nil {
		return false
	}
	return o.RequireMfaToApprove
}

func (o *CreateConfigurationTemplateInfo) GetRequireMfaToConnect() bool {
	if o == nil {
		return false
	}
	return o.RequireMfaToConnect
}

func (o *CreateConfigurationTemplateInfo) GetTicketPropagation() *TicketPropagationConfiguration {
	if o == nil {
		return nil
	}
	return o.TicketPropagation
}

func (o *CreateConfigurationTemplateInfo) GetVisibility() VisibilityInfo {
	if o == nil {
		return VisibilityInfo{}
	}
	return o.Visibility
}
