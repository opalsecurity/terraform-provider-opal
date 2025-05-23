// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/opalsecurity/terraform-provider-opal/internal/provider/types"
	"github.com/opalsecurity/terraform-provider-opal/internal/sdk/models/shared"
)

func (r *PaginatedBundleGroupListDataSourceModel) RefreshFromSharedPaginatedBundleGroupList(ctx context.Context, resp *shared.PaginatedBundleGroupList) diag.Diagnostics {
	var diags diag.Diagnostics

	if resp != nil {
		r.BundleGroups = []tfTypes.BundleGroup{}
		if len(r.BundleGroups) > len(resp.BundleGroups) {
			r.BundleGroups = r.BundleGroups[:len(resp.BundleGroups)]
		}
		for bundleGroupsCount, bundleGroupsItem := range resp.BundleGroups {
			var bundleGroups tfTypes.BundleGroup
			bundleGroups.BundleID = types.StringPointerValue(bundleGroupsItem.BundleID)
			bundleGroups.GroupID = types.StringPointerValue(bundleGroupsItem.GroupID)
			if bundleGroupsCount+1 > len(r.BundleGroups) {
				r.BundleGroups = append(r.BundleGroups, bundleGroups)
			} else {
				r.BundleGroups[bundleGroupsCount].BundleID = bundleGroups.BundleID
				r.BundleGroups[bundleGroupsCount].GroupID = bundleGroups.GroupID
			}
		}
		r.Next = types.StringPointerValue(resp.Next)
		r.Previous = types.StringPointerValue(resp.Previous)
		r.TotalCount = types.Int64PointerValue(resp.TotalCount)
	}

	return diags
}
