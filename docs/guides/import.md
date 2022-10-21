---
page_title: "Importing existing Opal infrastructure"
subcategory: "Import"
---

# Terraform  import
The current implementation of Terraform import can only import resources into the state. It does not generate configuration. See [this](https://developer.hashicorp.com/terraform/cli/import) for more details.

While it's totally possible to use `terraform import` in combination with writing manual configuration to import your infrastructure, we recommend using Terraformer (a tool that will import state + configuration).

# Terraformer
Note: We're currently waiting on Opal support being merged into the official [Terraformer](https://github.com/GoogleCloudPlatform/terraformer) repo. In the meantime, the setup to make imports work is slightly more tedious.

## Pre-requisites
- you have [golang](https://go.dev/doc/install) installed
- you have [terraform](https://learn.hashicorp.com/tutorials/terraform/install-cli) installed

## Installing  Terraformer

1. Clone the Opal fork of Terraformer: `git clone  https://github.com/opalsecurity/terraformer.git`
2. Go into the cloned repo: `cd terraformer/`
3. Build the Opal Terraformer tool: `go run ./build/main.go opal`
4. The above command build a binary called `terraformer-opal`. To expose this in your path, run `cp terraformer-opal /usr/local/bin/`
5. Refresh your shell: `source ~/.zshrc` (or equivalent for other shell)

Now `terraformer-opal` should  be available to you. Before we can start importing infrastructure, we need to install the Opal Terraform provider.

## Installing Opal Terraform provider

1. In the directory from which you want to import your Opal infrastructure, add a `versions.tf` file with the following content:
```hcl
terraform {
  required_providers {
    google = {
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
1. A read-only Opal Admin token: `export OPAL_OAUTH_TOKEN=XXX`
2. [Only needed for on-prem customers] The base url of your Opal instance: `export OPAL_BASE_URL=https://my.opal.corp.dev/v1`

### Importing everything

In order to import everything, you can run the following command (Note: In some shells you might have to escape the `*` by replacing it with `\*`):
```bash
$ terraformer-opal import opal --resources=* --path-pattern {output}/{provider} --no-sort
```

*NOTE:*
- `--no-sort` is needed for importing owners. The order of the users determines the escalation order if an escalation policy is set.
- Feel free to use any `path-pattern` that you'd like.

### Importing a specific resource by ID

```bash
$ terraformer-opal import opal --resources=resource --filter=resource=7900e913-81c2-4c3d-8d1e-1d37952ebcbf --path-pattern {output}/{provider} --no-sort
```

*NOTE:*
- running this for multiple resources will override previous imports. See below for how to import multiple resources by ID.

### Importing multiple resources by ID
The syntax to import multiple resources by their IDs is a bit cumbersome. To help generate an import command, you can use the following [python3 script](https://gist.github.com/jan-opal/44c796111763d1e5f11715741425e987).

To use the script, do the following:
1. Fill out the resource, group, and owner ids that you want to import. You can find them in the URL bar of the Opal web interface.
2. Run the script `python3 generate_import_command.py`.
3. The printed output is the `terraformer-opal` command that will import the specified resources, groups, and owners into Terraform.
4. Double-check that the command looks good and run it to complete the import.

### Inspect the imported terraform files

You should now see a `generated/` subdirectory with generated files. If you are using
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

> When running `terraformer-opal`, I get `open /Users/XXX/.terraform.d/plugins/darwin_arm64: no such file or directory`.

You did not correctly install the Opal Terraform provider or are not running the command from the directory in which you installed the provider. See [this](## Installing Opal Terraform provider).

> When running `terraform init`, I get `Error: Invalid legacy provider address`

You need to run a state migration. See [this](### Inspect the imported terraform files).