resource "opal_owner" "security" {
  name = "Security Team"

  user {
    id = data.opal_user.alice.id
  }

  user {
    id = data.opal_user.bob.id
  }
}
