// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/opalsecurity/terraform-provider-opal/internal/sdk/models/shared"
	"time"
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

func (r *TagResourceModel) RefreshFromSharedTag(resp *shared.Tag) {
	if resp != nil {
		if resp.CreatedAt != nil {
			r.CreatedAt = types.StringValue(resp.CreatedAt.Format(time.RFC3339Nano))
		} else {
			r.CreatedAt = types.StringNull()
		}
		r.ID = types.StringValue(resp.ID)
		r.Key = types.StringPointerValue(resp.Key)
		if resp.UpdatedAt != nil {
			r.UpdatedAt = types.StringValue(resp.UpdatedAt.Format(time.RFC3339Nano))
		} else {
			r.UpdatedAt = types.StringNull()
		}
		r.UserCreatorID = types.StringPointerValue(resp.UserCreatorID)
		r.Value = types.StringPointerValue(resp.Value)
	}
}
