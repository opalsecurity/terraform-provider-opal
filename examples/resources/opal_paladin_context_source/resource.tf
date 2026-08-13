resource "opal_paladin_context_source" "my_paladincontextsource" {
  name                 = "...my_name..."
  paladin_id           = "32acc112-21ff-4669-91c2-21e27683eaa1"
  remote_id            = "...my_remote_id..."
  source_kind          = "DOCUMENT"
  third_party_provider = "CONFLUENCE"
  url                  = "...my_url..."
}