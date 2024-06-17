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
	"github.com/opalsecurity/terraform-provider-opal/internal/sdk/models/shared"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &GroupListDataSource{}
var _ datasource.DataSourceWithConfigure = &GroupListDataSource{}

func NewGroupListDataSource() datasource.DataSource {
	return &GroupListDataSource{}
}

// GroupListDataSource is the data source implementation.
type GroupListDataSource struct {
	client *sdk.OpalAPI
}

// GroupListDataSourceModel describes the data model.
type GroupListDataSourceModel struct {
	GroupIds        []types.String  `tfsdk:"group_ids"`
	GroupName       types.String    `tfsdk:"group_name"`
	GroupTypeFilter types.String    `tfsdk:"group_type_filter"`
	PageSize        types.Int64     `tfsdk:"page_size"`
	Results         []tfTypes.Group `tfsdk:"results"`
}

// Metadata returns the data source type name.
func (r *GroupListDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_group_list"
}

// Schema defines the schema for the data source.
func (r *GroupListDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "GroupList DataSource",

		Attributes: map[string]schema.Attribute{
			"group_ids": schema.ListAttribute{
				Optional:    true,
				ElementType: types.StringType,
				Description: `The group ids to filter by.`,
			},
			"group_name": schema.StringAttribute{
				Optional:    true,
				Description: `Group name.`,
			},
			"group_type_filter": schema.StringAttribute{
				Optional:    true,
				Description: `The type of the group. must be one of ["ACTIVE_DIRECTORY_GROUP", "AWS_SSO_GROUP", "DUO_GROUP", "GIT_HUB_TEAM", "GIT_LAB_GROUP", "GOOGLE_GROUPS_GROUP", "LDAP_GROUP", "OKTA_GROUP", "OPAL_GROUP", "AZURE_AD_SECURITY_GROUP", "AZURE_AD_MICROSOFT_365_GROUP"]`,
			},
			"page_size": schema.Int64Attribute{
				Optional:    true,
				Description: `Number of results to return per page. Default is 200.`,
			},
			"results": schema.ListNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"admin_owner_id": schema.StringAttribute{
							Computed:    true,
							Description: `The ID of the owner of the group.`,
						},
						"app_id": schema.StringAttribute{
							Computed:    true,
							Description: `The ID of the group's app.`,
						},
						"description": schema.StringAttribute{
							Computed:    true,
							Description: `A description of the group.`,
						},
						"group_binding_id": schema.StringAttribute{
							Computed:    true,
							Description: `The ID of the associated group binding.`,
						},
						"group_leader_user_ids": schema.SetAttribute{
							Computed:    true,
							ElementType: types.StringType,
							Description: `A list of User IDs for the group leaders of the group`,
						},
						"group_type": schema.StringAttribute{
							Computed:    true,
							Description: `The type of the group. must be one of ["ACTIVE_DIRECTORY_GROUP", "AWS_SSO_GROUP", "DUO_GROUP", "GIT_HUB_TEAM", "GIT_LAB_GROUP", "GOOGLE_GROUPS_GROUP", "LDAP_GROUP", "OKTA_GROUP", "OPAL_GROUP", "AZURE_AD_SECURITY_GROUP", "AZURE_AD_MICROSOFT_365_GROUP"]`,
						},
						"id": schema.StringAttribute{
							Computed:    true,
							Description: `The ID of the group.`,
						},
						"name": schema.StringAttribute{
							Computed:    true,
							Description: `The name of the group.`,
						},
						"remote_info": schema.SingleNestedAttribute{
							Computed: true,
							Attributes: map[string]schema.Attribute{
								"active_directory_group": schema.SingleNestedAttribute{
									Computed: true,
									Attributes: map[string]schema.Attribute{
										"group_id": schema.StringAttribute{
											Computed:    true,
											Description: `The id of the Google group.`,
										},
									},
									Description: `Remote info for Active Directory group.`,
								},
								"azure_ad_microsoft_365_group": schema.SingleNestedAttribute{
									Computed: true,
									Attributes: map[string]schema.Attribute{
										"group_id": schema.StringAttribute{
											Computed:    true,
											Description: `The id of the Azure AD Microsoft 365 group.`,
										},
									},
									Description: `Remote info for Azure AD Microsoft 365 group.`,
								},
								"azure_ad_security_group": schema.SingleNestedAttribute{
									Computed: true,
									Attributes: map[string]schema.Attribute{
										"group_id": schema.StringAttribute{
											Computed:    true,
											Description: `The id of the Azure AD Security group.`,
										},
									},
									Description: `Remote info for Azure AD Security group.`,
								},
								"duo_group": schema.SingleNestedAttribute{
									Computed: true,
									Attributes: map[string]schema.Attribute{
										"group_id": schema.StringAttribute{
											Computed:    true,
											Description: `The id of the Duo Security group.`,
										},
									},
									Description: `Remote info for Duo Security group.`,
								},
								"github_team": schema.SingleNestedAttribute{
									Computed: true,
									Attributes: map[string]schema.Attribute{
										"team_slug": schema.StringAttribute{
											Computed:    true,
											Description: `The slug of the GitHub team.`,
										},
									},
									Description: `Remote info for GitHub team.`,
								},
								"gitlab_group": schema.SingleNestedAttribute{
									Computed: true,
									Attributes: map[string]schema.Attribute{
										"group_id": schema.StringAttribute{
											Computed:    true,
											Description: `The id of the Gitlab group.`,
										},
									},
									Description: `Remote info for Gitlab group.`,
								},
								"google_group": schema.SingleNestedAttribute{
									Computed: true,
									Attributes: map[string]schema.Attribute{
										"group_id": schema.StringAttribute{
											Computed:    true,
											Description: `The id of the Google group.`,
										},
									},
									Description: `Remote info for Google group.`,
								},
								"ldap_group": schema.SingleNestedAttribute{
									Computed: true,
									Attributes: map[string]schema.Attribute{
										"group_id": schema.StringAttribute{
											Computed:    true,
											Description: `The id of the LDAP group.`,
										},
									},
									Description: `Remote info for LDAP group.`,
								},
								"okta_group": schema.SingleNestedAttribute{
									Computed: true,
									Attributes: map[string]schema.Attribute{
										"group_id": schema.StringAttribute{
											Computed:    true,
											Description: `The id of the Okta Directory group.`,
										},
									},
									Description: `Remote info for Okta Directory group.`,
								},
							},
							Description: `Information that defines the remote group. This replaces the deprecated remote_id and metadata fields.`,
						},
						"remote_name": schema.StringAttribute{
							Computed:    true,
							Description: `The name of the remote.`,
						},
						"request_configurations": schema.ListNestedAttribute{
							Computed: true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"allow_requests": schema.BoolAttribute{
										Computed:    true,
										Description: `A bool representing whether or not to allow requests for this resource.`,
									},
									"auto_approval": schema.BoolAttribute{
										Computed:    true,
										Description: `A bool representing whether or not to automatically approve requests for this resource.`,
									},
									"condition": schema.SingleNestedAttribute{
										Computed: true,
										Attributes: map[string]schema.Attribute{
											"group_ids": schema.SetAttribute{
												Computed:    true,
												ElementType: types.StringType,
												Description: `The list of group IDs to match.`,
											},
											"role_remote_ids": schema.SetAttribute{
												Computed:    true,
												ElementType: types.StringType,
												Description: `The list of role remote IDs to match.`,
											},
										},
									},
									"max_duration": schema.Int64Attribute{
										Computed:    true,
										Description: `The maximum duration for which the resource can be requested (in minutes).`,
									},
									"priority": schema.Int64Attribute{
										Computed:    true,
										Description: `The priority of the request configuration.`,
									},
									"recommended_duration": schema.Int64Attribute{
										Computed:    true,
										Description: `The recommended duration for which the resource should be requested (in minutes). -1 represents an indefinite duration.`,
									},
									"request_template_id": schema.StringAttribute{
										Computed:    true,
										Description: `The ID of the associated request template.`,
									},
									"require_mfa_to_request": schema.BoolAttribute{
										Computed:    true,
										Description: `A bool representing whether or not to require MFA for requesting access to this resource.`,
									},
									"require_support_ticket": schema.BoolAttribute{
										Computed:    true,
										Description: `A bool representing whether or not access requests to the resource require an access ticket.`,
									},
									"reviewer_stages": schema.ListNestedAttribute{
										Computed: true,
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"operator": schema.StringAttribute{
													Computed:    true,
													Description: `The operator of the reviewer stage. Admin and manager approval are also treated as reviewers. must be one of ["AND", "OR"]`,
												},
												"owner_ids": schema.SetAttribute{
													Computed:    true,
													ElementType: types.StringType,
												},
												"require_admin_approval": schema.BoolAttribute{
													Computed:    true,
													Description: `Whether this reviewer stage should require admin approval.`,
												},
												"require_manager_approval": schema.BoolAttribute{
													Computed:    true,
													Description: `Whether this reviewer stage should require manager approval.`,
												},
											},
										},
										Description: `The list of reviewer stages for the request configuration.`,
									},
								},
							},
							Description: `A list of request configurations for this group.`,
						},
						"require_mfa_to_approve": schema.BoolAttribute{
							Computed:    true,
							Description: `A bool representing whether or not to require MFA for reviewers to approve requests for this group.`,
						},
					},
				},
			},
		},
	}
}

func (r *GroupListDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *GroupListDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *GroupListDataSourceModel
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

	var groupIds []string = []string{}
	for _, groupIdsItem := range data.GroupIds {
		groupIds = append(groupIds, groupIdsItem.ValueString())
	}
	groupName := new(string)
	if !data.GroupName.IsUnknown() && !data.GroupName.IsNull() {
		*groupName = data.GroupName.ValueString()
	} else {
		groupName = nil
	}
	groupTypeFilter := new(shared.GroupTypeEnum)
	if !data.GroupTypeFilter.IsUnknown() && !data.GroupTypeFilter.IsNull() {
		*groupTypeFilter = shared.GroupTypeEnum(data.GroupTypeFilter.ValueString())
	} else {
		groupTypeFilter = nil
	}
	pageSize := new(int64)
	if !data.PageSize.IsUnknown() && !data.PageSize.IsNull() {
		*pageSize = data.PageSize.ValueInt64()
	} else {
		pageSize = nil
	}
	request := operations.GetGroupsRequest{
		GroupIds:        groupIds,
		GroupName:       groupName,
		GroupTypeFilter: groupTypeFilter,
		PageSize:        pageSize,
	}
	res, err := r.client.Groups.List(ctx, request)
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
	if !(res.PaginatedGroupsList != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedPaginatedGroupsList(res.PaginatedGroupsList)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
