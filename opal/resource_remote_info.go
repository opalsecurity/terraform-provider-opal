package opal

import (
	"encoding/json"
	"errors"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/opalsecurity/opal-go"
)

// NOTE: Unfortunately, terraform go-sdk does not support nested object types natively. The only work-around
// is to have a schema.ListType with MaxItems=1 as a layer of indirection in between each level of nesting.
// This makes the implementation ugly, but it's thankfully mostly hidden from the client. If nested objects
// are ever natively supported it in the SDK, we should be able to update our code without a need for
// change in the HCL of clients.
func resourceRemoteInfoElem() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"aws_account": {
				Description: "The remote_info for an AWS account.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"account_id": {
							Description: "The ID of the AWS account.",
							Type:        schema.TypeString,
							Required:    true,
							ForceNew:    true,
						},
					},
				},
			},
			"aws_iam_role": {
				Description: "The remote_info for an AWS IAM role.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"arn": {
							Description: "The ARN of the IAM role.",
							Type:        schema.TypeString,
							Required:    true,
							ForceNew:    true,
						},
						"account_id": {
							Description: "The ID of the AWS account.",
							Type:        schema.TypeString,
							ForceNew:    true,
							Optional:    true,
						},
					},
				},
			},
			"aws_ec2_instance": {
				Description: "The remote_info for an AWS EC2 instance.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"instance_id": {
							Description: "The instanceId of the EC2 instance.",
							Type:        schema.TypeString,
							Required:    true,
							ForceNew:    true,
						},
						"region": {
							Description: "The region of the EC2 instance.",
							Type:        schema.TypeString,
							Required:    true,
							ForceNew:    true,
						},
						"account_id": {
							Description: "The ID of the AWS account.",
							Type:        schema.TypeString,
							ForceNew:    true,
							Optional:    true,
						},
					},
				},
			},
			"aws_rds_instance": {
				Description: "The remote_info for an AWS RDS instance.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"instance_id": {
							Description: "The instanceId of the RDS instance.",
							Type:        schema.TypeString,
							Required:    true,
							ForceNew:    true,
						},
						"resource_id": {
							Description: "The resourceId of the RDS instance.",
							Type:        schema.TypeString,
							Required:    true,
							ForceNew:    true,
						},
						"region": {
							Description: "The region of the RDS instance.",
							Type:        schema.TypeString,
							Required:    true,
							ForceNew:    true,
						},
						"account_id": {
							Description: "The ID of the AWS account.",
							Type:        schema.TypeString,
							ForceNew:    true,
							Optional:    true,
						},
					},
				},
			},
			"aws_eks_cluster": {
				Description: "The remote_info for an AWS EKS cluster.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"arn": {
							Description: "The ARN of the EKS cluster.",
							Type:        schema.TypeString,
							Required:    true,
							ForceNew:    true,
						},
						"account_id": {
							Description: "The ID of the AWS account.",
							Type:        schema.TypeString,
							ForceNew:    true,
							Optional:    true,
						},
					},
				},
			},
			"aws_permission_set": {
				Description: "The remote_info for an AWS permission set.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"arn": {
							Description: "The ARN of the permission set.",
							Type:        schema.TypeString,
							Required:    true,
							ForceNew:    true,
						},
						"account_id": {
							Description: "The ID of the AWS account.",
							Type:        schema.TypeString,
							Required:    true,
							ForceNew:    true,
						},
					},
				},
			},
			"gcp_bucket": {
				Description: "The remote_info for a GCP bucket.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"bucket_id": {
							Description: "The id of the bucket.",
							Type:        schema.TypeString,
							Required:    true,
							ForceNew:    true,
						},
					},
				},
			},
			"gcp_compute_instance": {
				Description: "The remote_info for a GCP compute instance.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"instance_id": {
							Description: "The id of the compute instance.",
							Type:        schema.TypeString,
							Required:    true,
							ForceNew:    true,
						},
						"project_id": {
							Description: "The id of the project the instance is in.",
							Type:        schema.TypeString,
							Required:    true,
							ForceNew:    true,
						},
						"zone": {
							Description: "The zone of the compute instance.",
							Type:        schema.TypeString,
							Required:    true,
							ForceNew:    true,
						},
					},
				},
			},
			"gcp_folder": {
				Description: "The remote_info for a GCP folder.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"folder_id": {
							Description: "The id of the folder.",
							Type:        schema.TypeString,
							Required:    true,
							ForceNew:    true,
						},
					},
				},
			},
			"gcp_gke_cluster": {
				Description: "The remote_info for a GCP GKE cluster.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cluster_name": {
							Description: "The name of the cluster.",
							Type:        schema.TypeString,
							Required:    true,
							ForceNew:    true,
						},
					},
				},
			},
			"gcp_project": {
				Description: "The remote_info for a GCP project.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"project_id": {
							Description: "The id of the project.",
							Type:        schema.TypeString,
							Required:    true,
							ForceNew:    true,
						},
					},
				},
			},
			"gcp_sql_instance": {
				Description: "The remote_info for a GCP SQL instance.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"instance_id": {
							Description: "The id of the sql instance.",
							Type:        schema.TypeString,
							Required:    true,
							ForceNew:    true,
						},
						"project_id": {
							Description: "The id of the project the instance is in.",
							Type:        schema.TypeString,
							Required:    true,
							ForceNew:    true,
						},
					},
				},
			},
			"github_repo": {
				Description: "The remote_info for a Github repo.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"repo_name": {
							Description: "The name of the repository.",
							Type:        schema.TypeString,
							Required:    true,
							ForceNew:    true,
						},
					},
				},
			},
			"gitlab_project": {
				Description: "The remote_info for a Gitlab project.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"project_id": {
							Description: "The id of the project.",
							Type:        schema.TypeString,
							Required:    true,
							ForceNew:    true,
						},
					},
				},
			},
			"okta_app": {
				Description: "The remote_info for an Okta app.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"app_id": {
							Description: "The id of the app.",
							Type:        schema.TypeString,
							Required:    true,
							ForceNew:    true,
						},
					},
				},
			},
			"okta_standard_role": {
				Description: "The remote_info for an Okta standard role.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"role_type": {
							Description: "The type of the role.",
							Type:        schema.TypeString,
							Required:    true,
							ForceNew:    true,
						},
					},
				},
			},
			"okta_custom_role": {
				Description: "The remote_info for an Okta custom role.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"role_id": {
							Description: "The id of the role.",
							Type:        schema.TypeString,
							Required:    true,
							ForceNew:    true,
						},
					},
				},
			},
			"pagerduty_role": {
				Description: "The remote_info for a Pagerduty role.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"role_name": {
							Description: "The name of the role.",
							Type:        schema.TypeString,
							Required:    true,
							ForceNew:    true,
						},
					},
				},
			},
			"salesforce_permission_set": {
				Description: "The remote_info for a Salesforce permission set.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"permission_set_id": {
							Description: "The id of the permission set.",
							Type:        schema.TypeString,
							Required:    true,
							ForceNew:    true,
						},
					},
				},
			},
			"salesforce_profile": {
				Description: "The remote_info for a Salesforce profile.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"profile_id": {
							Description: "The id of the profile.",
							Type:        schema.TypeString,
							Required:    true,
							ForceNew:    true,
						},
						"user_license_id": {
							Description: "The id of the user license.",
							Type:        schema.TypeString,
							Required:    true,
							ForceNew:    true,
						},
					},
				},
			},
			"salesforce_role": {
				Description: "The remote_info for a Salesforce role.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"role_id": {
							Description: "The id of the role.",
							Type:        schema.TypeString,
							Required:    true,
							ForceNew:    true,
						},
					},
				},
			},
			"teleport_role": {
				Description: "The remote_info for a Teleport role.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"role_name": {
							Description: "The name of the role.",
							Type:        schema.TypeString,
							Required:    true,
							ForceNew:    true,
						},
					},
				},
			},
		},
	}
}

func resourceRemoteInfoAPIToTerraform(remoteInfo *opal.ResourceRemoteInfo) (interface{}, error) {
	return remoteInfoAPIToTerraformInternal(remoteInfo)
}

// NOTE: See comment in `resourceRemoteInfoElem` for details on the structure we're parsing into
func remoteInfoAPIToTerraformInternal(remoteInfo interface{}) (interface{}, error) {
	var remoteInfoMap map[string]map[string]interface{}
	jsonRemoteInfo, err := json.Marshal(remoteInfo)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(jsonRemoteInfo, &remoteInfoMap)
	if err != nil {
		return nil, err
	}

	if len(remoteInfoMap) == 0 {
		return nil, nil
	}

	deprecatedKeysByApp := map[string]map[string]bool{
		"github_repo": {
			"repo_id": true,
		},
		"github_team": {
			"team_id": true,
		},
	}
	remoteInfoIList := make([]interface{}, 1)
	for appKey, remoteInfoRaw := range remoteInfoMap {
		itemRemoteInfo := map[string]interface{}{}
		for k, v := range remoteInfoRaw {
			if deprecatedKeys, ok := deprecatedKeysByApp[appKey]; ok {
				if _, ok := deprecatedKeys[k]; ok {
					continue
				}
			}
			itemRemoteInfo[k] = v
		}

		remoteInfoIList[0] = map[string]interface{}{
			appKey: []interface{}{
				itemRemoteInfo,
			},
		}
	}

	return remoteInfoIList, nil
}

// NOTE: See comment in `resourceRemoteInfoElem` for why the parsing is so convoluted.
func resourceRemoteInfoTerraformToAPI(remoteInfoI interface{}) (*opal.ResourceRemoteInfo, error) {
	remoteInfoIList := remoteInfoI.([]interface{})
	if len(remoteInfoIList) != 1 {
		return nil, errors.New("you cannot provide multiple remote_info blobs")
	}

	remoteInfoMap := remoteInfoIList[0].(map[string]interface{})
	if awsAccountI, ok := remoteInfoMap["aws_account"]; ok {
		awsAccountIList := awsAccountI.([]interface{})
		if len(awsAccountIList) == 1 {
			awsAccount := awsAccountIList[0].(map[string]any)
			return &opal.ResourceRemoteInfo{
				AwsAccount: &opal.ResourceRemoteInfoAwsAccount{
					AccountId: awsAccount["account_id"].(string),
				},
			}, nil
		}
	}
	if awsIamRoleI, ok := remoteInfoMap["aws_iam_role"]; ok {
		awsIamRoleIList := awsIamRoleI.([]interface{})

		if len(awsIamRoleIList) == 1 {
			awsIamRole := awsIamRoleIList[0].(map[string]any)
			remoteInfo := &opal.ResourceRemoteInfo{
				AwsIamRole: &opal.ResourceRemoteInfoAwsIamRole{
					Arn: awsIamRole["arn"].(string),
				},
			}
			if awsIamRoleAccountIdI, ok := awsIamRole["account_id"]; ok {
				awsIamRoleAccountId := awsIamRoleAccountIdI.(string)
				remoteInfo.AwsIamRole.AccountId = &awsIamRoleAccountId
			}
			return remoteInfo, nil
		}
	}
	if awsEc2InstanceI, ok := remoteInfoMap["aws_ec2_instance"]; ok {
		awsEc2InstanceIList := awsEc2InstanceI.([]interface{})

		if len(awsEc2InstanceIList) == 1 {
			awsEc2Instance := awsEc2InstanceIList[0].(map[string]any)
			remoteInfo := &opal.ResourceRemoteInfo{
				AwsEc2Instance: &opal.ResourceRemoteInfoAwsEc2Instance{
					InstanceId: awsEc2Instance["instance_id"].(string),
					Region:     awsEc2Instance["region"].(string),
				},
			}
			if awsEc2InstanceAccountIdI, ok := awsEc2Instance["account_id"]; ok {
				awsEc2InstanceAccountId := awsEc2InstanceAccountIdI.(string)
				remoteInfo.AwsEc2Instance.AccountId = &awsEc2InstanceAccountId
			}
			return remoteInfo, nil
		}
	}
	if awsRdsInstanceI, ok := remoteInfoMap["aws_rds_instance"]; ok {
		awsRdsInstanceIList := awsRdsInstanceI.([]interface{})

		if len(awsRdsInstanceIList) == 1 {
			awsRdsInstance := awsRdsInstanceIList[0].(map[string]any)
			remoteInfo := &opal.ResourceRemoteInfo{
				AwsRdsInstance: &opal.ResourceRemoteInfoAwsRdsInstance{
					InstanceId: awsRdsInstance["instance_id"].(string),
					ResourceId: awsRdsInstance["resource_id"].(string),
					Region:     awsRdsInstance["region"].(string),
				},
			}
			if awsRdsInstanceAccountIdI, ok := awsRdsInstance["account_id"]; ok {
				awsRdsInstanceAccountId := awsRdsInstanceAccountIdI.(string)
				remoteInfo.AwsRdsInstance.AccountId = &awsRdsInstanceAccountId
			}
			return remoteInfo, nil
		}
	}
	if awsEksClusterI, ok := remoteInfoMap["aws_eks_cluster"]; ok {
		awsEksClusterIList := awsEksClusterI.([]interface{})

		if len(awsEksClusterIList) == 1 {
			awsEksCluster := awsEksClusterIList[0].(map[string]any)
			remoteInfo := &opal.ResourceRemoteInfo{
				AwsEksCluster: &opal.ResourceRemoteInfoAwsEksCluster{
					Arn: awsEksCluster["arn"].(string),
				},
			}
			if awsEksClusterAccountIdI, ok := awsEksCluster["account_id"]; ok {
				awsEksClusterAccountId := awsEksClusterAccountIdI.(string)
				remoteInfo.AwsEksCluster.AccountId = &awsEksClusterAccountId
			}
			return remoteInfo, nil
		}
	}
	if awsPermissionSetI, ok := remoteInfoMap["aws_permission_set"]; ok {
		awsPermissionSetIList := awsPermissionSetI.([]interface{})
		if len(awsPermissionSetIList) == 1 {
			awsPermissionSet := awsPermissionSetIList[0].(map[string]any)
			return &opal.ResourceRemoteInfo{
				AwsPermissionSet: &opal.ResourceRemoteInfoAwsPermissionSet{
					Arn:       awsPermissionSet["arn"].(string),
					AccountId: awsPermissionSet["account_id"].(string),
				},
			}, nil
		}
	}
	if gcpBucketI, ok := remoteInfoMap["gcp_bucket"]; ok {
		gcpBucketIList := gcpBucketI.([]interface{})

		if len(gcpBucketIList) == 1 {
			gcpBucket := gcpBucketIList[0].(map[string]any)
			return &opal.ResourceRemoteInfo{
				GcpBucket: &opal.ResourceRemoteInfoGcpBucket{
					BucketId: gcpBucket["bucket_id"].(string),
				},
			}, nil
		}
	}
	if gcpComputeInstanceI, ok := remoteInfoMap["gcp_compute_instance"]; ok {
		gcpComputeInstanceIList := gcpComputeInstanceI.([]interface{})

		if len(gcpComputeInstanceIList) == 1 {
			gcpComputeInstance := gcpComputeInstanceIList[0].(map[string]any)
			return &opal.ResourceRemoteInfo{
				GcpComputeInstance: &opal.ResourceRemoteInfoGcpComputeInstance{
					InstanceId: gcpComputeInstance["instance_id"].(string),
					ProjectId:  gcpComputeInstance["project_id"].(string),
					Zone:       gcpComputeInstance["zone"].(string),
				},
			}, nil
		}
	}
	if gcpFolderI, ok := remoteInfoMap["gcp_folder"]; ok {
		gcpFolderIList := gcpFolderI.([]interface{})

		if len(gcpFolderIList) == 1 {
			gcpFolder := gcpFolderIList[0].(map[string]any)
			return &opal.ResourceRemoteInfo{
				GcpFolder: &opal.ResourceRemoteInfoGcpFolder{
					FolderId: gcpFolder["folder_id"].(string),
				},
			}, nil
		}
	}
	if gcpGkeClusterI, ok := remoteInfoMap["gcp_gke_cluster"]; ok {
		gcpGkeClusterIList := gcpGkeClusterI.([]interface{})

		if len(gcpGkeClusterIList) == 1 {
			gcpGkeCluster := gcpGkeClusterIList[0].(map[string]any)
			return &opal.ResourceRemoteInfo{
				GcpGkeCluster: &opal.ResourceRemoteInfoGcpGkeCluster{
					ClusterName: gcpGkeCluster["cluster_name"].(string),
				},
			}, nil
		}
	}
	if gcpProjectI, ok := remoteInfoMap["gcp_project"]; ok {
		gcpProjectIList := gcpProjectI.([]interface{})

		if len(gcpProjectIList) == 1 {
			gcpProject := gcpProjectIList[0].(map[string]any)
			return &opal.ResourceRemoteInfo{
				GcpProject: &opal.ResourceRemoteInfoGcpProject{
					ProjectId: gcpProject["project_id"].(string),
				},
			}, nil
		}
	}
	if gcpSqlInstanceI, ok := remoteInfoMap["gcp_sql_instance"]; ok {
		gcpSqlInstanceIList := gcpSqlInstanceI.([]interface{})

		if len(gcpSqlInstanceIList) == 1 {
			gcpSqlInstance := gcpSqlInstanceIList[0].(map[string]any)
			return &opal.ResourceRemoteInfo{
				GcpSqlInstance: &opal.ResourceRemoteInfoGcpSqlInstance{
					InstanceId: gcpSqlInstance["instance_id"].(string),
					ProjectId:  gcpSqlInstance["project_id"].(string),
				},
			}, nil
		}
	}
	if githubRepoI, ok := remoteInfoMap["github_repo"]; ok {
		githubRepoIList := githubRepoI.([]interface{})

		if len(githubRepoIList) == 1 {
			githubRepo := githubRepoIList[0].(map[string]any)
			return &opal.ResourceRemoteInfo{
				GithubRepo: &opal.ResourceRemoteInfoGithubRepo{
					RepoName: githubRepo["repo_name"].(string),
				},
			}, nil
		}
	}
	if gitlabProjectI, ok := remoteInfoMap["gitlab_project"]; ok {
		gitlabProjectIList := gitlabProjectI.([]interface{})

		if len(gitlabProjectIList) == 1 {
			gitlabProject := gitlabProjectIList[0].(map[string]any)
			return &opal.ResourceRemoteInfo{
				GitlabProject: &opal.ResourceRemoteInfoGitlabProject{
					ProjectId: gitlabProject["project_id"].(string),
				},
			}, nil
		}
	}
	if oktaAppI, ok := remoteInfoMap["okta_app"]; ok {
		oktaAppIList := oktaAppI.([]interface{})

		if len(oktaAppIList) == 1 {
			oktaApp := oktaAppIList[0].(map[string]any)
			return &opal.ResourceRemoteInfo{
				OktaApp: &opal.ResourceRemoteInfoOktaApp{
					AppId: oktaApp["app_id"].(string),
				},
			}, nil
		}
	}
	if oktaStandardRoleI, ok := remoteInfoMap["okta_standard_role"]; ok {
		oktaStandardRoleIList := oktaStandardRoleI.([]interface{})

		if len(oktaStandardRoleIList) == 1 {
			oktaStandardRole := oktaStandardRoleIList[0].(map[string]any)
			return &opal.ResourceRemoteInfo{
				OktaStandardRole: &opal.ResourceRemoteInfoOktaStandardRole{
					RoleType: oktaStandardRole["role_type"].(string),
				},
			}, nil
		}
	}
	if oktaCustomRoleI, ok := remoteInfoMap["okta_custom_role"]; ok {
		oktaCustomRoleIList := oktaCustomRoleI.([]interface{})

		if len(oktaCustomRoleIList) == 1 {
			oktaCustomRole := oktaCustomRoleIList[0].(map[string]any)
			return &opal.ResourceRemoteInfo{
				OktaCustomRole: &opal.ResourceRemoteInfoOktaCustomRole{
					RoleId: oktaCustomRole["role_id"].(string),
				},
			}, nil
		}
	}
	if pagerdutyRoleI, ok := remoteInfoMap["pagerduty_role"]; ok {
		pagerdutyRoleIList := pagerdutyRoleI.([]interface{})

		if len(pagerdutyRoleIList) == 1 {
			pagerdutyRole := pagerdutyRoleIList[0].(map[string]any)
			return &opal.ResourceRemoteInfo{
				PagerdutyRole: &opal.ResourceRemoteInfoPagerdutyRole{
					RoleName: pagerdutyRole["role_name"].(string),
				},
			}, nil
		}
	}
	if salesforcePermissionSetI, ok := remoteInfoMap["salesforce_permission_set"]; ok {
		salesforcePermissionSetIList := salesforcePermissionSetI.([]interface{})

		if len(salesforcePermissionSetIList) == 1 {
			salesforcePermissionSet := salesforcePermissionSetIList[0].(map[string]any)
			return &opal.ResourceRemoteInfo{
				SalesforcePermissionSet: &opal.ResourceRemoteInfoSalesforcePermissionSet{
					PermissionSetId: salesforcePermissionSet["permission_set_id"].(string),
				},
			}, nil
		}
	}
	if salesforceProfileI, ok := remoteInfoMap["salesforce_profile"]; ok {
		salesforceProfileIList := salesforceProfileI.([]interface{})

		if len(salesforceProfileIList) == 1 {
			salesforceProfile := salesforceProfileIList[0].(map[string]any)
			return &opal.ResourceRemoteInfo{
				SalesforceProfile: &opal.ResourceRemoteInfoSalesforceProfile{
					ProfileId:     salesforceProfile["profile_id"].(string),
					UserLicenseId: salesforceProfile["user_license_id"].(string),
				},
			}, nil
		}
	}
	if salesforceRoleI, ok := remoteInfoMap["salesforce_role"]; ok {
		salesforceRoleIList := salesforceRoleI.([]interface{})

		if len(salesforceRoleIList) == 1 {
			salesforceRole := salesforceRoleIList[0].(map[string]any)
			return &opal.ResourceRemoteInfo{
				SalesforceRole: &opal.ResourceRemoteInfoSalesforceRole{
					RoleId: salesforceRole["role_id"].(string),
				},
			}, nil
		}
	}
	if teleportRoleI, ok := remoteInfoMap["teleport_role"]; ok {
		teleportRoleIList := teleportRoleI.([]interface{})

		if len(teleportRoleIList) == 1 {
			teleportRole := teleportRoleIList[0].(map[string]any)
			return &opal.ResourceRemoteInfo{
				TeleportRole: &opal.ResourceRemoteInfoTeleportRole{
					RoleName: teleportRole["role_name"].(string),
				},
			}, nil
		}
	}

	return nil, errors.New("could not find supported remote_info type")
}
