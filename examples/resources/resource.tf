resource "opal_resource" "sensitive_resource" {
  name = "Sensitive Resource"
  description = "A sensitive resource."
  resource_type = "CUSTOM"
  app_id = data.opal_app.my_custom_app.id
  admin_owner_id = opal_owner.security.id
  require_mfa_to_approve = true

  request_configuration {
    auto_approval = false
    reviewer_stage {
      reviewer {
        id = opal_owner.security.id
      }
    }
  }
}
