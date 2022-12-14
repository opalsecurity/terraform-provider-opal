---
page_title: "opal_message_channel Resource - terraform-provider-opal"
subcategory: ""
description: |-
  An Opal MessageChannel resource.
---

# opal_message_channel (Resource)

An Opal MessageChannel resource.

## Example Usage

```terraform
resource "opal_message_channel" "security_channel" {
  third_party_provider = "SLACK"
  remote_id = "C03L80ABS1T"
}

# Example owner usage
resource "opal_owner" "security" {
  // ...

  reviewer_message_channel_id = opal_message_channel.security_channel.id
}

# Example group usage
resource "opal_group" "security" {
  // ...

  audit_message_channel {
    id = opal_message_channel.security_channel.id
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `remote_id` (String) The remote ID of the message_channel.
- `third_party_provider` (String) The provider of the message channel (i.e. SLACK).

### Read-Only

- `id` (String) The ID of the message_channel.

Please [file a ticket](https://github.com/opalsecurity/terraform-provider-opal/issues) to discuss use cases that are not yet supported in the provider.
