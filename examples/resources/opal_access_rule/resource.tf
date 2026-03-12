resource "opal_access_rule" "my_accessrule" {
  admin_owner_id = "7c86c85d-0651-43e2-a748-d69d658418e8"
  description    = "This access rule represents all platform engineers in the company."
  name           = "Platform Engineering"
  rule_clauses = {
    unless = {
      clauses = [
        {
          selectors = [
            {
              connection_id = "cb74c0c0-da9a-4b2e-b301-5872b3381de2"
              key           = "...my_key..."
              value         = "...my_value..."
            }
          ]
        }
      ]
    }
    when = {
      clauses = [
        {
          selectors = [
            {
              connection_id = "63e7b7e6-8efd-4577-8c2a-10982993b910"
              key           = "...my_key..."
              value         = "...my_value..."
            }
          ]
        }
      ]
    }
  }
  status = "ACTIVE"
}