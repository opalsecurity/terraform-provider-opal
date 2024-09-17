resource "opal_group" "my_group" {
  admin_owner_id              = "7c86c85d-0651-43e2-a748-d69d658418e8"
  app_id                      = "f454d283-ca87-4a8a-bdbb-df212eca5353"
  custom_request_notification = "Check your email to register your account."
  description                 = "Engineering team Okta group."
  group_leader_user_ids = [
    "2c3ea0cf-9888-42c2-a96f-64f3fe436133"
  ]
  group_type = "OPAL_GROUP"
  message_channel_ids = [
    "0610f601-dce1-4a71-b5b2-fde35be7f613"
  ]
  name = "mongo-db-prod"
  on_call_schedule_ids = [
    "64cfce08-523c-4560-a3dd-5aa18aa2ee9b"
  ]
  remote_info = {
    active_directory_group = {
      group_id = "01fa7402-01d8-103b-8deb-5f3a0ab7884"
    }
    azure_ad_microsoft_365_group = {
      group_id = "01fa7402-01d8-103b-8deb-5f3a0ab7884"
    }
    azure_ad_security_group = {
      group_id = "01fa7402-01d8-103b-8deb-5f3a0ab7884"
    }
    duo_group = {
      group_id = "DSRD8W89B9DNDBY4RHAC"
    }
    github_team = {
      team_slug = "opal-security"
    }
    gitlab_group = {
      group_id = 898931321
    }
    google_group = {
      group_id = "1y6w882181n7sg"
    }
    ldap_group = {
      group_id = "01fa7402-01d8-103b-8deb-5f3a0ab7884"
    }
    okta_group = {
      group_id = "00gjs33pe8rtmRrp3rd6"
    }
  }
  request_configurations = [
    {
      allow_requests = true
      auto_approval  = false
      condition = {
        group_ids = [
          "eda93353-0444-45a7-b67b-40933f2886c8"
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
            "f161583d-0998-47dc-a5cb-7764b89b1852"
          ]
          require_admin_approval   = false
          require_manager_approval = false
        }
      ]
    }
  ]
  require_mfa_to_approve = false
  visibility             = "GLOBAL"
  visibility_group_ids = [
    "efa92f26-fd60-4cff-987f-3db464ae0915"
  ]
}