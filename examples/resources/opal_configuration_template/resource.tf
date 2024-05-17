resource "opal_configuration_template" "my_configurationtemplate" {
    admin_owner_id = "7c86c85d-0651-43e2-a748-d69d658418e8"
            name = "Prod AWS Template"
            request_configurations = [
        {
            allow_requests = true
            auto_approval = false
            condition = {
                group_ids = {
                    "bc928154-552e-430c-848b-8c2b5ed5f0cf",
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
                        "49098c56-2b97-4419-b360-edc1d66b24bf",
                    }
                    require_admin_approval = false
                    require_manager_approval = false
                },
            ]
        },
    ]
            require_mfa_to_approve = false
            require_mfa_to_connect = false
            visibility = {
        visibility = "GLOBAL"
        visibility_group_ids = {
            "72161b24-a4a5-4fe0-9a5b-24ee84a3fbf4",
        }
    }
        }