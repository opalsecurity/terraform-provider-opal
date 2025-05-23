// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type BundleGroup struct {
	// The ID of the bundle containing the group.
	BundleID *string `json:"bundle_id,omitempty"`
	// The ID of the group within a bundle.
	GroupID *string `json:"group_id,omitempty"`
}

func (o *BundleGroup) GetBundleID() *string {
	if o == nil {
		return nil
	}
	return o.BundleID
}

func (o *BundleGroup) GetGroupID() *string {
	if o == nil {
		return nil
	}
	return o.GroupID
}
