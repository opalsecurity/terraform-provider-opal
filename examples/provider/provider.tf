terraform {
  required_providers {
    opal = {
      source  = "opalsecurity/opal"
      version = "3.11.13"
    }
  }
}

provider "opal" {
  server_url = "..." # Optional
}