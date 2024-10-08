---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "opal_message_channel_list Data Source - terraform-provider-opal"
subcategory: ""
description: |-
  MessageChannelList DataSource
---

# opal_message_channel_list (Data Source)

MessageChannelList DataSource

## Example Usage

```terraform
data "opal_message_channel_list" "my_messagechannellist" {
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Read-Only

- `channels` (Attributes List) (see [below for nested schema](#nestedatt--channels))

<a id="nestedatt--channels"></a>
### Nested Schema for `channels`

Read-Only:

- `id` (String) The ID of the message channel.
- `is_private` (Boolean) A bool representing whether or not the message channel is private.
- `name` (String) The name of the message channel.
- `remote_id` (String) The remote ID of the message channel
- `third_party_provider` (String) The third party provider of the message channel.
