// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/opalsecurity/terraform-provider-opal/internal/provider/typeconvert"
	tfTypes "github.com/opalsecurity/terraform-provider-opal/internal/provider/types"
	"github.com/opalsecurity/terraform-provider-opal/internal/sdk/models/shared"
)

func (r *ResourceTagsDataSourceModel) RefreshFromSharedTagsList(ctx context.Context, resp *shared.TagsList) diag.Diagnostics {
	var diags diag.Diagnostics

	if resp != nil {
		r.Tags = []tfTypes.Tag{}
		if len(r.Tags) > len(resp.Tags) {
			r.Tags = r.Tags[:len(resp.Tags)]
		}
		for tagsCount, tagsItem := range resp.Tags {
			var tags tfTypes.Tag
			tags.CreatedAt = types.StringPointerValue(typeconvert.TimePointerToStringPointer(tagsItem.CreatedAt))
			tags.ID = types.StringValue(tagsItem.ID)
			tags.Key = types.StringPointerValue(tagsItem.Key)
			tags.UpdatedAt = types.StringPointerValue(typeconvert.TimePointerToStringPointer(tagsItem.UpdatedAt))
			tags.UserCreatorID = types.StringPointerValue(tagsItem.UserCreatorID)
			tags.Value = types.StringPointerValue(tagsItem.Value)
			if tagsCount+1 > len(r.Tags) {
				r.Tags = append(r.Tags, tags)
			} else {
				r.Tags[tagsCount].CreatedAt = tags.CreatedAt
				r.Tags[tagsCount].ID = tags.ID
				r.Tags[tagsCount].Key = tags.Key
				r.Tags[tagsCount].UpdatedAt = tags.UpdatedAt
				r.Tags[tagsCount].UserCreatorID = tags.UserCreatorID
				r.Tags[tagsCount].Value = tags.Value
			}
		}
	}

	return diags
}
