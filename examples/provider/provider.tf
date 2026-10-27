terraform {
  required_providers {
    opal = {
      source  = "opalsecurity/opal"
      version = "3.10.3"
    }
  }
}

provider "opal" {
  server_url = "..." # Optional
}