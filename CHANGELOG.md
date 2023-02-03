# Changelog

## v0.0.17
- Add support for `on_call_schedules` in group resources. Example: 

```terraform
resource "opal_on_call_schedule" "security_oncall_rotation" {
  third_party_provider = "PAGER_DUTY"
  remote_id = "PNXHVAA"
}

# Example group usage
resource "opal_group" "security" {
  // ...

  on_call_schedule {
    id = "opal_on_call_schedule.security_oncall_rotation.id
  }
```

## v0.0.2
- Initial release with support for managing owners, groups, and resources as terraform resources.
- v0.0.1 was a pre-release build and should not be used.

## v0.0.3
- Adds data sources for opal apps and users
- Adds a more structured `remote_info` attribute to the resource and group resources in favor of `metadata` and `remote_id`
- Adds support for `require_mfa_to_connect`

## v0.0.4
- Fixes a bug for owner user parsing
