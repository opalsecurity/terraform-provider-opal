// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/opalsecurity/terraform-provider-opal/internal/sdk/models/shared"
	"net/http"
)

type SetOwnerUsersRequest struct {
	UserIDList shared.UserIDList `request:"mediaType=application/json"`
	// The ID of the owner.
	OwnerID string `pathParam:"style=simple,explode=false,name=owner_id"`
}

func (o *SetOwnerUsersRequest) GetUserIDList() shared.UserIDList {
	if o == nil {
		return shared.UserIDList{}
	}
	return o.UserIDList
}

func (o *SetOwnerUsersRequest) GetOwnerID() string {
	if o == nil {
		return ""
	}
	return o.OwnerID
}

type SetOwnerUsersResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// The updated users for the owner.
	UserList *shared.UserList
}

func (o *SetOwnerUsersResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *SetOwnerUsersResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *SetOwnerUsersResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *SetOwnerUsersResponse) GetUserList() *shared.UserList {
	if o == nil {
		return nil
	}
	return o.UserList
}
