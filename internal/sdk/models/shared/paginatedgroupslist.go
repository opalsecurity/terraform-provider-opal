// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type PaginatedGroupsList struct {
	Results []Group `json:"results"`
}

func (o *PaginatedGroupsList) GetResults() []Group {
	if o == nil {
		return []Group{}
	}
	return o.Results
}
