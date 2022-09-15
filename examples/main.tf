terraform {
  required_providers {
    opal = {
      version = "0.0.1"
      source = "registry.terraform.io/opalsecurity/opal"
    }
  }
}

provider "opal" {
  base_url = "http://localhost:3000/api"
}
