// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/opalsecurity/terraform-provider-opal/internal/provider/types"
	"github.com/opalsecurity/terraform-provider-opal/internal/sdk/models/shared"
)

func (r *IdpGroupMappingDataSourceModel) RefreshFromSharedIdpGroupMappingList(resp *shared.IdpGroupMappingList) {
	if resp != nil {
		r.IdpGroupMappings = []tfTypes.IdpGroupMapping{}
		if len(r.IdpGroupMappings) > len(resp.IdpGroupMappings) {
			r.IdpGroupMappings = r.IdpGroupMappings[:len(resp.IdpGroupMappings)]
		}
		for idpGroupMappingsCount, idpGroupMappingsItem := range resp.IdpGroupMappings {
			var idpGroupMappings1 tfTypes.IdpGroupMapping
			idpGroupMappings1.Alias = types.StringPointerValue(idpGroupMappingsItem.Alias)
			idpGroupMappings1.AppResourceID = types.StringValue(idpGroupMappingsItem.AppResourceID)
			idpGroupMappings1.GroupID = types.StringValue(idpGroupMappingsItem.GroupID)
			idpGroupMappings1.HiddenFromEndUser = types.BoolValue(idpGroupMappingsItem.HiddenFromEndUser)
			idpGroupMappings1.ID = types.StringValue(idpGroupMappingsItem.ID)
			idpGroupMappings1.OrganizationID = types.StringValue(idpGroupMappingsItem.OrganizationID)
			if idpGroupMappingsCount+1 > len(r.IdpGroupMappings) {
				r.IdpGroupMappings = append(r.IdpGroupMappings, idpGroupMappings1)
			} else {
				r.IdpGroupMappings[idpGroupMappingsCount].Alias = idpGroupMappings1.Alias
				r.IdpGroupMappings[idpGroupMappingsCount].AppResourceID = idpGroupMappings1.AppResourceID
				r.IdpGroupMappings[idpGroupMappingsCount].GroupID = idpGroupMappings1.GroupID
				r.IdpGroupMappings[idpGroupMappingsCount].HiddenFromEndUser = idpGroupMappings1.HiddenFromEndUser
				r.IdpGroupMappings[idpGroupMappingsCount].ID = idpGroupMappings1.ID
				r.IdpGroupMappings[idpGroupMappingsCount].OrganizationID = idpGroupMappings1.OrganizationID
			}
		}
	}
}
