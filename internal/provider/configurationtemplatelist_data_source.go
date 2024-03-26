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
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &ConfigurationTemplateListDataSource{}
var _ datasource.DataSourceWithConfigure = &ConfigurationTemplateListDataSource{}

func NewConfigurationTemplateListDataSource() datasource.DataSource {
	return &ConfigurationTemplateListDataSource{}
}

// ConfigurationTemplateListDataSource is the data source implementation.
type ConfigurationTemplateListDataSource struct {
	client *sdk.OpalAPI
}

// ConfigurationTemplateListDataSourceModel describes the data model.
type ConfigurationTemplateListDataSourceModel struct {
	Results []tfTypes.ConfigurationTemplate `tfsdk:"results"`
}

// Metadata returns the data source type name.
func (r *ConfigurationTemplateListDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_configuration_template_list"
}

// Schema defines the schema for the data source.
func (r *ConfigurationTemplateListDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "ConfigurationTemplateList DataSource",

		Attributes: map[string]schema.Attribute{
			"results": schema.ListNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"admin_owner_id": schema.StringAttribute{
							Computed:    true,
							Description: `The ID of the owner of the configuration template.`,
						},
						"break_glass_user_ids": schema.ListAttribute{
							Computed:    true,
							ElementType: types.StringType,
							Description: `The IDs of the break glass users linked to the configuration template.`,
						},
						"configuration_template_id": schema.StringAttribute{
							Computed:    true,
							Description: `The ID of the configuration template.`,
						},
						"linked_audit_message_channel_ids": schema.ListAttribute{
							Computed:    true,
							ElementType: types.StringType,
							Description: `The IDs of the audit message channels linked to the configuration template.`,
						},
						"member_oncall_schedule_ids": schema.ListAttribute{
							Computed:    true,
							ElementType: types.StringType,
							Description: `The IDs of the on-call schedules linked to the configuration template.`,
						},
						"name": schema.StringAttribute{
							Computed:    true,
							Description: `The name of the configuration template.`,
						},
						"request_configuration_id": schema.StringAttribute{
							Computed:    true,
							Description: `The ID of the request configuration linked to the configuration template.`,
						},
						"require_mfa_to_approve": schema.BoolAttribute{
							Computed:    true,
							Description: `A bool representing whether or not to require MFA for reviewers to approve requests for this configuration template.`,
						},
						"require_mfa_to_connect": schema.BoolAttribute{
							Computed:    true,
							Description: `A bool representing whether or not to require MFA to connect to resources associated with this configuration template.`,
						},
						"visibility": schema.SingleNestedAttribute{
							Computed: true,
							Attributes: map[string]schema.Attribute{
								"visibility": schema.StringAttribute{
									Computed:    true,
									Description: `The visibility level of the entity. must be one of ["GLOBAL", "LIMITED"]`,
								},
								"visibility_group_ids": schema.ListAttribute{
									Computed:    true,
									ElementType: types.StringType,
								},
							},
							Description: `Visibility infomation of an entity.`,
						},
					},
				},
			},
		},
	}
}

func (r *ConfigurationTemplateListDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *ConfigurationTemplateListDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *ConfigurationTemplateListDataSourceModel
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

	res, err := r.client.ConfigurationTemplates.Get(ctx)
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
	if res.PaginatedConfigurationTemplateList == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedPaginatedConfigurationTemplateList(res.PaginatedConfigurationTemplateList)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
