data "opal_resources_list" "my_resources_list" {
  cursor             = "...my_cursor..."
  page_size          = 0
  parent_resource_id = "4d5afe1a-d084-4b03-a516-169f5fefbfba"
  resource_ids = [
    "1b978423-db0a-4037-a4cf-f79c60cb67b3",
  ]
  resource_name        = "...my_resource_name..."
  resource_type_filter = "AWS_IAM_ROLE"
}