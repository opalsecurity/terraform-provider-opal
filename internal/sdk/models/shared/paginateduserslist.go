// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type PaginatedUsersList struct {
	// The cursor with which to continue pagination if additional result pages exist.
	Next *string `json:"next,omitempty"`
	// The cursor used to obtain the current result page.
	Previous *string `json:"previous,omitempty"`
	Results  []User  `json:"results"`
}

func (o *PaginatedUsersList) GetNext() *string {
	if o == nil {
		return nil
	}
	return o.Next
}

func (o *PaginatedUsersList) GetPrevious() *string {
	if o == nil {
		return nil
	}
	return o.Previous
}

func (o *PaginatedUsersList) GetResults() []User {
	if o == nil {
		return []User{}
	}
	return o.Results
}
