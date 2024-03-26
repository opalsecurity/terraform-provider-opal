// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// # CreateMessageChannelInfo Object
// ### Description
// The `CreateMessageChannelInfo` object is used to describe the message channel object to be created.
type CreateMessageChannelInfo struct {
	// The remote ID of the message channel
	RemoteID string `json:"remote_id"`
	// The third party provider of the message channel.
	ThirdPartyProvider MessageChannelProviderEnum `json:"third_party_provider"`
}

func (o *CreateMessageChannelInfo) GetRemoteID() string {
	if o == nil {
		return ""
	}
	return o.RemoteID
}

func (o *CreateMessageChannelInfo) GetThirdPartyProvider() MessageChannelProviderEnum {
	if o == nil {
		return MessageChannelProviderEnum("")
	}
	return o.ThirdPartyProvider
}