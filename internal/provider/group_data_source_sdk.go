// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/opalsecurity/terraform-provider-opal/internal/provider/types"
	"github.com/opalsecurity/terraform-provider-opal/internal/sdk/models/operations"
	"github.com/opalsecurity/terraform-provider-opal/internal/sdk/models/shared"
)

func (r *GroupDataSourceModel) RefreshFromSharedGroup(resp *shared.Group) {
	if resp != nil {
		r.AdminOwnerID = types.StringPointerValue(resp.AdminOwnerID)
		r.AppID = types.StringPointerValue(resp.AppID)
		r.AutoApproval = types.BoolPointerValue(resp.AutoApproval)
		r.Description = types.StringPointerValue(resp.Description)
		r.GroupBindingID = types.StringPointerValue(resp.GroupBindingID)
		if resp.GroupType != nil {
			r.GroupType = types.StringValue(string(*resp.GroupType))
		} else {
			r.GroupType = types.StringNull()
		}
		r.ID = types.StringPointerValue(resp.ID)
		r.IsRequestable = types.BoolPointerValue(resp.IsRequestable)
		r.MaxDuration = types.Int64PointerValue(resp.MaxDuration)
		r.Name = types.StringPointerValue(resp.Name)
		r.RecommendedDuration = types.Int64PointerValue(resp.RecommendedDuration)
		r.RemoteID = types.StringPointerValue(resp.RemoteID)
		if resp.RemoteInfo == nil {
			r.RemoteInfo = nil
		} else {
			r.RemoteInfo = &tfTypes.GroupRemoteInfo{}
			if resp.RemoteInfo.ActiveDirectoryGroup == nil {
				r.RemoteInfo.ActiveDirectoryGroup = nil
			} else {
				r.RemoteInfo.ActiveDirectoryGroup = &tfTypes.ActiveDirectoryGroup{}
				r.RemoteInfo.ActiveDirectoryGroup.GroupID = types.StringValue(resp.RemoteInfo.ActiveDirectoryGroup.GroupID)
			}
			if resp.RemoteInfo.AzureAdMicrosoft365Group == nil {
				r.RemoteInfo.AzureAdMicrosoft365Group = nil
			} else {
				r.RemoteInfo.AzureAdMicrosoft365Group = &tfTypes.ActiveDirectoryGroup{}
				r.RemoteInfo.AzureAdMicrosoft365Group.GroupID = types.StringValue(resp.RemoteInfo.AzureAdMicrosoft365Group.GroupID)
			}
			if resp.RemoteInfo.AzureAdSecurityGroup == nil {
				r.RemoteInfo.AzureAdSecurityGroup = nil
			} else {
				r.RemoteInfo.AzureAdSecurityGroup = &tfTypes.ActiveDirectoryGroup{}
				r.RemoteInfo.AzureAdSecurityGroup.GroupID = types.StringValue(resp.RemoteInfo.AzureAdSecurityGroup.GroupID)
			}
			if resp.RemoteInfo.DuoGroup == nil {
				r.RemoteInfo.DuoGroup = nil
			} else {
				r.RemoteInfo.DuoGroup = &tfTypes.ActiveDirectoryGroup{}
				r.RemoteInfo.DuoGroup.GroupID = types.StringValue(resp.RemoteInfo.DuoGroup.GroupID)
			}
			if resp.RemoteInfo.GithubTeam == nil {
				r.RemoteInfo.GithubTeam = nil
			} else {
				r.RemoteInfo.GithubTeam = &tfTypes.GithubTeam{}
				r.RemoteInfo.GithubTeam.TeamSlug = types.StringValue(resp.RemoteInfo.GithubTeam.TeamSlug)
			}
			if resp.RemoteInfo.GitlabGroup == nil {
				r.RemoteInfo.GitlabGroup = nil
			} else {
				r.RemoteInfo.GitlabGroup = &tfTypes.ActiveDirectoryGroup{}
				r.RemoteInfo.GitlabGroup.GroupID = types.StringValue(resp.RemoteInfo.GitlabGroup.GroupID)
			}
			if resp.RemoteInfo.GoogleGroup == nil {
				r.RemoteInfo.GoogleGroup = nil
			} else {
				r.RemoteInfo.GoogleGroup = &tfTypes.ActiveDirectoryGroup{}
				r.RemoteInfo.GoogleGroup.GroupID = types.StringValue(resp.RemoteInfo.GoogleGroup.GroupID)
			}
			if resp.RemoteInfo.LdapGroup == nil {
				r.RemoteInfo.LdapGroup = nil
			} else {
				r.RemoteInfo.LdapGroup = &tfTypes.ActiveDirectoryGroup{}
				r.RemoteInfo.LdapGroup.GroupID = types.StringValue(resp.RemoteInfo.LdapGroup.GroupID)
			}
			if resp.RemoteInfo.OktaGroup == nil {
				r.RemoteInfo.OktaGroup = nil
			} else {
				r.RemoteInfo.OktaGroup = &tfTypes.ActiveDirectoryGroup{}
				r.RemoteInfo.OktaGroup.GroupID = types.StringValue(resp.RemoteInfo.OktaGroup.GroupID)
			}
		}
		r.RemoteName = types.StringPointerValue(resp.RemoteName)
		if len(r.RequestConfigurations) > len(resp.RequestConfigurations) {
			r.RequestConfigurations = r.RequestConfigurations[:len(resp.RequestConfigurations)]
		}
		for requestConfigurationsCount, requestConfigurationsItem := range resp.RequestConfigurations {
			var requestConfigurations1 tfTypes.RequestConfiguration
			requestConfigurations1.AllowRequests = types.BoolValue(requestConfigurationsItem.AllowRequests)
			requestConfigurations1.AutoApproval = types.BoolValue(requestConfigurationsItem.AutoApproval)
			if requestConfigurationsItem.Condition == nil {
				requestConfigurations1.Condition = nil
			} else {
				requestConfigurations1.Condition = &tfTypes.Condition{}
				requestConfigurations1.Condition.GroupIds = []types.String{}
				for _, v := range requestConfigurationsItem.Condition.GroupIds {
					requestConfigurations1.Condition.GroupIds = append(requestConfigurations1.Condition.GroupIds, types.StringValue(v))
				}
				requestConfigurations1.Condition.RoleRemoteIds = []types.String{}
				for _, v := range requestConfigurationsItem.Condition.RoleRemoteIds {
					requestConfigurations1.Condition.RoleRemoteIds = append(requestConfigurations1.Condition.RoleRemoteIds, types.StringValue(v))
				}
			}
			requestConfigurations1.MaxDuration = types.Int64PointerValue(requestConfigurationsItem.MaxDuration)
			requestConfigurations1.Priority = types.Int64Value(requestConfigurationsItem.Priority)
			requestConfigurations1.RecommendedDuration = types.Int64PointerValue(requestConfigurationsItem.RecommendedDuration)
			requestConfigurations1.RequestTemplateID = types.StringPointerValue(requestConfigurationsItem.RequestTemplateID)
			requestConfigurations1.RequireMfaToRequest = types.BoolValue(requestConfigurationsItem.RequireMfaToRequest)
			requestConfigurations1.RequireSupportTicket = types.BoolValue(requestConfigurationsItem.RequireSupportTicket)
			for reviewerStagesCount, reviewerStagesItem := range requestConfigurationsItem.ReviewerStages {
				var reviewerStages1 tfTypes.ReviewerStage
				if reviewerStagesItem.Operator != nil {
					reviewerStages1.Operator = types.StringValue(string(*reviewerStagesItem.Operator))
				} else {
					reviewerStages1.Operator = types.StringNull()
				}
				reviewerStages1.OwnerIds = []types.String{}
				for _, v := range reviewerStagesItem.OwnerIds {
					reviewerStages1.OwnerIds = append(reviewerStages1.OwnerIds, types.StringValue(v))
				}
				reviewerStages1.RequireManagerApproval = types.BoolValue(reviewerStagesItem.RequireManagerApproval)
				if reviewerStagesCount+1 > len(requestConfigurations1.ReviewerStages) {
					requestConfigurations1.ReviewerStages = append(requestConfigurations1.ReviewerStages, reviewerStages1)
				} else {
					requestConfigurations1.ReviewerStages[reviewerStagesCount].Operator = reviewerStages1.Operator
					requestConfigurations1.ReviewerStages[reviewerStagesCount].OwnerIds = reviewerStages1.OwnerIds
					requestConfigurations1.ReviewerStages[reviewerStagesCount].RequireManagerApproval = reviewerStages1.RequireManagerApproval
				}
			}
			if requestConfigurationsCount+1 > len(r.RequestConfigurations) {
				r.RequestConfigurations = append(r.RequestConfigurations, requestConfigurations1)
			} else {
				r.RequestConfigurations[requestConfigurationsCount].AllowRequests = requestConfigurations1.AllowRequests
				r.RequestConfigurations[requestConfigurationsCount].AutoApproval = requestConfigurations1.AutoApproval
				r.RequestConfigurations[requestConfigurationsCount].Condition = requestConfigurations1.Condition
				r.RequestConfigurations[requestConfigurationsCount].MaxDuration = requestConfigurations1.MaxDuration
				r.RequestConfigurations[requestConfigurationsCount].Priority = requestConfigurations1.Priority
				r.RequestConfigurations[requestConfigurationsCount].RecommendedDuration = requestConfigurations1.RecommendedDuration
				r.RequestConfigurations[requestConfigurationsCount].RequestTemplateID = requestConfigurations1.RequestTemplateID
				r.RequestConfigurations[requestConfigurationsCount].RequireMfaToRequest = requestConfigurations1.RequireMfaToRequest
				r.RequestConfigurations[requestConfigurationsCount].RequireSupportTicket = requestConfigurations1.RequireSupportTicket
				r.RequestConfigurations[requestConfigurationsCount].ReviewerStages = requestConfigurations1.ReviewerStages
			}
		}
		r.RequestTemplateID = types.StringPointerValue(resp.RequestTemplateID)
		r.RequireMfaToApprove = types.BoolPointerValue(resp.RequireMfaToApprove)
	}
}

func (r *GroupDataSourceModel) RefreshFromOperationsGetGroupMessageChannelsResponseBody(resp *operations.GetGroupMessageChannelsResponseBody) {
	if resp != nil {
		if len(r.MessageChannels.Channels) > len(resp.Channels) {
			r.MessageChannels.Channels = r.MessageChannels.Channels[:len(resp.Channels)]
		}
		for channelsCount, channelsItem := range resp.Channels {
			var channels1 tfTypes.MessageChannel
			channels1.ID = types.StringPointerValue(channelsItem.ID)
			channels1.IsPrivate = types.BoolPointerValue(channelsItem.IsPrivate)
			channels1.Name = types.StringPointerValue(channelsItem.Name)
			channels1.RemoteID = types.StringPointerValue(channelsItem.RemoteID)
			if channelsItem.ThirdPartyProvider != nil {
				channels1.ThirdPartyProvider = types.StringValue(string(*channelsItem.ThirdPartyProvider))
			} else {
				channels1.ThirdPartyProvider = types.StringNull()
			}
			if channelsCount+1 > len(r.MessageChannels.Channels) {
				r.MessageChannels.Channels = append(r.MessageChannels.Channels, channels1)
			} else {
				r.MessageChannels.Channels[channelsCount].ID = channels1.ID
				r.MessageChannels.Channels[channelsCount].IsPrivate = channels1.IsPrivate
				r.MessageChannels.Channels[channelsCount].Name = channels1.Name
				r.MessageChannels.Channels[channelsCount].RemoteID = channels1.RemoteID
				r.MessageChannels.Channels[channelsCount].ThirdPartyProvider = channels1.ThirdPartyProvider
			}
		}
	}
}

func (r *GroupDataSourceModel) RefreshFromOperationsGetGroupVisibilityResponseBody(resp *operations.GetGroupVisibilityResponseBody) {
	if resp != nil {
		r.Visibility = types.StringValue(string(resp.Visibility))
		r.VisibilityGroupIds = []types.String{}
		for _, v := range resp.VisibilityGroupIds {
			r.VisibilityGroupIds = append(r.VisibilityGroupIds, types.StringValue(v))
		}
	}
}
