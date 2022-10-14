resource "opal_resource" "sensitive_resource" {
  name = "Sensitive Resource"
  description = "A sensitive resource."
  resource_type = "CUSTOM"
  app_id = data.opal_app.my_custom_app.id
  auto_approval = false
  require_mfa_to_approve = true

  reviewer {
    id = opal_owner.security.id
  }
}

resource "opal_resource" "aws_iam_role_example" {
  name = "AWS IAM role"
  description = "AWS IAM role created via terraform"
  resource_type = "AWS_IAM_ROLE"
  app_id = data.opal_app.aws.id

  remote_info {
    aws_iam_role {
      # Note: This can also be referenced from your AWS terraform files
      arn = "arn:aws:iam::2582003"
    }
  }
}