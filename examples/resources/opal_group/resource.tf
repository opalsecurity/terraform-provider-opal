resource "opal_group" "my_group" {
    admin_owner_id = "7c86c85d-0651-43e2-a748-d69d658418e8"
            app_id = "f454d283-ca87-4a8a-bdbb-df212eca5353"
            description = "This group represents Active Directory group \"Payments Production Admin\". We use this AD group to facilitate staging deployments and qualifying new releases."
            group_type = "OPAL_GROUP"
            message_channel_ids = {
        "e51c0009-dd65-4b4e-90c4-2893326c8d3b",
    }
            name = "api-group"
            on_call_schedule_ids = {
        "1053f385-d17c-4e27-958f-256147d92ea2",
    }
            request_configurations = [
        {
            allow_requests = true
            auto_approval = false
            condition = {
                group_ids = {
                    "3c124e14-6eb0-45cd-86c1-6b8673554554",
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
                        "37d5bf18-869a-4e72-ac0c-c018ec506c2a",
                    }
                    require_admin_approval = false
                    require_manager_approval = false
                },
            ]
        },
    ]
            require_mfa_to_approve = false
            visibility = "GLOBAL"
        }