// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/opalsecurity/terraform-provider-opal/v3/internal/sdk/models/operations"
)

func (r *ResourceReviewersListDataSourceModel) RefreshFromString(ctx context.Context, resp []string) diag.Diagnostics {
	var diags diag.Diagnostics

	r.Data = make([]types.String, 0, len(resp))
	for _, v := range resp {
		r.Data = append(r.Data, types.StringValue(v))
	}

	return diags
}

func (r *ResourceReviewersListDataSourceModel) ToOperationsGetResourceReviewersRequest(ctx context.Context) (*operations.GetResourceReviewersRequest, diag.Diagnostics) {
	var diags diag.Diagnostics

	var resourceID string
	resourceID = r.ResourceID.ValueString()

	out := operations.GetResourceReviewersRequest{
		ResourceID: resourceID,
	}

	return &out, diags
}
