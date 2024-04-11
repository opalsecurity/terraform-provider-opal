data "opal_resources_list" "my_resources_list" {
  cursor             = "...my_cursor..."
  page_size          = 6
  parent_resource_id = "4eea2a3b-bb60-42fd-9bde-daf753fdfec2"
  resource_ids = [
    "1b978423-db0a-4037-a4cf-f79c60cb67b3",
  ]
  resource_name        = "...my_resource_name..."
  resource_type_filter = "AWS_IAM_ROLE"
}