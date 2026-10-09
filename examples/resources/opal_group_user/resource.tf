resource "opal_group_user" "my_groupuser" {
  access_level_remote_id = "arn:aws:iam::590304332660:role/AdministratorAccess"
  duration_minutes       = 0
  group_id               = "4baf8423-db0a-4037-a4cf-f79c60cb67a5"
  user_id                = "f92aa855-cea9-4814-b9d8-f2a60d3e4a06"
}