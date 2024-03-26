// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/opal-dev/terraform-provider-opal/internal/sdk/models/shared"
	"net/http"
)

type GetUsersRequest struct {
	// The pagination cursor value.
	Cursor *string `queryParam:"style=form,explode=true,name=cursor"`
	// Number of results to return per page. Default is 200.
	PageSize *int64 `queryParam:"style=form,explode=true,name=page_size"`
}

func (o *GetUsersRequest) GetCursor() *string {
	if o == nil {
		return nil
	}
	return o.Cursor
}

func (o *GetUsersRequest) GetPageSize() *int64 {
	if o == nil {
		return nil
	}
	return o.PageSize
}

type GetUsersResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// One page worth users in your organization.
	PaginatedUsersList *shared.PaginatedUsersList
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetUsersResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetUsersResponse) GetPaginatedUsersList() *shared.PaginatedUsersList {
	if o == nil {
		return nil
	}
	return o.PaginatedUsersList
}

func (o *GetUsersResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetUsersResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}