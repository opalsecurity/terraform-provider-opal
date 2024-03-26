data "opal_apps" "my_apps" {
  app_type_filter = [
    "OKTA_DIRECTORY",
  ]
  owner_filter = "ae94bc89-fbd0-43a9-a242-f2e3005dda20"
}