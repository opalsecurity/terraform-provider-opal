resource "opal_script" "my_script" {
  name        = "MFA check"
  owner_id    = "7c86c85d-0651-43e2-a748-d69d658418e8"
  script      = "...my_script..."
  script_type = "REQUEST_REVIEW"
}