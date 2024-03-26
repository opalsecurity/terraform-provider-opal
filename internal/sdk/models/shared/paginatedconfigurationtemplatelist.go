// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// # PaginatedConfigurationTemplateList Object
// ### Description
// The `PaginatedConfigurationTemplateList` object is used to store a list of configuration templates.
//
// ### Usage Example
// Returned from the `GET Configuration Templates` endpoint.
type PaginatedConfigurationTemplateList struct {
	Results []ConfigurationTemplate `json:"results,omitempty"`
}

func (o *PaginatedConfigurationTemplateList) GetResults() []ConfigurationTemplate {
	if o == nil {
		return nil
	}
	return o.Results
}