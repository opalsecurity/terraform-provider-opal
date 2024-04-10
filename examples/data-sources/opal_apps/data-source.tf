data "opal_apps" "my_apps" {
  app_type_filter = [
    "OKTA_DIRECTORY",
  ]
  owner_filter = "4d8baf06-9fe5-45e2-ba36-74e7f570e314"
}