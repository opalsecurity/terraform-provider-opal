resource "opal_group" "oncall" {
  name = "On-Call Rotation"

  group_type = "OPAL_GROUP"

  # App IDs can be pulled from the URL in the Opal web app,
  # e.g. https://app.opal.dev/apps/dbe38d2d-9ce4-4d13-95a4-945716a257b4#overview
  app_id = "dbe38d2d-9ce4-4d13-95a4-945716a257b4"

  require_mfa_to_approve = true
}