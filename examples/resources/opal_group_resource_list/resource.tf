resource "opal_group_resource_list" "my_groupresourcelist" {
  group_id = "d5bf1886-9ae7-426c-8cc0-18ec506c2a39"
  resources = [
    {
      access_level_remote_id = "write"
      resource_id            = "b5a5ca27-0ea3-4d86-9199-2126d57d1fbd"
    },
  ]
}