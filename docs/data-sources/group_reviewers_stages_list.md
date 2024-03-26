---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "opal_group_reviewers_stages_list Data Source - terraform-provider-opal"
subcategory: ""
description: |-
  GroupReviewersStagesList DataSource
---

# opal_group_reviewers_stages_list (Data Source)

GroupReviewersStagesList DataSource

## Example Usage

```terraform
data "opal_group_reviewers_stages_list" "my_groupreviewersstageslist" {
  group_id = "f4c67e89-5185-4d4b-901a-07642ba38bc7"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `group_id` (String) The ID of the group.

### Read-Only

- `data` (Attributes List) The reviewer stages for this group. (see [below for nested schema](#nestedatt--data))

<a id="nestedatt--data"></a>
### Nested Schema for `data`

Read-Only:

- `operator` (String) The operator of the reviewer stage. must be one of ["AND", "OR"]
- `owner_ids` (List of String)
- `require_manager_approval` (Boolean) Whether this reviewer stage should require manager approval.

