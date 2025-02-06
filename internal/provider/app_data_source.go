// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	tfTypes "github.com/opalsecurity/terraform-provider-opal/internal/provider/types"
	"github.com/opalsecurity/terraform-provider-opal/internal/sdk"
	"github.com/opalsecurity/terraform-provider-opal/internal/sdk/models/operations"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &AppDataSource{}
var _ datasource.DataSourceWithConfigure = &AppDataSource{}

func NewAppDataSource() datasource.DataSource {
	return &AppDataSource{}
}

// AppDataSource is the data source implementation.
type AppDataSource struct {
	client *sdk.OpalAPI
}

// AppDataSourceModel describes the data model.
type AppDataSourceModel struct {
	AdminOwnerID types.String            `tfsdk:"admin_owner_id"`
	Description  types.String            `tfsdk:"description"`
	ID           types.String            `tfsdk:"id"`
	Name         types.String            `tfsdk:"name"`
	Type         types.String            `tfsdk:"type"`
	Validations  []tfTypes.AppValidation `tfsdk:"validations"`
}

// Metadata returns the data source type name.
func (r *AppDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_app"
}

// Schema defines the schema for the data source.
func (r *AppDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "App DataSource",

		Attributes: map[string]schema.Attribute{
			"admin_owner_id": schema.StringAttribute{
				Computed:    true,
				Description: `The ID of the owner of the app.`,
			},
			"description": schema.StringAttribute{
				Computed:    true,
				Description: `A description of the app.`,
			},
			"id": schema.StringAttribute{
				Required:    true,
				Description: `The ID of the app.`,
			},
			"name": schema.StringAttribute{
				Computed:    true,
				Description: `The name of the app.`,
			},
			"type": schema.StringAttribute{
				Computed:    true,
				Description: `The type of an app.`,
			},
			"validations": schema.ListNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"details": schema.StringAttribute{
							Computed:    true,
							Description: `Extra details regarding the validation. Could be an error message or restrictions on permissions.`,
						},
						"key": schema.StringAttribute{
							Computed:    true,
							Description: `The key of the app validation. These are not unique IDs between runs.`,
						},
						"name": schema.StringAttribute{
							Computed:    true,
							Description: `The human-readable description of whether the validation has the permissions. Parsed as JSON.`,
						},
						"severity": schema.StringAttribute{
							Computed:    true,
							Description: `The severity of an app validation.`,
						},
						"status": schema.StringAttribute{
							Computed:    true,
							Description: `The status of an app validation.`,
						},
						"updated_at": schema.StringAttribute{
							Computed:    true,
							Description: `The date and time the app validation was last run.`,
						},
						"usage_reason": schema.StringAttribute{
							Computed:    true,
							Description: `The reason for needing the validation.`,
						},
					},
				},
				Description: `Validation checks of an apps' configuration and permissions.`,
			},
		},
	}
}

func (r *AppDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *AppDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *AppDataSourceModel
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

	request := operations.GetAppIDRequest{
		ID: id,
	}
	res, err := r.client.Apps.GetID(ctx, request)
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
	if !(res.App != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedApp(res.App)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
