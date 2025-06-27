resource "opal_group" "my_group" {
  admin_owner_id              = "7c86c85d-0651-43e2-a748-d69d658418e8"
  app_id                      = "f454d283-ca87-4a8a-bdbb-df212eca5353"
  custom_request_notification = "Check your email to register your account."
  description                 = "Engineering team Okta group."
  group_leader_user_ids = [
    "23ac9822-9f43-4e31-a31d-6a6109f207ae"
  ]
  group_type = "OPAL_GROUP"
  message_channel_ids = [
    "01f0dea1-52d3-4b76-b362-1ee677e90fd2"
  ]
  name = "mongo-db-prod"
  on_call_schedule_ids = [
    "6cc05350-3da1-4a2e-bbeb-bd4bc4f9b06b"
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
    okta_group_rule = {
      rule_id = "0pr3f7zMZZHPgUoWO0g4"
    }
    snowflake_role = {
      role_id = "01fa7402-01d8-103b-8deb-5f3a0ab7884"
    }
    workday_user_security_group = {
      group_id = "123abc456def"
    }
  }
  request_configurations = [
    {
      allow_requests = true
      auto_approval  = false
      condition = {
        group_ids = [
          "ea350457-6b03-4f86-8810-12ec5b59be85"
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
            "f653097c-5b74-48b8-a26c-33571f9211ff"
          ]
          require_admin_approval   = false
          require_manager_approval = false
        }
      ]
    }
  ]
  require_mfa_to_approve    = false
  risk_sensitivity_override = "CRITICAL"
  visibility                = "GLOBAL"
  visibility_group_ids = [
    "ea22f6cf-8fd4-44e9-b53d-66a5731ab7da"
  ]
}