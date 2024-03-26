# Terraform provider for Opal

<div align="left">
    <a href="https://speakeasyapi.dev/"><img src="https://custom-icon-badges.demolab.com/badge/-Built%20By%20Speakeasy-212015?style=for-the-badge&logoColor=FBE331&logo=speakeasy&labelColor=545454" /></a>
<!--     <a href="https://github.com/speakeasy-sdks/terraform-provider-opal.git/actions"><img src="https://img.shields.io/github/actions/workflow/status/speakeasy-sdks/terraform-provider-opal/speakeasy_sdk_generation.yml?style=for-the-badge" /></a>
     -->
</div>

> ⚠️ This is a sample terraform provider that is not yet ready for production use or release to the TF registry. Resources being covered in this [POC](https://www.notion.so/speakeasyapi/Opal-Speakeasy-Partnership-Plan-f33b9f4c28aa4dcb9fbb4db6fc374ad1?pvs=4)

> To view the edits made to the spec for this generation please see [here](./terraform_overlay.yaml). This is an an [OpenAPI Overlay file](https://www.speakeasyapi.dev/docs/openapi/overlays) computed using the Speakeasy CLI `speakeasy overlay compare --schemas=./openapi_original.yaml --schemas=./openapi.yaml > terraform_overlay.yaml`. 
 

<!-- Start SDK SDK Installation -->
## SDK Installation

To install this provider, copy and paste this code into your Terraform configuration. Then, run `terraform init`.

```hcl
terraform {
  required_providers {
    opal = {
      source  = "opal-dev/opal"
      version = "0.1.0"
    }
  }
}

provider "opal" {
  # Configuration options
}
```
<!-- End SDK SDK Installation -->

## SDK Example Usage

<!-- Start SDK SDK Example Usage -->
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
<!-- End SDK SDK Example Usage -->


<!-- Start SDK SDK Available Operations -->

<!-- End SDK SDK Available Operations -->



<!-- Start SDK Installation [installation] -->
## SDK Installation

To install this provider, copy and paste this code into your Terraform configuration. Then, run `terraform init`.

```hcl
terraform {
  required_providers {
    opal = {
      source  = "opal-dev/opal"
      version = "0.13.29"
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

### SDK Created by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
