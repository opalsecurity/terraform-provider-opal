// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/opal-dev/terraform-provider-opal/internal/sdk/models/shared"
	"net/http"
)

type GetMessageChannelsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// A list of message channels for your organization.
	MessageChannelList *shared.MessageChannelList
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetMessageChannelsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetMessageChannelsResponse) GetMessageChannelList() *shared.MessageChannelList {
	if o == nil {
		return nil
	}
	return o.MessageChannelList
}

func (o *GetMessageChannelsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetMessageChannelsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}