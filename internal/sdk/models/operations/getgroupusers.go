// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/opalsecurity/terraform-provider-opal/v3/internal/sdk/models/shared"
	"net/http"
)

type GetGroupUsersRequest struct {
	// The ID of the group.
	GroupID string `pathParam:"style=simple,explode=false,name=group_id"`
}

func (o *GetGroupUsersRequest) GetGroupID() string {
	if o == nil {
		return ""
	}
	return o.GroupID
}

type GetGroupUsersResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// List of users with access to this group.
	GroupUserList *shared.GroupUserList
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetGroupUsersResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetGroupUsersResponse) GetGroupUserList() *shared.GroupUserList {
	if o == nil {
		return nil
	}
	return o.GroupUserList
}

func (o *GetGroupUsersResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetGroupUsersResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
