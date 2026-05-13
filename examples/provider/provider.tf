terraform {
  required_providers {
    opal = {
      source  = "opalsecurity/opal"
      version = "3.5.2"
    }
  }
}

provider "opal" {
  server_url = "..." # Optional
}