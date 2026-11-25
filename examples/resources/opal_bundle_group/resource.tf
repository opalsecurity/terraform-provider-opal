resource "opal_bundle_group" "my_bundlegroup" {
  access_level_name      = "AdministratorAccess"
  access_level_remote_id = "arn:aws:iam::590304332660:role/AdministratorAccess"
  bundle_id              = "32acc112-21ff-4669-91c2-21e27683eaa1"
  group_id               = "72e75a6f-7183-48c5-94ff-6013f213314b"
}