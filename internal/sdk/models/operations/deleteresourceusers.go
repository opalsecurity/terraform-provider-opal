// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type DeleteResourceUsersRequest struct {
	// The remote ID of the access level for which this user has direct access. If omitted, the default access level remote ID value (empty string) is assumed.
	AccessLevelRemoteID *string `queryParam:"style=form,explode=true,name=access_level_remote_id"`
	// The ID of the resource.
	ResourceID string `pathParam:"style=simple,explode=false,name=resource_id"`
	// The ID of a user to remove from this resource.
	UserID string `pathParam:"style=simple,explode=false,name=user_id"`
}

func (o *DeleteResourceUsersRequest) GetAccessLevelRemoteID() *string {
	if o == nil {
		return nil
	}
	return o.AccessLevelRemoteID
}

func (o *DeleteResourceUsersRequest) GetResourceID() string {
	if o == nil {
		return ""
	}
	return o.ResourceID
}

func (o *DeleteResourceUsersRequest) GetUserID() string {
	if o == nil {
		return ""
	}
	return o.UserID
}

type DeleteResourceUsersResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *DeleteResourceUsersResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteResourceUsersResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteResourceUsersResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}