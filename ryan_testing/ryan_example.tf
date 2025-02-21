terraform {
  required_providers {
    opal = {
      source  = "opalsecurity/opal"
      version = "0.31.0"
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
# bearer_auth = "V80X3b33CxErfGgF2FUMwOMNXAK6s6pOnapD6R9F956J_GItbiHoUdmNuwCtACORAo6RJ6WmvmobegxAcpj3YQ=="
  server_url = "http://localhost:3000/v1"
  # TF_VAR_server_url=http://localhost:3000/v1
}

data "opal_user" "ryan" {
    email = "ryan.riddle@opal.dev"
}

data "opal_app" "okta" {
  # App ids can be retrieved via the Opal web app or via the API (https://docs.opal.dev/reference/getapps)
  id = "9c429c9a-8dc0-42ad-806b-d872bd9b7331"
}

data "opal_group" "test_okta" {
  # Group ids can be retrieved via the Opal web app or via the API (https://docs.opal.dev/reference/getgroups)
  id = "406205b4-ce71-48e1-9735-41c0a9a889a0"
}

resource "opal_idp_group_mappings" "okta_test_okta" {
  # IdP group mapping ids can be retrieved via the Opal web app or via the API (https://docs.opal.dev/reference/getidpgroupmappings)
  app_resource_id = "4f8e70e4-b58c-4e6d-ae1f-8b39e164c2bf"
  mappings = [
    {
      group_id = data.opal_group.test_okta.id
      alias = "updates"
    },
    {
      group_id = "d6476f54-884f-4f1a-a03e-aa529b9ef277"
      alias = "updates 2"
    }
  ]
}
resource "opal_idp_group_mappings" "okta_test_okta2" {
  # IdP group mapping ids can be retrieved via the Opal web app or via the API (https://docs.opal.dev/reference/getidpgroupmappings)
  app_resource_id = "3426a446-362e-4ae3-a874-3a12948d8b5e"
  mappings = [
    {
      group_id = data.opal_group.test_okta.id
      hidden = true
      alias = "test"
    }
  ]
}





# data "opal_app" "opal" {
#   # App ids can be retrieved via the Opal web app or via the API (https://docs.opal.dev/reference/getapps)
#   id = "78472573-e422-4d2d-83dc-cd4facff6f2d"
# }
# data "opal_app" "custom" {
#   # App ids can be retrieved via the Opal web app or via the API (https://docs.opal.dev/reference/getapps)
#   id = "ba168eb6-ce81-4d6b-97e8-821bf713458b"
# }