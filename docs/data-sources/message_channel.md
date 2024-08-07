---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "opal_message_channel Data Source - terraform-provider-opal"
subcategory: ""
description: |-
  MessageChannel DataSource
---

# opal_message_channel (Data Source)

MessageChannel DataSource

## Example Usage

```terraform
data "opal_message_channel" "my_messagechannel" {
  id = "6670617d-e72a-47f5-a84c-693817ab4860"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) The ID of the message_channel.

### Read-Only

- `is_private` (Boolean) A bool representing whether or not the message channel is private.
- `name` (String) The name of the message channel.
- `remote_id` (String) The remote ID of the message channel
- `third_party_provider` (String) The third party provider of the message channel. must be one of ["SLACK"]
