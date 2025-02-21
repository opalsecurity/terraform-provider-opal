// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/opalsecurity/terraform-provider-opal/internal/sdk/models/shared"
	"net/http"
)

type GetAppRequest struct {
	// The ID of the app.
	AppID string `pathParam:"style=simple,explode=true,name=app_id"`
}

func (o *GetAppRequest) GetAppID() string {
	if o == nil {
		return ""
	}
	return o.AppID
}

type GetAppResponse struct {
	// The requested `App`.
	App *shared.App
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetAppResponse) GetApp() *shared.App {
	if o == nil {
		return nil
	}
	return o.App
}

func (o *GetAppResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetAppResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetAppResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
