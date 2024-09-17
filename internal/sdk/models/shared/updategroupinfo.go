// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// # UpdateGroupInfo Object
// ### Description
// The `UpdateGroupInfo` object is used as an input to the UpdateGroup API.
type UpdateGroupInfo struct {
	// The ID of the owner of the group.
	AdminOwnerID *string `json:"admin_owner_id,omitempty"`
	// Custom request notification sent upon request approval for this configuration template.
	CustomRequestNotification *string `json:"custom_request_notification,omitempty"`
	// A description of the group.
	Description *string `json:"description,omitempty"`
	// A list of User IDs for the group leaders of the group
	GroupLeaderUserIds []string `json:"group_leader_user_ids,omitempty"`
	// The ID of the group.
	ID string `json:"group_id"`
	// The name of the group.
	Name *string `json:"name,omitempty"`
	// The request configuration list of the configuration template. If not provided, the default request configuration will be used.
	RequestConfigurations []RequestConfiguration `json:"request_configurations"`
	// A bool representing whether or not to require MFA for reviewers to approve requests for this group.
	RequireMfaToApprove *bool `json:"require_mfa_to_approve,omitempty"`
}

func (o *UpdateGroupInfo) GetAdminOwnerID() *string {
	if o == nil {
		return nil
	}
	return o.AdminOwnerID
}

func (o *UpdateGroupInfo) GetCustomRequestNotification() *string {
	if o == nil {
		return nil
	}
	return o.CustomRequestNotification
}

func (o *UpdateGroupInfo) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *UpdateGroupInfo) GetGroupLeaderUserIds() []string {
	if o == nil {
		return nil
	}
	return o.GroupLeaderUserIds
}

func (o *UpdateGroupInfo) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *UpdateGroupInfo) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *UpdateGroupInfo) GetRequestConfigurations() []RequestConfiguration {
	if o == nil {
		return []RequestConfiguration{}
	}
	return o.RequestConfigurations
}

func (o *UpdateGroupInfo) GetRequireMfaToApprove() *bool {
	if o == nil {
		return nil
	}
	return o.RequireMfaToApprove
}
