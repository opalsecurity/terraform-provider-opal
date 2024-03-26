---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "opal_uars_list Data Source - terraform-provider-opal"
subcategory: ""
description: |-
  UARSList DataSource
---

# opal_uars_list (Data Source)

UARSList DataSource

## Example Usage

```terraform
data "opal_uars_list" "my_uars_list" {
  cursor    = "...my_cursor..."
  page_size = 6
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `cursor` (String) The pagination cursor value.
- `page_size` (Number) Number of results to return per page. Default is 200.

### Read-Only

- `next` (String) The cursor with which to continue pagination if additional result pages exist.
- `previous` (String) The cursor used to obtain the current result page.
- `results` (Attributes List) (see [below for nested schema](#nestedatt--results))

<a id="nestedatt--results"></a>
### Nested Schema for `results`

Read-Only:

- `deadline` (String) The last day for reviewers to complete their access reviews.
- `name` (String) The name of the UAR.
- `reviewer_assignment_policy` (String) A policy for auto-assigning reviewers. If auto-assignment is on, specific assignments can still be manually adjusted after the access review is started. Default is Manually. must be one of ["MANUALLY", "BY_OWNING_TEAM_ADMIN", "BY_MANAGER"]
- `self_review_allowed` (Boolean) A bool representing whether to present a warning when a user is the only reviewer for themself. Default is False.
- `send_reviewer_assignment_notification` (Boolean) A bool representing whether to send a notification to reviewers when they're assigned a new review. Default is False.
- `time_zone` (String) The time zone name (as defined by the IANA Time Zone database) used in the access review deadline and exported audit report. Default is America/Los_Angeles.
- `uar_id` (String) The ID of the UAR.
- `uar_scope` (Attributes) If set, the access review will only contain resources and groups that match at least one of the filters in scope. (see [below for nested schema](#nestedatt--results--uar_scope))

<a id="nestedatt--results--uar_scope"></a>
### Nested Schema for `results.uar_scope`

Read-Only:

- `admins` (List of String) This access review will include resources and groups who are owned by one of the owners corresponding to the given IDs.
- `names` (List of String) This access review will include resources and groups whose name contains one of the given strings.
- `tags` (Attributes List) This access review will include resources and groups who are tagged with one of the given tags. (see [below for nested schema](#nestedatt--results--uar_scope--tags))

<a id="nestedatt--results--uar_scope--tags"></a>
### Nested Schema for `results.uar_scope.tags`

Read-Only:

- `key` (String) The key of the tag.
- `value` (String) The value of the tag.

