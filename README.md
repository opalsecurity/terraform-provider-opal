# Terraform provider for Opal

## SDK Generation
Generate the new SDK using `speakeasy run`. This pulls the remote spec specified in `.speakeasy/workflow.yaml#6` and applies the overrides in `terraform_overlay.yaml`. Note the Makefile is only useful if you want to do development with a local OpenAPI spec and update the Speakeasy workflow config to reference that OpenAPI spec.

<!-- Start SDK Installation [installation] -->
## SDK Installation

To install this provider, copy and paste this code into your Terraform configuration. Then, run `terraform init`.

```hcl
terraform {
  required_providers {
    opal = {
      source  = "opalsecurity/opal"
      version = "0.18.4"
    }
  }
}

provider "opal" {
  # Configuration options
}
```
<!-- End SDK Installation [installation] -->



<!-- Start SDK Example Usage [usage] -->
## SDK Example Usage

### Testing the provider locally

Should you want to validate a change locally, the `--debug` flag allows you to execute the provider against a terraform instance locally.

This also allows for debuggers (e.g. delve) to be attached to the provider.

### Example

```sh
go run main.go --debug
# Copy the TF_REATTACH_PROVIDERS env var
# In a new terminal
cd examples/your-example
TF_REATTACH_PROVIDERS=... terraform init
TF_REATTACH_PROVIDERS=... terraform apply
```
<!-- End SDK Example Usage [usage] -->



<!-- Start Available Resources and Operations [operations] -->
## Available Resources and Operations


<!-- End Available Resources and Operations [operations] -->

<!-- Placeholder for Future Speakeasy SDK Sections -->

Terraform allows you to use local provider builds by setting a `dev_overrides` block in a configuration file called `.terraformrc`. This block overrides all other configured installation methods.

Terraform searches for the `.terraformrc` file in your home directory and applies any configuration settings you set.

```
provider_installation {

  dev_overrides {
      "registry.terraform.io/github.com/opal-dev/terraform-provider-opal/scaffolding" = "<PATH>"
  }

  # For all other providers, install them directly from their origin provider
  # registries as normal. If you omit this, Terraform will _only_ use
  # the dev_overrides block, and so no other providers will be available.
  direct {}
}
```

Your `<PATH>` may vary depending on how your Go environment variables are configured. Execute `go env GOBIN` to set it, then set the `<PATH>` to the value returned. If nothing is returned, set it to the default location, `$HOME/go/bin`.

### Contributions

While we value open-source contributions to this SDK, this library is generated programmatically.
Feel free to open a PR or a Github issue as a proof of concept and we'll do our best to include it in a future release!
