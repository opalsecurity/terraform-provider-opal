// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/opalsecurity/terraform-provider-opal/internal/sdk/models/shared"
	"net/http"
)

type GetIdpGroupMappingsRequest struct {
	// The ID of the Okta app.
	AppResourceID string `pathParam:"style=simple,explode=false,name=app_resource_id"`
}

func (o *GetIdpGroupMappingsRequest) GetAppResourceID() string {
	if o == nil {
		return ""
	}
	return o.AppResourceID
}

type GetIdpGroupMappingsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// The configured set of available `IdpGroupMapping` objects for an Okta app.
	IdpGroupMappingList *shared.IdpGroupMappingList
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetIdpGroupMappingsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetIdpGroupMappingsResponse) GetIdpGroupMappingList() *shared.IdpGroupMappingList {
	if o == nil {
		return nil
	}
	return o.IdpGroupMappingList
}

func (o *GetIdpGroupMappingsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetIdpGroupMappingsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
