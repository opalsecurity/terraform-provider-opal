---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "opal_on_call_schedule_list Data Source - terraform-provider-opal"
subcategory: ""
description: |-
  OnCallScheduleList DataSource
---

# opal_on_call_schedule_list (Data Source)

OnCallScheduleList DataSource

## Example Usage

```terraform
data "opal_on_call_schedule_list" "my_oncallschedule_list" {
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Read-Only

- `on_call_schedules` (Attributes List) (see [below for nested schema](#nestedatt--on_call_schedules))

<a id="nestedatt--on_call_schedules"></a>
### Nested Schema for `on_call_schedules`

Read-Only:

- `id` (String) The ID of the on-call schedule.
- `name` (String) The name of the on call schedule.
- `remote_id` (String) The remote ID of the on call schedule
- `third_party_provider` (String) The third party provider of the on call schedule.
