// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/opalsecurity/terraform-provider-opal/internal/sdk/models/shared"
	"net/http"
)

type SetBundleVisibilityRequest struct {
	VisibilityInfo shared.VisibilityInfo `request:"mediaType=application/json"`
	// The ID of the bundle.
	BundleID string `pathParam:"style=simple,explode=false,name=bundle_id"`
}

func (o *SetBundleVisibilityRequest) GetVisibilityInfo() shared.VisibilityInfo {
	if o == nil {
		return shared.VisibilityInfo{}
	}
	return o.VisibilityInfo
}

func (o *SetBundleVisibilityRequest) GetBundleID() string {
	if o == nil {
		return ""
	}
	return o.BundleID
}

type SetBundleVisibilityResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *SetBundleVisibilityResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *SetBundleVisibilityResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *SetBundleVisibilityResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
