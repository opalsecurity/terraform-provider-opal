// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type Condition struct {
	// The list of group IDs to match.
	GroupIds []string `json:"group_ids,omitempty"`
	// The list of role remote IDs to match.
	RoleRemoteIds []string `json:"role_remote_ids,omitempty"`
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
