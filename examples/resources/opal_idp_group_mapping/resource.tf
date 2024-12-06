resource "opal_idp_group_mapping" "my_idpgroupmapping" {
  app_resource_id = "4baf8423-db0a-4037-a4cf-f79c60cb67a5"
  mappings = [
    {
      alias    = "...my_alias..."
      group_id = "7734bec8-83b2-425e-a40b-cf91386448f8"
    }
  ]
}