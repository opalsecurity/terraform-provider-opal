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
