// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type UpdateGroupBindingInfoList struct {
	// A list of group bindings with information to update.
	GroupBindings []UpdateGroupBindingInfo `json:"group_bindings"`
}

func (o *UpdateGroupBindingInfoList) GetGroupBindings() []UpdateGroupBindingInfo {
	if o == nil {
		return []UpdateGroupBindingInfo{}
	}
	return o.GroupBindings
}
