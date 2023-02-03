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
