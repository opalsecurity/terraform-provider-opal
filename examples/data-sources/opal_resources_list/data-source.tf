data "opal_resources_list" "my_resources_list" {
  ancestor_resource_id   = ["4baf8423-db0a-4037-a4cf-f79c60cb67a5"]
  app_id                 = "f454d283-ca87-4a8a-bdbb-df212eca5353"
  cursor                 = "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw"
  page_size              = 200
  parent_resource_id     = ["4baf8423-db0a-4037-a4cf-f79c60cb67a5"]
  requestable_by_user_id = "29827fb8-f2dd-4e80-9576-28e31e9934ac"
  resource_ids = [
    "4baf8423-db0a-4037-a4cf-f79c60cb67a5",
    "1b978423-db0a-4037-a4cf-f79c60cb67b3",
  ]
  resource_name        = "example-name"
  resource_name_match  = "ENDS_WITH"
  resource_type_filter = "AWS_IAM_ROLE"
  tag_ids = [
    "d65a0178-2679-4631-9aed-b1ad19ca8a12"
  ]
}