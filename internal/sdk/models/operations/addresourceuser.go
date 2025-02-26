// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/opalsecurity/terraform-provider-opal/internal/sdk/models/shared"
	"net/http"
)

type AddResourceUserRequestBody struct {
	// The remote ID of the access level to grant to this user. If omitted, the default access level remote ID value (empty string) is used.
	AccessLevelRemoteID *string `json:"access_level_remote_id,omitempty"`
	// The duration for which the resource can be accessed (in minutes). Use 0 to set to indefinite.
	DurationMinutes int64 `json:"duration_minutes"`
}

func (o *AddResourceUserRequestBody) GetAccessLevelRemoteID() *string {
	if o == nil {
		return nil
	}
	return o.AccessLevelRemoteID
}

func (o *AddResourceUserRequestBody) GetDurationMinutes() int64 {
	if o == nil {
		return 0
	}
	return o.DurationMinutes
}

type AddResourceUserRequest struct {
	RequestBody *AddResourceUserRequestBody `request:"mediaType=application/json"`
	// The remote ID of the access level to grant to this user. If omitted, the default access level remote ID value (empty string) is used.
	//
	// Deprecated: This will be removed in a future release, please migrate away from it as soon as possible.
	AccessLevelRemoteID *string `queryParam:"style=form,explode=true,name=access_level_remote_id"`
	// The duration for which the resource can be accessed (in minutes). Use 0 to set to indefinite.
	//
	// Deprecated: This will be removed in a future release, please migrate away from it as soon as possible.
	DurationMinutes *int64 `queryParam:"style=form,explode=true,name=duration_minutes"`
	// The ID of the resource.
	ResourceID string `pathParam:"style=simple,explode=false,name=resource_id"`
	// The ID of the user to add.
	UserID string `pathParam:"style=simple,explode=false,name=user_id"`
}

func (o *AddResourceUserRequest) GetRequestBody() *AddResourceUserRequestBody {
	if o == nil {
		return nil
	}
	return o.RequestBody
}

func (o *AddResourceUserRequest) GetAccessLevelRemoteID() *string {
	if o == nil {
		return nil
	}
	return o.AccessLevelRemoteID
}

func (o *AddResourceUserRequest) GetDurationMinutes() *int64 {
	if o == nil {
		return nil
	}
	return o.DurationMinutes
}

func (o *AddResourceUserRequest) GetResourceID() string {
	if o == nil {
		return ""
	}
	return o.ResourceID
}

func (o *AddResourceUserRequest) GetUserID() string {
	if o == nil {
		return ""
	}
	return o.UserID
}

type AddResourceUserResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// The ResourceUser that was created.
	ResourceUser *shared.ResourceUser
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *AddResourceUserResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *AddResourceUserResponse) GetResourceUser() *shared.ResourceUser {
	if o == nil {
		return nil
	}
	return o.ResourceUser
}

func (o *AddResourceUserResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *AddResourceUserResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
