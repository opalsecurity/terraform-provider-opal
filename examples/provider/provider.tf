terraform {
  required_providers {
    opal = {
      source  = "opalsecurity/opal"
      version = "3.4.5"
    }
  }
}

provider "opal" {
  server_url = "..." # Optional
}