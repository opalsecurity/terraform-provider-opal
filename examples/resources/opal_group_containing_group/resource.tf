resource "opal_group_containing_group" "my_groupcontaininggroup" {
  access_level_remote_id = "arn:aws:iam::590304332660:role/ReadOnlyAccess"
  containing_group_id    = "f454d283-ca87-4a8a-bdbb-df212eca5353"
  duration_minutes       = 120
  group_id               = "4baf8423-db0a-4037-a4cf-f79c60cb67a5"
}