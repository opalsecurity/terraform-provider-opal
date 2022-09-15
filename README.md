# Terraform Opal Provider

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
# Get an auth token from https://opal.dev/settings#api or your Opal installation.
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
