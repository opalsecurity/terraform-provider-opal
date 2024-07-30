// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"github.com/opalsecurity/terraform-provider-opal/internal/sdk/internal/utils"
	"time"
)

// GroupBinding - # Group Binding Object
// ### Description
// The `GroupBinding` object is used to represent a group binding.
//
// ### Usage Example
// Get group bindings from the `GET Group Bindings` endpoint.
type GroupBinding struct {
	// The date the group binding was created.
	CreatedAt time.Time `json:"created_at"`
	// The ID of the user that created the group binding.
	CreatedByID string `json:"created_by_id"`
	// The ID of the group binding.
	GroupBindingID string `json:"group_binding_id"`
	// The list of groups.
	Groups []GroupBindingGroup `json:"groups"`
	// The ID of the source group.
	SourceGroupID string `json:"source_group_id"`
}

func (g GroupBinding) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GroupBinding) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *GroupBinding) GetCreatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreatedAt
}

func (o *GroupBinding) GetCreatedByID() string {
	if o == nil {
		return ""
	}
	return o.CreatedByID
}

func (o *GroupBinding) GetGroupBindingID() string {
	if o == nil {
		return ""
	}
	return o.GroupBindingID
}

func (o *GroupBinding) GetGroups() []GroupBindingGroup {
	if o == nil {
		return []GroupBindingGroup{}
	}
	return o.Groups
}

func (o *GroupBinding) GetSourceGroupID() string {
	if o == nil {
		return ""
	}
	return o.SourceGroupID
}
