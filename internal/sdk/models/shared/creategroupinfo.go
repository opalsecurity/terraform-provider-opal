// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// # CreateGroupInfo Object
// ### Description
// The `CreateGroupInfo` object is used to store creation info for a group.
//
// ### Usage Example
// Use in the `POST Groups` endpoint.
type CreateGroupInfo struct {
	// The ID of the app for the group.
	AppID string `json:"app_id"`
	// Custom request notification sent upon request approval.
	CustomRequestNotification *string `json:"custom_request_notification,omitempty"`
	// A description of the remote group.
	Description *string `json:"description,omitempty"`
	// The type of the group.
	GroupType GroupTypeEnum `json:"group_type"`
	// The name of the remote group.
	Name string `json:"name"`
	// Information that defines the remote group. This replaces the deprecated remote_id and metadata fields.
	RemoteInfo              *GroupRemoteInfo     `json:"remote_info,omitempty"`
	RiskSensitivityOverride *RiskSensitivityEnum `json:"risk_sensitivity_override,omitempty"`
}

func (o *CreateGroupInfo) GetAppID() string {
	if o == nil {
		return ""
	}
	return o.AppID
}

func (o *CreateGroupInfo) GetCustomRequestNotification() *string {
	if o == nil {
		return nil
	}
	return o.CustomRequestNotification
}

func (o *CreateGroupInfo) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *CreateGroupInfo) GetGroupType() GroupTypeEnum {
	if o == nil {
		return GroupTypeEnum("")
	}
	return o.GroupType
}

func (o *CreateGroupInfo) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *CreateGroupInfo) GetRemoteInfo() *GroupRemoteInfo {
	if o == nil {
		return nil
	}
	return o.RemoteInfo
}

func (o *CreateGroupInfo) GetRiskSensitivityOverride() *RiskSensitivityEnum {
	if o == nil {
		return nil
	}
	return o.RiskSensitivityOverride
}
