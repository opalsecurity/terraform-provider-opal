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

func (r *UARSListDataSourceModel) RefreshFromSharedPaginatedUARsList(ctx context.Context, resp *shared.PaginatedUARsList) diag.Diagnostics {
	var diags diag.Diagnostics

	if resp != nil {
		r.Next = types.StringPointerValue(resp.Next)
		r.Previous = types.StringPointerValue(resp.Previous)
		r.Results = []tfTypes.Uar{}
		if len(r.Results) > len(resp.Results) {
			r.Results = r.Results[:len(resp.Results)]
		}
		for resultsCount, resultsItem := range resp.Results {
			var results tfTypes.Uar
			results.Deadline = types.StringValue(typeconvert.TimeToString(resultsItem.Deadline))
			results.Name = types.StringValue(resultsItem.Name)
			results.ReviewerAssignmentPolicy = types.StringValue(string(resultsItem.ReviewerAssignmentPolicy))
			results.SelfReviewAllowed = types.BoolValue(resultsItem.SelfReviewAllowed)
			results.SendReviewerAssignmentNotification = types.BoolValue(resultsItem.SendReviewerAssignmentNotification)
			results.TimeZone = types.StringValue(resultsItem.TimeZone)
			results.UarID = types.StringValue(resultsItem.UarID)
			if resultsItem.UarScope == nil {
				results.UarScope = nil
			} else {
				results.UarScope = &tfTypes.UARScope{}
				results.UarScope.Admins = make([]types.String, 0, len(resultsItem.UarScope.Admins))
				for _, v := range resultsItem.UarScope.Admins {
					results.UarScope.Admins = append(results.UarScope.Admins, types.StringValue(v))
				}
				results.UarScope.Apps = make([]types.String, 0, len(resultsItem.UarScope.Apps))
				for _, v := range resultsItem.UarScope.Apps {
					results.UarScope.Apps = append(results.UarScope.Apps, types.StringValue(v))
				}
				results.UarScope.Entities = make([]types.String, 0, len(resultsItem.UarScope.Entities))
				for _, v := range resultsItem.UarScope.Entities {
					results.UarScope.Entities = append(results.UarScope.Entities, types.StringValue(v))
				}
				if resultsItem.UarScope.FilterOperator != nil {
					results.UarScope.FilterOperator = types.StringValue(string(*resultsItem.UarScope.FilterOperator))
				} else {
					results.UarScope.FilterOperator = types.StringNull()
				}
				results.UarScope.GroupTypes = make([]types.String, 0, len(resultsItem.UarScope.GroupTypes))
				for _, v := range resultsItem.UarScope.GroupTypes {
					results.UarScope.GroupTypes = append(results.UarScope.GroupTypes, types.StringValue(string(v)))
				}
				if resultsItem.UarScope.GroupVisibility != nil {
					results.UarScope.GroupVisibility = types.StringValue(string(*resultsItem.UarScope.GroupVisibility))
				} else {
					results.UarScope.GroupVisibility = types.StringNull()
				}
				results.UarScope.IncludeGroupBindings = types.BoolPointerValue(resultsItem.UarScope.IncludeGroupBindings)
				results.UarScope.Names = make([]types.String, 0, len(resultsItem.UarScope.Names))
				for _, v := range resultsItem.UarScope.Names {
					results.UarScope.Names = append(results.UarScope.Names, types.StringValue(v))
				}
				results.UarScope.ResourceTypes = make([]types.String, 0, len(resultsItem.UarScope.ResourceTypes))
				for _, v := range resultsItem.UarScope.ResourceTypes {
					results.UarScope.ResourceTypes = append(results.UarScope.ResourceTypes, types.StringValue(string(v)))
				}
				results.UarScope.Tags = []tfTypes.TagFilter{}
				for tagsCount, tagsItem := range resultsItem.UarScope.Tags {
					var tags tfTypes.TagFilter
					tags.Key = types.StringValue(tagsItem.Key)
					tags.Value = types.StringPointerValue(tagsItem.Value)
					if tagsCount+1 > len(results.UarScope.Tags) {
						results.UarScope.Tags = append(results.UarScope.Tags, tags)
					} else {
						results.UarScope.Tags[tagsCount].Key = tags.Key
						results.UarScope.Tags[tagsCount].Value = tags.Value
					}
				}
				results.UarScope.Users = make([]types.String, 0, len(resultsItem.UarScope.Users))
				for _, v := range resultsItem.UarScope.Users {
					results.UarScope.Users = append(results.UarScope.Users, types.StringValue(v))
				}
			}
			if resultsCount+1 > len(r.Results) {
				r.Results = append(r.Results, results)
			} else {
				r.Results[resultsCount].Deadline = results.Deadline
				r.Results[resultsCount].Name = results.Name
				r.Results[resultsCount].ReviewerAssignmentPolicy = results.ReviewerAssignmentPolicy
				r.Results[resultsCount].SelfReviewAllowed = results.SelfReviewAllowed
				r.Results[resultsCount].SendReviewerAssignmentNotification = results.SendReviewerAssignmentNotification
				r.Results[resultsCount].TimeZone = results.TimeZone
				r.Results[resultsCount].UarID = results.UarID
				r.Results[resultsCount].UarScope = results.UarScope
			}
		}
	}

	return diags
}

func (r *UARSListDataSourceModel) ToOperationsGetUARsRequest(ctx context.Context) (*operations.GetUARsRequest, diag.Diagnostics) {
	var diags diag.Diagnostics

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
	out := operations.GetUARsRequest{
		Cursor:   cursor,
		PageSize: pageSize,
	}

	return &out, diags
}
