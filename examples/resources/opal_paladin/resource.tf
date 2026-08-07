resource "opal_paladin" "my_paladin" {
  admin_view_only = false
  enabled_connectors = [
    "PAGER_DUTY"
  ]
  monitor_mode = true
  name         = "paladin-agent-1"
  owner_id     = "7c86c85d-0651-43e2-a748-d69d658418e8"
}