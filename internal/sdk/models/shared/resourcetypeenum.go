// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// ResourceTypeEnum - The type of the resource.
type ResourceTypeEnum string

const (
	ResourceTypeEnumAwsIamRole                  ResourceTypeEnum = "AWS_IAM_ROLE"
	ResourceTypeEnumAwsEc2Instance              ResourceTypeEnum = "AWS_EC2_INSTANCE"
	ResourceTypeEnumAwsEksCluster               ResourceTypeEnum = "AWS_EKS_CLUSTER"
	ResourceTypeEnumAwsRdsPostgresInstance      ResourceTypeEnum = "AWS_RDS_POSTGRES_INSTANCE"
	ResourceTypeEnumAwsRdsMysqlInstance         ResourceTypeEnum = "AWS_RDS_MYSQL_INSTANCE"
	ResourceTypeEnumAwsAccount                  ResourceTypeEnum = "AWS_ACCOUNT"
	ResourceTypeEnumAwsSsoPermissionSet         ResourceTypeEnum = "AWS_SSO_PERMISSION_SET"
	ResourceTypeEnumCustom                      ResourceTypeEnum = "CUSTOM"
	ResourceTypeEnumGcpOrganization             ResourceTypeEnum = "GCP_ORGANIZATION"
	ResourceTypeEnumGcpBucket                   ResourceTypeEnum = "GCP_BUCKET"
	ResourceTypeEnumGcpComputeInstance          ResourceTypeEnum = "GCP_COMPUTE_INSTANCE"
	ResourceTypeEnumGcpFolder                   ResourceTypeEnum = "GCP_FOLDER"
	ResourceTypeEnumGcpGkeCluster               ResourceTypeEnum = "GCP_GKE_CLUSTER"
	ResourceTypeEnumGcpProject                  ResourceTypeEnum = "GCP_PROJECT"
	ResourceTypeEnumGcpCloudSQLPostgresInstance ResourceTypeEnum = "GCP_CLOUD_SQL_POSTGRES_INSTANCE"
	ResourceTypeEnumGcpCloudSQLMysqlInstance    ResourceTypeEnum = "GCP_CLOUD_SQL_MYSQL_INSTANCE"
	ResourceTypeEnumGcpBigQueryDataset          ResourceTypeEnum = "GCP_BIG_QUERY_DATASET"
	ResourceTypeEnumGcpBigQueryTable            ResourceTypeEnum = "GCP_BIG_QUERY_TABLE"
	ResourceTypeEnumGcpServiceAccount           ResourceTypeEnum = "GCP_SERVICE_ACCOUNT"
	ResourceTypeEnumGitHubRepo                  ResourceTypeEnum = "GIT_HUB_REPO"
	ResourceTypeEnumGitLabProject               ResourceTypeEnum = "GIT_LAB_PROJECT"
	ResourceTypeEnumGoogleWorkspaceRole         ResourceTypeEnum = "GOOGLE_WORKSPACE_ROLE"
	ResourceTypeEnumMongoInstance               ResourceTypeEnum = "MONGO_INSTANCE"
	ResourceTypeEnumMongoAtlasInstance          ResourceTypeEnum = "MONGO_ATLAS_INSTANCE"
	ResourceTypeEnumOktaApp                     ResourceTypeEnum = "OKTA_APP"
	ResourceTypeEnumOktaRole                    ResourceTypeEnum = "OKTA_ROLE"
	ResourceTypeEnumOpalRole                    ResourceTypeEnum = "OPAL_ROLE"
	ResourceTypeEnumPagerdutyRole               ResourceTypeEnum = "PAGERDUTY_ROLE"
	ResourceTypeEnumTailscaleSSH                ResourceTypeEnum = "TAILSCALE_SSH"
	ResourceTypeEnumSalesforcePermissionSet     ResourceTypeEnum = "SALESFORCE_PERMISSION_SET"
	ResourceTypeEnumSalesforceProfile           ResourceTypeEnum = "SALESFORCE_PROFILE"
	ResourceTypeEnumSalesforceRole              ResourceTypeEnum = "SALESFORCE_ROLE"
	ResourceTypeEnumWorkdayRole                 ResourceTypeEnum = "WORKDAY_ROLE"
	ResourceTypeEnumMysqlInstance               ResourceTypeEnum = "MYSQL_INSTANCE"
	ResourceTypeEnumMariadbInstance             ResourceTypeEnum = "MARIADB_INSTANCE"
	ResourceTypeEnumTeleportRole                ResourceTypeEnum = "TELEPORT_ROLE"
)

func (e ResourceTypeEnum) ToPointer() *ResourceTypeEnum {
	return &e
}
func (e *ResourceTypeEnum) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "AWS_IAM_ROLE":
		fallthrough
	case "AWS_EC2_INSTANCE":
		fallthrough
	case "AWS_EKS_CLUSTER":
		fallthrough
	case "AWS_RDS_POSTGRES_INSTANCE":
		fallthrough
	case "AWS_RDS_MYSQL_INSTANCE":
		fallthrough
	case "AWS_ACCOUNT":
		fallthrough
	case "AWS_SSO_PERMISSION_SET":
		fallthrough
	case "CUSTOM":
		fallthrough
	case "GCP_ORGANIZATION":
		fallthrough
	case "GCP_BUCKET":
		fallthrough
	case "GCP_COMPUTE_INSTANCE":
		fallthrough
	case "GCP_FOLDER":
		fallthrough
	case "GCP_GKE_CLUSTER":
		fallthrough
	case "GCP_PROJECT":
		fallthrough
	case "GCP_CLOUD_SQL_POSTGRES_INSTANCE":
		fallthrough
	case "GCP_CLOUD_SQL_MYSQL_INSTANCE":
		fallthrough
	case "GCP_BIG_QUERY_DATASET":
		fallthrough
	case "GCP_BIG_QUERY_TABLE":
		fallthrough
	case "GCP_SERVICE_ACCOUNT":
		fallthrough
	case "GIT_HUB_REPO":
		fallthrough
	case "GIT_LAB_PROJECT":
		fallthrough
	case "GOOGLE_WORKSPACE_ROLE":
		fallthrough
	case "MONGO_INSTANCE":
		fallthrough
	case "MONGO_ATLAS_INSTANCE":
		fallthrough
	case "OKTA_APP":
		fallthrough
	case "OKTA_ROLE":
		fallthrough
	case "OPAL_ROLE":
		fallthrough
	case "PAGERDUTY_ROLE":
		fallthrough
	case "TAILSCALE_SSH":
		fallthrough
	case "SALESFORCE_PERMISSION_SET":
		fallthrough
	case "SALESFORCE_PROFILE":
		fallthrough
	case "SALESFORCE_ROLE":
		fallthrough
	case "WORKDAY_ROLE":
		fallthrough
	case "MYSQL_INSTANCE":
		fallthrough
	case "MARIADB_INSTANCE":
		fallthrough
	case "TELEPORT_ROLE":
		*e = ResourceTypeEnum(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ResourceTypeEnum: %v", v)
	}
}
