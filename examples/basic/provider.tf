terraform {
  required_providers {
    opal = {
      source = "registry.terraform.io/opalsecurity/opal"
    }
  }
}

provider "opal" {
}

resource "opal_owner" "security" {
  name = "Security Team"

  user {
    id = "f865f6a5-5be8-46a0-bc57-b9cadaf4d1e5"
  }
}

resource "opal_resource" "sensitive_resource" {
  name = "Sensitive Resource"
  description = "A sensitive resource that should be accessed for on-call only."

  resource_type = "CUSTOM"
  app_id = "03c06479-6ffa-45e1-9f65-cd470ff128b3" # Taken from the Opal web app.
  admin_owner_id = "${opal_owner.security.id}"

  visibility = "LIMITED"

  visibility_group {
    id = "${opal_group.oncall.id}"
  }

  reviewer {
    id = "${opal_owner.security.id}"
  }
}

resource "opal_group" "oncall" {
  name = "On-Call Rotation"

  group_type = "OPAL_GROUP"
  app_id = "dbe38d2d-9ce4-4d13-95a4-945716a257b4" # Taken from the Opal web app.

  require_mfa_to_approve = true
}
