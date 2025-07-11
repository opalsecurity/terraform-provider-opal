resource "opal_scoped_role_permission_list" "my_scopedrolepermissionlist" {
  permissions = [
    {
      allow_all       = true
      permission_name = "READ"
      target_ids = [
        "cfc2dc81-05b3-4559-9210-8a3a9661acd2"
      ]
      target_type = "RESOURCE"
    }
  ]
  resource_id = "1b978423-db0a-4037-a4cf-f79c60cb67b3"
}