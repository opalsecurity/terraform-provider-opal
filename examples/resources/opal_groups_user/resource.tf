resource "opal_groups_user" "my_groupsuser" {
  access_level_remote_id = "arn:aws:iam::590304332660:role/AdministratorAccess"
  duration_minutes       = 60
  group_id               = "040eb635-d1ef-4437-9d08-dbe759fb0ef1"
  user_id                = "dad9955d-0f34-49b3-b0ec-0cf728ce9e23"
}