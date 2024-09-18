resource "opal_configuration_template" "my_configurationtemplate" {
  admin_owner_id = "7c86c85d-0651-43e2-a748-d69d658418e8"
  break_glass_user_ids = [
    "346033c2-e0ba-40cc-9f5b-b0827f4601b9"
  ]
  custom_request_notification = "Check your email to register your account."
  linked_audit_message_channel_ids = [
    "8058be8e-1c01-436e-a690-783c4d24b732"
  ]
  member_oncall_schedule_ids = [
    "bf57d7aa-bd0e-446b-9c15-f7d7d0222969"
  ]
  name = "Prod AWS Template"
  request_configurations = [
    {
      allow_requests = true
      auto_approval  = false
      condition = {
        group_ids = [
          "16c559ff-9787-4082-8841-ac5cdfb6bde6"
        ]
        role_remote_ids = [
          "..."
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
            "bb3e66e9-5012-49e8-a32e-f155d7a63933"
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
    enabled_on_grant      = false
    enabled_on_revocation = true
    ticket_project_id     = "...my_ticket_project_id..."
    ticket_provider       = "LINEAR"
  }
  visibility = {
    visibility = "GLOBAL"
    visibility_group_ids = [
      "4dc2e0ea-6466-44ed-b9b7-d9c8803abe57"
    ]
  }
}