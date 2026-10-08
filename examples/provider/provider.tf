terraform {
  required_providers {
    opal = {
      source  = "opalsecurity/opal"
      version = "3.9.0"
    }
  }
}

provider "opal" {
  server_url = "..." # Optional
}