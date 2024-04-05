data "opal_events" "my_events" {
  actor_filter      = "c137cf6f-e2c7-4c19-8104-295311709f01"
  api_token_filter  = "...my_api_token_filter..."
  cursor            = "...my_cursor..."
  end_date_filter   = "...my_end_date_filter..."
  event_type_filter = "...my_event_type_filter..."
  object_filter     = "64bcc2fc-4397-4c34-8271-060114ab43bd"
  page_size         = 9
  start_date_filter = "...my_start_date_filter..."
}