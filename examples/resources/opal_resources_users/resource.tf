resource "opal_resources_users" "my_resourcesusers" {
  access_level_remote_id = "arn:aws:iam::590304332660:role/AdministratorAccess"
  duration_minutes       = 60
  resource_id            = "d9bb531a-6209-406b-ad03-356e6f799e16"
  user_id                = "b00b6f8b-aed1-4f8f-9d89-35bb4fec8046"
}