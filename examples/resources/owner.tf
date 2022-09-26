resource "opal_owner" "security" {
  name = "Security Team"

  user {
    # User IDs can be pulled from the URL in the Opal web app,
    # e.g. https://app.opal.dev/users/f865f6a5-5be8-46a0-bc57-b9cadaf4d1e5#overview
    id = "f865f6a5-5be8-46a0-bc57-b9cadaf4d1e5"
  }

  user {
    id = "b8059a7a-0e0c-46f2-a430-2666dce89620"
  }
}
