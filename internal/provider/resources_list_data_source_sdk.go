// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/opalsecurity/terraform-provider-opal/internal/provider/types"
	"github.com/opalsecurity/terraform-provider-opal/internal/sdk/models/shared"
)

func (r *ResourcesListDataSourceModel) RefreshFromSharedPaginatedResourcesList(ctx context.Context, resp *shared.PaginatedResourcesList) diag.Diagnostics {
	var diags diag.Diagnostics

	if resp != nil {
		r.Next = types.StringPointerValue(resp.Next)
		r.Previous = types.StringPointerValue(resp.Previous)
		r.Results = []tfTypes.Resource{}
		if len(r.Results) > len(resp.Results) {
			r.Results = r.Results[:len(resp.Results)]
		}
		for resultsCount, resultsItem := range resp.Results {
			var results tfTypes.Resource
			results.AdminOwnerID = types.StringPointerValue(resultsItem.AdminOwnerID)
			results.AppID = types.StringPointerValue(resultsItem.AppID)
			results.CustomRequestNotification = types.StringPointerValue(resultsItem.CustomRequestNotification)
			results.Description = types.StringPointerValue(resultsItem.Description)
			results.ID = types.StringValue(resultsItem.ID)
			results.Name = types.StringPointerValue(resultsItem.Name)
			results.ParentResourceID = types.StringPointerValue(resultsItem.ParentResourceID)
			if resultsItem.RemoteInfo == nil {
				results.RemoteInfo = nil
			} else {
				results.RemoteInfo = &tfTypes.ResourceRemoteInfo{}
				if resultsItem.RemoteInfo.AwsAccount == nil {
					results.RemoteInfo.AwsAccount = nil
				} else {
					results.RemoteInfo.AwsAccount = &tfTypes.AwsAccount{}
					results.RemoteInfo.AwsAccount.AccountID = types.StringValue(resultsItem.RemoteInfo.AwsAccount.AccountID)
				}
				if resultsItem.RemoteInfo.AwsEc2Instance == nil {
					results.RemoteInfo.AwsEc2Instance = nil
				} else {
					results.RemoteInfo.AwsEc2Instance = &tfTypes.AwsEc2Instance{}
					results.RemoteInfo.AwsEc2Instance.AccountID = types.StringPointerValue(resultsItem.RemoteInfo.AwsEc2Instance.AccountID)
					results.RemoteInfo.AwsEc2Instance.InstanceID = types.StringValue(resultsItem.RemoteInfo.AwsEc2Instance.InstanceID)
					results.RemoteInfo.AwsEc2Instance.Region = types.StringValue(resultsItem.RemoteInfo.AwsEc2Instance.Region)
				}
				if resultsItem.RemoteInfo.AwsEksCluster == nil {
					results.RemoteInfo.AwsEksCluster = nil
				} else {
					results.RemoteInfo.AwsEksCluster = &tfTypes.AwsEksCluster{}
					results.RemoteInfo.AwsEksCluster.AccountID = types.StringPointerValue(resultsItem.RemoteInfo.AwsEksCluster.AccountID)
					results.RemoteInfo.AwsEksCluster.Arn = types.StringValue(resultsItem.RemoteInfo.AwsEksCluster.Arn)
				}
				if resultsItem.RemoteInfo.AwsIamRole == nil {
					results.RemoteInfo.AwsIamRole = nil
				} else {
					results.RemoteInfo.AwsIamRole = &tfTypes.AwsEksCluster{}
					results.RemoteInfo.AwsIamRole.AccountID = types.StringPointerValue(resultsItem.RemoteInfo.AwsIamRole.AccountID)
					results.RemoteInfo.AwsIamRole.Arn = types.StringValue(resultsItem.RemoteInfo.AwsIamRole.Arn)
				}
				if resultsItem.RemoteInfo.AwsPermissionSet == nil {
					results.RemoteInfo.AwsPermissionSet = nil
				} else {
					results.RemoteInfo.AwsPermissionSet = &tfTypes.AwsPermissionSet{}
					results.RemoteInfo.AwsPermissionSet.AccountID = types.StringValue(resultsItem.RemoteInfo.AwsPermissionSet.AccountID)
					results.RemoteInfo.AwsPermissionSet.Arn = types.StringValue(resultsItem.RemoteInfo.AwsPermissionSet.Arn)
				}
				if resultsItem.RemoteInfo.AwsRdsInstance == nil {
					results.RemoteInfo.AwsRdsInstance = nil
				} else {
					results.RemoteInfo.AwsRdsInstance = &tfTypes.AwsRdsInstance{}
					results.RemoteInfo.AwsRdsInstance.AccountID = types.StringPointerValue(resultsItem.RemoteInfo.AwsRdsInstance.AccountID)
					results.RemoteInfo.AwsRdsInstance.InstanceID = types.StringValue(resultsItem.RemoteInfo.AwsRdsInstance.InstanceID)
					results.RemoteInfo.AwsRdsInstance.Region = types.StringValue(resultsItem.RemoteInfo.AwsRdsInstance.Region)
					results.RemoteInfo.AwsRdsInstance.ResourceID = types.StringValue(resultsItem.RemoteInfo.AwsRdsInstance.ResourceID)
				}
				if resultsItem.RemoteInfo.GcpBigQueryDataset == nil {
					results.RemoteInfo.GcpBigQueryDataset = nil
				} else {
					results.RemoteInfo.GcpBigQueryDataset = &tfTypes.GcpBigQueryDataset{}
					results.RemoteInfo.GcpBigQueryDataset.DatasetID = types.StringValue(resultsItem.RemoteInfo.GcpBigQueryDataset.DatasetID)
					results.RemoteInfo.GcpBigQueryDataset.ProjectID = types.StringValue(resultsItem.RemoteInfo.GcpBigQueryDataset.ProjectID)
				}
				if resultsItem.RemoteInfo.GcpBigQueryTable == nil {
					results.RemoteInfo.GcpBigQueryTable = nil
				} else {
					results.RemoteInfo.GcpBigQueryTable = &tfTypes.GcpBigQueryTable{}
					results.RemoteInfo.GcpBigQueryTable.DatasetID = types.StringValue(resultsItem.RemoteInfo.GcpBigQueryTable.DatasetID)
					results.RemoteInfo.GcpBigQueryTable.ProjectID = types.StringValue(resultsItem.RemoteInfo.GcpBigQueryTable.ProjectID)
					results.RemoteInfo.GcpBigQueryTable.TableID = types.StringValue(resultsItem.RemoteInfo.GcpBigQueryTable.TableID)
				}
				if resultsItem.RemoteInfo.GcpBucket == nil {
					results.RemoteInfo.GcpBucket = nil
				} else {
					results.RemoteInfo.GcpBucket = &tfTypes.GcpBucket{}
					results.RemoteInfo.GcpBucket.BucketID = types.StringValue(resultsItem.RemoteInfo.GcpBucket.BucketID)
				}
				if resultsItem.RemoteInfo.GcpComputeInstance == nil {
					results.RemoteInfo.GcpComputeInstance = nil
				} else {
					results.RemoteInfo.GcpComputeInstance = &tfTypes.GcpComputeInstance{}
					results.RemoteInfo.GcpComputeInstance.InstanceID = types.StringValue(resultsItem.RemoteInfo.GcpComputeInstance.InstanceID)
					results.RemoteInfo.GcpComputeInstance.ProjectID = types.StringValue(resultsItem.RemoteInfo.GcpComputeInstance.ProjectID)
					results.RemoteInfo.GcpComputeInstance.Zone = types.StringValue(resultsItem.RemoteInfo.GcpComputeInstance.Zone)
				}
				if resultsItem.RemoteInfo.GcpFolder == nil {
					results.RemoteInfo.GcpFolder = nil
				} else {
					results.RemoteInfo.GcpFolder = &tfTypes.GcpFolder{}
					results.RemoteInfo.GcpFolder.FolderID = types.StringValue(resultsItem.RemoteInfo.GcpFolder.FolderID)
				}
				if resultsItem.RemoteInfo.GcpGkeCluster == nil {
					results.RemoteInfo.GcpGkeCluster = nil
				} else {
					results.RemoteInfo.GcpGkeCluster = &tfTypes.GcpGkeCluster{}
					results.RemoteInfo.GcpGkeCluster.ClusterName = types.StringValue(resultsItem.RemoteInfo.GcpGkeCluster.ClusterName)
				}
				if resultsItem.RemoteInfo.GcpOrganization == nil {
					results.RemoteInfo.GcpOrganization = nil
				} else {
					results.RemoteInfo.GcpOrganization = &tfTypes.GcpOrganization{}
					results.RemoteInfo.GcpOrganization.OrganizationID = types.StringValue(resultsItem.RemoteInfo.GcpOrganization.OrganizationID)
				}
				if resultsItem.RemoteInfo.GcpProject == nil {
					results.RemoteInfo.GcpProject = nil
				} else {
					results.RemoteInfo.GcpProject = &tfTypes.GcpProject{}
					results.RemoteInfo.GcpProject.ProjectID = types.StringValue(resultsItem.RemoteInfo.GcpProject.ProjectID)
				}
				if resultsItem.RemoteInfo.GcpServiceAccount == nil {
					results.RemoteInfo.GcpServiceAccount = nil
				} else {
					results.RemoteInfo.GcpServiceAccount = &tfTypes.GcpServiceAccount{}
					results.RemoteInfo.GcpServiceAccount.Email = types.StringValue(resultsItem.RemoteInfo.GcpServiceAccount.Email)
					results.RemoteInfo.GcpServiceAccount.ProjectID = types.StringValue(resultsItem.RemoteInfo.GcpServiceAccount.ProjectID)
					results.RemoteInfo.GcpServiceAccount.ServiceAccountID = types.StringValue(resultsItem.RemoteInfo.GcpServiceAccount.ServiceAccountID)
				}
				if resultsItem.RemoteInfo.GcpSQLInstance == nil {
					results.RemoteInfo.GcpSQLInstance = nil
				} else {
					results.RemoteInfo.GcpSQLInstance = &tfTypes.GcpSQLInstance{}
					results.RemoteInfo.GcpSQLInstance.InstanceID = types.StringValue(resultsItem.RemoteInfo.GcpSQLInstance.InstanceID)
					results.RemoteInfo.GcpSQLInstance.ProjectID = types.StringValue(resultsItem.RemoteInfo.GcpSQLInstance.ProjectID)
				}
				if resultsItem.RemoteInfo.GithubRepo == nil {
					results.RemoteInfo.GithubRepo = nil
				} else {
					results.RemoteInfo.GithubRepo = &tfTypes.GithubRepo{}
					results.RemoteInfo.GithubRepo.RepoName = types.StringValue(resultsItem.RemoteInfo.GithubRepo.RepoName)
				}
				if resultsItem.RemoteInfo.GitlabProject == nil {
					results.RemoteInfo.GitlabProject = nil
				} else {
					results.RemoteInfo.GitlabProject = &tfTypes.GcpProject{}
					results.RemoteInfo.GitlabProject.ProjectID = types.StringValue(resultsItem.RemoteInfo.GitlabProject.ProjectID)
				}
				if resultsItem.RemoteInfo.OktaApp == nil {
					results.RemoteInfo.OktaApp = nil
				} else {
					results.RemoteInfo.OktaApp = &tfTypes.OktaApp{}
					results.RemoteInfo.OktaApp.AppID = types.StringValue(resultsItem.RemoteInfo.OktaApp.AppID)
				}
				if resultsItem.RemoteInfo.OktaCustomRole == nil {
					results.RemoteInfo.OktaCustomRole = nil
				} else {
					results.RemoteInfo.OktaCustomRole = &tfTypes.SnowflakeRole{}
					results.RemoteInfo.OktaCustomRole.RoleID = types.StringValue(resultsItem.RemoteInfo.OktaCustomRole.RoleID)
				}
				if resultsItem.RemoteInfo.OktaStandardRole == nil {
					results.RemoteInfo.OktaStandardRole = nil
				} else {
					results.RemoteInfo.OktaStandardRole = &tfTypes.OktaStandardRole{}
					results.RemoteInfo.OktaStandardRole.RoleType = types.StringValue(resultsItem.RemoteInfo.OktaStandardRole.RoleType)
				}
				if resultsItem.RemoteInfo.PagerdutyRole == nil {
					results.RemoteInfo.PagerdutyRole = nil
				} else {
					results.RemoteInfo.PagerdutyRole = &tfTypes.PagerdutyRole{}
					results.RemoteInfo.PagerdutyRole.RoleName = types.StringValue(resultsItem.RemoteInfo.PagerdutyRole.RoleName)
				}
				if resultsItem.RemoteInfo.SalesforcePermissionSet == nil {
					results.RemoteInfo.SalesforcePermissionSet = nil
				} else {
					results.RemoteInfo.SalesforcePermissionSet = &tfTypes.SalesforcePermissionSet{}
					results.RemoteInfo.SalesforcePermissionSet.PermissionSetID = types.StringValue(resultsItem.RemoteInfo.SalesforcePermissionSet.PermissionSetID)
				}
				if resultsItem.RemoteInfo.SalesforceProfile == nil {
					results.RemoteInfo.SalesforceProfile = nil
				} else {
					results.RemoteInfo.SalesforceProfile = &tfTypes.SalesforceProfile{}
					results.RemoteInfo.SalesforceProfile.ProfileID = types.StringValue(resultsItem.RemoteInfo.SalesforceProfile.ProfileID)
					results.RemoteInfo.SalesforceProfile.UserLicenseID = types.StringValue(resultsItem.RemoteInfo.SalesforceProfile.UserLicenseID)
				}
				if resultsItem.RemoteInfo.SalesforceRole == nil {
					results.RemoteInfo.SalesforceRole = nil
				} else {
					results.RemoteInfo.SalesforceRole = &tfTypes.SnowflakeRole{}
					results.RemoteInfo.SalesforceRole.RoleID = types.StringValue(resultsItem.RemoteInfo.SalesforceRole.RoleID)
				}
				if resultsItem.RemoteInfo.TeleportRole == nil {
					results.RemoteInfo.TeleportRole = nil
				} else {
					results.RemoteInfo.TeleportRole = &tfTypes.PagerdutyRole{}
					results.RemoteInfo.TeleportRole.RoleName = types.StringValue(resultsItem.RemoteInfo.TeleportRole.RoleName)
				}
			}
			results.RequestConfigurations = []tfTypes.RequestConfiguration{}
			for requestConfigurationsCount, requestConfigurationsItem := range resultsItem.RequestConfigurations {
				var requestConfigurations tfTypes.RequestConfiguration
				requestConfigurations.AllowRequests = types.BoolValue(requestConfigurationsItem.AllowRequests)
				requestConfigurations.AutoApproval = types.BoolValue(requestConfigurationsItem.AutoApproval)
				if requestConfigurationsItem.Condition == nil {
					requestConfigurations.Condition = nil
				} else {
					requestConfigurations.Condition = &tfTypes.Condition{}
					requestConfigurations.Condition.GroupIds = make([]types.String, 0, len(requestConfigurationsItem.Condition.GroupIds))
					for _, v := range requestConfigurationsItem.Condition.GroupIds {
						requestConfigurations.Condition.GroupIds = append(requestConfigurations.Condition.GroupIds, types.StringValue(v))
					}
					requestConfigurations.Condition.RoleRemoteIds = make([]types.String, 0, len(requestConfigurationsItem.Condition.RoleRemoteIds))
					for _, v := range requestConfigurationsItem.Condition.RoleRemoteIds {
						requestConfigurations.Condition.RoleRemoteIds = append(requestConfigurations.Condition.RoleRemoteIds, types.StringValue(v))
					}
				}
				requestConfigurations.MaxDuration = types.Int64PointerValue(requestConfigurationsItem.MaxDuration)
				requestConfigurations.Priority = types.Int64Value(requestConfigurationsItem.Priority)
				requestConfigurations.RecommendedDuration = types.Int64PointerValue(requestConfigurationsItem.RecommendedDuration)
				requestConfigurations.RequestTemplateID = types.StringPointerValue(requestConfigurationsItem.RequestTemplateID)
				requestConfigurations.RequireMfaToRequest = types.BoolValue(requestConfigurationsItem.RequireMfaToRequest)
				requestConfigurations.RequireSupportTicket = types.BoolValue(requestConfigurationsItem.RequireSupportTicket)
				requestConfigurations.ReviewerStages = []tfTypes.ReviewerStage{}
				for reviewerStagesCount, reviewerStagesItem := range requestConfigurationsItem.ReviewerStages {
					var reviewerStages tfTypes.ReviewerStage
					if reviewerStagesItem.Operator != nil {
						reviewerStages.Operator = types.StringValue(string(*reviewerStagesItem.Operator))
					} else {
						reviewerStages.Operator = types.StringNull()
					}
					reviewerStages.OwnerIds = make([]types.String, 0, len(reviewerStagesItem.OwnerIds))
					for _, v := range reviewerStagesItem.OwnerIds {
						reviewerStages.OwnerIds = append(reviewerStages.OwnerIds, types.StringValue(v))
					}
					reviewerStages.RequireAdminApproval = types.BoolPointerValue(reviewerStagesItem.RequireAdminApproval)
					reviewerStages.RequireManagerApproval = types.BoolValue(reviewerStagesItem.RequireManagerApproval)
					if reviewerStagesCount+1 > len(requestConfigurations.ReviewerStages) {
						requestConfigurations.ReviewerStages = append(requestConfigurations.ReviewerStages, reviewerStages)
					} else {
						requestConfigurations.ReviewerStages[reviewerStagesCount].Operator = reviewerStages.Operator
						requestConfigurations.ReviewerStages[reviewerStagesCount].OwnerIds = reviewerStages.OwnerIds
						requestConfigurations.ReviewerStages[reviewerStagesCount].RequireAdminApproval = reviewerStages.RequireAdminApproval
						requestConfigurations.ReviewerStages[reviewerStagesCount].RequireManagerApproval = reviewerStages.RequireManagerApproval
					}
				}
				if requestConfigurationsCount+1 > len(results.RequestConfigurations) {
					results.RequestConfigurations = append(results.RequestConfigurations, requestConfigurations)
				} else {
					results.RequestConfigurations[requestConfigurationsCount].AllowRequests = requestConfigurations.AllowRequests
					results.RequestConfigurations[requestConfigurationsCount].AutoApproval = requestConfigurations.AutoApproval
					results.RequestConfigurations[requestConfigurationsCount].Condition = requestConfigurations.Condition
					results.RequestConfigurations[requestConfigurationsCount].MaxDuration = requestConfigurations.MaxDuration
					results.RequestConfigurations[requestConfigurationsCount].Priority = requestConfigurations.Priority
					results.RequestConfigurations[requestConfigurationsCount].RecommendedDuration = requestConfigurations.RecommendedDuration
					results.RequestConfigurations[requestConfigurationsCount].RequestTemplateID = requestConfigurations.RequestTemplateID
					results.RequestConfigurations[requestConfigurationsCount].RequireMfaToRequest = requestConfigurations.RequireMfaToRequest
					results.RequestConfigurations[requestConfigurationsCount].RequireSupportTicket = requestConfigurations.RequireSupportTicket
					results.RequestConfigurations[requestConfigurationsCount].ReviewerStages = requestConfigurations.ReviewerStages
				}
			}
			results.RequireMfaToApprove = types.BoolPointerValue(resultsItem.RequireMfaToApprove)
			results.RequireMfaToConnect = types.BoolPointerValue(resultsItem.RequireMfaToConnect)
			if resultsItem.ResourceType != nil {
				results.ResourceType = types.StringValue(string(*resultsItem.ResourceType))
			} else {
				results.ResourceType = types.StringNull()
			}
			if resultsItem.RiskSensitivity != nil {
				results.RiskSensitivity = types.StringValue(string(*resultsItem.RiskSensitivity))
			} else {
				results.RiskSensitivity = types.StringNull()
			}
			if resultsItem.RiskSensitivityOverride != nil {
				results.RiskSensitivityOverride = types.StringValue(string(*resultsItem.RiskSensitivityOverride))
			} else {
				results.RiskSensitivityOverride = types.StringNull()
			}
			if resultsItem.TicketPropagation == nil {
				results.TicketPropagation = nil
			} else {
				results.TicketPropagation = &tfTypes.TicketPropagationConfiguration{}
				results.TicketPropagation.EnabledOnGrant = types.BoolValue(resultsItem.TicketPropagation.EnabledOnGrant)
				results.TicketPropagation.EnabledOnRevocation = types.BoolValue(resultsItem.TicketPropagation.EnabledOnRevocation)
				results.TicketPropagation.TicketProjectID = types.StringPointerValue(resultsItem.TicketPropagation.TicketProjectID)
				if resultsItem.TicketPropagation.TicketProvider != nil {
					results.TicketPropagation.TicketProvider = types.StringValue(string(*resultsItem.TicketPropagation.TicketProvider))
				} else {
					results.TicketPropagation.TicketProvider = types.StringNull()
				}
			}
			if resultsCount+1 > len(r.Results) {
				r.Results = append(r.Results, results)
			} else {
				r.Results[resultsCount].AdminOwnerID = results.AdminOwnerID
				r.Results[resultsCount].AppID = results.AppID
				r.Results[resultsCount].CustomRequestNotification = results.CustomRequestNotification
				r.Results[resultsCount].Description = results.Description
				r.Results[resultsCount].ID = results.ID
				r.Results[resultsCount].Name = results.Name
				r.Results[resultsCount].ParentResourceID = results.ParentResourceID
				r.Results[resultsCount].RemoteInfo = results.RemoteInfo
				r.Results[resultsCount].RequestConfigurations = results.RequestConfigurations
				r.Results[resultsCount].RequireMfaToApprove = results.RequireMfaToApprove
				r.Results[resultsCount].RequireMfaToConnect = results.RequireMfaToConnect
				r.Results[resultsCount].ResourceType = results.ResourceType
				r.Results[resultsCount].RiskSensitivity = results.RiskSensitivity
				r.Results[resultsCount].RiskSensitivityOverride = results.RiskSensitivityOverride
				r.Results[resultsCount].TicketPropagation = results.TicketPropagation
			}
		}
	}

	return diags
}
