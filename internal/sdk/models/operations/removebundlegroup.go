// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"net/http"
)

type RemoveBundleGroupRequest struct {
	// The ID of the bundle.
	BundleID string `pathParam:"style=simple,explode=false,name=bundle_id"`
	// The ID of the group to remove.
	GroupID string `pathParam:"style=simple,explode=false,name=group_id"`
}

func (o *RemoveBundleGroupRequest) GetBundleID() string {
	if o == nil {
		return ""
	}
	return o.BundleID
}

func (o *RemoveBundleGroupRequest) GetGroupID() string {
	if o == nil {
		return ""
	}
	return o.GroupID
}

type RemoveBundleGroupResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *RemoveBundleGroupResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *RemoveBundleGroupResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *RemoveBundleGroupResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
