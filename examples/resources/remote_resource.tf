resource "opal_resource" "aws_iam_role_example" {
  name = "AWS IAM role"
  description = "AWS IAM role created via terraform"
  resource_type = "AWS_IAM_ROLE"
  app_id = data.opal_app.aws.id
  admin_owner_id = data.opal_owner.security.id

  remote_info {
    aws_iam_role {
      # Note: This can reference your AWS terraform files
      arn = "arn:aws:iam::2582003"
    }
  }
}

resource "opal_resource" "okta_app_example" {
  name = "Okta app"
  // ...

  remote_info {
    okta_app {
      # Note: This can reference your Okta terraform files
      app_id = "0oa2aa0fcje6E2kXC5d7"
    }
  }
}

resource "opal_resource" "github_repo_example" {
  name = "GitHub repo"
  // ...

  remote_info {
    github_repo {
      # Note: This can reference your GitHub terraform files
      repo_id = "234432289"
      repo_name = "my-repo"
    }
  }
}
