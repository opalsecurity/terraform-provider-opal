// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// # Group Object
// ### Description
// The `Group` object is used to represent a group.
//
// ### Usage Example
// Update from the `UPDATE Groups` endpoint.
type Group struct {
	// The ID of the owner of the group.
	AdminOwnerID *string `json:"admin_owner_id,omitempty"`
	// The ID of the group's app.
	AppID *string `json:"app_id,omitempty"`
	// Custom request notification sent to the requester when the request is approved.
	CustomRequestNotification *string `json:"custom_request_notification,omitempty"`
	// A description of the group.
	Description *string `json:"description,omitempty"`
	// The ID of the associated group binding.
	GroupBindingID *string `json:"group_binding_id,omitempty"`
	// A list of User IDs for the group leaders of the group
	GroupLeaderUserIds []string `json:"group_leader_user_ids,omitempty"`
	// The type of the group.
	GroupType *GroupTypeEnum `json:"group_type,omitempty"`
	// The ID of the group.
	ID string `json:"group_id"`
	// The name of the group.
	Name *string `json:"name,omitempty"`
	// Information that defines the remote group. This replaces the deprecated remote_id and metadata fields.
	RemoteInfo *GroupRemoteInfo `json:"remote_info,omitempty"`
	// The name of the remote.
	RemoteName *string `json:"remote_name,omitempty"`
	// A list of request configurations for this group.
	RequestConfigurations []RequestConfiguration `json:"request_configurations,omitempty"`
	// A bool representing whether or not to require MFA for reviewers to approve requests for this group.
	RequireMfaToApprove *bool `json:"require_mfa_to_approve,omitempty"`
}

func (o *Group) GetAdminOwnerID() *string {
	if o == nil {
		return nil
	}
	return o.AdminOwnerID
}

func (o *Group) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

func (o *Group) GetCustomRequestNotification() *string {
	if o == nil {
		return nil
	}
	return o.CustomRequestNotification
}

func (o *Group) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *Group) GetGroupBindingID() *string {
	if o == nil {
		return nil
	}
	return o.GroupBindingID
}

func (o *Group) GetGroupLeaderUserIds() []string {
	if o == nil {
		return nil
	}
	return o.GroupLeaderUserIds
}

func (o *Group) GetGroupType() *GroupTypeEnum {
	if o == nil {
		return nil
	}
	return o.GroupType
}

func (o *Group) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *Group) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *Group) GetRemoteInfo() *GroupRemoteInfo {
	if o == nil {
		return nil
	}
	return o.RemoteInfo
}

func (o *Group) GetRemoteName() *string {
	if o == nil {
		return nil
	}
	return o.RemoteName
}

func (o *Group) GetRequestConfigurations() []RequestConfiguration {
	if o == nil {
		return nil
	}
	return o.RequestConfigurations
}

func (o *Group) GetRequireMfaToApprove() *bool {
	if o == nil {
		return nil
	}
	return o.RequireMfaToApprove
}
