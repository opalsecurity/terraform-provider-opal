terraform {
  required_providers {
    opal = {
      source  = "opalsecurity/opal"
      version = "3.0.13"
    }
  }
}

variable "auth_token" {
  type = string
}
variable "server_url" {
  type=string
}

variable "user_email" {
  type = string
}
# variable "okta_app_id" {
#   type=string
# }

# variable "opal_app_id" {
#   type = string
# }

# variable "aws_app_id" {
#   type = string
# }
# variable "on_call_schedule_id" {
#   type=string
# }

variable "custom_app_id" {
  type =string
}

# variable "known_group_id" {
#   type=string
# }

provider "opal" {
  bearer_auth = var.auth_token
  server_url = var.server_url
  http_headers = {
    "X-Example-Header" = "example-value"
  }

}

# data "opal_app" "okta" {
#   # App ids can be retrieved via the Opal web app or via the API (https://docs.opal.dev/reference/getapps)
#   id = var.okta_app_id
# }

# data "opal_app" "aws" {
#   id = var.aws_app_id
# }

# data "opal_app" "opal" {
#   # App ids can be retrieved via the Opal web app or via the API (https://docs.opal.dev/reference/getapps)
#   id = var.opal_app_id
# }
data "opal_app" "custom" {
  # App ids can be retrieved via the Opal web app or via the API (https://docs.opal.dev/reference/getapps)
  id = var.custom_app_id
}

data "opal_user" "amruth" {
  email = var.user_email
}

resource "opal_owner" "security" {
  name = "Security Team 2"

  user_ids = [
    # "8ee65686-6c2c-4fca-a996-8d4b88137e4e",
    data.opal_user.amruth.id
  ]
  description = "Test owner description"
  # reviewer_message_channel_id = opal_message_channel.my_messagechannel.id
  # source_group_id = "c95394cf-e1e8-4baa-b850-ab5fd27335dc"
  # access_request_escalation_period = 15
}

# resource "opal_resource" "sensitive_resource" {
#   name = "Sensitive Resource"
#   description = "A sensitive resource that should be accessed for on-call only."
#   resource_type = "CUSTOM"
#   app_id = data.opal_app.custom.id
#   admin_owner_id = opal_owner.security.id
#   visibility = "LIMITED"
#   request_configurations = [
#     {
#       allow_requests         = true
#       auto_approval          = false
#       priority               = 0
#       recommended_duration   = 30
#       max_duration           = 60
#       require_mfa_to_request = false
#       require_support_ticket = false
#       reviewer_stages = [
#         {
#           require_manager_approval = false
#           owner_ids = [
#             opal_owner.security.id
#           ]
#         }
#       ]
#     }
#   ]

# }

# resource "opal_message_channel" "my_messagechannel" {
#   remote_id            = "C03FJR97276"
#   third_party_provider = "SLACK"
# }

resource "opal_group" "oncall" {
  name = "On-Call Rotation 2"
  app_id = "6a52222a-65b8-41a5-a19a-f728962a576d"
  admin_owner_id = opal_owner.security.id
  group_type = "OKTA_GROUP"
  remote_info =  {
    okta_group = {
      group_id = "00gkls5cfeXzGI4Ep1d7"
    }
  }
  visibility = "GLOBAL"
  request_configurations=[
    {
        allow_requests = true
        auto_approval = true
        require_support_ticket = false
        require_mfa_to_request = true
        priority = 0
        reviewer_stages=[]
    }
  ]
}

# resource "opal_group" "okta" {
#   name = "Made up okta group"
#   app_id = data.opal_app.okta.id
#   group_type = "OKTA_GROUP"
#   admin_owner_id = opal_owner.security.id
#   message_channel_ids = [opal_message_channel.my_messagechannel.id]
#   on_call_schedule_ids = []
#   visibility = "GLOBAL"
#   visibility_group_ids = []
#   remote_info = {
#     okta_group = {
#       group_id = "00gjsjszbR3IqjjfO5d6"
#     }
#   }
#   request_configurations = [
#     {
#       allow_requests= true
#       auto_approval = true
#       max_duration = 10
#       priority = 0
#       recommended_duration = 15
#       request_template_id = "57fb47fa-afe3-4a4d-9d40-a644ca8279d3"
#       require_mfa_to_request =true
#       require_support_ticket = true
#       reviewer_stages = []
#     },
#     {
#       allow_requests= true
#       auto_approval = false
#       max_duration = 1
#       priority = 1
#       recommended_duration = 5
#       require_mfa_to_request =true
#       require_support_ticket = true
#       reviewer_stages = [
#         {
#           operator = "OR"
#           owner_ids = [opal_owner.security.id]
#           require_manager_approval = false
#         }
#       ]
#       condition = {
#         group_ids = [opal_group.oncall.id]
#         # role_remote_ids = ["full"]
#       }
#     },
#   ]
# }

# resource "opal_group" "foobarbaz" {
#   name = "foobarbaz"
#   description = ""
#   app_id = data.opal_app.okta.id
#   group_type = "OKTA_GROUP"
#   admin_owner_id = opal_owner.security.id
#   message_channel_ids = [opal_message_channel.my_messagechannel.id]
#   on_call_schedule_ids = []
#   visibility = "LIMITED"
#   visibility_group_ids = []
#   remote_info = {
#     okta_group = {
#       group_id = "00gjsjszbR3IqjjfO5d6"
#     }
#   }
#   require_mfa_to_approve = false
#   request_configurations=[
#     {
#       allow_requests= true
#       auto_approval = true
#       max_duration = 10
#       priority = 0
#       recommended_duration = 15
#       request_template_id = "57fb47fa-afe3-4a4d-9d40-a644ca8279d3"
#       require_mfa_to_request =true
#       require_support_ticket = true
#       reviewer_stages = []
#     }
#   ]
# }

# resource "opal_resource" "another_one" {
#   name = "Another Resource"
#   description = "Another One"
#   admin_owner_id = opal_owner.security.id
#   resource_type = "AWS_EC2_INSTANCE"
#   app_id = data.opal_app.aws.id
#   visibility = "LIMITED"
#   visibility_group_ids = [opal_group.foobarbaz.id]
#   remote_info = {
#     aws_ec2_instance = {
#       account_id = "123456789012"
#       instance_id = "i-012abcd34efghinew"
#       region = "us-east-1"
#     }
#   }
#   require_mfa_to_approve = false
#   require_mfa_to_connect = false
#   request_configurations=[
#     {
#       is_requestable=true
#       allow_requests= true
#       auto_approval = true
#       priority = 0
#       require_mfa_to_request =false
#       require_support_ticket = false
#       reviewer_stages = []
#     }
#   ]
# }

# resource "opal_group_user" "my_groupsuser" {
#   # access_level_remote_id = "arn:aws:iam::590304332660:role/AdministratorAccess"
#   group_id               = opal_group.okta.id
#   user_id                = data.opal_user.amruth.id
# }

# resource "opal_group_resource_list" "my_groupresourcelist" {
#   group_id = opal_group.okta.id
#   resources = [
#     {
#       resource_id            = opal_resource.sensitive_resource.id
#     },
#   ]
# }

# resource "opal_tag" "my_tag" {
#   # admin_owner_id = "94bc89fb-d03a-4922-82f2-e3005dda2041" # todo test
#   key        = "amruth-key"
#   value      = "amruth-value-updated2"
# }

# resource "opal_group_tag" "my_grouptag" {
#   group_id = opal_group.okta.id
#   tag_id   = opal_tag.my_tag.id
# }

# resource "opal_resource_tag" "my_resourcetag" {
#   resource_id = opal_resource.sensitive_resource.id
#   tag_id      = opal_tag.my_tag.id
# }

# resource "opal_tag_user" "my_taguser" {
#   tag_id      = opal_tag.my_tag.id
#   user_id                = data.opal_user.amruth.id
# }
