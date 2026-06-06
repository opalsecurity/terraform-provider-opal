resource "opal_resource_custom_access_level" "my_resourcecustomaccesslevel" {
  access_level = {
    access_level_name      = "AdminRole"
    access_level_remote_id = "arn:aws:iam::590304332660:role/AdministratorAccess"
  }
  access_level_remote_id      = "admin"
  policy                      = "...my_policy..."
  requestable_by_default      = true
  resource_id                 = "1b978423-db0a-4037-a4cf-f79c60cb67b3"
  stackable_sensitivity_index = 1
}