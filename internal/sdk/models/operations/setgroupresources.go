// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/opalsecurity/terraform-provider-opal/internal/sdk/models/shared"
	"net/http"
)

type SetGroupResourcesRequest struct {
	UpdateGroupResourcesInfo shared.UpdateGroupResourcesInfo `request:"mediaType=application/json"`
	// The ID of the group.
	GroupID string `pathParam:"style=simple,explode=false,name=group_id"`
}

func (o *SetGroupResourcesRequest) GetUpdateGroupResourcesInfo() shared.UpdateGroupResourcesInfo {
	if o == nil {
		return shared.UpdateGroupResourcesInfo{}
	}
	return o.UpdateGroupResourcesInfo
}

func (o *SetGroupResourcesRequest) GetGroupID() string {
	if o == nil {
		return ""
	}
	return o.GroupID
}

type SetGroupResourcesResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *SetGroupResourcesResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *SetGroupResourcesResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *SetGroupResourcesResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
