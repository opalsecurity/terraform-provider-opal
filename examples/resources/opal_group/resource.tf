resource "opal_group" "my_group" {
    admin_owner_id = "7c86c85d-0651-43e2-a748-d69d658418e8"
            app_id = "f454d283-ca87-4a8a-bdbb-df212eca5353"
            description = "Engineering team Okta group."
            group_type = "OPAL_GROUP"
            message_channel_ids = {
        "931861ef-161b-454f-b189-8e51c0009dd6",
    }
            name = "mongo-db-prod"
            on_call_schedule_ids = {
        "c4289332-6c8d-43b6-943a-d1053f385d17",
    }
            request_configurations = {
        {
            allow_requests = true
            auto_approval = false
            condition = {
                group_ids = {
                    "27d58f25-6147-4d92-aa24-6933c124e146",
                }
                role_remote_ids = {
                    "...",
                }
            }
            max_duration = 120
            priority = 1
            recommended_duration = 120
            request_template_id = "06851574-e50d-40ca-8c78-f72ae6ab4304"
            require_mfa_to_request = false
            require_support_ticket = false
            reviewer_stages = [
                {
                    operator = "AND"
                    owner_ids = {
                        "ec25e287-06e5-40ad-a559-d94490f51937",
                    }
                    require_admin_approval = false
                    require_manager_approval = false
                },
            ]
        },
    }
            require_mfa_to_approve = false
            visibility = "GLOBAL"
        }