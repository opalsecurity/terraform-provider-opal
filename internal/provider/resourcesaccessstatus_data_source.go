// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	tfTypes "github.com/opal-dev/terraform-provider-opal/internal/provider/types"
	"github.com/opal-dev/terraform-provider-opal/internal/sdk"
	"github.com/opal-dev/terraform-provider-opal/internal/sdk/models/operations"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &ResourcesAccessStatusDataSource{}
var _ datasource.DataSourceWithConfigure = &ResourcesAccessStatusDataSource{}

func NewResourcesAccessStatusDataSource() datasource.DataSource {
	return &ResourcesAccessStatusDataSource{}
}

// ResourcesAccessStatusDataSource is the data source implementation.
type ResourcesAccessStatusDataSource struct {
	client *sdk.OpalAPI
}

// ResourcesAccessStatusDataSourceModel describes the data model.
type ResourcesAccessStatusDataSourceModel struct {
	AccessLevel         *tfTypes.ResourceAccessLevel `tfsdk:"access_level"`
	AccessLevelRemoteID types.String                 `tfsdk:"access_level_remote_id"`
	Cursor              types.String                 `tfsdk:"cursor"`
	ExpirationDate      types.String                 `tfsdk:"expiration_date"`
	PageSize            types.Int64                  `tfsdk:"page_size"`
	ResourceID          types.String                 `tfsdk:"resource_id"`
	Status              types.String                 `tfsdk:"status"`
	UserID              types.String                 `tfsdk:"user_id"`
}

// Metadata returns the data source type name.
func (r *ResourcesAccessStatusDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_resources_access_status"
}

// Schema defines the schema for the data source.
func (r *ResourcesAccessStatusDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "ResourcesAccessStatus DataSource",

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
			"access_level_remote_id": schema.StringAttribute{
				Optional:    true,
				Description: `The remote ID of the access level that you wish to query for the resource. If omitted, the default access level remote ID value (empty string) is used.`,
			},
			"cursor": schema.StringAttribute{
				Optional:    true,
				Description: `The pagination cursor value.`,
			},
			"expiration_date": schema.StringAttribute{
				Computed:    true,
				Description: `The day and time the user's access will expire.`,
			},
			"page_size": schema.Int64Attribute{
				Optional:    true,
				Description: `Number of results to return per page. Default is 200.`,
			},
			"resource_id": schema.StringAttribute{
				Required:    true,
				Description: `The ID of the resource.`,
			},
			"status": schema.StringAttribute{
				Computed:    true,
				Description: `The status of the user's access to the resource. must be one of ["AUTHORIZED", "REQUESTED", "UNAUTHORIZED"]`,
			},
			"user_id": schema.StringAttribute{
				Required:    true,
				Description: `The ID of the user.`,
			},
		},
	}
}

func (r *ResourcesAccessStatusDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *ResourcesAccessStatusDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *ResourcesAccessStatusDataSourceModel
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

	accessLevelRemoteID := new(string)
	if !data.AccessLevelRemoteID.IsUnknown() && !data.AccessLevelRemoteID.IsNull() {
		*accessLevelRemoteID = data.AccessLevelRemoteID.ValueString()
	} else {
		accessLevelRemoteID = nil
	}
	cursor := new(string)
	if !data.Cursor.IsUnknown() && !data.Cursor.IsNull() {
		*cursor = data.Cursor.ValueString()
	} else {
		cursor = nil
	}
	pageSize := new(int64)
	if !data.PageSize.IsUnknown() && !data.PageSize.IsNull() {
		*pageSize = data.PageSize.ValueInt64()
	} else {
		pageSize = nil
	}
	resourceID := data.ResourceID.ValueString()
	userID := data.UserID.ValueString()
	request := operations.GetResourceUserAccessStatusRequest{
		AccessLevelRemoteID: accessLevelRemoteID,
		Cursor:              cursor,
		PageSize:            pageSize,
		ResourceID:          resourceID,
		UserID:              userID,
	}
	res, err := r.client.Resources.GetAccessStatus(ctx, request)
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
	if res.ResourceUserAccessStatus == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedResourceUserAccessStatus(res.ResourceUserAccessStatus)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
