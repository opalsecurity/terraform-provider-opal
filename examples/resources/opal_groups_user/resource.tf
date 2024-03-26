resource "opal_groups_user" "my_groupsuser" {
  access_level_remote_id = "arn:aws:iam::590304332660:role/AdministratorAccess"
  duration_minutes       = 60
  group_id               = "eaf8553e-c25e-4287-86e5-0ad6559d9449"
  user_id                = "0f51937d-5bf1-4886-9ae7-26c0cc018ec5"
}