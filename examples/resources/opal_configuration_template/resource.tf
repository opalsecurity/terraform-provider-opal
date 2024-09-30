resource "opal_configuration_template" "my_configurationtemplate" {
  admin_owner_id = "7c86c85d-0651-43e2-a748-d69d658418e8"
  break_glass_user_ids = [
    "363ceb0c-fb02-4f61-9943-5ac9e969aba2"
  ]
  custom_request_notification = "Check your email to register your account."
  linked_audit_message_channel_ids = [
    "85b8103e-608c-4d47-9207-1aa604564cf3"
  ]
  member_oncall_schedule_ids = [
    "b5dab04b-c577-4029-899d-37113cdd854c"
  ]
  name = "Prod AWS Template"
  request_configurations = [
    {
      allow_requests = true
      auto_approval  = false
      condition = {
        group_ids = [
          "1c5f9802-81cc-4f6d-a68f-50913fa8d0d4"
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