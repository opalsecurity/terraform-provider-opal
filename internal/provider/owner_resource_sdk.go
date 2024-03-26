// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/opal-dev/terraform-provider-opal/internal/sdk/models/shared"
)

func (r *OwnerResourceModel) ToSharedCreateOwnerInfo() *shared.CreateOwnerInfo {
	accessRequestEscalationPeriod := new(int64)
	if !r.AccessRequestEscalationPeriod.IsUnknown() && !r.AccessRequestEscalationPeriod.IsNull() {
		*accessRequestEscalationPeriod = r.AccessRequestEscalationPeriod.ValueInt64()
	} else {
		accessRequestEscalationPeriod = nil
	}
	description := new(string)
	if !r.Description.IsUnknown() && !r.Description.IsNull() {
		*description = r.Description.ValueString()
	} else {
		description = nil
	}
	name := r.Name.ValueString()
	reviewerMessageChannelID := new(string)
	if !r.ReviewerMessageChannelID.IsUnknown() && !r.ReviewerMessageChannelID.IsNull() {
		*reviewerMessageChannelID = r.ReviewerMessageChannelID.ValueString()
	} else {
		reviewerMessageChannelID = nil
	}
	sourceGroupID := new(string)
	if !r.SourceGroupID.IsUnknown() && !r.SourceGroupID.IsNull() {
		*sourceGroupID = r.SourceGroupID.ValueString()
	} else {
		sourceGroupID = nil
	}
	var userIds []string = nil
	for _, userIdsItem := range r.UserIds {
		userIds = append(userIds, userIdsItem.ValueString())
	}
	out := shared.CreateOwnerInfo{
		AccessRequestEscalationPeriod: accessRequestEscalationPeriod,
		Description:                   description,
		Name:                          name,
		ReviewerMessageChannelID:      reviewerMessageChannelID,
		SourceGroupID:                 sourceGroupID,
		UserIds:                       userIds,
	}
	return &out
}

func (r *OwnerResourceModel) RefreshFromSharedOwner(resp *shared.Owner) {
	if resp != nil {
		r.AccessRequestEscalationPeriod = types.Int64PointerValue(resp.AccessRequestEscalationPeriod)
		r.Description = types.StringPointerValue(resp.Description)
		r.ID = types.StringPointerValue(resp.ID)
		r.Name = types.StringPointerValue(resp.Name)
		r.ReviewerMessageChannelID = types.StringPointerValue(resp.ReviewerMessageChannelID)
		r.SourceGroupID = types.StringPointerValue(resp.SourceGroupID)
	}
}

func (r *OwnerResourceModel) ToSharedUpdateOwnerInfo() *shared.UpdateOwnerInfo {
	accessRequestEscalationPeriod := new(int64)
	if !r.AccessRequestEscalationPeriod.IsUnknown() && !r.AccessRequestEscalationPeriod.IsNull() {
		*accessRequestEscalationPeriod = r.AccessRequestEscalationPeriod.ValueInt64()
	} else {
		accessRequestEscalationPeriod = nil
	}
	description := new(string)
	if !r.Description.IsUnknown() && !r.Description.IsNull() {
		*description = r.Description.ValueString()
	} else {
		description = nil
	}
	id := new(string)
	if !r.ID.IsUnknown() && !r.ID.IsNull() {
		*id = r.ID.ValueString()
	} else {
		id = nil
	}
	name := new(string)
	if !r.Name.IsUnknown() && !r.Name.IsNull() {
		*name = r.Name.ValueString()
	} else {
		name = nil
	}
	reviewerMessageChannelID := new(string)
	if !r.ReviewerMessageChannelID.IsUnknown() && !r.ReviewerMessageChannelID.IsNull() {
		*reviewerMessageChannelID = r.ReviewerMessageChannelID.ValueString()
	} else {
		reviewerMessageChannelID = nil
	}
	sourceGroupID := new(string)
	if !r.SourceGroupID.IsUnknown() && !r.SourceGroupID.IsNull() {
		*sourceGroupID = r.SourceGroupID.ValueString()
	} else {
		sourceGroupID = nil
	}
	out := shared.UpdateOwnerInfo{
		AccessRequestEscalationPeriod: accessRequestEscalationPeriod,
		Description:                   description,
		ID:                            id,
		Name:                          name,
		ReviewerMessageChannelID:      reviewerMessageChannelID,
		SourceGroupID:                 sourceGroupID,
	}
	return &out
}

func (r *OwnerResourceModel) RefreshFromSharedUpdateOwnerInfo(resp shared.UpdateOwnerInfo) {
	r.AccessRequestEscalationPeriod = types.Int64PointerValue(resp.AccessRequestEscalationPeriod)
	r.Description = types.StringPointerValue(resp.Description)
	r.ID = types.StringPointerValue(resp.ID)
	r.Name = types.StringPointerValue(resp.Name)
	r.ReviewerMessageChannelID = types.StringPointerValue(resp.ReviewerMessageChannelID)
	r.SourceGroupID = types.StringPointerValue(resp.SourceGroupID)
}
