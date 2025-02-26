resource "opal_idp_group_mappings" "my_idpgroupmappings" {
  app_resource_id = "4baf8423-db0a-4037-a4cf-f79c60cb67a5"
  mappings = [
    {
      alias    = "finance-team"
      group_id = "6f99639b-7928-4043-8184-47cbc6766145"
    }
  ]
}