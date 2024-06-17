// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

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
var _ datasource.DataSource = &SessionsDataSource{}
var _ datasource.DataSourceWithConfigure = &SessionsDataSource{}

func NewSessionsDataSource() datasource.DataSource {
	return &SessionsDataSource{}
}

// SessionsDataSource is the data source implementation.
type SessionsDataSource struct {
	client *sdk.OpalAPI
}

// SessionsDataSourceModel describes the data model.
type SessionsDataSourceModel struct {
	Next       types.String      `tfsdk:"next"`
	Previous   types.String      `tfsdk:"previous"`
	ResourceID types.String      `tfsdk:"resource_id"`
	Results    []tfTypes.Session `tfsdk:"results"`
	UserID     types.String      `tfsdk:"user_id"`
}

// Metadata returns the data source type name.
func (r *SessionsDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_sessions"
}

// Schema defines the schema for the data source.
func (r *SessionsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Sessions DataSource",

		Attributes: map[string]schema.Attribute{
			"next": schema.StringAttribute{
				Computed:    true,
				Description: `The cursor with which to continue pagination if additional result pages exist.`,
			},
			"previous": schema.StringAttribute{
				Computed:    true,
				Description: `The cursor used to obtain the current result page.`,
			},
			"resource_id": schema.StringAttribute{
				Required:    true,
				Description: `The ID of the resource.`,
			},
			"results": schema.ListNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"access_level": schema.SingleNestedAttribute{
							Computed: true,
							Attributes: map[string]schema.Attribute{
								"access_level_name": schema.StringAttribute{
									Computed:    true,
									Description: `The human-readable name of the access level.`,
								},
								"access_level_remote_id": schema.StringAttribute{
									Computed:    true,
									Description: `The machine-readable identifier of the access level.`,
								},
							},
							MarkdownDescription: `# Access Level Object` + "\n" +
								`### Description` + "\n" +
								`The ` + "`" + `ResourceAccessLevel` + "`" + ` object is used to represent the level of access that a user has to a resource or a resource has to a group. The "default" access` + "\n" +
								`level is a ` + "`" + `ResourceAccessLevel` + "`" + ` object whose fields are all empty strings.` + "\n" +
								`` + "\n" +
								`### Usage Example` + "\n" +
								`View the ` + "`" + `ResourceAccessLevel` + "`" + ` of a resource/user or resource/group pair to see the level of access granted to the resource.`,
						},
						"connection_id": schema.StringAttribute{
							Computed:    true,
							Description: `The ID of the connection.`,
						},
						"expiration_date": schema.StringAttribute{
							Computed:    true,
							Description: `The day and time the user's access will expire.`,
						},
						"resource_id": schema.StringAttribute{
							Computed:    true,
							Description: `The ID of the resource.`,
						},
						"user_id": schema.StringAttribute{
							Computed:    true,
							Description: `The ID of the user.`,
						},
					},
				},
			},
			"user_id": schema.StringAttribute{
				Optional:    true,
				Description: `The ID of the user you wish to query sessions for.`,
			},
		},
	}
}

func (r *SessionsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *SessionsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *SessionsDataSourceModel
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

	resourceID := data.ResourceID.ValueString()
	userID := new(string)
	if !data.UserID.IsUnknown() && !data.UserID.IsNull() {
		*userID = data.UserID.ValueString()
	} else {
		userID = nil
	}
	request := operations.GetSessionsRequest{
		ResourceID: resourceID,
		UserID:     userID,
	}
	res, err := r.client.Sessions.Get(ctx, request)
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
	if !(res.SessionsList != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedSessionsList(res.SessionsList)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
