data "opal_group_list" "my_group_list" {
  group_ids = [
    "1b978423-db0a-4037-a4cf-f79c60cb67b3"
  ]
  group_name        = "example-name"
  group_type_filter = "OPAL_GROUP"
  page_size         = 200
}