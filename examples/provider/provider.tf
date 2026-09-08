terraform {
  required_providers {
    opal = {
      source  = "opalsecurity/opal"
      version = "3.6.4"
    }
  }
}

provider "opal" {
  server_url = "..." # Optional
}