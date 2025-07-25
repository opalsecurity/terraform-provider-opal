// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/opalsecurity/terraform-provider-opal/v3/internal/sdk/models/shared"
	"net/http"
)

type CreateGroupResourcesRequestBody struct {
	// The remote ID of the access level to grant to this user. If omitted, the default access level remote ID value (empty string) is used.
	AccessLevelRemoteID *string `json:"access_level_remote_id,omitempty"`
	// The duration for which the resource can be accessed (in minutes). Use 0 to set to indefinite.
	DurationMinutes *int64 `json:"duration_minutes,omitempty"`
}

func (o *CreateGroupResourcesRequestBody) GetAccessLevelRemoteID() *string {
	if o == nil {
		return nil
	}
	return o.AccessLevelRemoteID
}

func (o *CreateGroupResourcesRequestBody) GetDurationMinutes() *int64 {
	if o == nil {
		return nil
	}
	return o.DurationMinutes
}

type CreateGroupResourcesRequest struct {
	RequestBody *CreateGroupResourcesRequestBody `request:"mediaType=application/json"`
	// The remote ID of the access level to grant to this user. If omitted, the default access level remote ID value (empty string) is used.
	//
	// Deprecated: This will be removed in a future release, please migrate away from it as soon as possible.
	AccessLevelRemoteID *string `queryParam:"style=form,explode=true,name=access_level_remote_id"`
	// The ID of the group.
	GroupID string `pathParam:"style=simple,explode=false,name=group_id"`
	// The ID of the resource.
	ResourceID string `pathParam:"style=simple,explode=false,name=resource_id"`
}

func (o *CreateGroupResourcesRequest) GetRequestBody() *CreateGroupResourcesRequestBody {
	if o == nil {
		return nil
	}
	return o.RequestBody
}

func (o *CreateGroupResourcesRequest) GetAccessLevelRemoteID() *string {
	if o == nil {
		return nil
	}
	return o.AccessLevelRemoteID
}

func (o *CreateGroupResourcesRequest) GetGroupID() string {
	if o == nil {
		return ""
	}
	return o.GroupID
}

func (o *CreateGroupResourcesRequest) GetResourceID() string {
	if o == nil {
		return ""
	}
	return o.ResourceID
}

type CreateGroupResourcesResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// The created `GroupResource` object.
	GroupResource *shared.GroupResource
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *CreateGroupResourcesResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateGroupResourcesResponse) GetGroupResource() *shared.GroupResource {
	if o == nil {
		return nil
	}
	return o.GroupResource
}

func (o *CreateGroupResourcesResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateGroupResourcesResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
