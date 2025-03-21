---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "opal_idp_group_mappings Data Source - terraform-provider-opal"
subcategory: ""
description: |-
  IdpGroupMappings DataSource
---

# opal_idp_group_mappings (Data Source)

IdpGroupMappings DataSource

## Example Usage

```terraform
data "opal_idp_group_mappings" "my_idpgroupmappings" {
  app_resource_id = "4baf8423-db0a-4037-a4cf-f79c60cb67a5"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `app_resource_id` (String) The ID of the Okta app.

### Read-Only

- `mappings` (Attributes List) (see [below for nested schema](#nestedatt--mappings))

<a id="nestedatt--mappings"></a>
### Nested Schema for `mappings`

Read-Only:

- `alias` (String) The alias of the group.
- `group_id` (String) The ID of the group.
- `hidden_from_end_user` (Boolean) A bool representing whether or not the group is hidden from the end user.
