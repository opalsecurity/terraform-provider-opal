resource "opal_group_resource_list" "my_groupresourcelist" {
  group_id = "b8461398-455a-43ef-b4a2-e91ca0c78b64"
  resources = [
    {
      access_level_remote_id = "write"
      resource_id            = "b5a5ca27-0ea3-4d86-9199-2126d57d1fbd"
    },
  ]
}