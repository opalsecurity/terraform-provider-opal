data "opal_events" "my_events" {
  actor_filter      = "29827fb8-f2dd-4e80-9576-28e31e9934ac"
  api_token_filter  = "fullaccess:**************************M_g=="
  cursor            = "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw"
  end_date_filter   = "...my_end_date_filter..."
  event_type_filter = "USER_MFA_RESET"
  object_filter     = "29827fb8-f2dd-4e80-9576-28e31e9934ac"
  page_size         = 200
  start_date_filter = "...my_start_date_filter..."
}