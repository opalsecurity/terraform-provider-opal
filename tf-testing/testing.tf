terraform {
  required_providers {
    opal = {
      source  = "opalsecurity/opal"
      version = "0.32.2"
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

data "opal_app" "okta" {
  id = "11dad0bf-caf7-4f44-a6a8-05f67c8659db"
}

data "opal_app" "aws" {
  id = "87585b01-a031-42aa-8b18-10f95b231cf4"
}

data "opal_app" "opal" {
  id = "3f4356cc-c817-4ab6-9997-1c1284e315d0"
}
data "opal_app" "github" {
  id = "f55c1d83-7141-43b1-bf24-ffca67558329"
}
data "opal_app" "custom" {
  # App ids can be retrieved via the Opal web app or via the API (https://docs.opal.dev/reference/getapps)
  id = "b3e06823-36ec-4682-9d0a-834870a44813"
}

data "opal_user" "anna" {
  email = "anna@opal.dev"
}
resource "opal_resource" "sensitive_resource" {
  name = "Sensitive Resource"
  description = "A sensitive resource that should be accessed for on-call only."
  resource_type = "CUSTOM"
  app_id = data.opal_app.custom.id
  admin_owner_id = "dc461198-c338-4af9-a675-138415e9d569"
  visibility = "LIMITED"
  visibility_group_ids=[]#"c95394cf-e1e8-4baa-b850-ab5fd27335dc"]
  request_configurations=[
    {
      allow_requests= true
      auto_approval = true
      max_duration = 10
      priority = 0
      recommended_duration = 15
      require_mfa_to_request =true
      require_support_ticket = true
      reviewer_stages = []
    }
  ]
}

resource "opal_group" "terraform_group" {
  name = "Terraform Group"
  app_id = data.opal_app.opal.id
  group_type = "OPAL_GROUP"
  admin_owner_id = "6f56069b-e538-4062-8f3e-9e3414a92634"
  message_channel_ids = []
  # message_channel_ids = [opal_message_channel.my_messagechannel.id]
  on_call_schedule_ids = []
  # on_call_schedule_ids = ["d5d9099b-8ab1-4ed5-8ac6-7c874808dda1"]
  visibility = "LIMITED"
  # visibility_group_ids = ["c95394cf-e1e8-4baa-b850-ab5fd27335dc"]
  request_configurations=[
    {
      allow_requests= true
      auto_approval = true
      max_duration = 10
      priority = 0
      recommended_duration = 15
      require_mfa_to_request =true
      require_support_ticket = true
      reviewer_stages = []
    }
  ]
}

resource "opal_bundle" "terraform_bundle" {
    name = "Terraform Bundle"
    description = "This is a Terraform Bundle"
    admin_owner_id = "6f56069b-e538-4062-8f3e-9e3414a92634"
    visibility = "GLOBAL"
}

resource "opal_bundle" "terraform_bundle_2" {
  name = "Terraform Bundle 2"
  description = "This is a Terraform Bundle 2"
  admin_owner_id = "6f56069b-e538-4062-8f3e-9e3414a92634"
  visibility = "LIMITED"
  visibility_group_ids = [opal_group.terraform_group.id]
}

resource "opal_bundle_group" "terraform_bundle_group" {
  bundle_id = opal_bundle.terraform_bundle.id
  group_id = opal_group.terraform_group.id
}

resource "opal_bundle_resource" "terraform_bundle_resource" {
  bundle_id = opal_bundle.terraform_bundle_2.id
  resource_id = "72e75a6f-7183-48c5-94ff-6013f213314b"
  access_level_remote_id = "pull"
}