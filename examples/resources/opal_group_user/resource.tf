resource "opal_group_user" "my_groupuser" {
  access_level_remote_id = "arn:aws:iam::590304332660:role/AdministratorAccess"
  duration_minutes       = "0"
  group_id               = "1dad9955-d0f3-449b-bb0e-c0cf728ce9e2"
  user_id                = "3d7f3eba-df33-41c3-b087-92a4ec5fd6d7"
}