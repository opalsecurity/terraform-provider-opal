resource "opal_group" "opal_group_example" {
  name = "Opal group"
  description = "Opal group created via terraform"
  group_type = "OPAL_GROUP"
  app_id = data.opal_app.opal.id
  admin_owner_id = opal_owner.security.id
  require_mfa_to_approve = true
  auto_approval = false

  reviewer_stage {
    reviewer {
      id = opal_owner.security.id
    }
  }
}
