# Terraform provider for Opal

## SDK Generation
Generate the new SDK using `speakeasy run`. This pulls the remote spec specified in `.speakeasy/workflow.yaml#6` and applies the overrides in `terraform_overlay.yaml`. Note the Makefile is only useful if you want to do development with a local OpenAPI spec and update the Speakeasy workflow config to reference that OpenAPI spec.

<!-- Start SDK Installation [installation] -->
## Using the SDK

To install this provider in your Terraform usage, copy and paste this code into your Terraform configuration files. Then, run `terraform init`.

```hcl
terraform {
  required_providers {
    opal = {
      source  = "opalsecurity/opal"
      version = "3.0.1"  # or other later version
    }
  }
}

provider "opal" {
  bearer_auth = <auth-token>
  server_url = "https://api.opal.dev/v1"
}
```

<!-- End SDK Installation [installation] -->


<!-- Start SDK Example Usage [usage] -->
## SDK Example Usage

### Testing the provider locally
If you want to test the provider using a development version of this provider, you can run this provider locally by simply running

```sh
go run main.go --debug
```
This command should output a log line that looks like
```sh
TF_REATTACH_PROVIDERS='{"registry.terraform.io/opalsecurity/opal":{"Protocol":"grpc","ProtocolVersion":6,"Pid":55387,"Test":true,"Addr":{"Network":"unix","String":"/var/folders/rw/nppqqcz93r11_b8n3_q1tzsr0000gn/T/plugin2970912145"}}}'
```
This logline tells you the value of the environment variable to set wherever you invoke your Terraform operations (e.g. `plan`, `apply`, etc). You can either export `TF_REATTACH_PROVIDERS` or just prefix your commands with the envar.

If you would like to enable IDE debugging in VScode you can add the following launch profile.
```
{
    "version": "0.2.0",
    "configurations": [
        {
            "name": "Debug",
            "type": "go",
            "request": "launch",
            "mode": "auto",
            "program": "${workspaceFolder}/main.go",
            "args": ["--debug"],
        }
    ]
}
```
For the IDE to trigger any breakpoints you must run the debug process _within_ VSCode instead of a standalone terminal (e.g. Terminal, ITerm, etc). Take the `TF_REATTACH_PROVIDERS` like before and use it while applying the Terraform operations.


### Contributions

While we value open-source contributions to this SDK, this library is generated programmatically.
Feel free to open a PR or a Github issue as a proof of concept and we'll do our best to include it in a future release!

<!-- No SDK Installation -->
<!-- No SDK Example Usage -->
<!-- No SDK Available Operations -->
<!-- Placeholder for Future Speakeasy SDK Sections -->


