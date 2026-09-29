terraform {
  required_providers {
    opal = {
      source  = "opalsecurity/opal"
      version = "3.8.0"
    }
  }
}

provider "opal" {
  server_url = "..." # Optional
}