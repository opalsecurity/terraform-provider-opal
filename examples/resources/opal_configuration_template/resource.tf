resource "opal_configuration_template" "my_configurationtemplate" {
  admin_owner_id = "7c86c85d-0651-43e2-a748-d69d658418e8"
  break_glass_user_ids = [
    "37cb7e41-12ba-46da-92ff-030abe0450b1",
    "37cb7e41-12ba-46da-92ff-030abe0450b2",
  ]
  custom_request_notification = "Check your email to register your account."
  linked_audit_message_channel_ids = [
    "37cb7e41-12ba-46da-92ff-030abe0450b1",
    "37cb7e41-12ba-46da-92ff-030abe0450b2",
  ]
  member_on_call_schedule_ids = [
    "37cb7e41-12ba-46da-92ff-030abe0450b1",
    "37cb7e41-12ba-46da-92ff-030abe0450b2",
  ]
  name = "Prod AWS Template"
  request_configurations = [
    {
      allow_requests = true
      auto_approval  = false
      condition = {
        group_ids = [
          "1b978423-db0a-4037-a4cf-f79c60cb67b3",
        ]
        role_remote_ids = [
          "arn:aws:iam::590304332660:role/AdministratorAccess",
        ]
      }
      extensions_duration_in_minutes = 120
      max_duration                   = 120
      priority                       = 1
      recommended_duration           = 120
      request_template_id            = "06851574-e50d-40ca-8c78-f72ae6ab4304"
      require_mfa_to_request         = false
      require_support_ticket         = false
      reviewer_stages = [
        {
          operator = "AND"
          owner_ids = [
            "b36e5198-3e15-4769-a321-00db76ac9873"
          ]
          require_admin_approval   = false
          require_manager_approval = false
        }
      ]
    }
  ]
  require_mfa_to_approve = false
  require_mfa_to_connect = false
  ticket_propagation = {
    enabled_on_grant      = true
    enabled_on_revocation = false
    ticket_project_id     = "...my_ticket_project_id..."
    ticket_provider       = "LINEAR"
  }
  visibility = {
    visibility = "GLOBAL"
    visibility_group_ids = [
      "4cee664d-9798-40ae-97ab-eeb66b726920"
    ]
  }
}