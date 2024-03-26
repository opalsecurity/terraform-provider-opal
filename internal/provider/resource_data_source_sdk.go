// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/opal-dev/terraform-provider-opal/internal/provider/types"
	"github.com/opal-dev/terraform-provider-opal/internal/sdk/models/shared"
)

func (r *ResourceDataSourceModel) RefreshFromSharedResource(resp *shared.Resource) {
	if resp != nil {
		r.AdminOwnerID = types.StringPointerValue(resp.AdminOwnerID)
		r.AppID = types.StringPointerValue(resp.AppID)
		r.AutoApproval = types.BoolPointerValue(resp.AutoApproval)
		r.Description = types.StringPointerValue(resp.Description)
		r.ID = types.StringPointerValue(resp.ID)
		r.IsRequestable = types.BoolPointerValue(resp.IsRequestable)
		r.MaxDuration = types.Int64PointerValue(resp.MaxDuration)
		r.Name = types.StringPointerValue(resp.Name)
		r.ParentResourceID = types.StringPointerValue(resp.ParentResourceID)
		r.RecommendedDuration = types.Int64PointerValue(resp.RecommendedDuration)
		if resp.RemoteInfo == nil {
			r.RemoteInfo = nil
		} else {
			r.RemoteInfo = &tfTypes.ResourceRemoteInfo{}
			if resp.RemoteInfo.AwsAccount == nil {
				r.RemoteInfo.AwsAccount = nil
			} else {
				r.RemoteInfo.AwsAccount = &tfTypes.AwsAccount{}
				r.RemoteInfo.AwsAccount.AccountID = types.StringValue(resp.RemoteInfo.AwsAccount.AccountID)
			}
			if resp.RemoteInfo.AwsEc2Instance == nil {
				r.RemoteInfo.AwsEc2Instance = nil
			} else {
				r.RemoteInfo.AwsEc2Instance = &tfTypes.AwsEc2Instance{}
				r.RemoteInfo.AwsEc2Instance.AccountID = types.StringPointerValue(resp.RemoteInfo.AwsEc2Instance.AccountID)
				r.RemoteInfo.AwsEc2Instance.InstanceID = types.StringValue(resp.RemoteInfo.AwsEc2Instance.InstanceID)
				r.RemoteInfo.AwsEc2Instance.Region = types.StringValue(resp.RemoteInfo.AwsEc2Instance.Region)
			}
			if resp.RemoteInfo.AwsEksCluster == nil {
				r.RemoteInfo.AwsEksCluster = nil
			} else {
				r.RemoteInfo.AwsEksCluster = &tfTypes.AwsEksCluster{}
				r.RemoteInfo.AwsEksCluster.AccountID = types.StringPointerValue(resp.RemoteInfo.AwsEksCluster.AccountID)
				r.RemoteInfo.AwsEksCluster.Arn = types.StringValue(resp.RemoteInfo.AwsEksCluster.Arn)
			}
			if resp.RemoteInfo.AwsIamRole == nil {
				r.RemoteInfo.AwsIamRole = nil
			} else {
				r.RemoteInfo.AwsIamRole = &tfTypes.AwsEksCluster{}
				r.RemoteInfo.AwsIamRole.AccountID = types.StringPointerValue(resp.RemoteInfo.AwsIamRole.AccountID)
				r.RemoteInfo.AwsIamRole.Arn = types.StringValue(resp.RemoteInfo.AwsIamRole.Arn)
			}
			if resp.RemoteInfo.AwsPermissionSet == nil {
				r.RemoteInfo.AwsPermissionSet = nil
			} else {
				r.RemoteInfo.AwsPermissionSet = &tfTypes.AwsPermissionSet{}
				r.RemoteInfo.AwsPermissionSet.AccountID = types.StringValue(resp.RemoteInfo.AwsPermissionSet.AccountID)
				r.RemoteInfo.AwsPermissionSet.Arn = types.StringValue(resp.RemoteInfo.AwsPermissionSet.Arn)
			}
			if resp.RemoteInfo.AwsRdsInstance == nil {
				r.RemoteInfo.AwsRdsInstance = nil
			} else {
				r.RemoteInfo.AwsRdsInstance = &tfTypes.AwsRdsInstance{}
				r.RemoteInfo.AwsRdsInstance.AccountID = types.StringPointerValue(resp.RemoteInfo.AwsRdsInstance.AccountID)
				r.RemoteInfo.AwsRdsInstance.InstanceID = types.StringValue(resp.RemoteInfo.AwsRdsInstance.InstanceID)
				r.RemoteInfo.AwsRdsInstance.Region = types.StringValue(resp.RemoteInfo.AwsRdsInstance.Region)
				r.RemoteInfo.AwsRdsInstance.ResourceID = types.StringValue(resp.RemoteInfo.AwsRdsInstance.ResourceID)
			}
			if resp.RemoteInfo.GcpBigQueryDataset == nil {
				r.RemoteInfo.GcpBigQueryDataset = nil
			} else {
				r.RemoteInfo.GcpBigQueryDataset = &tfTypes.GcpBigQueryDataset{}
				r.RemoteInfo.GcpBigQueryDataset.DatasetID = types.StringValue(resp.RemoteInfo.GcpBigQueryDataset.DatasetID)
				r.RemoteInfo.GcpBigQueryDataset.ProjectID = types.StringValue(resp.RemoteInfo.GcpBigQueryDataset.ProjectID)
			}
			if resp.RemoteInfo.GcpBigQueryTable == nil {
				r.RemoteInfo.GcpBigQueryTable = nil
			} else {
				r.RemoteInfo.GcpBigQueryTable = &tfTypes.GcpBigQueryTable{}
				r.RemoteInfo.GcpBigQueryTable.DatasetID = types.StringValue(resp.RemoteInfo.GcpBigQueryTable.DatasetID)
				r.RemoteInfo.GcpBigQueryTable.ProjectID = types.StringValue(resp.RemoteInfo.GcpBigQueryTable.ProjectID)
				r.RemoteInfo.GcpBigQueryTable.TableID = types.StringValue(resp.RemoteInfo.GcpBigQueryTable.TableID)
			}
			if resp.RemoteInfo.GcpBucket == nil {
				r.RemoteInfo.GcpBucket = nil
			} else {
				r.RemoteInfo.GcpBucket = &tfTypes.GcpBucket{}
				r.RemoteInfo.GcpBucket.BucketID = types.StringValue(resp.RemoteInfo.GcpBucket.BucketID)
			}
			if resp.RemoteInfo.GcpComputeInstance == nil {
				r.RemoteInfo.GcpComputeInstance = nil
			} else {
				r.RemoteInfo.GcpComputeInstance = &tfTypes.GcpComputeInstance{}
				r.RemoteInfo.GcpComputeInstance.InstanceID = types.StringValue(resp.RemoteInfo.GcpComputeInstance.InstanceID)
				r.RemoteInfo.GcpComputeInstance.ProjectID = types.StringValue(resp.RemoteInfo.GcpComputeInstance.ProjectID)
				r.RemoteInfo.GcpComputeInstance.Zone = types.StringValue(resp.RemoteInfo.GcpComputeInstance.Zone)
			}
			if resp.RemoteInfo.GcpFolder == nil {
				r.RemoteInfo.GcpFolder = nil
			} else {
				r.RemoteInfo.GcpFolder = &tfTypes.GcpFolder{}
				r.RemoteInfo.GcpFolder.FolderID = types.StringValue(resp.RemoteInfo.GcpFolder.FolderID)
			}
			if resp.RemoteInfo.GcpGkeCluster == nil {
				r.RemoteInfo.GcpGkeCluster = nil
			} else {
				r.RemoteInfo.GcpGkeCluster = &tfTypes.GcpGkeCluster{}
				r.RemoteInfo.GcpGkeCluster.ClusterName = types.StringValue(resp.RemoteInfo.GcpGkeCluster.ClusterName)
			}
			if resp.RemoteInfo.GcpOrganization == nil {
				r.RemoteInfo.GcpOrganization = nil
			} else {
				r.RemoteInfo.GcpOrganization = &tfTypes.GcpOrganization{}
				r.RemoteInfo.GcpOrganization.OrganizationID = types.StringValue(resp.RemoteInfo.GcpOrganization.OrganizationID)
			}
			if resp.RemoteInfo.GcpProject == nil {
				r.RemoteInfo.GcpProject = nil
			} else {
				r.RemoteInfo.GcpProject = &tfTypes.GcpProject{}
				r.RemoteInfo.GcpProject.ProjectID = types.StringValue(resp.RemoteInfo.GcpProject.ProjectID)
			}
			if resp.RemoteInfo.GcpSQLInstance == nil {
				r.RemoteInfo.GcpSQLInstance = nil
			} else {
				r.RemoteInfo.GcpSQLInstance = &tfTypes.GcpSQLInstance{}
				r.RemoteInfo.GcpSQLInstance.InstanceID = types.StringValue(resp.RemoteInfo.GcpSQLInstance.InstanceID)
				r.RemoteInfo.GcpSQLInstance.ProjectID = types.StringValue(resp.RemoteInfo.GcpSQLInstance.ProjectID)
			}
			if resp.RemoteInfo.GithubRepo == nil {
				r.RemoteInfo.GithubRepo = nil
			} else {
				r.RemoteInfo.GithubRepo = &tfTypes.GithubRepo{}
				r.RemoteInfo.GithubRepo.RepoName = types.StringValue(resp.RemoteInfo.GithubRepo.RepoName)
			}
			if resp.RemoteInfo.GitlabProject == nil {
				r.RemoteInfo.GitlabProject = nil
			} else {
				r.RemoteInfo.GitlabProject = &tfTypes.GcpProject{}
				r.RemoteInfo.GitlabProject.ProjectID = types.StringValue(resp.RemoteInfo.GitlabProject.ProjectID)
			}
			if resp.RemoteInfo.OktaApp == nil {
				r.RemoteInfo.OktaApp = nil
			} else {
				r.RemoteInfo.OktaApp = &tfTypes.OktaApp{}
				r.RemoteInfo.OktaApp.AppID = types.StringValue(resp.RemoteInfo.OktaApp.AppID)
			}
			if resp.RemoteInfo.OktaCustomRole == nil {
				r.RemoteInfo.OktaCustomRole = nil
			} else {
				r.RemoteInfo.OktaCustomRole = &tfTypes.OktaCustomRole{}
				r.RemoteInfo.OktaCustomRole.RoleID = types.StringValue(resp.RemoteInfo.OktaCustomRole.RoleID)
			}
			if resp.RemoteInfo.OktaStandardRole == nil {
				r.RemoteInfo.OktaStandardRole = nil
			} else {
				r.RemoteInfo.OktaStandardRole = &tfTypes.OktaStandardRole{}
				r.RemoteInfo.OktaStandardRole.RoleType = types.StringValue(resp.RemoteInfo.OktaStandardRole.RoleType)
			}
			if resp.RemoteInfo.PagerdutyRole == nil {
				r.RemoteInfo.PagerdutyRole = nil
			} else {
				r.RemoteInfo.PagerdutyRole = &tfTypes.PagerdutyRole{}
				r.RemoteInfo.PagerdutyRole.RoleName = types.StringValue(resp.RemoteInfo.PagerdutyRole.RoleName)
			}
			if resp.RemoteInfo.SalesforcePermissionSet == nil {
				r.RemoteInfo.SalesforcePermissionSet = nil
			} else {
				r.RemoteInfo.SalesforcePermissionSet = &tfTypes.SalesforcePermissionSet{}
				r.RemoteInfo.SalesforcePermissionSet.PermissionSetID = types.StringValue(resp.RemoteInfo.SalesforcePermissionSet.PermissionSetID)
			}
			if resp.RemoteInfo.SalesforceProfile == nil {
				r.RemoteInfo.SalesforceProfile = nil
			} else {
				r.RemoteInfo.SalesforceProfile = &tfTypes.SalesforceProfile{}
				r.RemoteInfo.SalesforceProfile.ProfileID = types.StringValue(resp.RemoteInfo.SalesforceProfile.ProfileID)
				r.RemoteInfo.SalesforceProfile.UserLicenseID = types.StringValue(resp.RemoteInfo.SalesforceProfile.UserLicenseID)
			}
			if resp.RemoteInfo.SalesforceRole == nil {
				r.RemoteInfo.SalesforceRole = nil
			} else {
				r.RemoteInfo.SalesforceRole = &tfTypes.OktaCustomRole{}
				r.RemoteInfo.SalesforceRole.RoleID = types.StringValue(resp.RemoteInfo.SalesforceRole.RoleID)
			}
			if resp.RemoteInfo.TeleportRole == nil {
				r.RemoteInfo.TeleportRole = nil
			} else {
				r.RemoteInfo.TeleportRole = &tfTypes.PagerdutyRole{}
				r.RemoteInfo.TeleportRole.RoleName = types.StringValue(resp.RemoteInfo.TeleportRole.RoleName)
			}
		}
		r.RemoteResourceID = types.StringPointerValue(resp.RemoteResourceID)
		r.RemoteResourceName = types.StringPointerValue(resp.RemoteResourceName)
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
		r.RequireMfaToConnect = types.BoolPointerValue(resp.RequireMfaToConnect)
		r.RequireMfaToRequest = types.BoolPointerValue(resp.RequireMfaToRequest)
		r.RequireSupportTicket = types.BoolPointerValue(resp.RequireSupportTicket)
		if resp.ResourceType != nil {
			r.ResourceType = types.StringValue(string(*resp.ResourceType))
		} else {
			r.ResourceType = types.StringNull()
		}
	}
}
