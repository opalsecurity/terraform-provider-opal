// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type ResourceAccessUserList struct {
	Results []ResourceAccessUser `json:"results,omitempty"`
}

func (o *ResourceAccessUserList) GetResults() []ResourceAccessUser {
	if o == nil {
		return nil
	}
	return o.Results
}
