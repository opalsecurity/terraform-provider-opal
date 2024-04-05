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
  request_configurations = [
    {
      allow_requests = true
      auto_approval  = false
      condition = {
        group_ids = [
          "ce27d58f-2561-447d-92ea-246933c124e1",
        ]
        role_remote_ids = [
          "...",
        ]
      }
      max_duration           = 120
      priority               = 1
      recommended_duration   = 120
      request_template_id    = "06851574-e50d-40ca-8c78-f72ae6ab4304"
      require_mfa_to_request = false
      require_support_ticket = false
      reviewer_stages = [
        {
          operator = "AND"
          owner_ids = [
            "53ec25e2-8706-4e50-ad65-59d94490f519",
          ]
          require_manager_approval = false
        },
      ]
    },
  ]
  require_mfa_to_approve = false
  visibility             = "GLOBAL"
  visibility_group_ids = [
    "d5bf1886-9ae7-426c-8cc0-18ec506c2a39",
  ]
}