terraform {
  required_providers {
    opal = {
      source = "registry.terraform.io/opalsecurity/opal"
    }
  }
}

provider "opal" {
}

data "opal_app" "opal" {
  # App ids can be retrieved via the Opal web app or via the API (https://docs.opal.dev/reference/getapps)
  id = "ba0ec360-af48-4b98-be34-bceb58760626"
}

data "opal_app" "my_custom_app" {
  id = "077f6beb-956e-42be-815c-e0b17ef9077e"
}

data "opal_user" "alice" {
  email = "aashish@opal.dev"
}

data "opal_user" "bob" {
  id = "5c0a330c-dd4c-403b-aca5-888dd847fca4"
}

resource "opal_owner" "security" {
  name = "Andrew's TF owner 4"

  user {
    id = data.opal_user.alice.id
  }

  user {
    id = data.opal_user.bob.id
  }
}

resource "opal_resource" "sensitive_resource" {
  name = "Sensitive Resource"
  description = "A sensitive resource that should be accessed for on-call only."
  resource_type = "CUSTOM"
  app_id = data.opal_app.my_custom_app.id
  admin_owner_id = opal_owner.security.id
  visibility = "LIMITED"

  request_configuration {
    priority = 0
    is_requestable = true
    auto_approval = true
    require_mfa_to_request = false
    max_duration = 60
    recommended_duration = 60
    require_support_ticket = false
    reviewer_stage {
      reviewer {
        id = opal_owner.security.id
      }
    }
  }

  request_configuration {
    priority = 1
    group_ids = ["bd8a3b83-2bac-410d-af5c-6c67263077ea"]
    is_requestable = true
    auto_approval = true
    require_mfa_to_request = false
    max_duration = 60
    recommended_duration = 60
    require_support_ticket = false
    reviewer_stage {
      reviewer {
        id = opal_owner.security.id
      }
    }
  }
}

# resource "opal_group" "oncall" {
#   name = "On-Call Rotation"
#   group_type = "OPAL_GROUP"
#   app_id = data.opal_app.opal.id
#   require_mfa_to_approve = true
#   admin_owner_id = opal_owner.security.id

#   reviewer_stage {
#     reviewer {
#       id = opal_owner.security.id
#     }
#   }
# }
