---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "opal_resources_access_status Data Source - terraform-provider-opal"
subcategory: ""
description: |-
  ResourcesAccessStatus DataSource
---

# opal_resources_access_status (Data Source)

ResourcesAccessStatus DataSource

## Example Usage

```terraform
data "opal_resources_access_status" "my_resourcesaccessstatus" {
  access_level_remote_id = "...my_access_level_remote_id..."
  cursor                 = "...my_cursor..."
  page_size              = 1
  resource_id            = "56ac5d88-fc9d-4ba6-b988-1a6b52c1c934"
  user_id                = "d25038b6-4047-4648-8eea-2a3bbb602fdd"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `resource_id` (String) The ID of the resource.
- `user_id` (String) The ID of the user.

### Optional

- `access_level_remote_id` (String) The remote ID of the access level that you wish to query for the resource. If omitted, the default access level remote ID value (empty string) is used.
- `cursor` (String) The pagination cursor value.
- `page_size` (Number) Number of results to return per page. Default is 200.

### Read-Only

- `access_level` (Attributes) # Access Level Object
### Description
The `ResourceAccessLevel` object is used to represent the level of access that a user has to a resource or a resource has to a group. The "default" access
level is a `ResourceAccessLevel` object whose fields are all empty strings.

### Usage Example
View the `ResourceAccessLevel` of a resource/user or resource/group pair to see the level of access granted to the resource. (see [below for nested schema](#nestedatt--access_level))
- `expiration_date` (String) The day and time the user's access will expire.
- `status` (String) The status of the user's access to the resource. must be one of ["AUTHORIZED", "REQUESTED", "UNAUTHORIZED"]

<a id="nestedatt--access_level"></a>
### Nested Schema for `access_level`

Read-Only:

- `access_level_name` (String) The human-readable name of the access level.
- `access_level_remote_id` (String) The machine-readable identifier of the access level.

