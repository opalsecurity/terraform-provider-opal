// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type PaginatedBundleResourceList struct {
	BundleResources []BundleResource `json:"bundle_resources"`
	// The cursor with which to continue pagination if additional result pages exist.
	Next *string `json:"next,omitempty"`
	// The cursor used to obtain the current result page.
	Previous *string `json:"previous,omitempty"`
	// The total number of items in the result set.
	TotalCount *int64 `json:"total_count,omitempty"`
}

func (o *PaginatedBundleResourceList) GetBundleResources() []BundleResource {
	if o == nil {
		return []BundleResource{}
	}
	return o.BundleResources
}

func (o *PaginatedBundleResourceList) GetNext() *string {
	if o == nil {
		return nil
	}
	return o.Next
}

func (o *PaginatedBundleResourceList) GetPrevious() *string {
	if o == nil {
		return nil
	}
	return o.Previous
}

func (o *PaginatedBundleResourceList) GetTotalCount() *int64 {
	if o == nil {
		return nil
	}
	return o.TotalCount
}
