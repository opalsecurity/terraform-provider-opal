data "opal_requests" "my_requests" {
  cursor            = "...my_cursor..."
  page_size         = 7
  show_pending_only = true
}