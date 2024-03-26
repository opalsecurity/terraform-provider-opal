resource "opal_groups_user" "my_groupsuser" {
  access_level_remote_id = "arn:aws:iam::590304332660:role/AdministratorAccess"
  duration_minutes       = 60
  group_id               = "0f51937d-5bf1-4886-9ae7-26c0cc018ec5"
  user_id                = "06c2a39b-8461-4398-855a-3ef74a2e91ca"
}