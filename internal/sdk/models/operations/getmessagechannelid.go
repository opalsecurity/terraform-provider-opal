// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/opal-dev/terraform-provider-opal/internal/sdk/models/shared"
	"net/http"
)

type GetMessageChannelIDRequest struct {
	// The ID of the message_channel.
	ID string `pathParam:"style=simple,explode=false,name=message_channel_id"`
}

func (o *GetMessageChannelIDRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type GetMessageChannelIDResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// The requested message channel.
	MessageChannel *shared.MessageChannel
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetMessageChannelIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetMessageChannelIDResponse) GetMessageChannel() *shared.MessageChannel {
	if o == nil {
		return nil
	}
	return o.MessageChannel
}

func (o *GetMessageChannelIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetMessageChannelIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
