terraform {
  required_providers {
    opal = {
      version = "0.0.1"
      source = "registry.terraform.io/opalsecurity/opal"
    }
  }
}

provider "opal" {
  base_url = "http://localhost:3000"
}

resource "opal_owner" "test_owner" {
  name = "hi"
  user {
    id = "f865f6a5-5be8-46a0-bc57-b9cadaf4d1e5"
  }
}

resource "opal_resource" "test_resource" {
  name = "hi resource"
  resource_type = "CUSTOM"
  app_id = "03c06479-6ffa-45e1-9f65-cd470ff128b3"
  admin_owner_id = "${opal_owner.test_owner.id}"
}

resource "opal_group" "test_group" {
  name = "hello group"
  group_type = "OPAL_GROUP"
  app_id = "dbe38d2d-9ce4-4d13-95a4-945716a257b4"
}
