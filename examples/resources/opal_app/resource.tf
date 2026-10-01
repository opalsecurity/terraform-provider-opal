# Push-only app (CUSTOM)
resource "opal_app" "push_only" {
  name           = "My Push-only App"
  description    = "Bookkeeping app for internal tools."
  admin_owner_id = "7c86c85d-0651-43e2-a748-d69d658418e8"
  app_type       = "CUSTOM"
  visibility     = "GLOBAL"
}

# Custom Connector app (CUSTOM_CONNECTOR)
resource "opal_app" "custom_connector" {
  name           = "My Custom Connector"
  description    = "Custom connector for internal tools."
  admin_owner_id = "7c86c85d-0651-43e2-a748-d69d658418e8"
  app_type       = "CUSTOM_CONNECTOR"
  visibility     = "GLOBAL"
  custom_connector = {
    identifier     = "my-connector"
    base_url       = "https://my-connector.example.com"
    signing_secret = "...my_signing_secret..."
  }
}
