resource "opal_owner" "my_owner" {
  access_request_escalation_period = 120
  description                      = "This owner represents the API team owners."
  name                             = "API Owner"
  reviewer_message_channel_id      = "37cb7e41-12ba-46da-92ff-030abe0450b1"
  source_group_id                  = "1b978423-db0a-4037-a4cf-f79c60cb67b3"
  user_ids = [
    "b4274372-d8d8-4b2e-964f-eb68ae88222a",
  ]
}