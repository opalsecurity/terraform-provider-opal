package opal

import (
	"context"
	"errors"
	"github.com/hashicorp/go-multierror"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/opalsecurity/opal-go"
)

func dataSourceApp() *schema.Resource {
	return &schema.Resource{
		Description: "An Opal app data source.",
		ReadContext: dataSourceAppRead,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"id": {
				Description: "The ID of the app.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"name": {
				Description: "The name of the app.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"type": {
				Description: "The type of app.",
				Type:        schema.TypeString,
				Computed:    true,
			},
		},
	}
}

func dataSourceAppRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*opal.APIClient)

	id, idOk := d.GetOk("id")
	if !idOk {
		return diagFromErr(ctx, errors.New("must provide id for app data source"))
	}

	app, _, err := client.AppsApi.GetApp(ctx, id.(string)).Execute()
	if err != nil {
		return diagFromErr(ctx, err)
	}

	d.SetId(app.AppId)
	if err := multierror.Append(
		d.Set("name", app.Name),
		d.Set("type", app.AppType),
	); err.ErrorOrNil() != nil {
		return diagFromErr(ctx, err)
	}

	return nil
}
