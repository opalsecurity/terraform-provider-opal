// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/opalsecurity/terraform-provider-opal/internal/sdk/models/shared"
	"net/http"
)

type AddResourceNhiRequestBody struct {
	// The remote ID of the access level to grant. If omitted, the default access level remote ID value (empty string) is used.
	AccessLevelRemoteID *string `json:"access_level_remote_id,omitempty"`
	// The duration for which the resource can be accessed (in minutes). Use 0 to set to indefinite.
	DurationMinutes int64 `json:"duration_minutes"`
}

func (o *AddResourceNhiRequestBody) GetAccessLevelRemoteID() *string {
	if o == nil {
		return nil
	}
	return o.AccessLevelRemoteID
}

func (o *AddResourceNhiRequestBody) GetDurationMinutes() int64 {
	if o == nil {
		return 0
	}
	return o.DurationMinutes
}

type AddResourceNhiRequest struct {
	RequestBody *AddResourceNhiRequestBody `request:"mediaType=application/json"`
	// The resource ID of the non-human identity to add.
	NonHumanIdentityID string `pathParam:"style=simple,explode=false,name=non_human_identity_id"`
	// The ID of the resource.
	ResourceID string `pathParam:"style=simple,explode=false,name=resource_id"`
}

func (o *AddResourceNhiRequest) GetRequestBody() *AddResourceNhiRequestBody {
	if o == nil {
		return nil
	}
	return o.RequestBody
}

func (o *AddResourceNhiRequest) GetNonHumanIdentityID() string {
	if o == nil {
		return ""
	}
	return o.NonHumanIdentityID
}

func (o *AddResourceNhiRequest) GetResourceID() string {
	if o == nil {
		return ""
	}
	return o.ResourceID
}

type AddResourceNhiResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Details about the access that the non-human identity was granted to the resource.
	ResourceNHI *shared.ResourceNHI
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *AddResourceNhiResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *AddResourceNhiResponse) GetResourceNHI() *shared.ResourceNHI {
	if o == nil {
		return nil
	}
	return o.ResourceNHI
}

func (o *AddResourceNhiResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *AddResourceNhiResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}