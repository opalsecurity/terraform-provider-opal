// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/opalsecurity/terraform-provider-opal/v3/internal/provider/typeconvert"
	tfTypes "github.com/opalsecurity/terraform-provider-opal/v3/internal/provider/types"
	"github.com/opalsecurity/terraform-provider-opal/v3/internal/sdk/models/operations"
	"github.com/opalsecurity/terraform-provider-opal/v3/internal/sdk/models/shared"
)

func (r *PaginatedBundleListDataSourceModel) RefreshFromSharedPaginatedBundleList(ctx context.Context, resp *shared.PaginatedBundleList) diag.Diagnostics {
	var diags diag.Diagnostics

	if resp != nil {
		r.Bundles = []tfTypes.Bundle{}
		if len(r.Bundles) > len(resp.Bundles) {
			r.Bundles = r.Bundles[:len(resp.Bundles)]
		}
		for bundlesCount, bundlesItem := range resp.Bundles {
			var bundles tfTypes.Bundle
			bundles.AdminOwnerID = types.StringPointerValue(bundlesItem.AdminOwnerID)
			bundles.BundleID = types.StringPointerValue(bundlesItem.BundleID)
			bundles.CreatedAt = types.StringPointerValue(typeconvert.TimePointerToStringPointer(bundlesItem.CreatedAt))
			bundles.Description = types.StringPointerValue(bundlesItem.Description)
			bundles.Name = types.StringPointerValue(bundlesItem.Name)
			bundles.TotalNumGroups = types.Int64PointerValue(bundlesItem.TotalNumGroups)
			bundles.TotalNumItems = types.Int64PointerValue(bundlesItem.TotalNumItems)
			bundles.TotalNumResources = types.Int64PointerValue(bundlesItem.TotalNumResources)
			bundles.UpdatedAt = types.StringPointerValue(typeconvert.TimePointerToStringPointer(bundlesItem.UpdatedAt))
			if bundlesCount+1 > len(r.Bundles) {
				r.Bundles = append(r.Bundles, bundles)
			} else {
				r.Bundles[bundlesCount].AdminOwnerID = bundles.AdminOwnerID
				r.Bundles[bundlesCount].BundleID = bundles.BundleID
				r.Bundles[bundlesCount].CreatedAt = bundles.CreatedAt
				r.Bundles[bundlesCount].Description = bundles.Description
				r.Bundles[bundlesCount].Name = bundles.Name
				r.Bundles[bundlesCount].TotalNumGroups = bundles.TotalNumGroups
				r.Bundles[bundlesCount].TotalNumItems = bundles.TotalNumItems
				r.Bundles[bundlesCount].TotalNumResources = bundles.TotalNumResources
				r.Bundles[bundlesCount].UpdatedAt = bundles.UpdatedAt
			}
		}
		r.Next = types.StringPointerValue(resp.Next)
		r.Previous = types.StringPointerValue(resp.Previous)
		r.TotalCount = types.Int64PointerValue(resp.TotalCount)
	}

	return diags
}

func (r *PaginatedBundleListDataSourceModel) ToOperationsGetBundlesRequest(ctx context.Context) (*operations.GetBundlesRequest, diag.Diagnostics) {
	var diags diag.Diagnostics

	contains := new(string)
	if !r.Contains.IsUnknown() && !r.Contains.IsNull() {
		*contains = r.Contains.ValueString()
	} else {
		contains = nil
	}
	cursor := new(string)
	if !r.Cursor.IsUnknown() && !r.Cursor.IsNull() {
		*cursor = r.Cursor.ValueString()
	} else {
		cursor = nil
	}
	pageSize := new(int64)
	if !r.PageSize.IsUnknown() && !r.PageSize.IsNull() {
		*pageSize = r.PageSize.ValueInt64()
	} else {
		pageSize = nil
	}
	out := operations.GetBundlesRequest{
		Contains: contains,
		Cursor:   cursor,
		PageSize: pageSize,
	}

	return &out, diags
}
