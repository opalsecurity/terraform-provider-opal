// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// # CreateTagInfo Object
// ### Description
// The `CreateTagInfo` object is used to represent configuration for a new tag.
//
// ### Usage Example
// Use in the `POST Tag` endpoint.
type CreateTagInfo struct {
	// The key of the tag to create.
	Key *string `json:"tag_key,omitempty"`
	// The value of the tag to create.
	Value *string `json:"tag_value,omitempty"`
}

func (o *CreateTagInfo) GetKey() *string {
	if o == nil {
		return nil
	}
	return o.Key
}

func (o *CreateTagInfo) GetValue() *string {
	if o == nil {
		return nil
	}
	return o.Value
}
