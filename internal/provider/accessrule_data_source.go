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
var _ datasource.DataSource = &AccessRuleDataSource{}
var _ datasource.DataSourceWithConfigure = &AccessRuleDataSource{}

func NewAccessRuleDataSource() datasource.DataSource {
	return &AccessRuleDataSource{}
}

// AccessRuleDataSource is the data source implementation.
type AccessRuleDataSource struct {
	client *sdk.OpalAPI
}

// AccessRuleDataSourceModel describes the data model.
type AccessRuleDataSourceModel struct {
	AdminOwnerID types.String        `tfsdk:"admin_owner_id"`
	Description  types.String        `tfsdk:"description"`
	ID           types.String        `tfsdk:"id"`
	Name         types.String        `tfsdk:"name"`
	RuleClauses  tfTypes.RuleClauses `tfsdk:"rule_clauses"`
	Status       types.String        `tfsdk:"status"`
}

// Metadata returns the data source type name.
func (r *AccessRuleDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_access_rule"
}

// Schema defines the schema for the data source.
func (r *AccessRuleDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "AccessRule DataSource",

		Attributes: map[string]schema.Attribute{
			"admin_owner_id": schema.StringAttribute{
				Computed:    true,
				Description: `The ID of the owner of the group.`,
			},
			"description": schema.StringAttribute{
				Computed:    true,
				Description: `A description of the group.`,
			},
			"id": schema.StringAttribute{
				Computed:    true,
				Description: `The access rule ID (group ID) of the access rule.`,
			},
			"name": schema.StringAttribute{
				Computed:    true,
				Description: `The name of the access rule.`,
			},
			"rule_clauses": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"unless": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"clauses": schema.ListNestedAttribute{
								Computed: true,
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"selectors": schema.ListNestedAttribute{
											Computed: true,
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"connection_id": schema.StringAttribute{
														Computed: true,
													},
													"key": schema.StringAttribute{
														Computed: true,
													},
													"value": schema.StringAttribute{
														Computed: true,
													},
												},
											},
										},
									},
								},
							},
						},
					},
					"when": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"clauses": schema.ListNestedAttribute{
								Computed: true,
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"selectors": schema.ListNestedAttribute{
											Computed: true,
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"connection_id": schema.StringAttribute{
														Computed: true,
													},
													"key": schema.StringAttribute{
														Computed: true,
													},
													"value": schema.StringAttribute{
														Computed: true,
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			"status": schema.StringAttribute{
				Computed:    true,
				Description: `The status of the access rule.`,
			},
		},
	}
}

func (r *AccessRuleDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *AccessRuleDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *AccessRuleDataSourceModel
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

	request := operations.GetAccessRuleRequest{
		ID: id,
	}
	res, err := r.client.AccessRules.GetAccessRule(ctx, request)
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
	if !(res.AccessRule != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	resp.Diagnostics.Append(data.RefreshFromSharedAccessRule(ctx, res.AccessRule)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
