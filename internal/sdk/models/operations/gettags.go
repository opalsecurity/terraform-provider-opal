// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/opalsecurity/terraform-provider-opal/internal/sdk/models/shared"
	"net/http"
)

type GetTagsRequest struct {
	// The pagination cursor value.
	Cursor *string `queryParam:"style=form,explode=true,name=cursor"`
	// Number of results to return per page. Default is 200.
	PageSize *int64 `queryParam:"style=form,explode=true,name=page_size"`
}

func (o *GetTagsRequest) GetCursor() *string {
	if o == nil {
		return nil
	}
	return o.Cursor
}

func (o *GetTagsRequest) GetPageSize() *int64 {
	if o == nil {
		return nil
	}
	return o.PageSize
}

type GetTagsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// A list of tags created by your organization.
	PaginatedTagsList *shared.PaginatedTagsList
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetTagsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetTagsResponse) GetPaginatedTagsList() *shared.PaginatedTagsList {
	if o == nil {
		return nil
	}
	return o.PaginatedTagsList
}

func (o *GetTagsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetTagsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
