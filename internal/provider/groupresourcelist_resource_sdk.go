// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/opalsecurity/terraform-provider-opal/internal/provider/types"
	"github.com/opalsecurity/terraform-provider-opal/internal/sdk/models/shared"
)

func (r *GroupResourceListResourceModel) ToSharedUpdateGroupResourcesInfo() *shared.UpdateGroupResourcesInfo {
	var resources []shared.ResourceWithAccessLevel = []shared.ResourceWithAccessLevel{}
	for _, resourcesItem := range r.Resources {
		accessLevelRemoteID := new(string)
		if !resourcesItem.AccessLevelRemoteID.IsUnknown() && !resourcesItem.AccessLevelRemoteID.IsNull() {
			*accessLevelRemoteID = resourcesItem.AccessLevelRemoteID.ValueString()
		} else {
			accessLevelRemoteID = nil
		}
		var resourceID string
		resourceID = resourcesItem.ResourceID.ValueString()

		resources = append(resources, shared.ResourceWithAccessLevel{
			AccessLevelRemoteID: accessLevelRemoteID,
			ResourceID:          resourceID,
		})
	}
	out := shared.UpdateGroupResourcesInfo{
		Resources: resources,
	}
	return &out
}

func (r *GroupResourceListResourceModel) RefreshFromSharedGroupResourceList(resp *shared.GroupResourceList) {
	if resp != nil {
		r.GroupResources = []tfTypes.GroupResource{}
		if len(r.GroupResources) > len(resp.GroupResources) {
			r.GroupResources = r.GroupResources[:len(resp.GroupResources)]
		}
		for groupResourcesCount, groupResourcesItem := range resp.GroupResources {
			var groupResources1 tfTypes.GroupResource
			groupResources1.AccessLevel.AccessLevelName = types.StringValue(groupResourcesItem.AccessLevel.AccessLevelName)
			groupResources1.AccessLevel.AccessLevelRemoteID = types.StringValue(groupResourcesItem.AccessLevel.AccessLevelRemoteID)
			groupResources1.GroupID = types.StringValue(groupResourcesItem.GroupID)
			groupResources1.ResourceID = types.StringValue(groupResourcesItem.ResourceID)
			if groupResourcesCount+1 > len(r.GroupResources) {
				r.GroupResources = append(r.GroupResources, groupResources1)
			} else {
				r.GroupResources[groupResourcesCount].AccessLevel = groupResources1.AccessLevel
				r.GroupResources[groupResourcesCount].GroupID = groupResources1.GroupID
				r.GroupResources[groupResourcesCount].ResourceID = groupResources1.ResourceID
			}
		}
	}
}
