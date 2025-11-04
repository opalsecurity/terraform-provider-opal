resource "opal_resource" "my_resource" {
  admin_owner_id                 = "7c86c85d-0651-43e2-a748-d69d658418e8"
  app_id                         = "f454d283-ca87-4a8a-bdbb-df212eca5353"
  custom_request_notification    = "Check your email to register your account."
  description                    = "Engineering team Okta role."
  extensions_duration_in_minutes = 120
  name                           = "mongo-db-prod"
  remote_info = {
    anthropic_workspace = {
      workspace_id = "ws-123456"
    }
    aws_account = {
      account_id             = 234234234234
      organizational_unit_id = "ou-1234"
    }
    aws_ec2_instance = {
      account_id  = 234234234234
      instance_id = "i-13f1a1e2899f9e93a"
      region      = "us-east-2"
    }
    aws_eks_cluster = {
      account_id = 234234234234
      arn        = "arn:aws:eks:us-east-2:234234234234:cluster/testcluster"
    }
    aws_iam_role = {
      account_id = 234234234234
      arn        = "arn:aws:iam::179308207300:role/MyRole"
    }
    aws_organizational_unit = {
      organizational_unit_id = "ou-1234"
      parent_id              = "ou-1234"
    }
    aws_permission_set = {
      account_id = 234234234234
      arn        = "arn:aws:sso:::permissionSet/asdf-32139302d201d32/ps-f03323201211e1b9"
    }
    aws_rds_cluster = {
      account_id    = 234234234234
      cluster_id    = "demo-mysql-cluster"
      database_name = "mydatabase"
      engine        = "MYSQL"
      region        = "us-east-2"
      resource_id   = "cluster-AOO8V0XUCNU13XLZXQDQRSN0NQ"
    }
    aws_rds_instance = {
      account_id  = 234234234234
      instance_id = "demo-mysql-db"
      region      = "us-east-2"
      resource_id = "db-AOO8V0XUCNU13XLZXQDQRSN0NQ"
    }
    azure_enterprise_app = {
      resource_id = "aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee"
    }
    azure_entra_id_role = {
      resource_id = "11111111-2222-3333-4444-555555555555"
    }
    azure_management_group = {
      resource_id = "/providers/Microsoft.Management/managementGroups/my-mg"
    }
    azure_resource_group = {
      resource_id = "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/my-rg"
    }
    azure_sql_database = {
      resource_id = "/subscriptions/0000/resourceGroups/rg/providers/Microsoft.Sql/servers/sqldev01/databases/db01"
    }
    azure_sql_managed_database = {
      resource_id = "/subscriptions/0000/resourceGroups/rg/providers/Microsoft.Sql/managedInstances/mi01/databases/midb01"
    }
    azure_sql_managed_instance = {
      resource_id = "/subscriptions/0000/resourceGroups/rg/providers/Microsoft.Sql/managedInstances/mi01"
    }
    azure_sql_server = {
      resource_id = "/subscriptions/0000/resourceGroups/rg/providers/Microsoft.Sql/servers/sqldev01"
    }
    azure_storage_account = {
      resource_id = "/subscriptions/0000/resourceGroups/rg/providers/Microsoft.Storage/storageAccounts/sa01"
    }
    azure_storage_container = {
      resource_id = "/subscriptions/0000/resourceGroups/rg/providers/Microsoft.Storage/storageAccounts/sa01/blobServices/default/containers/c01"
    }
    azure_subscription = {
      resource_id = "/subscriptions/00000000-0000-0000-0000-000000000000"
    }
    azure_user_assigned_managed_identity = {
      resource_id = "/subscriptions/0000/resourceGroups/rg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/id01"
    }
    azure_virtual_machine = {
      resource_id = "/subscriptions/0000/resourceGroups/rg/providers/Microsoft.Compute/virtualMachines/vm01"
    }
    coupa_role = {
      role_id = 999
    }
    cursor_organization = {
      org_id = "123e4567-e89b-12d3-a456-426614174000"
    }
    custom_connector = {
      can_have_usage_events = false
      remote_resource_id    = "01fa7402-01d8-103b-8deb-5f3a0ab7884"
    }
    databricks_account_service_principal = {
      application_id = "00000000-0000-0000-0000-000000000000"
      resource_id    = "00000000-0000-0000-0000-000000000000"
    }
    datastax_astra_role = {
      role_id = "123e4567-e89b-12d3-a456-426614174000"
    }
    gcp_big_query_dataset = {
      dataset_id = "example-dataset-898931321"
      project_id = "example-project-898931321"
    }
    gcp_big_query_table = {
      dataset_id = "example-dataset-898931321"
      project_id = "example-project-898931321"
      table_id   = "example-table-898931321"
    }
    gcp_bucket = {
      bucket_id = "example-bucket-898931321"
    }
    gcp_compute_instance = {
      instance_id = "example-instance-898931321"
      project_id  = "example-project-898931321"
      zone        = "us-central1-a"
    }
    gcp_folder = {
      folder_id = "folder/898931321"
    }
    gcp_gke_cluster = {
      cluster_name = "example-cluster-898931321"
    }
    gcp_organization = {
      organization_id = "organizations/898931321"
    }
    gcp_project = {
      project_id = "example-project-898931321"
    }
    gcp_service_account = {
      email              = "production@project.iam.gserviceaccount.com"
      project_id         = "example-project-898931321"
      service_account_id = 103561576023829460000
    }
    gcp_sql_instance = {
      instance_id = "example-sql-898931321"
      project_id  = "example-project-898931321"
    }
    github_org = {
      org_name = "Opal Security"
    }
    github_org_role = {
      role_id = 112233
    }
    github_repo = {
      repo_name = "Opal Security"
    }
    gitlab_project = {
      project_id = 898931321
    }
    google_workspace_role = {
      role_id = "google-workspace-role:01234567890123456"
    }
    ilevel_advanced_role = {
      role_name = "IT User"
    }
    okta_app = {
      app_id = "a9dfas0f678asdf67867"
    }
    okta_custom_role = {
      role_id = "a9dfas0f678asdf67867"
    }
    okta_standard_role = {
      role_type = "ORG_ADMIN"
    }
    openai_platform_project = {
      project_id = "proj_abc"
    }
    openai_platform_service_account = {
      project_id         = "proj_abc"
      service_account_id = "svc_acct_abc"
    }
    oracle_fusion_role = {
      role_id = "123e4567-e89b-12d3-a456-426614174000"
    }
    pagerduty_role = {
      role_name = "owner"
    }
    salesforce_permission_set = {
      permission_set_id = "0PS5Y090202wOV7WAM"
    }
    salesforce_profile = {
      profile_id      = "0PS5Y090202wOV7WAM"
      user_license_id = "1005Y030081Qb5XJHS"
    }
    salesforce_role = {
      role_id = "0PS5Y090202wOV7WAM"
    }
    snowflake_database = {
      database_name = "mydatabase"
    }
    snowflake_schema = {
      database_name = "mydatabase"
      schema_name   = "mycatalogschema"
    }
    snowflake_table = {
      database_name = "mydatabase"
      schema_name   = "mycatalogschema"
      table_name    = "myitemstable"
    }
    tailscale_ssh = {
      tag_name = "admin"
    }
    teleport_role = {
      role_name = "admin_role"
    }
    workday_role = {
      role_id = "123abc456def"
    }
  }
  request_configurations = [
    {
      allow_requests = true
      auto_approval  = false
      condition = {
        group_ids = [
          "1b978423-db0a-4037-a4cf-f79c60cb67b3",
        ]
        role_remote_ids = [
          "arn:aws:iam::590304332660:role/AdministratorAccess",
        ]
      }
      extensions_duration_in_minutes = 120
      max_duration                   = 120
      priority                       = 1
      recommended_duration           = 120
      request_template_id            = "06851574-e50d-40ca-8c78-f72ae6ab4304"
      require_mfa_to_request         = false
      require_support_ticket         = false
      reviewer_stages = [
        {
          operator = "AND"
          owner_ids = [
            "c1fddd27-1944-4f29-a2c5-cd206276bb44"
          ]
          require_admin_approval   = false
          require_manager_approval = false
        }
      ]
    }
  ]
  require_mfa_to_approve    = false
  require_mfa_to_connect    = false
  resource_type             = "AWS_IAM_ROLE"
  risk_sensitivity_override = "HIGH"
  ticket_propagation = {
    enabled_on_grant      = false
    enabled_on_revocation = false
    ticket_project_id     = "...my_ticket_project_id..."
    ticket_provider       = "LINEAR"
  }
  visibility = "GLOBAL"
  visibility_group_ids = [
    "c20519cc-5d81-4468-891e-3dd6093e4e5e"
  ]
}