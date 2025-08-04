data "opal_requests" "my_requests" {
  cursor            = "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw"
  end_date_filter   = "...my_end_date_filter..."
  page_size         = 200
  requester_id      = "37cb7e41-12ba-46da-92ff-030abe0450b1"
  show_pending_only = false
  start_date_filter = "...my_start_date_filter..."
  target_user_id    = "37cb7e41-12ba-46da-92ff-030abe0450b1"
}