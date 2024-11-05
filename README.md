# Terraform provider for Opal

## SDK Generation
Generate the new SDK using `speakeasy run`. This pulls the remote spec specified in `.speakeasy/workflow.yaml#6` and applies the overrides in `terraform_overlay.yaml`. Note the Makefile is only useful if you want to do development with a local OpenAPI spec and update the Speakeasy workflow config to reference that OpenAPI spec.

<!-- Start Summary [summary] -->
## Summary

Opal API: Your Home For Developer Resources.
<!-- End Summary [summary] -->

<!-- Start Table of Contents [toc] -->
## Table of Contents

* [Installation](#installation)
* [Available Resources and Data Sources](#available-resources-and-data-sources)
* [Testing the provider locally](#testing-the-provider-locally)
<!-- End Table of Contents [toc] -->

<!-- Start Installation [installation] -->
## Installation

To install this provider, copy and paste this code into your Terraform configuration. Then, run `terraform init`.

```hcl
terraform {
  required_providers {
    opal = {
      source  = "opalsecurity/opal"
      version = "0.27.0"
    }
  }
}

provider "opal" {
  # Configuration options
}
```
<!-- End Installation [installation] -->


<!-- Start Available Resources and Data Sources [operations] -->
## Available Resources and Data Sources

### Resources

* [opal_configuration_template](docs/resources/configuration_template.md)
* [opal_group](docs/resources/group.md)
* [opal_group_resource_list](docs/resources/group_resource_list.md)
* [opal_group_tag](docs/resources/group_tag.md)
* [opal_group_user](docs/resources/group_user.md)
* [opal_message_channel](docs/resources/message_channel.md)
* [opal_on_call_schedule](docs/resources/on_call_schedule.md)
* [opal_owner](docs/resources/owner.md)
* [opal_resource](docs/resources/resource.md)
* [opal_resource_tag](docs/resources/resource_tag.md)
* [opal_tag](docs/resources/tag.md)
* [opal_tag_user](docs/resources/tag_user.md)
### Data Sources

* [opal_app](docs/data-sources/app.md)
* [opal_apps](docs/data-sources/apps.md)
* [opal_configuration_template_list](docs/data-sources/configuration_template_list.md)
* [opal_events](docs/data-sources/events.md)
* [opal_group](docs/data-sources/group.md)
* [opal_group_list](docs/data-sources/group_list.md)
* [opal_group_resource_list](docs/data-sources/group_resource_list.md)
* [opal_group_reviewers_stages_list](docs/data-sources/group_reviewers_stages_list.md)
* [opal_group_tags](docs/data-sources/group_tags.md)
* [opal_group_users](docs/data-sources/group_users.md)
* [opal_message_channel](docs/data-sources/message_channel.md)
* [opal_message_channel_list](docs/data-sources/message_channel_list.md)
* [opal_on_call_schedule](docs/data-sources/on_call_schedule.md)
* [opal_on_call_schedule_list](docs/data-sources/on_call_schedule_list.md)
* [opal_owner](docs/data-sources/owner.md)
* [opal_owner_from_name](docs/data-sources/owner_from_name.md)
* [opal_owners](docs/data-sources/owners.md)
* [opal_requests](docs/data-sources/requests.md)
* [opal_resource](docs/data-sources/resource.md)
* [opal_resource_message_channel_list](docs/data-sources/resource_message_channel_list.md)
* [opal_resource_reviewers_list](docs/data-sources/resource_reviewers_list.md)
* [opal_resources_list](docs/data-sources/resources_list.md)
* [opal_resources_access_status](docs/data-sources/resources_access_status.md)
* [opal_resources_users_list](docs/data-sources/resources_users_list.md)
* [opal_resource_tags](docs/data-sources/resource_tags.md)
* [opal_resource_visibility](docs/data-sources/resource_visibility.md)
* [opal_sessions](docs/data-sources/sessions.md)
* [opal_tag](docs/data-sources/tag.md)
* [opal_tags_list](docs/data-sources/tags_list.md)
* [opal_uar](docs/data-sources/uar.md)
* [opal_uars_list](docs/data-sources/uars_list.md)
* [opal_user](docs/data-sources/user.md)
* [opal_users](docs/data-sources/users.md)
* [opal_user_tags](docs/data-sources/user_tags.md)
<!-- End Available Resources and Data Sources [operations] -->

<!-- Start Testing the provider locally [usage] -->
## Testing the provider locally

#### Local Provider

Should you want to validate a change locally, the `--debug` flag allows you to execute the provider against a terraform instance locally.

This also allows for debuggers (e.g. delve) to be attached to the provider.

```sh
go run main.go --debug
# Copy the TF_REATTACH_PROVIDERS env var
# In a new terminal
cd examples/your-example
TF_REATTACH_PROVIDERS=... terraform init
TF_REATTACH_PROVIDERS=... terraform apply
```

#### Compiled Provider

Terraform allows you to use local provider builds by setting a `dev_overrides` block in a configuration file called `.terraformrc`. This block overrides all other configured installation methods.

1. Execute `go build` to construct a binary called `terraform-provider-opal`
2. Ensure that the `.terraformrc` file is configured with a `dev_overrides` section such that your local copy of terraform can see the provider binary

Terraform searches for the `.terraformrc` file in your home directory and applies any configuration settings you set.

```
provider_installation {

  dev_overrides {
      "registry.terraform.io/opalsecurity/opal" = "<PATH>"
  }

  # For all other providers, install them directly from their origin provider
  # registries as normal. If you omit this, Terraform will _only_ use
  # the dev_overrides block, and so no other providers will be available.
  direct {}
}
```
<!-- End Testing the provider locally [usage] -->

<!-- Placeholder for Future Speakeasy SDK Sections -->


