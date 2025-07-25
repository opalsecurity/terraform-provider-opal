// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/opalsecurity/terraform-provider-opal/v3/internal/sdk"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &BundleDataSource{}
var _ datasource.DataSourceWithConfigure = &BundleDataSource{}

func NewBundleDataSource() datasource.DataSource {
	return &BundleDataSource{}
}

// BundleDataSource is the data source implementation.
type BundleDataSource struct {
	// Provider configured SDK client.
	client *sdk.OpalAPI
}

// BundleDataSourceModel describes the data model.
type BundleDataSourceModel struct {
	AdminOwnerID       types.String   `tfsdk:"admin_owner_id"`
	BundleID           types.String   `tfsdk:"bundle_id"`
	CreatedAt          types.String   `tfsdk:"created_at"`
	Description        types.String   `tfsdk:"description"`
	Name               types.String   `tfsdk:"name"`
	TotalNumGroups     types.Int64    `tfsdk:"total_num_groups"`
	TotalNumItems      types.Int64    `tfsdk:"total_num_items"`
	TotalNumResources  types.Int64    `tfsdk:"total_num_resources"`
	UpdatedAt          types.String   `tfsdk:"updated_at"`
	Visibility         types.String   `tfsdk:"visibility"`
	VisibilityGroupIds []types.String `tfsdk:"visibility_group_ids"`
}

// Metadata returns the data source type name.
func (r *BundleDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_bundle"
}

// Schema defines the schema for the data source.
func (r *BundleDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Bundle DataSource",

		Attributes: map[string]schema.Attribute{
			"admin_owner_id": schema.StringAttribute{
				Computed:    true,
				Description: `The ID of the owner of the bundle.`,
			},
			"bundle_id": schema.StringAttribute{
				Computed:    true,
				Description: `The ID of the bundle.`,
			},
			"created_at": schema.StringAttribute{
				Computed:    true,
				Description: `The creation timestamp of the bundle, in ISO 8601 format`,
			},
			"description": schema.StringAttribute{
				Computed:    true,
				Description: `The description of the bundle.`,
			},
			"name": schema.StringAttribute{
				Computed:    true,
				Description: `The name of the bundle.`,
			},
			"total_num_groups": schema.Int64Attribute{
				Computed:    true,
				Description: `The total number of groups in the bundle.`,
			},
			"total_num_items": schema.Int64Attribute{
				Computed:    true,
				Description: `The total number of items in the bundle.`,
			},
			"total_num_resources": schema.Int64Attribute{
				Computed:    true,
				Description: `The total number of resources in the bundle.`,
			},
			"updated_at": schema.StringAttribute{
				Computed:    true,
				Description: `The last updated timestamp of the bundle, in ISO 8601 format`,
			},
			"visibility": schema.StringAttribute{
				Computed:    true,
				Description: `The visibility level of the entity.`,
			},
			"visibility_group_ids": schema.SetAttribute{
				Computed:    true,
				ElementType: types.StringType,
			},
		},
	}
}

func (r *BundleDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *BundleDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *BundleDataSourceModel
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

	request, requestDiags := data.ToOperationsGetBundleRequest(ctx)
	resp.Diagnostics.Append(requestDiags...)

	if resp.Diagnostics.HasError() {
		return
	}
	res, err := r.client.Bundles.GetBundle(ctx, *request)
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
	if !(res.Bundle != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	resp.Diagnostics.Append(data.RefreshFromSharedBundle(ctx, res.Bundle)...)

	if resp.Diagnostics.HasError() {
		return
	}
	request1, request1Diags := data.ToOperationsGetBundleVisibilityRequest(ctx)
	resp.Diagnostics.Append(request1Diags...)

	if resp.Diagnostics.HasError() {
		return
	}
	res1, err := r.client.Bundles.GetBundleVisibility(ctx, *request1)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res1 != nil && res1.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(res1.RawResponse))
		}
		return
	}
	if res1 == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res1))
		return
	}
	if res1.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res1.StatusCode), debugResponse(res1.RawResponse))
		return
	}
	if !(res1.VisibilityInfo != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res1.RawResponse))
		return
	}
	resp.Diagnostics.Append(data.RefreshFromSharedVisibilityInfo(ctx, res1.VisibilityInfo)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
