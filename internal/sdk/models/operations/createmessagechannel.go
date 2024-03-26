// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/opal-dev/terraform-provider-opal/internal/sdk/models/shared"
	"net/http"
)

type CreateMessageChannelResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// The message channel that was created.
	MessageChannel *shared.MessageChannel
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *CreateMessageChannelResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateMessageChannelResponse) GetMessageChannel() *shared.MessageChannel {
	if o == nil {
		return nil
	}
	return o.MessageChannel
}

func (o *CreateMessageChannelResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateMessageChannelResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
