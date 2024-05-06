resource "opal_group_resource_list" "my_groupresourcelist" {
  group_id = "4baf8423-db0a-4037-a4cf-f79c60cb67a5"
  resources = [
    {
      access_level_remote_id = "write"
      resource_id            = "b5a5ca27-0ea3-4d86-9199-2126d57d1fbd"
    },
  ]
}