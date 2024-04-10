resource "opal_resource" "my_resource" {
    admin_owner_id = "7c86c85d-0651-43e2-a748-d69d658418e8"
            app_id = "f454d283-ca87-4a8a-bdbb-df212eca5353"
            description = "Engineering team Okta role."
            name = "mongo-db-prod"
            request_configurations = {
        {
            allow_requests = true
            auto_approval = false
            condition = {
                group_ids = [
                    "b32b9f65-3f84-49e2-bffa-28b2affecaf8",
                ]
                role_remote_ids = [
                    "...",
                ]
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
                    owner_ids = [
                        "fb82a285-fd0c-4387-b2b3-1cd922ce15f9",
                    ]
                    require_manager_approval = false
                },
            ]
        },
    }
            require_mfa_to_approve = false
            require_mfa_to_connect = false
            resource_type = "AWS_IAM_ROLE"
            visibility = "GLOBAL"
        }