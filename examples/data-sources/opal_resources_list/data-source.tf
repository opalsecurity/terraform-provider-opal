data "opal_resources_list" "my_resources_list" {
  ancestor_resource_id = ["4baf8423-db0a-4037-a4cf-f79c60cb67a5"]
  cursor               = "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw"
  page_size            = 200
  parent_resource_id   = ["4baf8423-db0a-4037-a4cf-f79c60cb67a5"]
  resource_ids = [
    "1b978423-db0a-4037-a4cf-f79c60cb67b3"
  ]
  resource_name        = "example-name"
  resource_type_filter = "AWS_IAM_ROLE"
}