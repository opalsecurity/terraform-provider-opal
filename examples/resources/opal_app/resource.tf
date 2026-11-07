resource "opal_app" "my_app" {
  admin_owner_id = "7c86c85d-0651-43e2-a748-d69d658418e8"
  app_type       = "CUSTOM_CONNECTOR"
  custom_connector = {
    base_url                  = "https://my-connector.example.com"
    identifier                = "my-connector"
    idp_protocol              = "OPAL"
    signing_secret            = "...my_signing_secret..."
    supports_event_ingestion  = false
    supports_groups           = true
    supports_idp              = false
    supports_nested_groups    = false
    supports_nested_resources = true
    tls_ca_cert_content       = "...my_tls_ca_cert_content..."
    tls_mode                  = false
  }
  description       = "Bookkeeping app for internal tools."
  import_visibility = "GLOBAL"
  name              = "My Custom Connector App"
  visibility        = "GLOBAL"
}
