resource "opal_resources_users" "my_resourcesusers" {
  access_level_remote_id = "arn:aws:iam::590304332660:role/AdministratorAccess"
  duration_minutes       = 60
  resource_id            = "b4274372-d8d8-4b2e-964f-eb68ae88222a"
  user_id                = "7d9bb531-a620-4906-bed0-3356e6f799e1"
}