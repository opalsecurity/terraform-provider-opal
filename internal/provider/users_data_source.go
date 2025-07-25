// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	tfTypes "github.com/opalsecurity/terraform-provider-opal/v3/internal/provider/types"
	"github.com/opalsecurity/terraform-provider-opal/v3/internal/sdk"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &UsersDataSource{}
var _ datasource.DataSourceWithConfigure = &UsersDataSource{}

func NewUsersDataSource() datasource.DataSource {
	return &UsersDataSource{}
}

// UsersDataSource is the data source implementation.
type UsersDataSource struct {
	// Provider configured SDK client.
	client *sdk.OpalAPI
}

// UsersDataSourceModel describes the data model.
type UsersDataSourceModel struct {
	Cursor   types.String   `queryParam:"style=form,explode=true,name=cursor" tfsdk:"cursor"`
	Next     types.String   `tfsdk:"next"`
	PageSize types.Int64    `queryParam:"style=form,explode=true,name=page_size" tfsdk:"page_size"`
	Previous types.String   `tfsdk:"previous"`
	Results  []tfTypes.User `tfsdk:"results"`
}

// Metadata returns the data source type name.
func (r *UsersDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_users"
}

// Schema defines the schema for the data source.
func (r *UsersDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Users DataSource",

		Attributes: map[string]schema.Attribute{
			"cursor": schema.StringAttribute{
				Optional:    true,
				Description: `The pagination cursor value.`,
			},
			"next": schema.StringAttribute{
				Computed:    true,
				Description: `The cursor with which to continue pagination if additional result pages exist.`,
			},
			"page_size": schema.Int64Attribute{
				Optional:    true,
				Description: `Number of results to return per page. Default is 200.`,
				Validators: []validator.Int64{
					int64validator.AtMost(1000),
				},
			},
			"previous": schema.StringAttribute{
				Computed:    true,
				Description: `The cursor used to obtain the current result page.`,
			},
			"results": schema.ListNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"email": schema.StringAttribute{
							Computed:    true,
							Description: `The email of the user.`,
						},
						"first_name": schema.StringAttribute{
							Computed:    true,
							Description: `The first name of the user.`,
						},
						"hr_idp_status": schema.StringAttribute{
							Computed:    true,
							Description: `User status pulled from an HR/IDP provider.`,
						},
						"id": schema.StringAttribute{
							Computed:    true,
							Description: `The ID of the user.`,
						},
						"last_name": schema.StringAttribute{
							Computed:    true,
							Description: `The last name of the user.`,
						},
						"name": schema.StringAttribute{
							Computed:    true,
							Description: `The full name of the user.`,
						},
						"position": schema.StringAttribute{
							Computed:    true,
							Description: `The user's position.`,
						},
					},
				},
			},
		},
	}
}

func (r *UsersDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *UsersDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *UsersDataSourceModel
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

	request, requestDiags := data.ToOperationsGetUsersRequest(ctx)
	resp.Diagnostics.Append(requestDiags...)

	if resp.Diagnostics.HasError() {
		return
	}
	res, err := r.client.Users.GetUsers(ctx, *request)
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
	if res.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}
	if !(res.PaginatedUsersList != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	resp.Diagnostics.Append(data.RefreshFromSharedPaginatedUsersList(ctx, res.PaginatedUsersList)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
