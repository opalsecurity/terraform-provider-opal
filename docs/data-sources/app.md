---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "opal_app Data Source - terraform-provider-opal"
subcategory: ""
description: |-
  App DataSource
---

# opal_app (Data Source)

App DataSource

## Example Usage

```terraform
data "opal_app" "my_app" {
  id = "32acc112-21ff-4669-91c2-21e27683eaa1"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) The ID of the app.

### Read-Only

- `admin_owner_id` (String) The ID of the owner of the app.
- `description` (String) A description of the app.
- `name` (String) The name of the app.
- `type` (String) The type of an app.
