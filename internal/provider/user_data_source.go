// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/opal-dev/terraform-provider-opal/internal/sdk"
	"github.com/opal-dev/terraform-provider-opal/internal/sdk/models/operations"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &UserDataSource{}
var _ datasource.DataSourceWithConfigure = &UserDataSource{}

func NewUserDataSource() datasource.DataSource {
	return &UserDataSource{}
}

// UserDataSource is the data source implementation.
type UserDataSource struct {
	client *sdk.OpalAPI
}

// UserDataSourceModel describes the data model.
type UserDataSourceModel struct {
	Email       types.String `tfsdk:"email"`
	FirstName   types.String `tfsdk:"first_name"`
	HrIdpStatus types.String `tfsdk:"hr_idp_status"`
	ID          types.String `tfsdk:"id"`
	LastName    types.String `tfsdk:"last_name"`
	Name        types.String `tfsdk:"name"`
	Position    types.String `tfsdk:"position"`
}

// Metadata returns the data source type name.
func (r *UserDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_user"
}

// Schema defines the schema for the data source.
func (r *UserDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "User DataSource",

		Attributes: map[string]schema.Attribute{
			"email": schema.StringAttribute{
				Computed:    true,
				Optional:    true,
				Description: `The email of the user. If both user ID and email are provided, user ID will take precedence. If neither are provided, an error will occur.`,
			},
			"first_name": schema.StringAttribute{
				Computed:    true,
				Description: `The first name of the user.`,
			},
			"hr_idp_status": schema.StringAttribute{
				Computed:    true,
				Description: `User status pulled from an HR/IDP provider. must be one of ["ACTIVE", "SUSPENDED", "DEPROVISIONED", "DELETED", "NOT_FOUND"]`,
			},
			"id": schema.StringAttribute{
				Computed:    true,
				Optional:    true,
				Description: `The user ID of the user.`,
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
	}
}

func (r *UserDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *UserDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *UserDataSourceModel
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

	email := new(string)
	if !data.Email.IsUnknown() && !data.Email.IsNull() {
		*email = data.Email.ValueString()
	} else {
		email = nil
	}
	id := new(string)
	if !data.ID.IsUnknown() && !data.ID.IsNull() {
		*id = data.ID.ValueString()
	} else {
		id = nil
	}
	request := operations.GetUserRequest{
		Email: email,
		ID:    id,
	}
	res, err := r.client.Users.Get(ctx, request)
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
	if res.User == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedUser(res.User)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
