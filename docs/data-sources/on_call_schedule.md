---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "opal_on_call_schedule Data Source - terraform-provider-opal"
subcategory: ""
description: |-
  OnCallSchedule DataSource
---

# opal_on_call_schedule (Data Source)

OnCallSchedule DataSource

## Example Usage

```terraform
data "opal_on_call_schedule" "my_oncallschedule" {
  id = "6b6f0baa-3e2c-4f50-88df-6c7ef0f714d9"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) The ID of the on_call_schedule.

### Read-Only

- `name` (String) The name of the on call schedule.
- `remote_id` (String) The remote ID of the on call schedule
- `third_party_provider` (String) The third party provider of the on call schedule. must be one of ["OPSGENIE", "PAGER_DUTY"]

