// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"net/http"
)

type AddGroupTagRequest struct {
	// The ID of the group to apply the tag to.
	GroupID string `pathParam:"style=simple,explode=false,name=group_id"`
	// The ID of the tag to apply.
	TagID string `pathParam:"style=simple,explode=false,name=tag_id"`
}

func (o *AddGroupTagRequest) GetGroupID() string {
	if o == nil {
		return ""
	}
	return o.GroupID
}

func (o *AddGroupTagRequest) GetTagID() string {
	if o == nil {
		return ""
	}
	return o.TagID
}

type AddGroupTagResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *AddGroupTagResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *AddGroupTagResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *AddGroupTagResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
