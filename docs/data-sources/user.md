---
page_title: "opal_user Data Source - terraform-provider-opal"
subcategory: ""
description: |-
  An Opal User data source.
---

# opal_user (Data Source)

An Opal User data source.

## Example Usage

```terraform
package data_sources

data "opal_user" "alice" {
  email = "alice@mycompany.com"
}

data "opal_user" "bob" {
  id = "e5e5ba2b-e126-4699-a8bc-dc186d490b6e"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `email` (String) The email of the user.
- `id` (String) The ID of the user.

### Read-Only

- `name` (String) The name of the user.
- `position` (String) The position of the user.

Please [file a ticket](https://github.com/opalsecurity/terraform-provider-opal/issues) to discuss use cases that are not yet supported in the provider.