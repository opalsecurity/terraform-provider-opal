terraform {
  required_providers {
    opal = {
      source  = "opalsecurity/opal"
      version = "3.11.8"
    }
  }
}

provider "opal" {
  server_url = "..." # Optional
}