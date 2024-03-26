data "opal_events" "my_events" {
  actor_filter      = "41a0d761-a939-488b-9ee1-94539e33fdc8"
  api_token_filter  = "...my_api_token_filter..."
  cursor            = "...my_cursor..."
  end_date_filter   = "...my_end_date_filter..."
  event_type_filter = "...my_event_type_filter..."
  object_filter     = "0a5e4af3-e79e-4632-ae60-59909fc8f1d1"
  page_size         = 7
  start_date_filter = "...my_start_date_filter..."
}