data "opal_group_list" "my_group_list" {
  app_id = "f454d283-ca87-4a8a-bdbb-df212eca5353"
  group_ids = [
    "4baf8423-db0a-4037-a4cf-f79c60cb67a5",
    "1b978423-db0a-4037-a4cf-f79c60cb67b3",
  ]
  group_name             = "example-name"
  group_type_filter      = "OPAL_GROUP"
  page_size              = 200
  remote_id              = "...my_remote_id..."
  requestable            = true
  requestable_by_user_id = "29827fb8-f2dd-4e80-9576-28e31e9934ac"
  tag_ids = [
    "3c34a212-2dee-4019-8657-e4e961fa0e85"
  ]
}