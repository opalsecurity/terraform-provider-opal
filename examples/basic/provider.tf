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
  id = "dbe38d2d-9ce4-4d13-95a4-945716a257b4"
}

data "opal_app" "my_custom_app" {
  id = "03c06479-6ffa-45e1-9f65-cd470ff128b3"
}

data "opal_user" "alice" {
  email = "alice@mycompany.com"
}

data "opal_user" "bob" {
  id = "e5e5ba2b-e126-4699-a8bc-dc186d490b6e"
}

resource "opal_owner" "security" {
  name = "Security Team"

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

resource "opal_group" "oncall" {
  name = "On-Call Rotation"
  group_type = "OPAL_GROUP"
  app_id = data.opal_app.opal.id
  require_mfa_to_approve = true
  admin_owner_id = opal_owner.security.id

  reviewer_stage {
    reviewer {
      id = opal_owner.security.id
    }
  }
}