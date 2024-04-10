data "opal_resources_list" "my_resources_list" {
  cursor             = "...my_cursor..."
  page_size          = 2
  parent_resource_id = "84eea2a3-bbb6-402f-9dbd-edaf753fdfec"
  resource_ids = [
    "1b978423-db0a-4037-a4cf-f79c60cb67b3",
  ]
  resource_name        = "...my_resource_name..."
  resource_type_filter = "AWS_IAM_ROLE"
}