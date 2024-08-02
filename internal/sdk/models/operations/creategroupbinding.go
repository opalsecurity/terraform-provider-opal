// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/opalsecurity/terraform-provider-opal/internal/sdk/models/shared"
	"net/http"
)

type CreateGroupBindingResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// The group binding just created.
	GroupBinding *shared.GroupBinding
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *CreateGroupBindingResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateGroupBindingResponse) GetGroupBinding() *shared.GroupBinding {
	if o == nil {
		return nil
	}
	return o.GroupBinding
}

func (o *CreateGroupBindingResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateGroupBindingResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
