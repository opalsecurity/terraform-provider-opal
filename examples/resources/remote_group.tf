resource "opal_group" "okta_group_example" {
  name = "Okta group"
  description = "Okta group created via terraform"
  group_type = "OKTA_GROUP"
  app_id = data.opal_app.okta.id
  admin_owner_id = opal_owner.security.id

  remote_info {
    okta_group {
      # Note: This can reference your Okta terraform files
      group_id = "00gd7wmwj7hfT3wTH8d6"
    }
  }
}

resource "opal_group" "google_group_example" {
  name = "Google group"
  // ...

  remote_info {
    google_group {
      # Note: This can reference your Google terraform files
      group_id = "081mfhsl1l8rxqv"
    }
  }
}

resource "opal_group" "azure_ad_365_example" {
  name = "ms365group"
  // ...

  remote_info {
    azure_ad_microsoft_365_group {
      group_id = "70ef8380-1e43-47cb-80a3-2afd16fe1e96"
    }
  }
}


resource "opal_group" "azure_ad_security_group_example" {
  name = "another group"
  // ...

  remote_info {
    azure_ad_security_group {
      group_id = "265a2b67-7bcd-4ef5-9325-4ebb21254efb"
    }
  }
}
