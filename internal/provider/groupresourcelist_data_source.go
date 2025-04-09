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
var _ datasource.DataSource = &GroupResourceListDataSource{}
var _ datasource.DataSourceWithConfigure = &GroupResourceListDataSource{}

func NewGroupResourceListDataSource() datasource.DataSource {
	return &GroupResourceListDataSource{}
}

// GroupResourceListDataSource is the data source implementation.
type GroupResourceListDataSource struct {
	client *sdk.OpalAPI
}

// GroupResourceListDataSourceModel describes the data model.
type GroupResourceListDataSourceModel struct {
	GroupID        types.String            `tfsdk:"group_id"`
	GroupResources []tfTypes.GroupResource `tfsdk:"group_resources"`
}

// Metadata returns the data source type name.
func (r *GroupResourceListDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_group_resource_list"
}

// Schema defines the schema for the data source.
func (r *GroupResourceListDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "GroupResourceList DataSource",

		Attributes: map[string]schema.Attribute{
			"group_id": schema.StringAttribute{
				Required:    true,
				Description: `The ID of the group.`,
			},
			"group_resources": schema.ListNestedAttribute{
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
						"group_id": schema.StringAttribute{
							Computed:    true,
							Description: `The ID of the group.`,
						},
						"resource_id": schema.StringAttribute{
							Computed:    true,
							Description: `The ID of the resource.`,
						},
					},
				},
			},
		},
	}
}

func (r *GroupResourceListDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *GroupResourceListDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *GroupResourceListDataSourceModel
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

	var groupID string
	groupID = data.GroupID.ValueString()

	request := operations.GetGroupResourcesRequest{
		GroupID: groupID,
	}
	res, err := r.client.Groups.GetResources(ctx, request)
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
	if !(res.GroupResourceList != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	resp.Diagnostics.Append(data.RefreshFromSharedGroupResourceList(ctx, res.GroupResourceList)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
