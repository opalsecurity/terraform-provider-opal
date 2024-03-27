resource "opal_resources_users" "my_resourcesusers" {
  access_level_remote_id = "arn:aws:iam::590304332660:role/AdministratorAccess"
  duration_minutes       = 60
  resource_id            = "242f2e30-05dd-4a20-81a0-d761a93988b1"
  user_id                = "ee194539-e33f-4dc8-8a5e-4af3e79e632e"
}