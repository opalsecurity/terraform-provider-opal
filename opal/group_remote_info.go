package opal

import (
	"errors"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/opalsecurity/opal-go"
)

// NOTE: Unfortunately, terraform go-sdk does not support nested object types natively. The only work-around
// is to have a schema.ListType with MaxItems=1 as a layer of indirection in between each level of nesting.
// This makes the implementation ugly, but it's thankfully mostly hidden from the client. If nested objects
// are ever natively supported it in the SDK, we should be able to update our code without a need for
// change in the HCL of clients.
func groupRemoteInfoElem() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"active_directory_group": {
				Description: "The remote_info for an Active Directory group.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"group_id": {
							Description: "The id of the Active Directory group.",
							Type:        schema.TypeString,
							Required:    true,
							ForceNew:    true,
						},
					},
				},
			},
			"duo_group": {
				Description: "The remote_info for an Duo Security group.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"group_id": {
							Description: "The id of the Duo Security group.",
							Type:        schema.TypeString,
							Required:    true,
							ForceNew:    true,
						},
					},
				},
			},
			"github_team": {
				Description: "The remote_info for a GitHub team.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"team_slug": {
							Description: "The slug of the GitHub team.",
							Type:        schema.TypeString,
							Required:    true,
							ForceNew:    true,
						},
					},
				},
			},
			"gitlab_group": {
				Description: "The remote_info for a Gitlab group.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"group_id": {
							Description: "The id of the Gitlab group.",
							Type:        schema.TypeString,
							Required:    true,
							ForceNew:    true,
						},
					},
				},
			},
			"google_group": {
				Description: "The remote_info for a Google group.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"group_id": {
							Description: "The id of the Google group.",
							Type:        schema.TypeString,
							Required:    true,
							ForceNew:    true,
						},
					},
				},
			},
			"ldap_group": {
				Description: "The remote_info for a LDAP group.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"group_id": {
							Description: "The id of the LDAP group.",
							Type:        schema.TypeString,
							Required:    true,
							ForceNew:    true,
						},
					},
				},
			},
			"okta_group": {
				Description: "The remote_info for an Okta group.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"group_id": {
							Description: "The id of the Okta group.",
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
func parseGroupRemoteInfo(remoteInfoI interface{}) (*opal.GroupRemoteInfo, error) {
	remoteInfoIList := remoteInfoI.([]interface{})
	if len(remoteInfoIList) != 1 {
		return nil, errors.New("you cannot provide multiple remote_info blobs")
	}

	remoteInfoMap := remoteInfoIList[0].(map[string]interface{})
	if activeDirectoryGroupI, ok := remoteInfoMap["active_directory_group"]; ok {
		activeDirectoryGroupIList := activeDirectoryGroupI.([]interface{})

		if len(activeDirectoryGroupIList) == 1 {
			activeDirectoryGroup := activeDirectoryGroupIList[0].(map[string]any)
			return &opal.GroupRemoteInfo{
				ActiveDirectoryGroup: &opal.GroupRemoteInfoActiveDirectoryGroup{
					GroupId: activeDirectoryGroup["group_id"].(string),
				},
			}, nil
		}
	}
	if duoGroupI, ok := remoteInfoMap["duo_group"]; ok {
		duoGroupIList := duoGroupI.([]interface{})

		if len(duoGroupIList) == 1 {
			duoGroup := duoGroupIList[0].(map[string]any)
			return &opal.GroupRemoteInfo{
				DuoGroup: &opal.GroupRemoteInfoDuoGroup{
					GroupId: duoGroup["group_id"].(string),
				},
			}, nil
		}
	}
	if githubTeamI, ok := remoteInfoMap["github_team"]; ok {
		githubTeamIList := githubTeamI.([]interface{})

		if len(githubTeamIList) == 1 {
			githubTeam := githubTeamIList[0].(map[string]any)
			return &opal.GroupRemoteInfo{
				GithubTeam: &opal.GroupRemoteInfoGithubTeam{
					TeamSlug: githubTeam["team_slug"].(string),
				},
			}, nil
		}
	}
	if gitlabGroupI, ok := remoteInfoMap["gitlab_group"]; ok {
		gitlabGroupIList := gitlabGroupI.([]interface{})

		if len(gitlabGroupIList) == 1 {
			gitlabGroup := gitlabGroupIList[0].(map[string]any)
			return &opal.GroupRemoteInfo{
				GitlabGroup: &opal.GroupRemoteInfoGitlabGroup{
					GroupId: gitlabGroup["group_id"].(string),
				},
			}, nil
		}
	}
	if googleGroupI, ok := remoteInfoMap["google_group"]; ok {
		googleGroupIList := googleGroupI.([]interface{})

		if len(googleGroupIList) == 1 {
			googleGroup := googleGroupIList[0].(map[string]any)
			return &opal.GroupRemoteInfo{
				GoogleGroup: &opal.GroupRemoteInfoGoogleGroup{
					GroupId: googleGroup["group_id"].(string),
				},
			}, nil
		}
	}
	if ldapGroupI, ok := remoteInfoMap["ldap_group"]; ok {
		ldapGroupIList := ldapGroupI.([]interface{})

		if len(ldapGroupIList) == 1 {
			ldapGroup := ldapGroupIList[0].(map[string]any)
			return &opal.GroupRemoteInfo{
				LdapGroup: &opal.GroupRemoteInfoLdapGroup{
					GroupId: ldapGroup["group_id"].(string),
				},
			}, nil
		}
	}
	if oktaGroupI, ok := remoteInfoMap["okta_group"]; ok {
		oktaGroupIList := oktaGroupI.([]interface{})

		if len(oktaGroupIList) == 1 {
			oktaGroup := oktaGroupIList[0].(map[string]any)
			return &opal.GroupRemoteInfo{
				OktaGroup: &opal.GroupRemoteInfoOktaGroup{
					GroupId: oktaGroup["group_id"].(string),
				},
			}, nil
		}
	}

	return nil, errors.New("could not find supported remote_info type")
}
