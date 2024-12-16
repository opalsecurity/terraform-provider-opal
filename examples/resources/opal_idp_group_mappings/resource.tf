resource "opal_idp_group_mappings" "my_idpgroupmappings" {
  app_resource_id = "4baf8423-db0a-4037-a4cf-f79c60cb67a5"
  mappings = [
    {
      alias    = "...my_alias..."
      group_id = "4550782b-1f2d-4a4d-aa02-b9ad69c6f983"
    }
  ]
}