package opal

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/hashicorp/go-multierror"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/opalsecurity/opal-go"
)

var allowedOnCallScheduleProviders = enumSliceToStringSlice(opal.AllowedOnCallScheduleProviderEnumEnumValues)

func resourceOnCallSchedule() *schema.Resource {
	return &schema.Resource{
		Description:   "An Opal OnCallSchedule resource.",
		CreateContext: resourceOnCallScheduleCreate,
		ReadContext:   resourceOnCallScheduleRead,
		DeleteContext: resourceOnCallScheduleDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"id": {
				Description: "The ID of the on call schedule.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"third_party_provider": {
				Description:  "The provider of the on call schedule (i.e. PAGER_DUTY, OPSGENIE).",
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice(allowedOnCallScheduleProviders, false),
				ForceNew:     true,
				Required:     true,
			},
			"remote_id": {
				Description: "The remote ID of the on call schedule.",
				Type:        schema.TypeString,
				ForceNew:    true,
				Required:    true,
			},
		},
	}
}

func resourceOnCallScheduleCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*opal.APIClient)

	provider := opal.OnCallScheduleProviderEnum(d.Get("third_party_provider").(string))
	remoteID := d.Get("remote_id").(string)

	createInfo := opal.NewCreateOnCallScheduleInfo(provider, remoteID)

	onCallSchedule, _, err := client.OnCallSchedulesApi.CreateOnCallSchedule(ctx).CreateOnCallScheduleInfo(*createInfo).Execute()
	if err != nil {
		return diagFromErr(ctx, err)
	}
	tflog.Debug(ctx, "Created on call schedule", map[string]any{
		"provider": provider,
		"id":       onCallSchedule.OnCallScheduleId,
		"remoteID": remoteID,
	})

	d.SetId(onCallSchedule.GetOnCallScheduleId())
	return resourceOnCallScheduleRead(ctx, d, m)
}

func resourceOnCallScheduleRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*opal.APIClient)

	id := d.Get("id").(string)
	onCallSchedule, _, err := client.OnCallSchedulesApi.GetOnCallSchedule(ctx, id).Execute()
	if err != nil {
		return diagFromErr(ctx, err)
	}

	d.SetId(onCallSchedule.GetOnCallScheduleId())
	if err := multierror.Append(
		d.Set("third_party_provider", onCallSchedule.ThirdPartyProvider),
		d.Set("remote_id", onCallSchedule.RemoteId),
	); err.ErrorOrNil() != nil {
		return diagFromErr(ctx, err)
	}

	return nil
}

func resourceOnCallScheduleDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// TODO: Implement
	return nil
}
