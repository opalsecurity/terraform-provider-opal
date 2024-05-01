data "opal_events" "my_events" {
  actor_filter      = "d8baf069-fe55-4e23-a367-4e7f570e3140"
  api_token_filter  = "...my_api_token_filter..."
  cursor            = "...my_cursor..."
  end_date_filter   = "...my_end_date_filter..."
  event_type_filter = "...my_event_type_filter..."
  object_filter     = "589e4478-1ce6-4ef6-835f-4c67e895185d"
  page_size         = 3
  start_date_filter = "...my_start_date_filter..."
}