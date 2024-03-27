resource "opal_owner" "my_owner" {
  access_request_escalation_period = 120
  description                      = "This owner represents the API team owners."
  name                             = "API Owner"
  reviewer_message_channel_id      = "37cb7e41-12ba-46da-92ff-030abe0450b1"
  source_group_id                  = "1b978423-db0a-4037-a4cf-f79c60cb67b3"
  user_ids = [
    "7d9bb531-a620-4906-bed0-3356e6f799e1",
  ]
}