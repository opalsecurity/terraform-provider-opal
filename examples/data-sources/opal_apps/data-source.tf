data "opal_apps" "my_apps" {
  app_type_filter = [
    "OKTA_DIRECTORY"
  ]
  owner_filter = "29827fb8-f2dd-4e80-9576-28e31e9934ac"
}