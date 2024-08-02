// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/opalsecurity/terraform-provider-opal/internal/sdk"
	"github.com/opalsecurity/terraform-provider-opal/internal/sdk/models/operations"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &MessageChannelDataSource{}
var _ datasource.DataSourceWithConfigure = &MessageChannelDataSource{}

func NewMessageChannelDataSource() datasource.DataSource {
	return &MessageChannelDataSource{}
}

// MessageChannelDataSource is the data source implementation.
type MessageChannelDataSource struct {
	client *sdk.OpalAPI
}

// MessageChannelDataSourceModel describes the data model.
type MessageChannelDataSourceModel struct {
	ID                 types.String `tfsdk:"id"`
	IsPrivate          types.Bool   `tfsdk:"is_private"`
	Name               types.String `tfsdk:"name"`
	RemoteID           types.String `tfsdk:"remote_id"`
	ThirdPartyProvider types.String `tfsdk:"third_party_provider"`
}

// Metadata returns the data source type name.
func (r *MessageChannelDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_message_channel"
}

// Schema defines the schema for the data source.
func (r *MessageChannelDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "MessageChannel DataSource",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Required:    true,
				Description: `The ID of the message_channel.`,
			},
			"is_private": schema.BoolAttribute{
				Computed:    true,
				Description: `A bool representing whether or not the message channel is private.`,
			},
			"name": schema.StringAttribute{
				Computed:    true,
				Description: `The name of the message channel.`,
			},
			"remote_id": schema.StringAttribute{
				Computed:    true,
				Description: `The remote ID of the message channel`,
			},
			"third_party_provider": schema.StringAttribute{
				Computed:    true,
				Description: `The third party provider of the message channel. must be one of ["SLACK"]`,
			},
		},
	}
}

func (r *MessageChannelDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*sdk.OpalAPI)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected DataSource Configure Type",
			fmt.Sprintf("Expected *sdk.OpalAPI, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *MessageChannelDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *MessageChannelDataSourceModel
	var item types.Object

	resp.Diagnostics.Append(req.Config.Get(ctx, &item)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(item.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}

	var id string
	id = data.ID.ValueString()

	request := operations.GetMessageChannelIDRequest{
		ID: id,
	}
	res, err := r.client.MessageChannels.GetID(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res != nil && res.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(res.RawResponse))
		}
		return
	}
	if res == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res))
		return
	}
	if res.StatusCode == 404 {
		resp.State.RemoveResource(ctx)
		return
	}
	if res.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}
	if !(res.MessageChannel != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedMessageChannel(res.MessageChannel)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
