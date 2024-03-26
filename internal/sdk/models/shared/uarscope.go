// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// UARScope - If set, the access review will only contain resources and groups that match at least one of the filters in scope.
type UARScope struct {
	// This access review will include resources and groups who are owned by one of the owners corresponding to the given IDs.
	Admins []string `json:"admins,omitempty"`
	// This access review will include resources and groups whose name contains one of the given strings.
	Names []string `json:"names,omitempty"`
	// This access review will include resources and groups who are tagged with one of the given tags.
	Tags []TagFilter `json:"tags,omitempty"`
}

func (o *UARScope) GetAdmins() []string {
	if o == nil {
		return nil
	}
	return o.Admins
}

func (o *UARScope) GetNames() []string {
	if o == nil {
		return nil
	}
	return o.Names
}

func (o *UARScope) GetTags() []TagFilter {
	if o == nil {
		return nil
	}
	return o.Tags
}
