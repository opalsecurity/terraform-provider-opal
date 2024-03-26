// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/opal-dev/terraform-provider-opal/internal/sdk/models/shared"
	"net/http"
)

type GetGroupsRequest struct {
	// The group ids to filter by.
	GroupIds []string `queryParam:"style=form,explode=false,name=group_ids"`
	// Group name.
	GroupName *string `queryParam:"style=form,explode=true,name=group_name"`
	// The group type to filter by.
	GroupTypeFilter *shared.GroupTypeEnum `queryParam:"style=form,explode=true,name=group_type_filter"`
	// Number of results to return per page. Default is 200.
	PageSize *int64 `queryParam:"style=form,explode=true,name=page_size"`
}

func (o *GetGroupsRequest) GetGroupIds() []string {
	if o == nil {
		return nil
	}
	return o.GroupIds
}

func (o *GetGroupsRequest) GetGroupName() *string {
	if o == nil {
		return nil
	}
	return o.GroupName
}

func (o *GetGroupsRequest) GetGroupTypeFilter() *shared.GroupTypeEnum {
	if o == nil {
		return nil
	}
	return o.GroupTypeFilter
}

func (o *GetGroupsRequest) GetPageSize() *int64 {
	if o == nil {
		return nil
	}
	return o.PageSize
}

type GetGroupsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// One page worth groups associated with your organization.
	PaginatedGroupsList *shared.PaginatedGroupsList
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetGroupsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetGroupsResponse) GetPaginatedGroupsList() *shared.PaginatedGroupsList {
	if o == nil {
		return nil
	}
	return o.PaginatedGroupsList
}

func (o *GetGroupsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetGroupsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}