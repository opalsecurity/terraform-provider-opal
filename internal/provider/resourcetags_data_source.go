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
var _ datasource.DataSource = &ResourceTagsDataSource{}
var _ datasource.DataSourceWithConfigure = &ResourceTagsDataSource{}

func NewResourceTagsDataSource() datasource.DataSource {
	return &ResourceTagsDataSource{}
}

// ResourceTagsDataSource is the data source implementation.
type ResourceTagsDataSource struct {
	client *sdk.OpalAPI
}

// ResourceTagsDataSourceModel describes the data model.
type ResourceTagsDataSourceModel struct {
	ResourceID types.String  `tfsdk:"resource_id"`
	Tags       []tfTypes.Tag `tfsdk:"tags"`
}

// Metadata returns the data source type name.
func (r *ResourceTagsDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_resource_tags"
}

// Schema defines the schema for the data source.
func (r *ResourceTagsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "ResourceTags DataSource",

		Attributes: map[string]schema.Attribute{
			"resource_id": schema.StringAttribute{
				Required:    true,
				Description: `The ID of the resource whose tags to return.`,
			},
			"tags": schema.ListNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"created_at": schema.StringAttribute{
							Computed:    true,
							Description: `The date the tag was created.`,
						},
						"id": schema.StringAttribute{
							Computed:    true,
							Description: `The ID of the tag.`,
						},
						"key": schema.StringAttribute{
							Computed:    true,
							Description: `The key of the tag.`,
						},
						"updated_at": schema.StringAttribute{
							Computed:    true,
							Description: `The date the tag was last updated.`,
						},
						"user_creator_id": schema.StringAttribute{
							Computed:    true,
							Description: `The ID of the user that created the tag.`,
						},
						"value": schema.StringAttribute{
							Computed:    true,
							Description: `The value of the tag.`,
						},
					},
				},
			},
		},
	}
}

func (r *ResourceTagsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *ResourceTagsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *ResourceTagsDataSourceModel
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

	var resourceID string
	resourceID = data.ResourceID.ValueString()

	request := operations.GetResourceTagsRequest{
		ResourceID: resourceID,
	}
	res, err := r.client.Resources.GetTags(ctx, request)
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
	if !(res.TagsList != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedTagsList(res.TagsList)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
