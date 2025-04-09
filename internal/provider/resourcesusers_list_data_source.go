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
var _ datasource.DataSource = &ResourcesUsersListDataSource{}
var _ datasource.DataSourceWithConfigure = &ResourcesUsersListDataSource{}

func NewResourcesUsersListDataSource() datasource.DataSource {
	return &ResourcesUsersListDataSource{}
}

// ResourcesUsersListDataSource is the data source implementation.
type ResourcesUsersListDataSource struct {
	client *sdk.OpalAPI
}

// ResourcesUsersListDataSourceModel describes the data model.
type ResourcesUsersListDataSourceModel struct {
	Limit      types.Int64                  `queryParam:"style=form,explode=true,name=limit" tfsdk:"limit"`
	ResourceID types.String                 `tfsdk:"resource_id"`
	Results    []tfTypes.ResourceAccessUser `tfsdk:"results"`
}

// Metadata returns the data source type name.
func (r *ResourcesUsersListDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_resources_users_list"
}

// Schema defines the schema for the data source.
func (r *ResourcesUsersListDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "ResourcesUsersList DataSource",

		Attributes: map[string]schema.Attribute{
			"limit": schema.Int64Attribute{
				Optional:    true,
				Description: `Limit the number of results returned.`,
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
								`The ` + "`" + `AccessLevel` + "`" + ` object is used to represent the level of access that a principal has. The "default" access` + "\n" +
								`level is a ` + "`" + `AccessLevel` + "`" + ` object whose fields are all empty strings.` + "\n" +
								`` + "\n" +
								`### Usage Example` + "\n" +
								`View the ` + "`" + `AccessLevel` + "`" + ` of a resource/user or resource/group pair to see the level of access granted to the resource.`,
						},
						"email": schema.StringAttribute{
							Computed:    true,
							Description: `The user's email.`,
						},
						"expiration_date": schema.StringAttribute{
							Computed:    true,
							Description: `The day and time the user's access will expire.`,
						},
						"full_name": schema.StringAttribute{
							Computed:    true,
							Description: `The user's full name.`,
						},
						"has_direct_access": schema.BoolAttribute{
							Computed:    true,
							Description: `The user has direct access to this resources (vs. indirectly, like through a group).`,
						},
						"num_access_paths": schema.Int32Attribute{
							Computed:    true,
							Description: `The number of ways in which the user has access through this resource (directly and indirectly).`,
						},
						"propagation_status": schema.SingleNestedAttribute{
							Computed: true,
							Attributes: map[string]schema.Attribute{
								"status": schema.StringAttribute{
									Computed:    true,
									Description: `The status of whether the user has been synced to the group or resource in the remote system.`,
								},
							},
							Description: `The state of whether the push action was propagated to the remote system. If this is null, the access was synced from the remote system.`,
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
		},
	}
}

func (r *ResourcesUsersListDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *ResourcesUsersListDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *ResourcesUsersListDataSourceModel
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

	limit := new(int64)
	if !data.Limit.IsUnknown() && !data.Limit.IsNull() {
		*limit = data.Limit.ValueInt64()
	} else {
		limit = nil
	}
	var resourceID string
	resourceID = data.ResourceID.ValueString()

	request := operations.GetResourceUsersRequest{
		Limit:      limit,
		ResourceID: resourceID,
	}
	res, err := r.client.Resources.GetUsers(ctx, request)
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
	if !(res.ResourceAccessUserList != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	resp.Diagnostics.Append(data.RefreshFromSharedResourceAccessUserList(ctx, res.ResourceAccessUserList)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
