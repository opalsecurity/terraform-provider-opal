data "opal_events" "my_events" {
  actor_filter      = "0589e447-81ce-46ef-a435-f4c67e895185"
  api_token_filter  = "...my_api_token_filter..."
  cursor            = "...my_cursor..."
  end_date_filter   = "...my_end_date_filter..."
  event_type_filter = "...my_event_type_filter..."
  object_filter     = "d4b101a0-7642-4ba3-8bc7-37d2bdc6d83a"
  page_size         = 5
  start_date_filter = "...my_start_date_filter..."
}