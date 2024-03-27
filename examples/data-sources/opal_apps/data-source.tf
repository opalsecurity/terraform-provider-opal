data "opal_apps" "my_apps" {
  app_type_filter = [
    "OKTA_DIRECTORY",
  ]
  owner_filter = "4b101a07-642b-4a38-bc73-7d2bdc6d83a7"
}