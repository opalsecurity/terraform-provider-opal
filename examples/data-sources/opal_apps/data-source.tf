data "opal_apps" "my_apps" {
  app_type_filter = [
    "OKTA_DIRECTORY",
  ]
  owner_filter = "dbb0d5f3-dec6-4379-a1a5-f422d4f1a434"
}