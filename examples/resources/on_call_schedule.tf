resource "opal_on_call_schedule" "security_oncall_rotation" {
  third_party_provider = "PAGER_DUTY"
  remote_id            = "PNXHVAA"
}

# Example group usage
resource "opal_group" "security" {
  // ...

  on_call_schedule {
    id = opal_on_call_schedule.security_oncall_rotation.id
  }

  // or if an UUID is already present in Opal
  on_call_schedule {
    id = "878ba05b-33f0-4dd5-a199-09efc06abcf7"
  }
}
