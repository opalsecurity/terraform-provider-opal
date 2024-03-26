// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type UpdateResourceInfoList struct {
	// A list of resources with information to update.
	Resources []UpdateResourceInfo `json:"resources"`
}

func (o *UpdateResourceInfoList) GetResources() []UpdateResourceInfo {
	if o == nil {
		return []UpdateResourceInfo{}
	}
	return o.Resources
}
