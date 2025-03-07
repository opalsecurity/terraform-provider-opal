---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "opal_group_user Resource - terraform-provider-opal"
subcategory: ""
description: |-
  GroupUser Resource
---

# opal_group_user (Resource)

GroupUser Resource

## Example Usage

```terraform
resource "opal_group_user" "my_groupuser" {
  access_level_remote_id = "arn:aws:iam::590304332660:role/AdministratorAccess"
  duration_minutes       = 0
  group_id               = "4baf8423-db0a-4037-a4cf-f79c60cb67a5"
  user_id                = "f92aa855-cea9-4814-b9d8-f2a60d3e4a06"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `group_id` (String) The ID of the group. Requires replacement if changed.
- `user_id` (String) The ID of the user to add. Requires replacement if changed.

### Optional

- `access_level_remote_id` (String) The remote ID of the access level to grant to this user. If omitted, the default access level remote ID value (empty string) is used. Requires replacement if changed.
- `duration_minutes` (Number) Must be set to 0. Any nonzerovalue in terraform does not make sense. Default: 0; must be "0"; Requires replacement if changed.

### Read-Only

- `access_level` (Attributes) # Access Level Object
### Description
The `GroupAccessLevel` object is used to represent the level of access that a user has to a group or a group has to a group. The "default" access
level is a `GroupAccessLevel` object whose fields are all empty strings.

### Usage Example
View the `GroupAccessLevel` of a group/user or group/group pair to see the level of access granted to the group. (see [below for nested schema](#nestedatt--access_level))
- `email` (String) The user's email.
- `expiration_date` (String) The day and time the user's access will expire.
- `full_name` (String) The user's full name.
- `propagation_status` (Attributes) The state of whether the push action was propagated to the remote system. If this is null, the access was synced from the remote system. (see [below for nested schema](#nestedatt--propagation_status))

<a id="nestedatt--access_level"></a>
### Nested Schema for `access_level`

Read-Only:

- `access_level_name` (String) The human-readable name of the access level.
- `access_level_remote_id` (String) The machine-readable identifier of the access level.


<a id="nestedatt--propagation_status"></a>
### Nested Schema for `propagation_status`

Read-Only:

- `status` (String) The status of whether the user has been synced to the group or resource in the remote system. must be one of ["SUCCESS", "ERR_REMOTE_INTERNAL_ERROR", "ERR_REMOTE_USER_NOT_FOUND", "ERR_REMOTE_USER_NOT_LINKED", "ERR_REMOTE_RESOURCE_NOT_FOUND", "ERR_REMOTE_THROTTLE", "ERR_NOT_AUTHORIZED_TO_QUERY_RESOURCE", "ERR_REMOTE_PROVISIONING_VIA_IDP_FAILED", "ERR_IDP_EMAIL_UPDATE_CONFLICT", "ERR_TIMEOUT", "ERR_UNKNOWN", "ERR_OPAL_INTERNAL_ERROR", "ERR_ORG_READ_ONLY", "ERR_OPERATION_UNSUPPORTED", "PENDING", "PENDING_MANUAL_PROPAGATION", "PENDING_TICKET_CREATION", "ERR_TICKET_CREATION_SKIPPED", "ERR_DRY_RUN_MODE_ENABLED", "ERR_HR_IDP_PROVIDER_NOT_LINKED", "ERR_REMOTE_UNRECOVERABLE_ERROR"]
