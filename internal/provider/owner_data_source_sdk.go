// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/opalsecurity/terraform-provider-opal/internal/sdk/models/shared"
)

func (r *OwnerDataSourceModel) RefreshFromSharedOwner(resp *shared.Owner) {
	if resp != nil {
		r.AccessRequestEscalationPeriod = types.Int64PointerValue(resp.AccessRequestEscalationPeriod)
		r.Description = types.StringPointerValue(resp.Description)
		r.ID = types.StringValue(resp.ID)
		r.Name = types.StringPointerValue(resp.Name)
		r.ReviewerMessageChannelID = types.StringPointerValue(resp.ReviewerMessageChannelID)
		r.SourceGroupID = types.StringPointerValue(resp.SourceGroupID)
	}
}
