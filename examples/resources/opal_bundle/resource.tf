resource "opal_bundle" "my_bundle" {
  admin_owner_id = "7c86c85d-0651-43e2-a748-d69d658418e8"
  description    = "This is a test bundle"
  name           = "Test Bundle"
  visibility     = "GLOBAL"
  visibility_group_ids = [
    "c07edab4-0137-4fe5-91de-09be56d2304d"
  ]
}