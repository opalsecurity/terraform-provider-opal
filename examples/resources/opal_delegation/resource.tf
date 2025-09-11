resource "opal_delegation" "my_delegation" {
  delegate_user_id  = "7c86c85d-0651-43e2-a748-d69d658418e8"
  delegator_user_id = "123e4567-e89b-12d3-a456-426614174000"
  end_time          = "2023-10-01T12:00:00Z"
  reason            = "I need to be out of the office"
  start_time        = "2023-10-01T12:00:00Z"
}