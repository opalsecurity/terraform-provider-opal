# Terraform Opal Provider
[![Terraform Provider Tests](https://github.com/opalsecurity/terraform-provider-opal/actions/workflows/test.yml/badge.svg)](https://github.com/opalsecurity/terraform-provider-opal/actions/workflows/test.yml)

This project is under **active development** and is not yet ready for use.

## Installation
```hcl
terraform {
  required_providers {
    opal = {
      source = "opalsecurity/opal"
    }
  }
}

provider "opal" {
  # Configuration options
}
```

## Development

Go `>= 1.18` and terraform `>= 0.14` is required for development. It's recommended that you use a [`dev_overrides` block](https://www.terraform.io/cli/config/config-file) while developing:
```hcl
provider_installation {
  dev_overrides {
    "registry.terraform.io/opalsecurity/opal" = "/Users/you/src/terraform-provider-opal/bin"
  }
}
```

You can also source your local `OPAL_AUTH_TOKEN` while developing by using [direnv](https://direnv.net) (installable via homebrew) and creating a `.envrc.local` file:
```bash
# Get an auth token from https://app.opal.dev/settings#api or your Opal installation.
export OPAL_AUTH_TOKEN=YOUR_TOKEN_HERE
```

You can build the plugin using:
```
make build
```

Your `dev_overrides` configured above should tell your local terraform installation how to resolve the plugin:
```
$ cd examples/
$ terraform apply
╷
│ Warning: Provider development overrides are in effect
│
│ The following provider development overrides are set in the CLI configuration:
│  - opalsecurity/opal in /Users/user/src/terraform-provider-opal/bin
│
│ The behavior may therefore not match any released version of the provider and applying changes may cause the state to become incompatible
│ with published releases.
╵
```

If you don't see the above warning when running terraform commands, something is misconfigured.

### Writing Documentation

The `docs/` folder is entirely generated. Make changes to `templates/` or the go source files instead. `docs/`content is generated from:

- The source code, i.e. `Description` and `Name` fields in the resource schema
- The `templates/` folder, which serves as the basis for `docs/`. See [tfplugindocs](https://github.com/hashicorp/terraform-plugin-docs#templates) for more on the templating fields.