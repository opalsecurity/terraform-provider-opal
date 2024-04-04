terraform {
  required_providers {
    opal = {
      source  = "opalsecurity/opal"
      version = "0.3.0"
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
  # App ids can be retrieved via the Opal web app or via the API (https://docs.opal.dev/reference/getapps)
  id = "b6e7837c-87a3-48b1-9567-ac3a8fee9531"
}

data "opal_app" "aws" {
  id = "9d661e74-c0bf-4ea3-b272-b97a4e0361db"
}

data "opal_app" "opal" {
  # App ids can be retrieved via the Opal web app or via the API (https://docs.opal.dev/reference/getapps)
  id = "aea343c8-23f8-406e-8c55-525b38aad011"
}
data "opal_app" "custom" {
  # App ids can be retrieved via the Opal web app or via the API (https://docs.opal.dev/reference/getapps)
  id = "7f6f2ab7-7c6f-43f4-9c59-3ed6b6e2025b"
}

data "opal_user" "amruth" {
  email = "amruth@opal.dev"
}

resource "opal_owner" "security" {
  name = "Security Team 2"

  user_ids = [
    data.opal_user.amruth.id,
  ]
  description = "Test owner description"
  reviewer_message_channel_id = opal_message_channel.my_messagechannel.id
  # source_group_id = "c95394cf-e1e8-4baa-b850-ab5fd27335dc"
  access_request_escalation_period = 15
}

resource "opal_resource" "sensitive_resource" {
  name = "Sensitive Resource"
  description = "A sensitive resource that should be accessed for on-call only."
  resource_type = "CUSTOM"
  app_id = data.opal_app.custom.id
  admin_owner_id = opal_owner.security.id
  visibility = "LIMITED"
  visibility_group_ids=[]#"c95394cf-e1e8-4baa-b850-ab5fd27335dc"]
  request_configurations=[]
}

resource "opal_message_channel" "my_messagechannel" {
  remote_id            = "C03FJR97276"
  third_party_provider = "SLACK"
}

resource "opal_group" "oncall" {
  name = "On-Call Rotation 2"
  app_id = data.opal_app.opal.id
  group_type = "OPAL_GROUP"
  admin_owner_id = opal_owner.security.id
  message_channel_ids = []
  # message_channel_ids = [opal_message_channel.my_messagechannel.id]
  on_call_schedule_ids = []
  # on_call_schedule_ids = ["d5d9099b-8ab1-4ed5-8ac6-7c874808dda1"]
  visibility = "LIMITED"
  # visibility_group_ids = ["c95394cf-e1e8-4baa-b850-ab5fd27335dc"]
  request_configurations=[]
}

resource "opal_group" "okta" {
  name = "Made up okta group"
  app_id = data.opal_app.okta.id
  group_type = "OKTA_GROUP"
  admin_owner_id = opal_owner.security.id
  message_channel_ids = [] # [opal_message_channel.my_messagechannel.id]
  on_call_schedule_ids = []
  visibility = "GLOBAL"
  visibility_group_ids = []
  remote_info = {
    okta_group = {
      group_id = "dummy value"
    }
  }
  request_configurations = [
    {
      allow_requests= true
      auto_approval = true
      max_duration = 10
      priority = 0
      recommended_duration = 15
      request_template_id = "57fb47fa-afe3-4a4d-9d40-a644ca8279d3"
      require_mfa_to_request =true
      require_support_ticket = true
      reviewer_stages = []
    },
    {
      allow_requests= true
      auto_approval = false
      max_duration = 1
      priority = 1
      recommended_duration = 5
      require_mfa_to_request =true
      require_support_ticket = true
      reviewer_stages = [
        {
          operator = "OR"
          owner_ids = [opal_owner.security.id]
          require_manager_approval = false
        }
      ]
      condition = {
        group_ids = [opal_group.oncall.id]
        role_remote_ids = ["full"]
      }
    },
  ]
}

resource "opal_group" "foobarbaz" {
  name = "foobarbaz"
  app_id = data.opal_app.okta.id
  group_type = "OKTA_GROUP"
  admin_owner_id = opal_owner.security.id
  message_channel_ids = [] # [opal_message_channel.my_messagechannel.id]
  on_call_schedule_ids = []
  visibility = "GLOBAL"
  visibility_group_ids = []
  remote_info = {
    okta_group = {
      group_id = "dummy value #2"
    }
  }
  request_configurations=[]
}

resource "opal_resource" "another_one" {
  name = "Another Resource"
  description = "Another One"
  admin_owner_id = opal_owner.security.id
  resource_type = "AWS_EC2_INSTANCE"
  app_id = data.opal_app.aws.id
  visibility = "LIMITED"
  # visibility_group_ids = ["c95394cf-e1e8-4baa-b850-ab5fd27335dc"]
  remote_info = {
    aws_ec2_instance = {
      account_id = "123456789012"
      instance_id = "i-012abcd34efghinew"
      region = "us-east-1"
    }
  }
  require_mfa_to_approve = false
  require_mfa_to_connect = false
  request_configurations=[]
  # admin_owner_id = opal_owner.security.id
  # visibility = "LIMITED"
  # visibility_group_ids=["c95394cf-e1e8-4baa-b850-ab5fd27335dc"]
}

resource "opal_groups_user" "my_groupsuser" {
  # access_level_remote_id = "arn:aws:iam::590304332660:role/AdministratorAccess"
  duration_minutes       = 120
  group_id               = opal_group.okta.id
  user_id                = data.opal_user.amruth.id
}

resource "opal_group_resource_list" "my_groupresourcelist" {
  group_id = opal_group.okta.id
  resources = [
    {
      resource_id            = opal_resource.sensitive_resource.id
    },
  ]
}

resource "opal_resources_users" "my_resourcesusers" {
  # access_level_remote_id = "arn:aws:iam::590304332660:role/AdministratorAccess"
  duration_minutes       = 4
  resource_id            = opal_resource.sensitive_resource.id
  user_id                = data.opal_user.amruth.id
}

resource "opal_tag" "my_tag" {
  # admin_owner_id = "94bc89fb-d03a-4922-82f2-e3005dda2041" # todo test
  key        = "amruth-key"
  value      = "amruth-value-updated2"
}

resource "opal_group_tag" "my_grouptag" {
  group_id = opal_group.okta.id
  tag_id   = opal_tag.my_tag.id
}

resource "opal_resource_tag" "my_resourcetag" {
  resource_id = opal_resource.sensitive_resource.id
  tag_id      = opal_tag.my_tag.id
}

resource "opal_tag_user" "my_taguser" {
  tag_id      = opal_tag.my_tag.id
  user_id                = data.opal_user.amruth.id
}
