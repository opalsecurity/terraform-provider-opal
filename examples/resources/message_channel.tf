resource "opal_message_channel" "security_channel" {
  third_party_provider = "SLACK"
  remote_id = "C03L80ABS1T"
}

# Example owner usage
resource "opal_owner" "security" {
  // ...

  reviewer_message_channel_id = opal_message_channel.security_channel.id
}

# Example group usage
resource "opal_group" "security" {
  // ...

  audit_message_channel {
    id = opal_message_channel.security_channel.id
  }
}
