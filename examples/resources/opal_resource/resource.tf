resource "opal_resource" "my_resource" {
  admin_owner_id = "7c86c85d-0651-43e2-a748-d69d658418e8"
  app_id         = "f454d283-ca87-4a8a-bdbb-df212eca5353"
  description    = "Engineering team Okta role."
  name           = "mongo-db-prod"
  request_configurations = [
    {
      allow_requests = true
      auto_approval  = false
      condition = {
        group_ids = [
          "baed1f8f-1d89-435b-b4fe-c8046cdd06b0",
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
            "c57b8161-728c-4d36-87b3-ad9a192be13a",
          ]
          require_manager_approval = false
        },
      ]
    },
  ]
  require_mfa_to_approve = false
  require_mfa_to_connect = false
  resource_type          = "AWS_IAM_ROLE"
  visibility             = "GLOBAL"
  visibility_group_ids = [
    "2a285fd0-c387-4f2b-b1cd-922ce15f9509",
  ]
}