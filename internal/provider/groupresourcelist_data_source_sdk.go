// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/opalsecurity/terraform-provider-opal/internal/provider/types"
	"github.com/opalsecurity/terraform-provider-opal/internal/sdk/models/shared"
)

func (r *GroupResourceListDataSourceModel) RefreshFromSharedGroupResourceList(ctx context.Context, resp *shared.GroupResourceList) diag.Diagnostics {
	var diags diag.Diagnostics

	if resp != nil {
		r.GroupResources = []tfTypes.GroupResource{}
		if len(r.GroupResources) > len(resp.GroupResources) {
			r.GroupResources = r.GroupResources[:len(resp.GroupResources)]
		}
		for groupResourcesCount, groupResourcesItem := range resp.GroupResources {
			var groupResources tfTypes.GroupResource
			groupResources.AccessLevel.AccessLevelName = types.StringValue(groupResourcesItem.AccessLevel.AccessLevelName)
			groupResources.AccessLevel.AccessLevelRemoteID = types.StringValue(groupResourcesItem.AccessLevel.AccessLevelRemoteID)
			groupResources.GroupID = types.StringValue(groupResourcesItem.GroupID)
			groupResources.ResourceID = types.StringValue(groupResourcesItem.ResourceID)
			if groupResourcesCount+1 > len(r.GroupResources) {
				r.GroupResources = append(r.GroupResources, groupResources)
			} else {
				r.GroupResources[groupResourcesCount].AccessLevel = groupResources.AccessLevel
				r.GroupResources[groupResourcesCount].GroupID = groupResources.GroupID
				r.GroupResources[groupResourcesCount].ResourceID = groupResources.ResourceID
			}
		}
	}

	return diags
}
