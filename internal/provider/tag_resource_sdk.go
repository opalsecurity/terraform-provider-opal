// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/opalsecurity/terraform-provider-opal/internal/provider/typeconvert"
	"github.com/opalsecurity/terraform-provider-opal/internal/sdk/models/shared"
)

func (r *TagResourceModel) ToSharedCreateTagInfo() *shared.CreateTagInfo {
	var key string
	key = r.Key.ValueString()

	value := new(string)
	if !r.Value.IsUnknown() && !r.Value.IsNull() {
		*value = r.Value.ValueString()
	} else {
		value = nil
	}
	out := shared.CreateTagInfo{
		Key:   key,
		Value: value,
	}
	return &out
}

func (r *TagResourceModel) RefreshFromSharedTag(ctx context.Context, resp *shared.Tag) diag.Diagnostics {
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
