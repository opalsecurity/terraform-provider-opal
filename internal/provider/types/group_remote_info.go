// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package types

type GroupRemoteInfo struct {
	ActiveDirectoryGroup     *ActiveDirectoryGroup `tfsdk:"active_directory_group"`
	AzureAdMicrosoft365Group *ActiveDirectoryGroup `tfsdk:"azure_ad_microsoft_365_group"`
	AzureAdSecurityGroup     *ActiveDirectoryGroup `tfsdk:"azure_ad_security_group"`
	DuoGroup                 *ActiveDirectoryGroup `tfsdk:"duo_group"`
	GithubTeam               *GithubTeam           `tfsdk:"github_team"`
	GitlabGroup              *ActiveDirectoryGroup `tfsdk:"gitlab_group"`
	GoogleGroup              *ActiveDirectoryGroup `tfsdk:"google_group"`
	LdapGroup                *ActiveDirectoryGroup `tfsdk:"ldap_group"`
	OktaGroup                *ActiveDirectoryGroup `tfsdk:"okta_group"`
}
