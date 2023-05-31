# Changelog

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
