data "opal_apps" "my_apps" {
  app_type_filter = [
    "OKTA_DIRECTORY",
  ]
  owner_filter = "1a0d761a-9398-48b1-ae19-4539e33fdc80"
}