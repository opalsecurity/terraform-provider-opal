---
page_title: "Authenticating with your Opal instance"
subcategory: "Authentication"
---

# Authentication
Opal administrators can generate new API tokens in the Admin view. See [this guide](https://docs.opal.dev/reference/authentication) for more details.

For on-prem installations, you will also need to specify the base URL of your on-prem installation.

```terraform
provider "opal" {
  token = "YOUR_OPAL_API_TOKEN" # Or the OPAL_AUTH_TOKEN environment variable.

  # Optionally, you can specify the base url of your Opal on-prem installation.
  base_url = "https://my.opal.corp.dev"
}
```
