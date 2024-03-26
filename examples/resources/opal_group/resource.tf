resource "opal_group" "my_group" {
  admin_owner_id = "7c86c85d-0651-43e2-a748-d69d658418e8"
  app_id         = "f454d283-ca87-4a8a-bdbb-df212eca5353"
  description    = "Engineering team Okta group."
  group_type     = "OPAL_GROUP"
  message_channel_ids = [
    "e931861e-f161-4b54-bb18-98e51c0009dd",
  ]
  name = "mongo-db-prod"
  on_call_schedule_ids = [
    "10c42893-326c-48d3-b654-3ad1053f385d",
  ]
  require_mfa_to_approve = false
  visibility             = "GLOBAL"
  visibility_group_ids = [
    "7ce27d58-f256-4147-992e-a246933c124e",
  ]
}