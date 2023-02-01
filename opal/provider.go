package opal

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"path"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/opalsecurity/opal-go"
)

func init() {
	// Terraform docs say to set this value but I don't see any change
	// in output from tfplugindocs. It's possible this value is used somewhere else.
	schema.DescriptionKind = schema.StringMarkdown
}

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
				Description: "The base Opal API url in the format `https://[hostname]`. The default value is `https://api.opal.dev`. The value must be provided when working with on-prem",
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("OPAL_BASE_URL", "https://api.opal.dev"),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"opal_owner":            resourceOwner(),
			"opal_resource":         resourceResource(),
			"opal_group":            resourceGroup(),
			"opal_message_channel":  resourceMessageChannel(),
			"opal_on_call_schedule": resourceOnCallSchedule(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"opal_app":  dataSourceApp(),
			"opal_user": dataSourceUser(),
		},
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
			return nil, diag.Diagnostics{{
				Severity: diag.Error,
				Summary:  "Unable to parse provided base_url",
			}}
		}
		u.Path = path.Join(u.Path, "/v1")
		conf.Servers = opal.ServerConfigurations{{
			URL: u.String(),
		}}
	}

	return opal.NewAPIClient(conf), nil
}

// diagFromErr is a small wrapper around diagFromErr that attempts
// to pull more data out of Opal API errors.
func diagFromErr(ctx context.Context, err error) diag.Diagnostics {
	var gErr *opal.GenericOpenAPIError
	if errors.As(err, &gErr) {
		body := make(map[string]any)
		if unmarshalErr := json.Unmarshal(gErr.Body(), &body); unmarshalErr != nil {
			tflog.Error(ctx, "Could not unmarshal response body into json", map[string]any{
				"body": string(gErr.Body()),
			})
			return diag.FromErr(err)
		}

		if body["Message"] != nil {
			return diagFromErr(ctx, fmt.Errorf("opal api error: %s: %s", err.Error(), body["Message"]))
		}
	}
	return diag.FromErr(err)
}
