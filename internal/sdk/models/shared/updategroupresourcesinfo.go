// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type UpdateGroupResourcesInfo struct {
	Resources []ResourceWithAccessLevel `json:"resources"`
}

func (o *UpdateGroupResourcesInfo) GetResources() []ResourceWithAccessLevel {
	if o == nil {
		return []ResourceWithAccessLevel{}
	}
	return o.Resources
}
