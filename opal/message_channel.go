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

var allowedProviders = enumSliceToStringSlice(opal.AllowedMessageChannelProviderEnumEnumValues)

func resourceMessageChannel() *schema.Resource {
	return &schema.Resource{
		Description:   "An Opal MessageChannel resource.",
		CreateContext: resourceMessageChannelCreate,
		ReadContext:   resourceMessageChannelRead,
		DeleteContext: resourceMessageChannelDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"id": {
				Description: "The ID of the message_channel.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"third_party_provider": {
				Description:  "The provider of the message channel (i.e. SLACK).",
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice(allowedProviders, false),
				ForceNew:     true,
				Required:     true,
			},
			"remote_id": {
				Description: "The remote ID of the message_channel.",
				Type:        schema.TypeString,
				ForceNew:    true,
				Required:    true,
			},
		},
	}
}

func resourceMessageChannelCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*opal.APIClient)

	provider := opal.MessageChannelProviderEnum(d.Get("third_party_provider").(string))
	remoteID := d.Get("remote_id").(string)

	createInfo := opal.NewCreateMessageChannelInfo(provider, remoteID)

	messageChannel, _, err := client.MessageChannelsApi.CreateMessageChannel(ctx).CreateMessageChannelInfo(*createInfo).Execute()
	if err != nil {
		return diagFromErr(ctx, err)
	}
	tflog.Debug(ctx, "Created message channel", map[string]any{
		"provider": provider,
		"id":       messageChannel.MessageChannelId,
		"remoteID": remoteID,
	})

	d.SetId(messageChannel.MessageChannelId)
	return resourceMessageChannelRead(ctx, d, m)
}

func resourceMessageChannelRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*opal.APIClient)

	id := d.Get("id").(string)
	messageChannel, _, err := client.MessageChannelsApi.GetMessageChannel(ctx, id).Execute()
	if err != nil {
		return diagFromErr(ctx, err)
	}

	d.SetId(messageChannel.MessageChannelId)
	if err := multierror.Append(
		d.Set("third_party_provider", messageChannel.ThirdPartyProvider),
		d.Set("remote_id", messageChannel.RemoteId),
	); err.ErrorOrNil() != nil {
		return diagFromErr(ctx, err)
	}

	return nil
}

func resourceMessageChannelDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// TODO: Implement
	return nil
}
