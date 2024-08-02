// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/opalsecurity/terraform-provider-opal/internal/sdk/models/shared"
	"net/http"
)

type GetResourceUsersRequest struct {
	// Limit the number of results returned.
	Limit *int64 `queryParam:"style=form,explode=true,name=limit"`
	// The ID of the resource.
	ResourceID string `pathParam:"style=simple,explode=false,name=resource_id"`
}

func (o *GetResourceUsersRequest) GetLimit() *int64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *GetResourceUsersRequest) GetResourceID() string {
	if o == nil {
		return ""
	}
	return o.ResourceID
}

type GetResourceUsersResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// List of users with access to this resource.
	ResourceAccessUserList *shared.ResourceAccessUserList
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetResourceUsersResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetResourceUsersResponse) GetResourceAccessUserList() *shared.ResourceAccessUserList {
	if o == nil {
		return nil
	}
	return o.ResourceAccessUserList
}

func (o *GetResourceUsersResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetResourceUsersResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
