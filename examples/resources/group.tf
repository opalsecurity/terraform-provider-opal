resource "opal_group" "opal_group_example" {
  name = "Opal group"
  description = "Opal group created via terraform"
  group_type = "OPAL_GROUP"
  app_id = data.opal_app.opal.id
  require_mfa_to_approve = true
  admin_owner_id = data.opal_owner.security.id
  auto_approval = false

  reviewer {
    id = data.opal_owner.security.id
  }
}

resource "opal_group" "okta_group_example" {
  name = "Okta group"
  description = "Okta group created via terraform"
  group_type = "OKTA_GROUP"
  app_id = data.opal_app.okta.id

  remote_info {
    okta_group {
      # Note: This can also be referenced from your Okta terraform files
      group_id = "00gd7wmwj7hfT3wTH8d6"
    }
  }
}