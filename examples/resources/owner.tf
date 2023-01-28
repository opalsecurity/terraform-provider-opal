resource "opal_owner" "security" {
  name = "Security Team"

  user {
    id = data.opal_user.alice.id
  }

  user {
    id = data.opal_user.bob.id
  }
}

resource "opal_owner" "ops" {
  name            = "Ops Team"
  source_group_id = data.opal_group.opal_group_example.id
}
