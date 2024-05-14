// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// VisibilityInfo - Visibility infomation of an entity.
type VisibilityInfo struct {
	// The visibility level of the entity.
	Visibility         VisibilityTypeEnum `json:"visibility"`
	VisibilityGroupIds *[]string          `json:"visibility_group_ids,omitempty"`
}

func (o *VisibilityInfo) GetVisibility() VisibilityTypeEnum {
	if o == nil {
		return VisibilityTypeEnum("")
	}
	return o.Visibility
}

func (o *VisibilityInfo) GetVisibilityGroupIds() *[]string {
	if o == nil {
		return nil
	}
	return o.VisibilityGroupIds
}
