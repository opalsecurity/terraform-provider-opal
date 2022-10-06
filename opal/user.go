package opal

import (
	"context"
	"errors"
	"github.com/hashicorp/go-multierror"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/opalsecurity/opal-go"
)

func dataSourceUser() *schema.Resource {
	return &schema.Resource{
		Description: "An Opal User data source.",
		ReadContext: dataSourceUserRead,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"id": {
				Description: "The ID of the user.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"email": {
				Description: "The email of the user.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"name": {
				Description: "The name of the user.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"position": {
				Description: "The position of the user.",
				Type:        schema.TypeString,
				Computed:    true,
			},
		},
	}
}

func dataSourceUserRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*opal.APIClient)

	userApiCall := client.UsersApi.User(ctx)

	id, idOk := d.GetOk("id")
	email, emailOk := d.GetOk("email")

	if idOk {
		userApiCall = userApiCall.UserId(id.(string))
	} else if emailOk {
		userApiCall = userApiCall.Email(email.(string))
	} else {
		return diagFromErr(ctx, errors.New("must provide either id or email for user data source"))
	}

	user, _, err := userApiCall.Execute()
	if err != nil {
		return diagFromErr(ctx, err)
	}

	d.SetId(user.UserId)
	if err := multierror.Append(
		d.Set("name", user.FullName),
		d.Set("email", user.Email),
		d.Set("position", user.Position),
	); err.ErrorOrNil() != nil {
		return diagFromErr(ctx, err)
	}

	return nil
}
