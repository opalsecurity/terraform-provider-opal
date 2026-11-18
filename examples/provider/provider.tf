terraform {
  required_providers {
    opal = {
      source  = "opalsecurity/opal"
      version = "3.11.10"
    }
  }
}

provider "opal" {
  server_url = "..." # Optional
}