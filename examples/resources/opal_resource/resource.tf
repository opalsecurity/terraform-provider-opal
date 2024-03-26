resource "opal_resource" "my_resource" {
  admin_owner_id         = "7c86c85d-0651-43e2-a748-d69d658418e8"
  app_id                 = "f454d283-ca87-4a8a-bdbb-df212eca5353"
  description            = "Engineering team Okta role."
  name                   = "mongo-db-prod"
  require_mfa_to_approve = false
  require_mfa_to_connect = false
  resource_type          = "AWS_IAM_ROLE"
  visibility             = "GLOBAL"
}