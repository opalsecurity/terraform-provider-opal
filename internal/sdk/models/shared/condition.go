// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"github.com/opalsecurity/terraform-provider-opal/internal/sdk/internal/utils"
)

type Condition struct {
	// The list of group IDs to match.
	GroupIds []string `json:"group_ids"`
	// The list of role remote IDs to match.
	RoleRemoteIds []string `json:"role_remote_ids"`
}

func (c Condition) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *Condition) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *Condition) GetGroupIds() []string {
	if o == nil {
		return nil
	}
	return o.GroupIds
}

func (o *Condition) GetRoleRemoteIds() []string {
	if o == nil {
		return nil
	}
	return o.RoleRemoteIds
}
