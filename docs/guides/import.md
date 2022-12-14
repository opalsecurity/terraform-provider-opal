---
page_title: "Importing existing Opal infrastructure"
subcategory: "Import"
---

# Terraform  import
The current implementation of Terraform import can only import resources into the state. It does not generate configuration. See [this](https://developer.hashicorp.com/terraform/cli/import) for more details.

While it's totally possible to use `terraform import` in combination with writing manual configuration to import your infrastructure, we recommend using Terraformer (a tool that will import state + configuration).

# Terraformer
Note: We're currently waiting on Opal support being merged into the official [Terraformer](https://github.com/GoogleCloudPlatform/terraformer) repo.

In the meantime, we have pre-built binaries [available for download](https://github.com/opalsecurity/terraformer/releases/latest). For Apple Silicon macs, use the arm64 build.

After download, `terraformer` (i.e. `~/Downloads/terraformer-all-darwin-arm64`) should be available to you. Before we can start importing infrastructure, we also need to install the Opal Terraform provider.

## Installing Opal Terraform provider

1. In the directory from which you want to import your Opal infrastructure, add a `versions.tf` file with the following content:
```hcl
terraform {
  required_providers {
    opal = {
      source = "opalsecurity/opal"
    }
  }
  required_version = ">= 0.13"
}
```
2. Run `terraform init` to download the Opal provider.

Now we're ready to start importing Opal infrastructure.

## Importing Opal Infrastructure

### Environment variables

In order to run the import commands, the following environment variables are required:

1. A read-only Opal Admin token. You can generate it the admin settings of the Opal web UI
2. The base url of your Opal instance

```bash
$ export OPAL_AUTH_TOKEN=XXX
# NOTE: OPAL_BASE_URL is only needed if you have an on-prem installation
$ export OPAL_BASE_URL=https://my.opal.corp.dev
```

### Importing everything

In order to import everything, you can run the following command:
```bash
$ terraformer import opal --resources="*" --path-pattern {output}/{provider} --no-sort
```

*NOTE:*
- `--no-sort` is needed for importing owners. The order of the users determines the escalation order if an escalation policy is set.
- Feel free to use any `path-pattern` that you'd like.
- Make sure to follow the steps in "Inspect the imported terraform files" below to use your terraform files.

### Importing a specific resource by ID

```bash
$ terraformer import opal --resources=resource --filter=resource=7900e913-81c2-4c3d-8d1e-1d37952ebcbf --path-pattern {output}/{provider} --no-sort
```

*NOTE:*
- running this for multiple resources will override previous imports. See below for how to import multiple resources by ID.
- Make sure to follow the steps in "Inspect the imported terraform files" below to use your terraform files.

### Importing multiple resources by ID
The syntax to import multiple resources by their IDs is a bit cumbersome. To help generate an import command, you can use the following [python3 script](https://gist.github.com/jan-opal/44c796111763d1e5f11715741425e987).

To use the script, do the following:
1. Fill out the resource, group, and owner ids that you want to import. You can find them in the URL bar of the Opal web interface.
2. Run the script `python3 generate_import_command.py`.
3. The printed output is the `terraformer` command that will import the specified resources, groups, and owners into Terraform.
4. Double-check that the command looks good and run it to complete the import.
5. Make sure to follow the steps in "Inspect the imported terraform files" below to use your terraform files.

### Inspect the imported terraform files

You should now see a `generated/opal/` subdirectory with generated files. If you are using
terraform version `>= 0.13`, you will need to run a state migration:
```bash
$ cd generated/opal/
$ terraform state replace-provider -auto-approve "registry.terraform.io/-/opal" "opalsecurity/opal"
```

You can now initialize and use your new generated resources:
```bash
$ terraform init
$ terraform plan # No changes. Your infrastructure matches the configuration.
```

# Help guide

1. _All my imported terraform names have `tfer--` in them._

This is automatically done by the terraformer tool. You can read more about it [here](https://github.com/GoogleCloudPlatform/terraformer/pull/220). If you want to get rid of them, you can open the `generated/opal` directory in your favorite editor and `ReplaceAll("tfer--", "")`.

2. _When running `terraformer`, I get `open /Users/XXX/.terraform.d/plugins/darwin_arm64: no such file or directory`._

You did not correctly install the Opal Terraform provider or are not running the command from the directory in which you installed the provider. See the "Installing Opal Terraform provider" section.

3. _When running `terraform init`, I get `Error: Invalid legacy provider address`_

You need to run a state migration. See the "Inspect the imported terraform files" section.

# Building terraformer from source
Instead of using our pre-built binaries, you can build from source yourself.

## Pre-requisites
- you have [golang](https://go.dev/doc/install) installed
- you have [terraform](https://learn.hashicorp.com/tutorials/terraform/install-cli) installed

## Build

```bash
# clone the Opal fork of Terraformer
$ git clone https://github.com/opalsecurity/terraformer.git && cd terraformer

# build terraformer-opal binary
$ go run ./build/main.go opal

# optional: expose the terraformer-opal command in your path
$ cp terraformer-opal /usr/local/bin/

# refresh your shell (use equivalent command for non-zsh shell)
$ source ~/.zshrc
```
