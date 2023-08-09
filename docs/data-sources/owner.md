---
page_title: "opal_owner Data Source - terraform-provider-opal"
subcategory: ""
description: |-
  An Opal owner data source.
---

# opal_owner (Data Source)

An Opal owner data source.

## Example Usage

```terraform
package data_sources

data "opal_owner" "design" {
  name = "Design Owner"
}

data "opal_owner" "devops" {
  id = "e5e5ba2b-e126-4699-a8bc-dc186d490b6e"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `id` (String) The ID of the owner.
- `name` (String) The name of the owner.

Please [file a ticket](https://github.com/opalsecurity/terraform-provider-opal/issues) to discuss use cases that are not yet supported in the provider.