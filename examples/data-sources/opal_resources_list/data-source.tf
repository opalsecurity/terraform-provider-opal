data "opal_resources_list" "my_resources_list" {
  cursor             = "...my_cursor..."
  page_size          = 5
  parent_resource_id = "a39c4c5c-ef6b-47ff-bf73-4610ff591a73"
  resource_ids = [
    "1b978423-db0a-4037-a4cf-f79c60cb67b3",
  ]
  resource_name        = "...my_resource_name..."
  resource_type_filter = "AWS_IAM_ROLE"
}