// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// AppTypeEnum - The type of an app.
type AppTypeEnum string

const (
	AppTypeEnumActiveDirectory AppTypeEnum = "ACTIVE_DIRECTORY"
	AppTypeEnumAzureAd         AppTypeEnum = "AZURE_AD"
	AppTypeEnumAws             AppTypeEnum = "AWS"
	AppTypeEnumAwsSso          AppTypeEnum = "AWS_SSO"
	AppTypeEnumCustom          AppTypeEnum = "CUSTOM"
	AppTypeEnumDuo             AppTypeEnum = "DUO"
	AppTypeEnumGcp             AppTypeEnum = "GCP"
	AppTypeEnumGitHub          AppTypeEnum = "GIT_HUB"
	AppTypeEnumGitLab          AppTypeEnum = "GIT_LAB"
	AppTypeEnumGoogleGroups    AppTypeEnum = "GOOGLE_GROUPS"
	AppTypeEnumGoogleWorkspace AppTypeEnum = "GOOGLE_WORKSPACE"
	AppTypeEnumLdap            AppTypeEnum = "LDAP"
	AppTypeEnumMariadb         AppTypeEnum = "MARIADB"
	AppTypeEnumMongo           AppTypeEnum = "MONGO"
	AppTypeEnumMongoAtlas      AppTypeEnum = "MONGO_ATLAS"
	AppTypeEnumMysql           AppTypeEnum = "MYSQL"
	AppTypeEnumOktaDirectory   AppTypeEnum = "OKTA_DIRECTORY"
	AppTypeEnumOpal            AppTypeEnum = "OPAL"
	AppTypeEnumPagerduty       AppTypeEnum = "PAGERDUTY"
	AppTypeEnumSalesforce      AppTypeEnum = "SALESFORCE"
	AppTypeEnumTailscale       AppTypeEnum = "TAILSCALE"
	AppTypeEnumTeleport        AppTypeEnum = "TELEPORT"
	AppTypeEnumWorkday         AppTypeEnum = "WORKDAY"
)

func (e AppTypeEnum) ToPointer() *AppTypeEnum {
	return &e
}

func (e *AppTypeEnum) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "ACTIVE_DIRECTORY":
		fallthrough
	case "AZURE_AD":
		fallthrough
	case "AWS":
		fallthrough
	case "AWS_SSO":
		fallthrough
	case "CUSTOM":
		fallthrough
	case "DUO":
		fallthrough
	case "GCP":
		fallthrough
	case "GIT_HUB":
		fallthrough
	case "GIT_LAB":
		fallthrough
	case "GOOGLE_GROUPS":
		fallthrough
	case "GOOGLE_WORKSPACE":
		fallthrough
	case "LDAP":
		fallthrough
	case "MARIADB":
		fallthrough
	case "MONGO":
		fallthrough
	case "MONGO_ATLAS":
		fallthrough
	case "MYSQL":
		fallthrough
	case "OKTA_DIRECTORY":
		fallthrough
	case "OPAL":
		fallthrough
	case "PAGERDUTY":
		fallthrough
	case "SALESFORCE":
		fallthrough
	case "TAILSCALE":
		fallthrough
	case "TELEPORT":
		fallthrough
	case "WORKDAY":
		*e = AppTypeEnum(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AppTypeEnum: %v", v)
	}
}