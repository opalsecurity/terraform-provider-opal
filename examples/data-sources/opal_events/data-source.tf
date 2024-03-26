data "opal_events" "my_events" {
  actor_filter      = "a5e4af3e-79e6-432e-a605-9909fc8f1d1a"
  api_token_filter  = "...my_api_token_filter..."
  cursor            = "...my_cursor..."
  end_date_filter   = "...my_end_date_filter..."
  event_type_filter = "...my_event_type_filter..."
  object_filter     = "174cc096-b23d-46f0-a459-12c11553e8ef"
  page_size         = 2
  start_date_filter = "...my_start_date_filter..."
}