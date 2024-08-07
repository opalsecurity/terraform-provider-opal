// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type Groups struct {
	// The ID of the group.
	GroupID string `json:"group_id"`
}

func (o *Groups) GetGroupID() string {
	if o == nil {
		return ""
	}
	return o.GroupID
}

// # CreateGroupBindingInfo Object
// ### Description
// The `CreateGroupBindingInfo` object is used as an input to the CreateGroupBinding API.
type CreateGroupBindingInfo struct {
	// The list of groups.
	Groups []Groups `json:"groups"`
	// The ID of the source group.
	SourceGroupID string `json:"source_group_id"`
}

func (o *CreateGroupBindingInfo) GetGroups() []Groups {
	if o == nil {
		return []Groups{}
	}
	return o.Groups
}

func (o *CreateGroupBindingInfo) GetSourceGroupID() string {
	if o == nil {
		return ""
	}
	return o.SourceGroupID
}
