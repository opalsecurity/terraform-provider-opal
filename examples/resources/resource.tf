resource "opal_resource" "sensitive_resource" {
  name = "Sensitive Resource"
  description = "A sensitive resource that should be accessed for on-call only."

  resource_type = "CUSTOM"

  # App IDs can be pulled from the URL in the Opal web app,
  # e.g. https://app.opal.dev/apps/03c06479-6ffa-45e1-9f65-cd470ff128b3#overview
  app_id = "03c06479-6ffa-45e1-9f65-cd470ff128b3"

  reviewer {
    id = "${opal_owner.security.id}"
  }
}
