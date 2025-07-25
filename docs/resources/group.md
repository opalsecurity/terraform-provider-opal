---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "opal_group Resource - terraform-provider-opal"
subcategory: ""
description: |-
  Group Resource
---

# opal_group (Resource)

Group Resource

## Example Usage

```terraform
resource "opal_group" "my_group" {
  admin_owner_id              = "7c86c85d-0651-43e2-a748-d69d658418e8"
  app_id                      = "f454d283-ca87-4a8a-bdbb-df212eca5353"
  custom_request_notification = "Check your email to register your account."
  description                 = "Engineering team Okta group."
  group_leader_user_ids = [
    "23ac9822-9f43-4e31-a31d-6a6109f207ae"
  ]
  group_type = "OPAL_GROUP"
  message_channel_ids = [
    "01f0dea1-52d3-4b76-b362-1ee677e90fd2"
  ]
  name = "mongo-db-prod"
  on_call_schedule_ids = [
    "6cc05350-3da1-4a2e-bbeb-bd4bc4f9b06b"
  ]
  remote_info = {
    active_directory_group = {
      group_id = "01fa7402-01d8-103b-8deb-5f3a0ab7884"
    }
    azure_ad_microsoft_365_group = {
      group_id = "01fa7402-01d8-103b-8deb-5f3a0ab7884"
    }
    azure_ad_security_group = {
      group_id = "01fa7402-01d8-103b-8deb-5f3a0ab7884"
    }
    duo_group = {
      group_id = "DSRD8W89B9DNDBY4RHAC"
    }
    github_team = {
      team_slug = "opal-security"
    }
    gitlab_group = {
      group_id = 898931321
    }
    google_group = {
      group_id = "1y6w882181n7sg"
    }
    ldap_group = {
      group_id = "01fa7402-01d8-103b-8deb-5f3a0ab7884"
    }
    okta_group = {
      group_id = "00gjs33pe8rtmRrp3rd6"
    }
    okta_group_rule = {
      rule_id = "0pr3f7zMZZHPgUoWO0g4"
    }
    snowflake_role = {
      role_id = "01fa7402-01d8-103b-8deb-5f3a0ab7884"
    }
    workday_user_security_group = {
      group_id = "123abc456def"
    }
  }
  request_configurations = [
    {
      allow_requests = true
      auto_approval  = false
      condition = {
        group_ids = [
          "ea350457-6b03-4f86-8810-12ec5b59be85"
        ]
        role_remote_ids = [
          "..."
        ]
      }
      max_duration           = 120
      priority               = 1
      recommended_duration   = 120
      request_template_id    = "06851574-e50d-40ca-8c78-f72ae6ab4304"
      require_mfa_to_request = false
      require_support_ticket = false
      reviewer_stages = [
        {
          operator = "AND"
          owner_ids = [
            "f653097c-5b74-48b8-a26c-33571f9211ff"
          ]
          require_admin_approval   = false
          require_manager_approval = false
        }
      ]
    }
  ]
  require_mfa_to_approve    = false
  risk_sensitivity_override = "CRITICAL"
  visibility                = "GLOBAL"
  visibility_group_ids = [
    "ea22f6cf-8fd4-44e9-b53d-66a5731ab7da"
  ]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `app_id` (String) The ID of the app for the group. Requires replacement if changed.
- `group_type` (String) The type of the group. must be one of ["ACTIVE_DIRECTORY_GROUP", "AWS_SSO_GROUP", "DATABRICKS_ACCOUNT_GROUP", "DUO_GROUP", "GIT_HUB_TEAM", "GIT_LAB_GROUP", "GOOGLE_GROUPS_GROUP", "GOOGLE_GROUPS_GKE_GROUP", "LDAP_GROUP", "OKTA_GROUP", "OKTA_GROUP_RULE", "TAILSCALE_GROUP", "OPAL_GROUP", "OPAL_ACCESS_RULE", "AZURE_AD_SECURITY_GROUP", "AZURE_AD_MICROSOFT_365_GROUP", "CONNECTOR_GROUP", "SNOWFLAKE_ROLE", "WORKDAY_USER_SECURITY_GROUP"]; Requires replacement if changed.
- `name` (String) The name of the remote group.
- `request_configurations` (Attributes List) The request configuration list of the configuration template. If not provided, the default request configuration will be used. (see [below for nested schema](#nestedatt--request_configurations))
- `visibility` (String) The visibility level of the entity. must be one of ["GLOBAL", "LIMITED"]

### Optional

- `admin_owner_id` (String) The ID of the owner of the group.
- `custom_request_notification` (String) Custom request notification sent upon request approval.
- `description` (String) A description of the remote group.
- `group_leader_user_ids` (Set of String) A list of User IDs for the group leaders of the group
- `message_channel_ids` (Set of String)
- `on_call_schedule_ids` (Set of String)
- `remote_info` (Attributes) Information that defines the remote group. This replaces the deprecated remote_id and metadata fields. Requires replacement if changed. (see [below for nested schema](#nestedatt--remote_info))
- `require_mfa_to_approve` (Boolean) A bool representing whether or not to require MFA for reviewers to approve requests for this group. Default: false
- `risk_sensitivity_override` (String) Indicates the level of potential impact misuse or unauthorized access may incur. must be one of ["UNKNOWN", "CRITICAL", "HIGH", "MEDIUM", "LOW", "NONE"]
- `visibility_group_ids` (Set of String)

### Read-Only

- `group_binding_id` (String) The ID of the associated group binding.
- `id` (String) The ID of the group.
- `message_channels` (Attributes) The audit and reviewer message channels attached to the group. (see [below for nested schema](#nestedatt--message_channels))
- `oncall_schedules` (Attributes) The on call schedules attached to the group. (see [below for nested schema](#nestedatt--oncall_schedules))
- `remote_name` (String) The name of the remote.
- `risk_sensitivity` (String) The risk sensitivity level for the group. When an override is set, this field will match that. must be one of ["UNKNOWN", "CRITICAL", "HIGH", "MEDIUM", "LOW", "NONE"]

<a id="nestedatt--request_configurations"></a>
### Nested Schema for `request_configurations`

Optional:

- `allow_requests` (Boolean) A bool representing whether or not to allow requests for this resource. Not Null
- `auto_approval` (Boolean) A bool representing whether or not to automatically approve requests for this resource. Not Null
- `condition` (Attributes) (see [below for nested schema](#nestedatt--request_configurations--condition))
- `max_duration` (Number) The maximum duration for which the resource can be requested (in minutes).
- `priority` (Number) The priority of the request configuration. Not Null
- `recommended_duration` (Number) The recommended duration for which the resource should be requested (in minutes). -1 represents an indefinite duration.
- `request_template_id` (String) The ID of the associated request template.
- `require_mfa_to_request` (Boolean) A bool representing whether or not to require MFA for requesting access to this resource. Not Null
- `require_support_ticket` (Boolean) A bool representing whether or not access requests to the resource require an access ticket. Not Null
- `reviewer_stages` (Attributes List) The list of reviewer stages for the request configuration. (see [below for nested schema](#nestedatt--request_configurations--reviewer_stages))

<a id="nestedatt--request_configurations--condition"></a>
### Nested Schema for `request_configurations.condition`

Optional:

- `group_ids` (Set of String) The list of group IDs to match.
- `role_remote_ids` (Set of String) The list of role remote IDs to match.


<a id="nestedatt--request_configurations--reviewer_stages"></a>
### Nested Schema for `request_configurations.reviewer_stages`

Optional:

- `operator` (String) The operator of the reviewer stage. Admin and manager approval are also treated as reviewers. Default: "AND"; must be one of ["AND", "OR"]
- `owner_ids` (Set of String) Not Null
- `require_admin_approval` (Boolean) Whether this reviewer stage should require admin approval. Default: false
- `require_manager_approval` (Boolean) Whether this reviewer stage should require manager approval. Default: false



<a id="nestedatt--remote_info"></a>
### Nested Schema for `remote_info`

Optional:

- `active_directory_group` (Attributes) Remote info for Active Directory group. Requires replacement if changed. (see [below for nested schema](#nestedatt--remote_info--active_directory_group))
- `azure_ad_microsoft_365_group` (Attributes) Remote info for Microsoft Entra ID Microsoft 365 group. Requires replacement if changed. (see [below for nested schema](#nestedatt--remote_info--azure_ad_microsoft_365_group))
- `azure_ad_security_group` (Attributes) Remote info for Microsoft Entra ID Security group. Requires replacement if changed. (see [below for nested schema](#nestedatt--remote_info--azure_ad_security_group))
- `duo_group` (Attributes) Remote info for Duo Security group. Requires replacement if changed. (see [below for nested schema](#nestedatt--remote_info--duo_group))
- `github_team` (Attributes) Remote info for GitHub team. Requires replacement if changed. (see [below for nested schema](#nestedatt--remote_info--github_team))
- `gitlab_group` (Attributes) Remote info for Gitlab group. Requires replacement if changed. (see [below for nested schema](#nestedatt--remote_info--gitlab_group))
- `google_group` (Attributes) Remote info for Google group. Requires replacement if changed. (see [below for nested schema](#nestedatt--remote_info--google_group))
- `ldap_group` (Attributes) Remote info for LDAP group. Requires replacement if changed. (see [below for nested schema](#nestedatt--remote_info--ldap_group))
- `okta_group` (Attributes) Remote info for Okta Directory group. Requires replacement if changed. (see [below for nested schema](#nestedatt--remote_info--okta_group))
- `okta_group_rule` (Attributes) Remote info for Okta Directory group rule. Requires replacement if changed. (see [below for nested schema](#nestedatt--remote_info--okta_group_rule))
- `snowflake_role` (Attributes) Remote info for Snowflake role. Requires replacement if changed. (see [below for nested schema](#nestedatt--remote_info--snowflake_role))
- `workday_user_security_group` (Attributes) Remote info for Workday User Security group. Requires replacement if changed. (see [below for nested schema](#nestedatt--remote_info--workday_user_security_group))

<a id="nestedatt--remote_info--active_directory_group"></a>
### Nested Schema for `remote_info.active_directory_group`

Optional:

- `group_id` (String) The id of the Google group. Not Null; Requires replacement if changed.


<a id="nestedatt--remote_info--azure_ad_microsoft_365_group"></a>
### Nested Schema for `remote_info.azure_ad_microsoft_365_group`

Optional:

- `group_id` (String) The id of the Microsoft Entra ID Microsoft 365 group. Not Null; Requires replacement if changed.


<a id="nestedatt--remote_info--azure_ad_security_group"></a>
### Nested Schema for `remote_info.azure_ad_security_group`

Optional:

- `group_id` (String) The id of the Microsoft Entra ID Security group. Not Null; Requires replacement if changed.


<a id="nestedatt--remote_info--duo_group"></a>
### Nested Schema for `remote_info.duo_group`

Optional:

- `group_id` (String) The id of the Duo Security group. Not Null; Requires replacement if changed.


<a id="nestedatt--remote_info--github_team"></a>
### Nested Schema for `remote_info.github_team`

Optional:

- `team_slug` (String) The slug of the GitHub team. Not Null; Requires replacement if changed.


<a id="nestedatt--remote_info--gitlab_group"></a>
### Nested Schema for `remote_info.gitlab_group`

Optional:

- `group_id` (String) The id of the Gitlab group. Not Null; Requires replacement if changed.


<a id="nestedatt--remote_info--google_group"></a>
### Nested Schema for `remote_info.google_group`

Optional:

- `group_id` (String) The id of the Google group. Not Null; Requires replacement if changed.


<a id="nestedatt--remote_info--ldap_group"></a>
### Nested Schema for `remote_info.ldap_group`

Optional:

- `group_id` (String) The id of the LDAP group. Not Null; Requires replacement if changed.


<a id="nestedatt--remote_info--okta_group"></a>
### Nested Schema for `remote_info.okta_group`

Optional:

- `group_id` (String) The id of the Okta Directory group. Not Null; Requires replacement if changed.


<a id="nestedatt--remote_info--okta_group_rule"></a>
### Nested Schema for `remote_info.okta_group_rule`

Optional:

- `rule_id` (String) The id of the Okta group rule. Not Null; Requires replacement if changed.


<a id="nestedatt--remote_info--snowflake_role"></a>
### Nested Schema for `remote_info.snowflake_role`

Optional:

- `role_id` (String) The id of the Snowflake role. Not Null; Requires replacement if changed.


<a id="nestedatt--remote_info--workday_user_security_group"></a>
### Nested Schema for `remote_info.workday_user_security_group`

Optional:

- `group_id` (String) The id of the Workday User Security group. Not Null; Requires replacement if changed.



<a id="nestedatt--message_channels"></a>
### Nested Schema for `message_channels`

Read-Only:

- `channels` (Attributes List) (see [below for nested schema](#nestedatt--message_channels--channels))

<a id="nestedatt--message_channels--channels"></a>
### Nested Schema for `message_channels.channels`

Read-Only:

- `id` (String) The ID of the message channel.
- `is_private` (Boolean) A bool representing whether or not the message channel is private.
- `name` (String) The name of the message channel.
- `remote_id` (String) The remote ID of the message channel
- `third_party_provider` (String) The third party provider of the message channel. must be "SLACK"



<a id="nestedatt--oncall_schedules"></a>
### Nested Schema for `oncall_schedules`

Read-Only:

- `id` (String) The ID of the on-call schedule.
- `name` (String) The name of the on call schedule.
- `remote_id` (String) The remote ID of the on call schedule
- `third_party_provider` (String) The third party provider of the on call schedule. must be one of ["OPSGENIE", "PAGER_DUTY"]

## Import

Import is supported using the following syntax:

```shell
terraform import opal_group.my_opal_group "32acc112-21ff-4669-91c2-21e27683eaa1"
```
