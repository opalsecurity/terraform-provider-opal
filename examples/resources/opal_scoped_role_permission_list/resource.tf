resource "opal_scoped_role_permission_list" "my_scopedrolepermissionlist" {
  permissions = [
    {
      allow_all       = true
      permission_name = "READ"
      target_ids = [
        "a381e7a3-e5e0-4c48-b1d6-4ccb4c191bc1",
        "8294e9c9-deb6-48e9-9c99-da2a1e04a87f",
      ]
      target_type = "RESOURCE"
    }
  ]
  resource_id = "1b978423-db0a-4037-a4cf-f79c60cb67b3"
}