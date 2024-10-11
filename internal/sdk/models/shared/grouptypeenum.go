// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// GroupTypeEnum - The type of the group.
type GroupTypeEnum string

const (
	GroupTypeEnumActiveDirectoryGroup     GroupTypeEnum = "ACTIVE_DIRECTORY_GROUP"
	GroupTypeEnumAwsSsoGroup              GroupTypeEnum = "AWS_SSO_GROUP"
	GroupTypeEnumDuoGroup                 GroupTypeEnum = "DUO_GROUP"
	GroupTypeEnumGitHubTeam               GroupTypeEnum = "GIT_HUB_TEAM"
	GroupTypeEnumGitLabGroup              GroupTypeEnum = "GIT_LAB_GROUP"
	GroupTypeEnumGoogleGroupsGroup        GroupTypeEnum = "GOOGLE_GROUPS_GROUP"
	GroupTypeEnumGoogleGroupsGkeGroup     GroupTypeEnum = "GOOGLE_GROUPS_GKE_GROUP"
	GroupTypeEnumLdapGroup                GroupTypeEnum = "LDAP_GROUP"
	GroupTypeEnumOktaGroup                GroupTypeEnum = "OKTA_GROUP"
	GroupTypeEnumTailscaleGroup           GroupTypeEnum = "TAILSCALE_GROUP"
	GroupTypeEnumOpalGroup                GroupTypeEnum = "OPAL_GROUP"
	GroupTypeEnumAzureAdSecurityGroup     GroupTypeEnum = "AZURE_AD_SECURITY_GROUP"
	GroupTypeEnumAzureAdMicrosoft365Group GroupTypeEnum = "AZURE_AD_MICROSOFT_365_GROUP"
	GroupTypeEnumConnectorGroup           GroupTypeEnum = "CONNECTOR_GROUP"
	GroupTypeEnumSnowflakeRole            GroupTypeEnum = "SNOWFLAKE_ROLE"
	GroupTypeEnumWorkdayUserSecurityGroup GroupTypeEnum = "WORKDAY_USER_SECURITY_GROUP"
)

func (e GroupTypeEnum) ToPointer() *GroupTypeEnum {
	return &e
}
func (e *GroupTypeEnum) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "ACTIVE_DIRECTORY_GROUP":
		fallthrough
	case "AWS_SSO_GROUP":
		fallthrough
	case "DUO_GROUP":
		fallthrough
	case "GIT_HUB_TEAM":
		fallthrough
	case "GIT_LAB_GROUP":
		fallthrough
	case "GOOGLE_GROUPS_GROUP":
		fallthrough
	case "GOOGLE_GROUPS_GKE_GROUP":
		fallthrough
	case "LDAP_GROUP":
		fallthrough
	case "OKTA_GROUP":
		fallthrough
	case "TAILSCALE_GROUP":
		fallthrough
	case "OPAL_GROUP":
		fallthrough
	case "AZURE_AD_SECURITY_GROUP":
		fallthrough
	case "AZURE_AD_MICROSOFT_365_GROUP":
		fallthrough
	case "CONNECTOR_GROUP":
		fallthrough
	case "SNOWFLAKE_ROLE":
		fallthrough
	case "WORKDAY_USER_SECURITY_GROUP":
		*e = GroupTypeEnum(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GroupTypeEnum: %v", v)
	}
}
