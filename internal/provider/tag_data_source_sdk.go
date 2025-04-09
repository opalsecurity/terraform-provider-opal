// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/opalsecurity/terraform-provider-opal/internal/provider/typeconvert"
	"github.com/opalsecurity/terraform-provider-opal/internal/sdk/models/shared"
)

func (r *TagDataSourceModel) RefreshFromSharedTag(ctx context.Context, resp *shared.Tag) diag.Diagnostics {
	var diags diag.Diagnostics

	if resp != nil {
		r.CreatedAt = types.StringPointerValue(typeconvert.TimePointerToStringPointer(resp.CreatedAt))
		r.ID = types.StringValue(resp.ID)
		r.Key = types.StringPointerValue(resp.Key)
		r.UpdatedAt = types.StringPointerValue(typeconvert.TimePointerToStringPointer(resp.UpdatedAt))
		r.UserCreatorID = types.StringPointerValue(resp.UserCreatorID)
		r.Value = types.StringPointerValue(resp.Value)
	}

	return diags
}
