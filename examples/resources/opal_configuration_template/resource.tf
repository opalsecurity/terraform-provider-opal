resource "opal_configuration_template" "my_configurationtemplate" {
    admin_owner_id = "7c86c85d-0651-43e2-a748-d69d658418e8"
            name = "Prod AWS Template"
            request_configurations = [
        {
            allow_requests = true
            auto_approval = false
            condition = {
                group_ids = {
                    "8154552e-30c0-448b-8c2b-5ed5f0cf07c6",
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
                        "7419b360-edc1-4d66-b24b-f8cfb72161b2",
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
            "a5b24ee8-4a3f-4bf4-9cbc-cae3a741f65f",
        }
    }
        }