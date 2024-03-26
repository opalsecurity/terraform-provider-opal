// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// GroupAccessLevel - # Access Level Object
// ### Description
// The `GroupAccessLevel` object is used to represent the level of access that a user has to a group or a group has to a group. The "default" access
// level is a `GroupAccessLevel` object whose fields are all empty strings.
//
// ### Usage Example
// View the `GroupAccessLevel` of a group/user or group/group pair to see the level of access granted to the group.
type GroupAccessLevel struct {
	// The human-readable name of the access level.
	AccessLevelName string `json:"access_level_name"`
	// The machine-readable identifier of the access level.
	AccessLevelRemoteID string `json:"access_level_remote_id"`
}

func (o *GroupAccessLevel) GetAccessLevelName() string {
	if o == nil {
		return ""
	}
	return o.AccessLevelName
}

func (o *GroupAccessLevel) GetAccessLevelRemoteID() string {
	if o == nil {
		return ""
	}
	return o.AccessLevelRemoteID
}
