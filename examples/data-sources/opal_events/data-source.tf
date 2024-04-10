data "opal_events" "my_events" {
  actor_filter      = "589e4478-1ce6-4ef6-835f-4c67e895185d"
  api_token_filter  = "...my_api_token_filter..."
  cursor            = "...my_cursor..."
  end_date_filter   = "...my_end_date_filter..."
  event_type_filter = "...my_event_type_filter..."
  object_filter     = "4b101a07-642b-4a38-bc73-7d2bdc6d83a7"
  page_size         = 8
  start_date_filter = "...my_start_date_filter..."
}