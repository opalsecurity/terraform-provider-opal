terraform {
  required_providers {
    opal = {
      source  = "opalsecurity/opal"
      version = "3.0.6"
    }
  }
}

provider "opal" {
  bearer_auth = ""
  server_url = "http://localhost:4000/v1"
}

data "opal_app" "custom" {
  id = "9a47a3a6-0efc-4c19-8094-01c79bb82254"
}

data "opal_app" "opal" {
  id = "f32e15c8-98ce-4c73-846e-478d42fd5755"
}

resource "opal_resource" "sensitive_resource" {
  name = "Sensitive Resource"
  description = "A sensitive resource that should be accessed for on-call only."
  resource_type = "CUSTOM"
  app_id = data.opal_app.custom.id
  admin_owner_id = "1967ad3b-5402-4f6f-9dfb-ae6176f1e033"
  visibility = "LIMITED"
  visibility_group_ids=[]#"c95394cf-e1e8-4baa-b850-ab5fd27335dc"]
  request_configurations = [
    {
      allow_requests = true
      auto_approval  = true
      max_duration           = 120
      priority               = 0
      recommended_duration   = 120
      require_mfa_to_request = false
      require_support_ticket = false
    }
  ]
  risk_sensitivity_override = "HIGH"
}

resource "opal_group" "sensitive_group" {
  name = "Sensitive Group"
  description = "A sensitive group that should be accessed for on-call only."
  app_id = data.opal_app.opal.id
  group_type                = "OPAL_GROUP"
  admin_owner_id = "1967ad3b-5402-4f6f-9dfb-ae6176f1e033"
  visibility = "LIMITED"
  visibility_group_ids=[]#"c95394cf-e1e8-4baa-b850-ab5fd27335dc"]
  request_configurations = [
    {
      allow_requests = true
      auto_approval  = true
      max_duration           = 120
      priority               = 0
      recommended_duration   = 120
      require_mfa_to_request = false
      require_support_ticket = false
    }
  ]
  on_call_schedule_ids = []
  message_channel_ids = []
  risk_sensitivity_override = "MEDIUM"
}


