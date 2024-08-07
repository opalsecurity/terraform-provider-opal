// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// # MessageChannel Object
// ### Description
// The `MessageChannel` object is used to represent a message channel.
//
// ### Usage Example
// Update a groups message channel from the `UPDATE Groups` endpoint.
type MessageChannel struct {
	// The ID of the message channel.
	ID string `json:"message_channel_id"`
	// A bool representing whether or not the message channel is private.
	IsPrivate *bool `json:"is_private,omitempty"`
	// The name of the message channel.
	Name *string `json:"name,omitempty"`
	// The remote ID of the message channel
	RemoteID *string `json:"remote_id,omitempty"`
	// The third party provider of the message channel.
	ThirdPartyProvider *MessageChannelProviderEnum `json:"third_party_provider,omitempty"`
}

func (o *MessageChannel) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *MessageChannel) GetIsPrivate() *bool {
	if o == nil {
		return nil
	}
	return o.IsPrivate
}

func (o *MessageChannel) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *MessageChannel) GetRemoteID() *string {
	if o == nil {
		return nil
	}
	return o.RemoteID
}

func (o *MessageChannel) GetThirdPartyProvider() *MessageChannelProviderEnum {
	if o == nil {
		return nil
	}
	return o.ThirdPartyProvider
}
