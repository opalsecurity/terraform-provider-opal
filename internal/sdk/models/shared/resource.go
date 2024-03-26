// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// # Resource Object
// ### Description
// The `Resource` object is used to represent a resource.
//
// ### Usage Example
// Update from the `UPDATE Resources` endpoint.
type Resource struct {
	// The ID of the owner of the resource.
	AdminOwnerID *string `json:"admin_owner_id,omitempty"`
	// The ID of the app.
	AppID *string `json:"app_id,omitempty"`
	// A bool representing whether or not to automatically approve requests to this resource.
	AutoApproval *bool `json:"auto_approval,omitempty"`
	// The ID of the associated configuration template.
	ConfigurationTemplateID *string `json:"configuration_template_id,omitempty"`
	// A description of the resource.
	Description *string `json:"description,omitempty"`
	// The ID of the resource.
	ID *string `json:"resource_id,omitempty"`
	// A bool representing whether or not to allow access requests to this resource.
	IsRequestable *bool `json:"is_requestable,omitempty"`
	// The maximum duration for which the resource can be requested (in minutes).
	MaxDuration *int64 `json:"max_duration,omitempty"`
	// The name of the resource.
	Name *string `json:"name,omitempty"`
	// The ID of the parent resource.
	ParentResourceID *string `json:"parent_resource_id,omitempty"`
	// The recommended duration for which the resource should be requested (in minutes). -1 represents an indefinite duration.
	RecommendedDuration *int64 `json:"recommended_duration,omitempty"`
	// Information that defines the remote resource. This replaces the deprecated remote_id and metadata fields.
	RemoteInfo *ResourceRemoteInfo `json:"remote_info,omitempty"`
	// The ID of the resource on the remote system.
	RemoteResourceID *string `json:"remote_resource_id,omitempty"`
	// The name of the resource on the remote system.
	RemoteResourceName *string `json:"remote_resource_name,omitempty"`
	// A list of configurations for requests to this resource.
	RequestConfigurations []RequestConfiguration `json:"request_configurations,omitempty"`
	// The ID of the associated request template.
	RequestTemplateID *string `json:"request_template_id,omitempty"`
	// A bool representing whether or not to require MFA for reviewers to approve requests for this resource.
	RequireMfaToApprove *bool `json:"require_mfa_to_approve,omitempty"`
	// A bool representing whether or not to require MFA to connect to this resource.
	RequireMfaToConnect *bool `json:"require_mfa_to_connect,omitempty"`
	// A bool representing whether or not to require MFA for requesting access to this resource.
	RequireMfaToRequest *bool `json:"require_mfa_to_request,omitempty"`
	// A bool representing whether or not access requests to the resource require an access ticket.
	RequireSupportTicket *bool `json:"require_support_ticket,omitempty"`
	// The type of the resource.
	ResourceType *ResourceTypeEnum `json:"resource_type,omitempty"`
}

func (o *Resource) GetAdminOwnerID() *string {
	if o == nil {
		return nil
	}
	return o.AdminOwnerID
}

func (o *Resource) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

func (o *Resource) GetAutoApproval() *bool {
	if o == nil {
		return nil
	}
	return o.AutoApproval
}

func (o *Resource) GetConfigurationTemplateID() *string {
	if o == nil {
		return nil
	}
	return o.ConfigurationTemplateID
}

func (o *Resource) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *Resource) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *Resource) GetIsRequestable() *bool {
	if o == nil {
		return nil
	}
	return o.IsRequestable
}

func (o *Resource) GetMaxDuration() *int64 {
	if o == nil {
		return nil
	}
	return o.MaxDuration
}

func (o *Resource) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *Resource) GetParentResourceID() *string {
	if o == nil {
		return nil
	}
	return o.ParentResourceID
}

func (o *Resource) GetRecommendedDuration() *int64 {
	if o == nil {
		return nil
	}
	return o.RecommendedDuration
}

func (o *Resource) GetRemoteInfo() *ResourceRemoteInfo {
	if o == nil {
		return nil
	}
	return o.RemoteInfo
}

func (o *Resource) GetRemoteResourceID() *string {
	if o == nil {
		return nil
	}
	return o.RemoteResourceID
}

func (o *Resource) GetRemoteResourceName() *string {
	if o == nil {
		return nil
	}
	return o.RemoteResourceName
}

func (o *Resource) GetRequestConfigurations() []RequestConfiguration {
	if o == nil {
		return nil
	}
	return o.RequestConfigurations
}

func (o *Resource) GetRequestTemplateID() *string {
	if o == nil {
		return nil
	}
	return o.RequestTemplateID
}

func (o *Resource) GetRequireMfaToApprove() *bool {
	if o == nil {
		return nil
	}
	return o.RequireMfaToApprove
}

func (o *Resource) GetRequireMfaToConnect() *bool {
	if o == nil {
		return nil
	}
	return o.RequireMfaToConnect
}

func (o *Resource) GetRequireMfaToRequest() *bool {
	if o == nil {
		return nil
	}
	return o.RequireMfaToRequest
}

func (o *Resource) GetRequireSupportTicket() *bool {
	if o == nil {
		return nil
	}
	return o.RequireSupportTicket
}

func (o *Resource) GetResourceType() *ResourceTypeEnum {
	if o == nil {
		return nil
	}
	return o.ResourceType
}
