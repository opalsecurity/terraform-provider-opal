package opal

import (
	"context"
	"fmt"
	"net/url"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/opalsecurity/opal-go"
)

// NewProvider returns a *schema.Provider.
func NewProvider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"token": {
				Description: "The authentication token used to connect to Opal. The value can be sourced OPAL_AUTH_TOKEN.",
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{"OPAL_AUTH_TOKEN"}, nil),
				Sensitive:   true,
			},
			"base_url": {
				Description: "The base Opal API url in the format `https://[hostname]/v1`. The default value is `https://api.opal.dev/v1`. The value must be provided when working with on-prem",
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("OPAL_BASE_URL", "https://api.opal.dev/v1"),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"opal_owner": resourceOwner(),
		},
		DataSourcesMap:       map[string]*schema.Resource{},
		ConfigureContextFunc: configure,
	}
}

func configure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	conf := opal.NewConfiguration()

	tokenT, ok := d.GetOk("token")
	if !ok {
		return nil, diag.Diagnostics{{
			Severity: diag.Error,
			Summary:  "Unable to create Opal client",
			Detail:   "Token must be provided in provider configuration or via the OPAL_AUTH_TOKEN environment variable",
		}}
	}
	conf.DefaultHeader["Authorization"] = fmt.Sprintf("Bearer %s", tokenT.(string))

	baseUrlT, ok := d.GetOk("base_url")
	if ok {
		u, err := url.Parse(baseUrlT.(string))
		if err != nil {
			return nil, diag.FromErr(err)
		}
		conf.Host = u.Host
		conf.Scheme = u.Scheme
	}

	return opal.NewAPIClient(conf), nil
}
