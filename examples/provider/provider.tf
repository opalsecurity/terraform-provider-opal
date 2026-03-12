terraform {
  required_providers {
    opal = {
      source  = "opalsecurity/opal"
      version = "3.4.3"
    }
  }
}

provider "opal" {
  server_url = "..." # Optional
}