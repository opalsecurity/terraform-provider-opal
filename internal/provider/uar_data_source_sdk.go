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

func (r *UarDataSourceModel) RefreshFromSharedUar(ctx context.Context, resp *shared.Uar) diag.Diagnostics {
	var diags diag.Diagnostics

	if resp != nil {
		r.Deadline = types.StringValue(typeconvert.TimeToString(resp.Deadline))
		r.Name = types.StringValue(resp.Name)
		r.ReviewerAssignmentPolicy = types.StringValue(string(resp.ReviewerAssignmentPolicy))
		r.SelfReviewAllowed = types.BoolValue(resp.SelfReviewAllowed)
		r.SendReviewerAssignmentNotification = types.BoolValue(resp.SendReviewerAssignmentNotification)
		r.TimeZone = types.StringValue(resp.TimeZone)
		r.UarID = types.StringValue(resp.UarID)
		if resp.UarScope == nil {
			r.UarScope = nil
		} else {
			r.UarScope = &tfTypes.UARScope{}
			r.UarScope.Admins = make([]types.String, 0, len(resp.UarScope.Admins))
			for _, v := range resp.UarScope.Admins {
				r.UarScope.Admins = append(r.UarScope.Admins, types.StringValue(v))
			}
			r.UarScope.Apps = make([]types.String, 0, len(resp.UarScope.Apps))
			for _, v := range resp.UarScope.Apps {
				r.UarScope.Apps = append(r.UarScope.Apps, types.StringValue(v))
			}
			r.UarScope.Entities = make([]types.String, 0, len(resp.UarScope.Entities))
			for _, v := range resp.UarScope.Entities {
				r.UarScope.Entities = append(r.UarScope.Entities, types.StringValue(v))
			}
			if resp.UarScope.FilterOperator != nil {
				r.UarScope.FilterOperator = types.StringValue(string(*resp.UarScope.FilterOperator))
			} else {
				r.UarScope.FilterOperator = types.StringNull()
			}
			r.UarScope.GroupTypes = make([]types.String, 0, len(resp.UarScope.GroupTypes))
			for _, v := range resp.UarScope.GroupTypes {
				r.UarScope.GroupTypes = append(r.UarScope.GroupTypes, types.StringValue(string(v)))
			}
			if resp.UarScope.GroupVisibility != nil {
				r.UarScope.GroupVisibility = types.StringValue(string(*resp.UarScope.GroupVisibility))
			} else {
				r.UarScope.GroupVisibility = types.StringNull()
			}
			r.UarScope.IncludeGroupBindings = types.BoolPointerValue(resp.UarScope.IncludeGroupBindings)
			r.UarScope.Names = make([]types.String, 0, len(resp.UarScope.Names))
			for _, v := range resp.UarScope.Names {
				r.UarScope.Names = append(r.UarScope.Names, types.StringValue(v))
			}
			r.UarScope.ResourceTypes = make([]types.String, 0, len(resp.UarScope.ResourceTypes))
			for _, v := range resp.UarScope.ResourceTypes {
				r.UarScope.ResourceTypes = append(r.UarScope.ResourceTypes, types.StringValue(string(v)))
			}
			r.UarScope.Tags = []tfTypes.TagFilter{}
			if len(r.UarScope.Tags) > len(resp.UarScope.Tags) {
				r.UarScope.Tags = r.UarScope.Tags[:len(resp.UarScope.Tags)]
			}
			for tagsCount, tagsItem := range resp.UarScope.Tags {
				var tags tfTypes.TagFilter
				tags.Key = types.StringValue(tagsItem.Key)
				tags.Value = types.StringPointerValue(tagsItem.Value)
				if tagsCount+1 > len(r.UarScope.Tags) {
					r.UarScope.Tags = append(r.UarScope.Tags, tags)
				} else {
					r.UarScope.Tags[tagsCount].Key = tags.Key
					r.UarScope.Tags[tagsCount].Value = tags.Value
				}
			}
			r.UarScope.Users = make([]types.String, 0, len(resp.UarScope.Users))
			for _, v := range resp.UarScope.Users {
				r.UarScope.Users = append(r.UarScope.Users, types.StringValue(v))
			}
		}
	}

	return diags
}

func (r *UarDataSourceModel) ToOperationsGetUARIDRequest(ctx context.Context) (*operations.GetUARIDRequest, diag.Diagnostics) {
	var diags diag.Diagnostics

	var uarID string
	uarID = r.UarID.ValueString()

	out := operations.GetUARIDRequest{
		UarID: uarID,
	}

	return &out, diags
}
