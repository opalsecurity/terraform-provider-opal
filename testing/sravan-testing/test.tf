terraform {
  required_providers {
    opal = {
      source  = "opalsecurity/opal"
      version = "3.0.2"
    }
  }
}

variable "auth_token" {
  type = string
}
variable "server_url" {
  type=string
}

provider "opal" {
  bearer_auth = var.auth_token
  server_url = var.server_url
}

data "opal_app" "custom" {
  id = "9a47a3a6-0efc-4c19-8094-01c79bb82254"
}

resource "opal_resource" "sensitive_resource" {
  name = "Sensitive Resource"
  description = "A sensitive resource that should be accessed for on-call only."
  resource_type = "CUSTOM"
  app_id = data.opal_app.custom.id
  admin_owner_id = "1967ad3b-5402-4f6f-9dfb-ae6176f1e033"
  visibility = "LIMITED"
  visibility_group_ids=[]#"c95394cf-e1e8-4baa-b850-ab5fd27335dc"]
  request_configurations=[]
}
