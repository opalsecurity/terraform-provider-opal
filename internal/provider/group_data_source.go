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
var _ datasource.DataSource = &GroupDataSource{}
var _ datasource.DataSourceWithConfigure = &GroupDataSource{}

func NewGroupDataSource() datasource.DataSource {
	return &GroupDataSource{}
}

// GroupDataSource is the data source implementation.
type GroupDataSource struct {
	client *sdk.OpalAPI
}

// GroupDataSourceModel describes the data model.
type GroupDataSourceModel struct {
	AdminOwnerID          types.String                                `tfsdk:"admin_owner_id"`
	AppID                 types.String                                `tfsdk:"app_id"`
	Description           types.String                                `tfsdk:"description"`
	GroupBindingID        types.String                                `tfsdk:"group_binding_id"`
	GroupLeaderUserIds    []types.String                              `tfsdk:"group_leader_user_ids"`
	GroupType             types.String                                `tfsdk:"group_type"`
	ID                    types.String                                `tfsdk:"id"`
	MessageChannels       tfTypes.GetGroupMessageChannelsResponseBody `tfsdk:"message_channels"`
	Name                  types.String                                `tfsdk:"name"`
	OncallSchedules       tfTypes.GetGroupOnCallSchedulesResponseBody `tfsdk:"oncall_schedules"`
	RemoteInfo            *tfTypes.GroupRemoteInfo                    `tfsdk:"remote_info"`
	RemoteName            types.String                                `tfsdk:"remote_name"`
	RequestConfigurations []tfTypes.RequestConfiguration              `tfsdk:"request_configurations"`
	RequireMfaToApprove   types.Bool                                  `tfsdk:"require_mfa_to_approve"`
	Visibility            types.String                                `tfsdk:"visibility"`
	VisibilityGroupIds    []types.String                              `tfsdk:"visibility_group_ids"`
}

// Metadata returns the data source type name.
func (r *GroupDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_group"
}

// Schema defines the schema for the data source.
func (r *GroupDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Group DataSource",

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
				Required:    true,
				Description: `The ID of the group.`,
			},
			"message_channels": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"channels": schema.ListNestedAttribute{
						Computed: true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"id": schema.StringAttribute{
									Computed:    true,
									Description: `The ID of the message channel.`,
								},
								"is_private": schema.BoolAttribute{
									Computed:    true,
									Description: `A bool representing whether or not the message channel is private.`,
								},
								"name": schema.StringAttribute{
									Computed:    true,
									Description: `The name of the message channel.`,
								},
								"remote_id": schema.StringAttribute{
									Computed:    true,
									Description: `The remote ID of the message channel`,
								},
								"third_party_provider": schema.StringAttribute{
									Computed:    true,
									Description: `The third party provider of the message channel. must be one of ["SLACK"]`,
								},
							},
						},
					},
				},
				Description: `The audit and reviewer message channels attached to the group.`,
			},
			"name": schema.StringAttribute{
				Computed:    true,
				Description: `The name of the group.`,
			},
			"oncall_schedules": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Computed:    true,
						Description: `The ID of the on-call schedule.`,
					},
					"name": schema.StringAttribute{
						Computed:    true,
						Description: `The name of the on call schedule.`,
					},
					"remote_id": schema.StringAttribute{
						Computed:    true,
						Description: `The remote ID of the on call schedule`,
					},
					"third_party_provider": schema.StringAttribute{
						Computed:    true,
						Description: `The third party provider of the on call schedule. must be one of ["OPSGENIE", "PAGER_DUTY"]`,
					},
				},
				Description: `The on call schedules attached to the group.`,
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
								Description: `The id of the Microsoft Entra ID Microsoft 365 group.`,
							},
						},
						Description: `Remote info for Microsoft Entra ID Microsoft 365 group.`,
					},
					"azure_ad_security_group": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"group_id": schema.StringAttribute{
								Computed:    true,
								Description: `The id of the Microsoft Entra ID Security group.`,
							},
						},
						Description: `Remote info for Microsoft Entra ID Security group.`,
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
			"visibility": schema.StringAttribute{
				Computed:    true,
				Description: `The visibility level of the entity. must be one of ["GLOBAL", "LIMITED"]`,
			},
			"visibility_group_ids": schema.SetAttribute{
				Computed:    true,
				ElementType: types.StringType,
			},
		},
	}
}

func (r *GroupDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *GroupDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *GroupDataSourceModel
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

	request := operations.GetGroupRequest{
		ID: id,
	}
	res, err := r.client.Groups.GetGroup(ctx, request)
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
	if !(res.Group != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedGroup(res.Group)
	var id1 string
	id1 = data.ID.ValueString()

	request1 := operations.GetGroupMessageChannelsRequest{
		ID: id1,
	}
	res1, err := r.client.Groups.GetMessageChannels(ctx, request1)
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
	if res1.StatusCode == 404 {
		resp.State.RemoveResource(ctx)
		return
	}
	if res1.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res1.StatusCode), debugResponse(res1.RawResponse))
		return
	}
	if !(res1.Object != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res1.RawResponse))
		return
	}
	data.RefreshFromOperationsGetGroupMessageChannelsResponseBody(res1.Object)
	var id2 string
	id2 = data.ID.ValueString()

	request2 := operations.GetGroupOnCallSchedulesRequest{
		ID: id2,
	}
	res2, err := r.client.Groups.GetOnCallSchedule(ctx, request2)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res2 != nil && res2.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(res2.RawResponse))
		}
		return
	}
	if res2 == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res2))
		return
	}
	if res2.StatusCode == 404 {
		resp.State.RemoveResource(ctx)
		return
	}
	if res2.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res2.StatusCode), debugResponse(res2.RawResponse))
		return
	}
	var id3 string
	id3 = data.ID.ValueString()

	request3 := operations.GetGroupVisibilityRequest{
		ID: id3,
	}
	res3, err := r.client.Groups.GetVisibility(ctx, request3)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res3 != nil && res3.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(res3.RawResponse))
		}
		return
	}
	if res3 == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res3))
		return
	}
	if res3.StatusCode == 404 {
		resp.State.RemoveResource(ctx)
		return
	}
	if res3.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res3.StatusCode), debugResponse(res3.RawResponse))
		return
	}
	if !(res3.Object != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res3.RawResponse))
		return
	}
	data.RefreshFromOperationsGetGroupVisibilityResponseBody(res3.Object)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
