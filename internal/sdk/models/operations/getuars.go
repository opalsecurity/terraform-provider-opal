// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/opal-dev/terraform-provider-opal/internal/sdk/models/shared"
	"net/http"
)

type GetUARsRequest struct {
	// The pagination cursor value.
	Cursor *string `queryParam:"style=form,explode=true,name=cursor"`
	// Number of results to return per page. Default is 200.
	PageSize *int64 `queryParam:"style=form,explode=true,name=page_size"`
}

func (o *GetUARsRequest) GetCursor() *string {
	if o == nil {
		return nil
	}
	return o.Cursor
}

func (o *GetUARsRequest) GetPageSize() *int64 {
	if o == nil {
		return nil
	}
	return o.PageSize
}

type GetUARsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// A list of UARs for your organization.
	PaginatedUARsList *shared.PaginatedUARsList
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetUARsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetUARsResponse) GetPaginatedUARsList() *shared.PaginatedUARsList {
	if o == nil {
		return nil
	}
	return o.PaginatedUARsList
}

func (o *GetUARsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetUARsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
