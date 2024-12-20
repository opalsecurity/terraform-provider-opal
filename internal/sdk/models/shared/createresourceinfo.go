// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// # CreateResourceInfo Object
// ### Description
// The `CreateResourceInfo` object is used to store creation info for a resource.
//
// ### Usage Example
// Use in the `POST Resources` endpoint.
type CreateResourceInfo struct {
	// The ID of the app for the resource.
	AppID string `json:"app_id"`
	// Custom request notification sent upon request approval.
	CustomRequestNotification *string `json:"custom_request_notification,omitempty"`
	// A description of the remote resource.
	Description *string `json:"description,omitempty"`
	// The name of the remote resource.
	Name string `json:"name"`
	// Information that defines the remote resource. This replaces the deprecated remote_id and metadata fields.
	RemoteInfo *ResourceRemoteInfo `json:"remote_info,omitempty"`
	// The type of the resource.
	ResourceType            ResourceTypeEnum     `json:"resource_type"`
	RiskSensitivityOverride *RiskSensitivityEnum `json:"risk_sensitivity_override,omitempty"`
}

func (o *CreateResourceInfo) GetAppID() string {
	if o == nil {
		return ""
	}
	return o.AppID
}

func (o *CreateResourceInfo) GetCustomRequestNotification() *string {
	if o == nil {
		return nil
	}
	return o.CustomRequestNotification
}

func (o *CreateResourceInfo) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *CreateResourceInfo) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *CreateResourceInfo) GetRemoteInfo() *ResourceRemoteInfo {
	if o == nil {
		return nil
	}
	return o.RemoteInfo
}

func (o *CreateResourceInfo) GetResourceType() ResourceTypeEnum {
	if o == nil {
		return ResourceTypeEnum("")
	}
	return o.ResourceType
}

func (o *CreateResourceInfo) GetRiskSensitivityOverride() *RiskSensitivityEnum {
	if o == nil {
		return nil
	}
	return o.RiskSensitivityOverride
}
