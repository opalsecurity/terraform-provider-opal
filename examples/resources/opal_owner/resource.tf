resource "opal_owner" "my_owner" {
    access_request_escalation_period = 120
            description = "This owner represents the API team owners."
            name = "API Owner"
            reviewer_message_channel_id = "37cb7e41-12ba-46da-92ff-030abe0450b1"
            source_group_id = "1b978423-db0a-4037-a4cf-f79c60cb67b3"
            user_ids = {
        "d056b2fc-18e8-4c3c-9f83-f4e778f2db42",
    }
        }