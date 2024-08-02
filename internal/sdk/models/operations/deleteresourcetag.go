// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"net/http"
)

type DeleteResourceTagRequest struct {
	// The ID of the resource to remove the tag from.
	ResourceID string `pathParam:"style=simple,explode=false,name=resource_id"`
	// The ID of the tag to remove.
	TagID string `pathParam:"style=simple,explode=false,name=tag_id"`
}

func (o *DeleteResourceTagRequest) GetResourceID() string {
	if o == nil {
		return ""
	}
	return o.ResourceID
}

func (o *DeleteResourceTagRequest) GetTagID() string {
	if o == nil {
		return ""
	}
	return o.TagID
}

type DeleteResourceTagResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *DeleteResourceTagResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteResourceTagResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteResourceTagResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
