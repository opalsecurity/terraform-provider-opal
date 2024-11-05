// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/opalsecurity/terraform-provider-opal/internal/provider/types"
	"github.com/opalsecurity/terraform-provider-opal/internal/sdk/models/operations"
	"github.com/opalsecurity/terraform-provider-opal/internal/sdk/models/shared"
)

func (r *GroupResourceModel) ToSharedCreateGroupInfo() *shared.CreateGroupInfo {
	var appID string
	appID = r.AppID.ValueString()

	customRequestNotification := new(string)
	if !r.CustomRequestNotification.IsUnknown() && !r.CustomRequestNotification.IsNull() {
		*customRequestNotification = r.CustomRequestNotification.ValueString()
	} else {
		customRequestNotification = nil
	}
	description := new(string)
	if !r.Description.IsUnknown() && !r.Description.IsNull() {
		*description = r.Description.ValueString()
	} else {
		description = nil
	}
	groupType := shared.GroupTypeEnum(r.GroupType.ValueString())
	var name string
	name = r.Name.ValueString()

	var remoteInfo *shared.GroupRemoteInfo
	if r.RemoteInfo != nil {
		var activeDirectoryGroup *shared.ActiveDirectoryGroup
		if r.RemoteInfo.ActiveDirectoryGroup != nil {
			var groupID string
			groupID = r.RemoteInfo.ActiveDirectoryGroup.GroupID.ValueString()

			activeDirectoryGroup = &shared.ActiveDirectoryGroup{
				GroupID: groupID,
			}
		}
		var azureAdMicrosoft365Group *shared.AzureAdMicrosoft365Group
		if r.RemoteInfo.AzureAdMicrosoft365Group != nil {
			var groupId1 string
			groupId1 = r.RemoteInfo.AzureAdMicrosoft365Group.GroupID.ValueString()

			azureAdMicrosoft365Group = &shared.AzureAdMicrosoft365Group{
				GroupID: groupId1,
			}
		}
		var azureAdSecurityGroup *shared.AzureAdSecurityGroup
		if r.RemoteInfo.AzureAdSecurityGroup != nil {
			var groupId2 string
			groupId2 = r.RemoteInfo.AzureAdSecurityGroup.GroupID.ValueString()

			azureAdSecurityGroup = &shared.AzureAdSecurityGroup{
				GroupID: groupId2,
			}
		}
		var duoGroup *shared.DuoGroup
		if r.RemoteInfo.DuoGroup != nil {
			var groupId3 string
			groupId3 = r.RemoteInfo.DuoGroup.GroupID.ValueString()

			duoGroup = &shared.DuoGroup{
				GroupID: groupId3,
			}
		}
		var githubTeam *shared.GithubTeam
		if r.RemoteInfo.GithubTeam != nil {
			var teamSlug string
			teamSlug = r.RemoteInfo.GithubTeam.TeamSlug.ValueString()

			githubTeam = &shared.GithubTeam{
				TeamSlug: teamSlug,
			}
		}
		var gitlabGroup *shared.GitlabGroup
		if r.RemoteInfo.GitlabGroup != nil {
			var groupId4 string
			groupId4 = r.RemoteInfo.GitlabGroup.GroupID.ValueString()

			gitlabGroup = &shared.GitlabGroup{
				GroupID: groupId4,
			}
		}
		var googleGroup *shared.GoogleGroup
		if r.RemoteInfo.GoogleGroup != nil {
			var groupId5 string
			groupId5 = r.RemoteInfo.GoogleGroup.GroupID.ValueString()

			googleGroup = &shared.GoogleGroup{
				GroupID: groupId5,
			}
		}
		var ldapGroup *shared.LdapGroup
		if r.RemoteInfo.LdapGroup != nil {
			var groupId6 string
			groupId6 = r.RemoteInfo.LdapGroup.GroupID.ValueString()

			ldapGroup = &shared.LdapGroup{
				GroupID: groupId6,
			}
		}
		var oktaGroup *shared.OktaGroup
		if r.RemoteInfo.OktaGroup != nil {
			var groupId7 string
			groupId7 = r.RemoteInfo.OktaGroup.GroupID.ValueString()

			oktaGroup = &shared.OktaGroup{
				GroupID: groupId7,
			}
		}
		remoteInfo = &shared.GroupRemoteInfo{
			ActiveDirectoryGroup:     activeDirectoryGroup,
			AzureAdMicrosoft365Group: azureAdMicrosoft365Group,
			AzureAdSecurityGroup:     azureAdSecurityGroup,
			DuoGroup:                 duoGroup,
			GithubTeam:               githubTeam,
			GitlabGroup:              gitlabGroup,
			GoogleGroup:              googleGroup,
			LdapGroup:                ldapGroup,
			OktaGroup:                oktaGroup,
		}
	}
	riskSensitivityOverride := new(shared.RiskSensitivityEnum)
	if !r.RiskSensitivityOverride.IsUnknown() && !r.RiskSensitivityOverride.IsNull() {
		*riskSensitivityOverride = shared.RiskSensitivityEnum(r.RiskSensitivityOverride.ValueString())
	} else {
		riskSensitivityOverride = nil
	}
	out := shared.CreateGroupInfo{
		AppID:                     appID,
		CustomRequestNotification: customRequestNotification,
		Description:               description,
		GroupType:                 groupType,
		Name:                      name,
		RemoteInfo:                remoteInfo,
		RiskSensitivityOverride:   riskSensitivityOverride,
	}
	return &out
}

func (r *GroupResourceModel) RefreshFromSharedGroup(resp *shared.Group) {
	if resp != nil {
		r.AdminOwnerID = types.StringPointerValue(resp.AdminOwnerID)
		r.AppID = types.StringPointerValue(resp.AppID)
		r.CustomRequestNotification = types.StringPointerValue(resp.CustomRequestNotification)
		r.Description = types.StringPointerValue(resp.Description)
		r.GroupBindingID = types.StringPointerValue(resp.GroupBindingID)
		r.GroupLeaderUserIds = []types.String{}
		for _, v := range resp.GroupLeaderUserIds {
			r.GroupLeaderUserIds = append(r.GroupLeaderUserIds, types.StringValue(v))
		}
		if resp.GroupType != nil {
			r.GroupType = types.StringValue(string(*resp.GroupType))
		} else {
			r.GroupType = types.StringNull()
		}
		r.ID = types.StringValue(resp.ID)
		r.Name = types.StringPointerValue(resp.Name)
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
		r.RequestConfigurations = []tfTypes.RequestConfiguration{}
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
			requestConfigurations1.ReviewerStages = []tfTypes.ReviewerStage{}
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
				reviewerStages1.RequireAdminApproval = types.BoolPointerValue(reviewerStagesItem.RequireAdminApproval)
				reviewerStages1.RequireManagerApproval = types.BoolValue(reviewerStagesItem.RequireManagerApproval)
				if reviewerStagesCount+1 > len(requestConfigurations1.ReviewerStages) {
					requestConfigurations1.ReviewerStages = append(requestConfigurations1.ReviewerStages, reviewerStages1)
				} else {
					requestConfigurations1.ReviewerStages[reviewerStagesCount].Operator = reviewerStages1.Operator
					requestConfigurations1.ReviewerStages[reviewerStagesCount].OwnerIds = reviewerStages1.OwnerIds
					requestConfigurations1.ReviewerStages[reviewerStagesCount].RequireAdminApproval = reviewerStages1.RequireAdminApproval
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
		r.RequireMfaToApprove = types.BoolPointerValue(resp.RequireMfaToApprove)
		if resp.RiskSensitivity != nil {
			r.RiskSensitivity = types.StringValue(string(*resp.RiskSensitivity))
		} else {
			r.RiskSensitivity = types.StringNull()
		}
		if resp.RiskSensitivityOverride != nil {
			r.RiskSensitivityOverride = types.StringValue(string(*resp.RiskSensitivityOverride))
		} else {
			r.RiskSensitivityOverride = types.StringNull()
		}
	}
}

func (r *GroupResourceModel) ToSharedUpdateGroupInfo() *shared.UpdateGroupInfo {
	adminOwnerID := new(string)
	if !r.AdminOwnerID.IsUnknown() && !r.AdminOwnerID.IsNull() {
		*adminOwnerID = r.AdminOwnerID.ValueString()
	} else {
		adminOwnerID = nil
	}
	customRequestNotification := new(string)
	if !r.CustomRequestNotification.IsUnknown() && !r.CustomRequestNotification.IsNull() {
		*customRequestNotification = r.CustomRequestNotification.ValueString()
	} else {
		customRequestNotification = nil
	}
	description := new(string)
	if !r.Description.IsUnknown() && !r.Description.IsNull() {
		*description = r.Description.ValueString()
	} else {
		description = nil
	}
	var groupLeaderUserIds []string = []string{}
	for _, groupLeaderUserIdsItem := range r.GroupLeaderUserIds {
		groupLeaderUserIds = append(groupLeaderUserIds, groupLeaderUserIdsItem.ValueString())
	}
	var id string
	id = r.ID.ValueString()

	name := new(string)
	if !r.Name.IsUnknown() && !r.Name.IsNull() {
		*name = r.Name.ValueString()
	} else {
		name = nil
	}
	var requestConfigurations []shared.RequestConfiguration = []shared.RequestConfiguration{}
	for _, requestConfigurationsItem := range r.RequestConfigurations {
		var allowRequests bool
		allowRequests = requestConfigurationsItem.AllowRequests.ValueBool()

		var autoApproval bool
		autoApproval = requestConfigurationsItem.AutoApproval.ValueBool()

		var condition *shared.Condition
		if requestConfigurationsItem.Condition != nil {
			var groupIds []string = []string{}
			for _, groupIdsItem := range requestConfigurationsItem.Condition.GroupIds {
				groupIds = append(groupIds, groupIdsItem.ValueString())
			}
			var roleRemoteIds []string = []string{}
			for _, roleRemoteIdsItem := range requestConfigurationsItem.Condition.RoleRemoteIds {
				roleRemoteIds = append(roleRemoteIds, roleRemoteIdsItem.ValueString())
			}
			condition = &shared.Condition{
				GroupIds:      groupIds,
				RoleRemoteIds: roleRemoteIds,
			}
		}
		maxDuration := new(int64)
		if !requestConfigurationsItem.MaxDuration.IsUnknown() && !requestConfigurationsItem.MaxDuration.IsNull() {
			*maxDuration = requestConfigurationsItem.MaxDuration.ValueInt64()
		} else {
			maxDuration = nil
		}
		var priority int64
		priority = requestConfigurationsItem.Priority.ValueInt64()

		recommendedDuration := new(int64)
		if !requestConfigurationsItem.RecommendedDuration.IsUnknown() && !requestConfigurationsItem.RecommendedDuration.IsNull() {
			*recommendedDuration = requestConfigurationsItem.RecommendedDuration.ValueInt64()
		} else {
			recommendedDuration = nil
		}
		requestTemplateID := new(string)
		if !requestConfigurationsItem.RequestTemplateID.IsUnknown() && !requestConfigurationsItem.RequestTemplateID.IsNull() {
			*requestTemplateID = requestConfigurationsItem.RequestTemplateID.ValueString()
		} else {
			requestTemplateID = nil
		}
		var requireMfaToRequest bool
		requireMfaToRequest = requestConfigurationsItem.RequireMfaToRequest.ValueBool()

		var requireSupportTicket bool
		requireSupportTicket = requestConfigurationsItem.RequireSupportTicket.ValueBool()

		var reviewerStages []shared.ReviewerStage = []shared.ReviewerStage{}
		for _, reviewerStagesItem := range requestConfigurationsItem.ReviewerStages {
			operator := new(shared.Operator)
			if !reviewerStagesItem.Operator.IsUnknown() && !reviewerStagesItem.Operator.IsNull() {
				*operator = shared.Operator(reviewerStagesItem.Operator.ValueString())
			} else {
				operator = nil
			}
			var ownerIds []string = []string{}
			for _, ownerIdsItem := range reviewerStagesItem.OwnerIds {
				ownerIds = append(ownerIds, ownerIdsItem.ValueString())
			}
			requireAdminApproval := new(bool)
			if !reviewerStagesItem.RequireAdminApproval.IsUnknown() && !reviewerStagesItem.RequireAdminApproval.IsNull() {
				*requireAdminApproval = reviewerStagesItem.RequireAdminApproval.ValueBool()
			} else {
				requireAdminApproval = nil
			}
			var requireManagerApproval bool
			requireManagerApproval = reviewerStagesItem.RequireManagerApproval.ValueBool()

			reviewerStages = append(reviewerStages, shared.ReviewerStage{
				Operator:               operator,
				OwnerIds:               ownerIds,
				RequireAdminApproval:   requireAdminApproval,
				RequireManagerApproval: requireManagerApproval,
			})
		}
		requestConfigurations = append(requestConfigurations, shared.RequestConfiguration{
			AllowRequests:        allowRequests,
			AutoApproval:         autoApproval,
			Condition:            condition,
			MaxDuration:          maxDuration,
			Priority:             priority,
			RecommendedDuration:  recommendedDuration,
			RequestTemplateID:    requestTemplateID,
			RequireMfaToRequest:  requireMfaToRequest,
			RequireSupportTicket: requireSupportTicket,
			ReviewerStages:       reviewerStages,
		})
	}
	requireMfaToApprove := new(bool)
	if !r.RequireMfaToApprove.IsUnknown() && !r.RequireMfaToApprove.IsNull() {
		*requireMfaToApprove = r.RequireMfaToApprove.ValueBool()
	} else {
		requireMfaToApprove = nil
	}
	riskSensitivityOverride := new(shared.RiskSensitivityEnum)
	if !r.RiskSensitivityOverride.IsUnknown() && !r.RiskSensitivityOverride.IsNull() {
		*riskSensitivityOverride = shared.RiskSensitivityEnum(r.RiskSensitivityOverride.ValueString())
	} else {
		riskSensitivityOverride = nil
	}
	out := shared.UpdateGroupInfo{
		AdminOwnerID:              adminOwnerID,
		CustomRequestNotification: customRequestNotification,
		Description:               description,
		GroupLeaderUserIds:        groupLeaderUserIds,
		ID:                        id,
		Name:                      name,
		RequestConfigurations:     requestConfigurations,
		RequireMfaToApprove:       requireMfaToApprove,
		RiskSensitivityOverride:   riskSensitivityOverride,
	}
	return &out
}

func (r *GroupResourceModel) RefreshFromSharedUpdateGroupInfo(resp *shared.UpdateGroupInfo) {
	r.AdminOwnerID = types.StringPointerValue(resp.AdminOwnerID)
	r.CustomRequestNotification = types.StringPointerValue(resp.CustomRequestNotification)
	r.Description = types.StringPointerValue(resp.Description)
	r.GroupLeaderUserIds = []types.String{}
	for _, v := range resp.GroupLeaderUserIds {
		r.GroupLeaderUserIds = append(r.GroupLeaderUserIds, types.StringValue(v))
	}
	r.ID = types.StringValue(resp.ID)
	r.Name = types.StringPointerValue(resp.Name)
	r.RequestConfigurations = []tfTypes.RequestConfiguration{}
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
		requestConfigurations1.ReviewerStages = []tfTypes.ReviewerStage{}
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
			reviewerStages1.RequireAdminApproval = types.BoolPointerValue(reviewerStagesItem.RequireAdminApproval)
			reviewerStages1.RequireManagerApproval = types.BoolValue(reviewerStagesItem.RequireManagerApproval)
			if reviewerStagesCount+1 > len(requestConfigurations1.ReviewerStages) {
				requestConfigurations1.ReviewerStages = append(requestConfigurations1.ReviewerStages, reviewerStages1)
			} else {
				requestConfigurations1.ReviewerStages[reviewerStagesCount].Operator = reviewerStages1.Operator
				requestConfigurations1.ReviewerStages[reviewerStagesCount].OwnerIds = reviewerStages1.OwnerIds
				requestConfigurations1.ReviewerStages[reviewerStagesCount].RequireAdminApproval = reviewerStages1.RequireAdminApproval
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
	r.RequireMfaToApprove = types.BoolPointerValue(resp.RequireMfaToApprove)
	if resp.RiskSensitivityOverride != nil {
		r.RiskSensitivityOverride = types.StringValue(string(*resp.RiskSensitivityOverride))
	} else {
		r.RiskSensitivityOverride = types.StringNull()
	}
}

func (r *GroupResourceModel) ToSharedMessageChannelIDList() *shared.MessageChannelIDList {
	var messageChannelIds []string = []string{}
	for _, messageChannelIdsItem := range r.MessageChannelIds {
		messageChannelIds = append(messageChannelIds, messageChannelIdsItem.ValueString())
	}
	out := shared.MessageChannelIDList{
		MessageChannelIds: messageChannelIds,
	}
	return &out
}

func (r *GroupResourceModel) ToSharedVisibilityInfo() *shared.VisibilityInfo {
	visibility := shared.VisibilityTypeEnum(r.Visibility.ValueString())
	var visibilityGroupIds []string = []string{}
	for _, visibilityGroupIdsItem := range r.VisibilityGroupIds {
		visibilityGroupIds = append(visibilityGroupIds, visibilityGroupIdsItem.ValueString())
	}
	out := shared.VisibilityInfo{
		Visibility:         visibility,
		VisibilityGroupIds: visibilityGroupIds,
	}
	return &out
}

func (r *GroupResourceModel) RefreshFromOperationsGetGroupMessageChannelsResponseBody(resp *operations.GetGroupMessageChannelsResponseBody) {
	if resp != nil {
		r.MessageChannels.Channels = []tfTypes.MessageChannel{}
		if len(r.MessageChannels.Channels) > len(resp.Channels) {
			r.MessageChannels.Channels = r.MessageChannels.Channels[:len(resp.Channels)]
		}
		for channelsCount, channelsItem := range resp.Channels {
			var channels1 tfTypes.MessageChannel
			channels1.ID = types.StringValue(channelsItem.ID)
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

func (r *GroupResourceModel) RefreshFromOperationsGetGroupVisibilityResponseBody(resp *operations.GetGroupVisibilityResponseBody) {
	if resp != nil {
		r.Visibility = types.StringValue(string(resp.Visibility))
		r.VisibilityGroupIds = []types.String{}
		for _, v := range resp.VisibilityGroupIds {
			r.VisibilityGroupIds = append(r.VisibilityGroupIds, types.StringValue(v))
		}
	}
}
