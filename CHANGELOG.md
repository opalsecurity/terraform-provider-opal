# Changelog
## v.3.1.0
- Bumping minor version due to improvements in release cycle
- Improvements to documentation
  
## v.3.0.19
- Add support for Scoped Admin Roles

## v.3.0.17
- Add support for AWS Organizational Units
- Add support for hierarchical data to resources
- Add default values for require_admin_approval and require_manager_approval

## v.3.0.16
- Add support for Access Rules

## v.3.0.15
- Add support for group <> containing group pairs
- Add support for Snowflake Roles to Group creation

## v.3.0.14
- Add support for bundles

## v3.0.12
- Add default values for condition and require_mfa_to_approve fields
- Remove diff suppression on condition field.

## v3.0.11
- Allow sending custom HTTP headers on requests to the Opal backend

## v3.0.10
- Suppress "known after apply" output for risk_sensitivity fields

## v3.0.9
- Add risk sensitivity and overrides
- Add default empty arrays for a few fields (visibility_group_ids, message_channel_ids, member_oncall_schedule_ids, on_call_schedule_ids)

## v3.0.8
- Fix state upgrader when going from v2.x to latest version.

## v3.0.7
- Fix bugs in request configuration validation causing spurious changes. Request configurations should be passed in ascending order of priority.
- Suppress "known after apply" output for optional fields on resources (description, admin_owner_id, ticket_propogation, and require_mfa_to_connect)

## v3.0.4
- Suppress "known after apply" output for optional fields on groups (visibility_group_ids, description, admin_owner_id, group_leader_user_ids, and oncall_schedules)

## v3.0.3
- Fix various bugs in request configuration validation leading to spurious changes and errors

## v3.0.2
- No longer need to explicitly pass the auth token to the Opal provider. The provider will look for the OPAL_AUTH_TOKEN envar and use it if set.
- 404s on refreshing state for Terrafpr, resources will result in the resource being recreated on a subsequent apply rather than erroring.

## v3.0.0

*Compared to 2.0.2*

`opal_group`
- message_channel_ids now required (can provide empty list [])
- on_call_schedule_ids now required (can provide empty list [])
- visibility now required
- audit_message_channel => message_channel_ids (List of string ids)
- on_call_schedule => on_call_schedule_ids (List of string ids)
- visibility_group => visibility_group_ids (List of string ids)
- manage_resources => removed in favor of optional declaration of group <> resource relationship
- resource => moved to separate resource `opal_group_resource_list`
- request_configuration => request_configurations. List of configurations with at minimum a default configuration. Optionally specify extra configurations to apply to targeted groups

`opal_owner`
- user (Block list) => user_ids (List of strings)
- user_ids required instead of optional

`opal_resource`
- admin_owner_id now optional
- visibility now required
- visibility_group => visibility_group_ids (List of string ids)
- request_configurations now required
- request_configuration => request_configurations. List of configurations with at minimum a default configuration. Optionally specify extra configurations to apply to targeted groups

#### New capabilities
`opal_resources_users`
- Grant access to a Resource for a specific User

`opal_group_tag`
- Associate a Group and a Tag

`opal_resource_tag`
- Associate a Resource and a Tag

`opal_tag`
- Create an Opal tag to use with other Opal objects

`opal_tag_user`
- Associate a User and a Tag

## v1.0.4

NEW FEATURES:

- adds support for AWS IAM Identity Center

CHANGES:

- require Go 1.20, up from 1.18

## v1.0.3

NEW FEATURES:

- adds creation support for Gitlab and Teleport

## v1.0.2

BUG FIXES:

- adds a boolean flag to turn management of group <-> resource relationships off by default to avoid accidental access changes

NEW FEATURES:

- adds opal_owner data source

## v1.0.1

BUG FIXES:

- prevents resource / groups created without description to have an immediate diff from default description generation

## v1.0.0

BREAKING CHANGES:

- the `require_manager_approval` attribute was removed in favor of `reviewer_stage`
- the `reviewer` attribute was removed in favor of `reviewer_stage`

NEW FEATURES:

- adds support for multi-stage approvals
- adds support for `on_call_schedules` in group resources. Example:

```terraform
resource "opal_on_call_schedule" "security_oncall_rotation" {
  third_party_provider = "PAGER_DUTY"
  remote_id = "PNXHVAA"
}

# Example group usage
resource "opal_group" "security" {
  // ...

  on_call_schedule {
    id = opal_on_call_schedule.security_oncall_rotation.id
  }
```

## v0.0.4

- Fixes a bug for owner user parsing

## v0.0.3

- Adds data sources for opal apps and users
- Adds a more structured `remote_info` attribute to the resource and group resources in favor of `metadata` and `remote_id`
- Adds support for `require_mfa_to_connect`

## v0.0.2

- Initial release with support for managing owners, groups, and resources as terraform resources.
- v0.0.1 was a pre-release build and should not be used.
