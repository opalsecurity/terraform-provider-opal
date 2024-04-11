data "opal_apps" "my_apps" {
  app_type_filter = [
    "OKTA_DIRECTORY",
  ]
  owner_filter = "d8baf069-fe55-4e23-a367-4e7f570e3140"
}