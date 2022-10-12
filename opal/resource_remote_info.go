package opal

import (
	"errors"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/opalsecurity/opal-go"
)

// NOTE: Unfortunately, terraform go-sdk does not support nested object types natively. The only work-around
//       is to have a schema.ListType with MaxItems=1 as a layer of indirection in between each level of nesting.
//       This makes the implementation ugly, but it's thankfully mostly hidden from the client. If nested objects
//       are ever natively supported it in the SDK, we should be able to update our code without a need for
//       change in the HCL of clients.
func resourceRemoteInfoElem() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
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
						"repo_id": {
							Description: "The id of the repository.",
							Type:        schema.TypeString,
							Required:    true,
							ForceNew:    true,
						},
						"repo_name": {
							Description: "The name of the repository.",
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
		},
	}
}

// NOTE: See comment in `resourceRemoteInfoElem` for why the parsing is so convoluted.
func parseResourceRemoteInfo(remoteInfoI interface{}) (*opal.ResourceRemoteInfo, error) {
	remoteInfoIList := remoteInfoI.([]interface{})
	if len(remoteInfoIList) != 1 {
		return nil, errors.New("you cannot provide multiple remote_info blobs")
	}

	remoteInfoMap := remoteInfoIList[0].(map[string]interface{})
	if awsIamRoleI, ok := remoteInfoMap["aws_iam_role"]; ok {
		awsIamRoleIList := awsIamRoleI.([]interface{})

		if len(awsIamRoleIList) == 1 {
			awsIamRole := awsIamRoleIList[0].(map[string]any)
			return &opal.ResourceRemoteInfo{
				AwsIamRole: &opal.ResourceRemoteInfoAwsIamRole{
					Arn: awsIamRole["arn"].(string),
				},
			}, nil
		}
	}
	if awsEc2InstanceI, ok := remoteInfoMap["aws_ec2_instance"]; ok {
		awsEc2InstanceIList := awsEc2InstanceI.([]interface{})

		if len(awsEc2InstanceIList) == 1 {
			awsEc2Instance := awsEc2InstanceIList[0].(map[string]any)
			return &opal.ResourceRemoteInfo{
				AwsEc2Instance: &opal.ResourceRemoteInfoAwsEc2Instance{
					InstanceId: awsEc2Instance["instance_id"].(string),
					Region:     awsEc2Instance["region"].(string),
				},
			}, nil
		}
	}
	if awsRdsInstanceI, ok := remoteInfoMap["aws_rds_instance"]; ok {
		awsRdsInstanceIList := awsRdsInstanceI.([]interface{})

		if len(awsRdsInstanceIList) == 1 {
			awsRdsInstance := awsRdsInstanceIList[0].(map[string]any)
			return &opal.ResourceRemoteInfo{
				AwsRdsInstance: &opal.ResourceRemoteInfoAwsRdsInstance{
					InstanceId: awsRdsInstance["instance_id"].(string),
					ResourceId: awsRdsInstance["resource_id"].(string),
					Region:     awsRdsInstance["region"].(string),
				},
			}, nil
		}
	}
	if awsEksClusterI, ok := remoteInfoMap["aws_eks_cluster"]; ok {
		awsEksClusterIList := awsEksClusterI.([]interface{})

		if len(awsEksClusterIList) == 1 {
			awsEksCluster := awsEksClusterIList[0].(map[string]any)
			return &opal.ResourceRemoteInfo{
				AwsEksCluster: &opal.ResourceRemoteInfoAwsEksCluster{
					Arn: awsEksCluster["arn"].(string),
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
					RepoId:   githubRepo["repo_id"].(string),
					RepoName: githubRepo["repo_name"].(string),
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

	return nil, errors.New("could not find supported remote_info type")
}
