// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/opal-dev/terraform-provider-opal/internal/sdk/models/shared"
	"net/http"
)

type GetGroupRequest struct {
	// The ID of the group.
	ID string `pathParam:"style=simple,explode=true,name=group_id"`
}

func (o *GetGroupRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type GetGroupResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// The requested `Group`.
	Group *shared.Group
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetGroupResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetGroupResponse) GetGroup() *shared.Group {
	if o == nil {
		return nil
	}
	return o.Group
}

func (o *GetGroupResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetGroupResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
