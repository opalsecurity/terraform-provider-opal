resource "opal_group_user" "my_groupuser" {
  access_level_remote_id = "arn:aws:iam::590304332660:role/AdministratorAccess"
  duration_minutes       = "0"
  group_id               = "3d7f3eba-df33-41c3-b087-92a4ec5fd6d7"
  user_id                = "5e2e61d1-ef5c-40fe-be25-453aaabbdc74"
}