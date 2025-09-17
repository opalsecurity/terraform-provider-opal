terraform {
  required_providers {
    opal = {
      source  = "opalsecurity/opal"
      version = "3.2.2"
    }
  }
}

variable "auth_token" {
  type = string
}
variable "server_url" {
  type = string
}

provider "opal" {
  bearer_auth = var.auth_token
  server_url  = var.server_url
}


resource "opal_group" "my_group" {
  name       = "my_group"
  app_id     = "c63745be-041f-4dcc-beb9-da6447bbdbae"
  group_type = "GOOGLE_GROUPS_GROUP"
  visibility = "GLOBAL"
  remote_info = {
    google_group = {
      group_id = "037m2jsg218b2wb"
    }
  }
  request_configurations = [
    {
      allow_requests         = true
      auto_approval          = false
      max_duration           = -1
      priority               = 0
      recommended_duration   = -1
      require_mfa_to_request = false
      require_support_ticket = false
      reviewer_stages = [
        {
          operator = "AND"
          owner_ids = [
            "c57db31e-b5af-408a-a55f-134121faa678"
          ]
          require_admin_approval   = false
          require_manager_approval = false
        }
      ]
    },
    # {
    #   allow_requests         = true
    #   auto_approval          = false
    #   max_duration           = 1
    #   priority               = 1
    #   recommended_duration   = 5
    #   require_mfa_to_request = true
    #   require_support_ticket = true
    #   reviewer_stages = [
    #     {
    #       operator                 = "OR"
    #       owner_ids                = [opal_owner.security.id]
    #       require_manager_approval = false
    #     }
    #   ]
    #   condition = {
    #     group_ids       = [opal_group.oncall.id]
    #     role_remote_ids = ["full"]
    #   }
    # },
  ]
}
