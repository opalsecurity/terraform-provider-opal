// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/opalsecurity/terraform-provider-opal/internal/provider/types"
	"github.com/opalsecurity/terraform-provider-opal/internal/sdk/models/operations"
	"github.com/opalsecurity/terraform-provider-opal/internal/sdk/models/shared"
)

func (r *IdpGroupMappingsResourceModel) ToOperationsUpdateIdpGroupMappingsRequestBody() *operations.UpdateIdpGroupMappingsRequestBody {
	var mappings []operations.Mappings = []operations.Mappings{}
	for _, mappingsItem := range r.Mappings {
		alias := new(string)
		if !mappingsItem.Alias.IsUnknown() && !mappingsItem.Alias.IsNull() {
			*alias = mappingsItem.Alias.ValueString()
		} else {
			alias = nil
		}
		groupID := new(string)
		if !mappingsItem.GroupID.IsUnknown() && !mappingsItem.GroupID.IsNull() {
			*groupID = mappingsItem.GroupID.ValueString()
		} else {
			groupID = nil
		}
		hiddenFromEndUser := new(bool)
		if !mappingsItem.HiddenFromEndUser.IsUnknown() && !mappingsItem.HiddenFromEndUser.IsNull() {
			*hiddenFromEndUser = mappingsItem.HiddenFromEndUser.ValueBool()
		} else {
			hiddenFromEndUser = nil
		}
		mappings = append(mappings, operations.Mappings{
			Alias:             alias,
			GroupID:           groupID,
			HiddenFromEndUser: hiddenFromEndUser,
		})
	}
	out := operations.UpdateIdpGroupMappingsRequestBody{
		Mappings: mappings,
	}
	return &out
}

func (r *IdpGroupMappingsResourceModel) RefreshFromSharedIdpGroupMappingList(resp *shared.IdpGroupMappingList) {
	if resp != nil {
		r.Mappings = []tfTypes.Mappings{}
		if len(r.Mappings) > len(resp.Mappings) {
			r.Mappings = r.Mappings[:len(resp.Mappings)]
		}
		for mappingsCount, mappingsItem := range resp.Mappings {
			var mappings1 tfTypes.Mappings
			mappings1.Alias = types.StringPointerValue(mappingsItem.Alias)
			mappings1.GroupID = types.StringValue(mappingsItem.GroupID)
			mappings1.HiddenFromEndUser = types.BoolValue(mappingsItem.HiddenFromEndUser)
			if mappingsCount+1 > len(r.Mappings) {
				r.Mappings = append(r.Mappings, mappings1)
			} else {
				r.Mappings[mappingsCount].Alias = mappings1.Alias
				r.Mappings[mappingsCount].GroupID = mappings1.GroupID
				r.Mappings[mappingsCount].HiddenFromEndUser = mappings1.HiddenFromEndUser
			}
		}
	}
}
