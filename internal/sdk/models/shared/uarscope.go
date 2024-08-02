// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// FilterOperator - Specifies whether entities must match all (AND) or any (OR) of the filters.
type FilterOperator string

const (
	FilterOperatorAny FilterOperator = "ANY"
	FilterOperatorAll FilterOperator = "ALL"
)

func (e FilterOperator) ToPointer() *FilterOperator {
	return &e
}
func (e *FilterOperator) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "ANY":
		fallthrough
	case "ALL":
		*e = FilterOperator(v)
		return nil
	default:
		return fmt.Errorf("invalid value for FilterOperator: %v", v)
	}
}

// GroupVisibility - Specifies what users can see during an Access Review
type GroupVisibility string

const (
	GroupVisibilityStrict                 GroupVisibility = "STRICT"
	GroupVisibilityViewVisibleAndAssigned GroupVisibility = "VIEW_VISIBLE_AND_ASSIGNED"
	GroupVisibilityViewAll                GroupVisibility = "VIEW_ALL"
)

func (e GroupVisibility) ToPointer() *GroupVisibility {
	return &e
}
func (e *GroupVisibility) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "STRICT":
		fallthrough
	case "VIEW_VISIBLE_AND_ASSIGNED":
		fallthrough
	case "VIEW_ALL":
		*e = GroupVisibility(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GroupVisibility: %v", v)
	}
}

// UARScope - If set, the access review will only contain resources and groups that match at least one of the filters in scope.
type UARScope struct {
	// This access review will include resources and groups who are owned by one of the owners corresponding to the given IDs.
	Admins []string `json:"admins,omitempty"`
	// This access review will include items in the specified applications
	Apps []string `json:"apps,omitempty"`
	// This access review will include resources and groups with ids in the given strings.
	Entities []string `json:"entities,omitempty"`
	// Specifies whether entities must match all (AND) or any (OR) of the filters.
	FilterOperator *FilterOperator `json:"filter_operator,omitempty"`
	// This access review will include items of the specified group types
	GroupTypes []GroupTypeEnum `json:"group_types,omitempty"`
	// Specifies what users can see during an Access Review
	GroupVisibility      *GroupVisibility `json:"group_visibility,omitempty"`
	IncludeGroupBindings *bool            `json:"include_group_bindings,omitempty"`
	// This access review will include resources and groups whose name contains one of the given strings.
	Names []string `json:"names,omitempty"`
	// This access review will include items of the specified resource types
	ResourceTypes []ResourceTypeEnum `json:"resource_types,omitempty"`
	// This access review will include resources and groups who are tagged with one of the given tags.
	Tags []TagFilter `json:"tags,omitempty"`
	// The access review will only include the following users. If any users are selected, any entity filters will be applied to only the entities that the selected users have access to.
	Users []string `json:"users,omitempty"`
}

func (o *UARScope) GetAdmins() []string {
	if o == nil {
		return nil
	}
	return o.Admins
}

func (o *UARScope) GetApps() []string {
	if o == nil {
		return nil
	}
	return o.Apps
}

func (o *UARScope) GetEntities() []string {
	if o == nil {
		return nil
	}
	return o.Entities
}

func (o *UARScope) GetFilterOperator() *FilterOperator {
	if o == nil {
		return nil
	}
	return o.FilterOperator
}

func (o *UARScope) GetGroupTypes() []GroupTypeEnum {
	if o == nil {
		return nil
	}
	return o.GroupTypes
}

func (o *UARScope) GetGroupVisibility() *GroupVisibility {
	if o == nil {
		return nil
	}
	return o.GroupVisibility
}

func (o *UARScope) GetIncludeGroupBindings() *bool {
	if o == nil {
		return nil
	}
	return o.IncludeGroupBindings
}

func (o *UARScope) GetNames() []string {
	if o == nil {
		return nil
	}
	return o.Names
}

func (o *UARScope) GetResourceTypes() []ResourceTypeEnum {
	if o == nil {
		return nil
	}
	return o.ResourceTypes
}

func (o *UARScope) GetTags() []TagFilter {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *UARScope) GetUsers() []string {
	if o == nil {
		return nil
	}
	return o.Users
}
