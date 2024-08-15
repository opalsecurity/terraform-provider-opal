resource "opal_resource" "my_resource" {
    admin_owner_id = "7c86c85d-0651-43e2-a748-d69d658418e8"
            app_id = "f454d283-ca87-4a8a-bdbb-df212eca5353"
            description = "This resource represents AWS IAM role \"SupportUser\"."
            name = "my-mongo-db"
            request_configurations = {
        {
            allow_requests = true
            auto_approval = false
            condition = {
                group_ids = {
                    "caf8a4c7-ec57-4b81-a172-8cd3687b3ad9",
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
                        "66c2bdfc-b3e7-4464-ae94-bc89fbd03a92",
                    }
                    require_admin_approval = false
                    require_manager_approval = false
                },
            ]
        },
    }
            require_mfa_to_approve = false
            require_mfa_to_connect = false
            resource_type = "AWS_IAM_ROLE"
            visibility = "GLOBAL"
            visibility_group_ids = {
        "dda2041a-0d76-41a9-b988-b1ee194539e3",
    }
        }