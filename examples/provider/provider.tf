terraform {
  required_providers {
    opal = {
      source  = "opalsecurity/opal"
      version = "3.11.16"
    }
  }
}

provider "opal" {
  server_url = "..." # Optional
}