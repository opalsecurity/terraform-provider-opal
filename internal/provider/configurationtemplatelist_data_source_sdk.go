// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/opal-dev/terraform-provider-opal/internal/provider/types"
	"github.com/opal-dev/terraform-provider-opal/internal/sdk/models/shared"
)

func (r *ConfigurationTemplateListDataSourceModel) RefreshFromSharedPaginatedConfigurationTemplateList(resp *shared.PaginatedConfigurationTemplateList) {
	if resp != nil {
		if len(r.Results) > len(resp.Results) {
			r.Results = r.Results[:len(resp.Results)]
		}
		for resultsCount, resultsItem := range resp.Results {
			var results1 tfTypes.ConfigurationTemplate
			results1.AdminOwnerID = types.StringPointerValue(resultsItem.AdminOwnerID)
			results1.BreakGlassUserIds = []types.String{}
			for _, v := range resultsItem.BreakGlassUserIds {
				results1.BreakGlassUserIds = append(results1.BreakGlassUserIds, types.StringValue(v))
			}
			results1.ConfigurationTemplateID = types.StringPointerValue(resultsItem.ConfigurationTemplateID)
			results1.LinkedAuditMessageChannelIds = []types.String{}
			for _, v := range resultsItem.LinkedAuditMessageChannelIds {
				results1.LinkedAuditMessageChannelIds = append(results1.LinkedAuditMessageChannelIds, types.StringValue(v))
			}
			results1.MemberOncallScheduleIds = []types.String{}
			for _, v := range resultsItem.MemberOncallScheduleIds {
				results1.MemberOncallScheduleIds = append(results1.MemberOncallScheduleIds, types.StringValue(v))
			}
			results1.Name = types.StringPointerValue(resultsItem.Name)
			results1.RequestConfigurationID = types.StringPointerValue(resultsItem.RequestConfigurationID)
			results1.RequireMfaToApprove = types.BoolPointerValue(resultsItem.RequireMfaToApprove)
			results1.RequireMfaToConnect = types.BoolPointerValue(resultsItem.RequireMfaToConnect)
			if resultsItem.Visibility == nil {
				results1.Visibility = nil
			} else {
				results1.Visibility = &tfTypes.VisibilityInfo{}
				results1.Visibility.Visibility = types.StringValue(string(resultsItem.Visibility.Visibility))
				results1.Visibility.VisibilityGroupIds = []types.String{}
				for _, v := range resultsItem.Visibility.VisibilityGroupIds {
					results1.Visibility.VisibilityGroupIds = append(results1.Visibility.VisibilityGroupIds, types.StringValue(v))
				}
			}
			if resultsCount+1 > len(r.Results) {
				r.Results = append(r.Results, results1)
			} else {
				r.Results[resultsCount].AdminOwnerID = results1.AdminOwnerID
				r.Results[resultsCount].BreakGlassUserIds = results1.BreakGlassUserIds
				r.Results[resultsCount].ConfigurationTemplateID = results1.ConfigurationTemplateID
				r.Results[resultsCount].LinkedAuditMessageChannelIds = results1.LinkedAuditMessageChannelIds
				r.Results[resultsCount].MemberOncallScheduleIds = results1.MemberOncallScheduleIds
				r.Results[resultsCount].Name = results1.Name
				r.Results[resultsCount].RequestConfigurationID = results1.RequestConfigurationID
				r.Results[resultsCount].RequireMfaToApprove = results1.RequireMfaToApprove
				r.Results[resultsCount].RequireMfaToConnect = results1.RequireMfaToConnect
				r.Results[resultsCount].Visibility = results1.Visibility
			}
		}
	}
}
