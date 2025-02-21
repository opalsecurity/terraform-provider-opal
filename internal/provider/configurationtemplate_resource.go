// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	speakeasy_boolplanmodifier "github.com/opalsecurity/terraform-provider-opal/internal/planmodifiers/boolplanmodifier"
	speakeasy_objectplanmodifier "github.com/opalsecurity/terraform-provider-opal/internal/planmodifiers/objectplanmodifier"
	speakeasy_setplanmodifier "github.com/opalsecurity/terraform-provider-opal/internal/planmodifiers/setplanmodifier"
	speakeasy_stringplanmodifier "github.com/opalsecurity/terraform-provider-opal/internal/planmodifiers/stringplanmodifier"
	tfTypes "github.com/opalsecurity/terraform-provider-opal/internal/provider/types"
	"github.com/opalsecurity/terraform-provider-opal/internal/sdk"
	"github.com/opalsecurity/terraform-provider-opal/internal/sdk/models/operations"
	speakeasy_boolvalidators "github.com/opalsecurity/terraform-provider-opal/internal/validators/boolvalidators"
	custom_objectvalidators "github.com/opalsecurity/terraform-provider-opal/internal/validators/objectvalidators"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &ConfigurationTemplateResource{}
var _ resource.ResourceWithImportState = &ConfigurationTemplateResource{}

func NewConfigurationTemplateResource() resource.Resource {
	return &ConfigurationTemplateResource{}
}

// ConfigurationTemplateResource defines the resource implementation.
type ConfigurationTemplateResource struct {
	client *sdk.OpalAPI
}

// ConfigurationTemplateResourceModel describes the resource data model.
type ConfigurationTemplateResourceModel struct {
	AdminOwnerID                 types.String                            `tfsdk:"admin_owner_id"`
	BreakGlassUserIds            []types.String                          `tfsdk:"break_glass_user_ids"`
	ConfigurationTemplateID      types.String                            `tfsdk:"configuration_template_id"`
	CustomRequestNotification    types.String                            `tfsdk:"custom_request_notification"`
	LinkedAuditMessageChannelIds []types.String                          `tfsdk:"linked_audit_message_channel_ids"`
	MemberOncallScheduleIds      []types.String                          `tfsdk:"member_oncall_schedule_ids"`
	Name                         types.String                            `tfsdk:"name"`
	RequestConfigurationID       types.String                            `tfsdk:"request_configuration_id"`
	RequestConfigurations        []tfTypes.RequestConfiguration          `tfsdk:"request_configurations"`
	RequireMfaToApprove          types.Bool                              `tfsdk:"require_mfa_to_approve"`
	RequireMfaToConnect          types.Bool                              `tfsdk:"require_mfa_to_connect"`
	TicketPropagation            *tfTypes.TicketPropagationConfiguration `tfsdk:"ticket_propagation"`
	Visibility                   tfTypes.VisibilityInfo                  `tfsdk:"visibility"`
}

func (r *ConfigurationTemplateResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_configuration_template"
}

func (r *ConfigurationTemplateResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "ConfigurationTemplate Resource",
		Attributes: map[string]schema.Attribute{
			"admin_owner_id": schema.StringAttribute{
				Required:    true,
				Description: `The ID of the owner of the configuration template.`,
			},
			"break_glass_user_ids": schema.SetAttribute{
				Computed:    true,
				Optional:    true,
				ElementType: types.StringType,
				Description: `The IDs of the break glass users linked to the configuration template.`,
			},
			"configuration_template_id": schema.StringAttribute{
				Computed:    true,
				Description: `The ID of the configuration template.`,
			},
			"custom_request_notification": schema.StringAttribute{
				Computed:    true,
				Optional:    true,
				Description: `Custom request notification sent upon request approval for this configuration template.`,
				Validators: []validator.String{
					stringvalidator.UTF8LengthAtMost(800),
				},
			},
			"linked_audit_message_channel_ids": schema.SetAttribute{
				Computed:    true,
				Optional:    true,
				ElementType: types.StringType,
				Description: `The IDs of the audit message channels linked to the configuration template.`,
			},
			"member_oncall_schedule_ids": schema.SetAttribute{
				Computed:    true,
				Optional:    true,
				Default:     setdefault.StaticValue(types.SetValueMust(types.StringType, []attr.Value{})),
				ElementType: types.StringType,
				Description: `The IDs of the on-call schedules linked to the configuration template.`,
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: `The name of the configuration template.`,
			},
			"request_configuration_id": schema.StringAttribute{
				Computed:    true,
				Description: `The ID of the request configuration linked to the configuration template.`,
			},
			"request_configurations": schema.ListNestedAttribute{
				Optional: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"allow_requests": schema.BoolAttribute{
							Required:    true,
							Description: `A bool representing whether or not to allow requests for this resource.`,
						},
						"auto_approval": schema.BoolAttribute{
							Required:    true,
							Description: `A bool representing whether or not to automatically approve requests for this resource.`,
						},
						"condition": schema.SingleNestedAttribute{
							Optional: true,
							Attributes: map[string]schema.Attribute{
								"group_ids": schema.SetAttribute{
									Computed:    true,
									Optional:    true,
									Default:     setdefault.StaticValue(types.SetValueMust(types.StringType, []attr.Value{})),
									ElementType: types.StringType,
									Description: `The list of group IDs to match.`,
								},
								"role_remote_ids": schema.SetAttribute{
									Computed:    true,
									Optional:    true,
									Default:     setdefault.StaticValue(types.SetValueMust(types.StringType, []attr.Value{})),
									ElementType: types.StringType,
									Description: `The list of role remote IDs to match.`,
								},
							},
						},
						"max_duration": schema.Int64Attribute{
							Optional:    true,
							Description: `The maximum duration for which the resource can be requested (in minutes).`,
						},
						"priority": schema.Int64Attribute{
							Required:    true,
							Description: `The priority of the request configuration.`,
						},
						"recommended_duration": schema.Int64Attribute{
							Optional:    true,
							Description: `The recommended duration for which the resource should be requested (in minutes). -1 represents an indefinite duration.`,
						},
						"request_template_id": schema.StringAttribute{
							Optional:    true,
							Description: `The ID of the associated request template.`,
						},
						"require_mfa_to_request": schema.BoolAttribute{
							Required:    true,
							Description: `A bool representing whether or not to require MFA for requesting access to this resource.`,
						},
						"require_support_ticket": schema.BoolAttribute{
							Required:    true,
							Description: `A bool representing whether or not access requests to the resource require an access ticket.`,
						},
						"reviewer_stages": schema.ListNestedAttribute{
							Optional: true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"operator": schema.StringAttribute{
										Computed:    true,
										Optional:    true,
										Default:     stringdefault.StaticString(`AND`),
										Description: `The operator of the reviewer stage. Admin and manager approval are also treated as reviewers. Default: "AND"; must be one of ["AND", "OR"]`,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"AND",
												"OR",
											),
										},
									},
									"owner_ids": schema.SetAttribute{
										Required:    true,
										ElementType: types.StringType,
									},
									"require_admin_approval": schema.BoolAttribute{
										Optional:    true,
										Description: `Whether this reviewer stage should require admin approval.`,
									},
									"require_manager_approval": schema.BoolAttribute{
										Required:    true,
										Description: `Whether this reviewer stage should require manager approval.`,
									},
								},
							},
							Description: `The list of reviewer stages for the request configuration.`,
						},
					},
				},
				Description: `The request configuration list of the configuration template. If not provided, the default request configuration will be used.`,
			},
			"require_mfa_to_approve": schema.BoolAttribute{
				Computed:    true,
				Optional:    true,
				Default:     booldefault.StaticBool(false),
				Description: `A bool representing whether or not to require MFA for reviewers to approve requests for this configuration template. Default: false`,
			},
			"require_mfa_to_connect": schema.BoolAttribute{
				Required:    true,
				Description: `A bool representing whether or not to require MFA to connect to resources associated with this configuration template.`,
			},
			"ticket_propagation": schema.SingleNestedAttribute{
				Computed: true,
				Optional: true,
				PlanModifiers: []planmodifier.Object{
					speakeasy_objectplanmodifier.SuppressDiff(speakeasy_objectplanmodifier.ExplicitSuppress),
				},
				Attributes: map[string]schema.Attribute{
					"enabled_on_grant": schema.BoolAttribute{
						Computed: true,
						Optional: true,
						PlanModifiers: []planmodifier.Bool{
							speakeasy_boolplanmodifier.SuppressDiff(speakeasy_boolplanmodifier.ExplicitSuppress),
						},
						Description: `Not Null`,
						Validators: []validator.Bool{
							speakeasy_boolvalidators.NotNull(),
						},
					},
					"enabled_on_revocation": schema.BoolAttribute{
						Computed: true,
						Optional: true,
						PlanModifiers: []planmodifier.Bool{
							speakeasy_boolplanmodifier.SuppressDiff(speakeasy_boolplanmodifier.ExplicitSuppress),
						},
						Description: `Not Null`,
						Validators: []validator.Bool{
							speakeasy_boolvalidators.NotNull(),
						},
					},
					"ticket_project_id": schema.StringAttribute{
						Computed: true,
						Optional: true,
						PlanModifiers: []planmodifier.String{
							speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
						},
					},
					"ticket_provider": schema.StringAttribute{
						Computed: true,
						Optional: true,
						PlanModifiers: []planmodifier.String{
							speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
						},
						Description: `The third party ticketing platform provider. must be one of ["JIRA", "LINEAR", "SERVICE_NOW"]`,
						Validators: []validator.String{
							stringvalidator.OneOf(
								"JIRA",
								"LINEAR",
								"SERVICE_NOW",
							),
						},
					},
				},
				Description: `Configuration for ticket propagation, when enabled, a ticket will be created for access changes related to the users in this resource.`,
			},
			"visibility": schema.SingleNestedAttribute{
				Required: true,
				Attributes: map[string]schema.Attribute{
					"visibility": schema.StringAttribute{
						Required:    true,
						Description: `The visibility level of the entity. must be one of ["GLOBAL", "LIMITED"]`,
						Validators: []validator.String{
							stringvalidator.OneOf(
								"GLOBAL",
								"LIMITED",
							),
						},
					},
					"visibility_group_ids": schema.SetAttribute{
						Computed: true,
						Optional: true,
						Default:  setdefault.StaticValue(types.SetValueMust(types.StringType, []attr.Value{})),
						PlanModifiers: []planmodifier.Set{
							speakeasy_setplanmodifier.SuppressDiff(speakeasy_setplanmodifier.ExplicitSuppress),
						},
						ElementType: types.StringType,
					},
				},
				Description: `Visibility infomation of an entity.`,
				Validators: []validator.Object{
					custom_objectvalidators.VisibilityInfo(),
				},
			},
		},
	}
}

func (r *ConfigurationTemplateResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*sdk.OpalAPI)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *sdk.OpalAPI, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *ConfigurationTemplateResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *ConfigurationTemplateResourceModel
	var plan types.Object

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(plan.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}

	request := *data.ToSharedCreateConfigurationTemplateInfo()
	res, err := r.client.ConfigurationTemplates.Create(ctx, request)
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
	if !(res.ConfigurationTemplate != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedConfigurationTemplate(res.ConfigurationTemplate)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ConfigurationTemplateResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *ConfigurationTemplateResourceModel
	var item types.Object

	resp.Diagnostics.Append(req.State.Get(ctx, &item)...)
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

	// Not Implemented; we rely entirely on CREATE API request response

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ConfigurationTemplateResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *ConfigurationTemplateResourceModel
	var plan types.Object

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	merge(ctx, req, resp, &data)
	if resp.Diagnostics.HasError() {
		return
	}

	request := *data.ToSharedUpdateConfigurationTemplateInfo()
	res, err := r.client.ConfigurationTemplates.Update(ctx, request)
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
	if !(res.ConfigurationTemplate != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedConfigurationTemplate(res.ConfigurationTemplate)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ConfigurationTemplateResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *ConfigurationTemplateResourceModel
	var item types.Object

	resp.Diagnostics.Append(req.State.Get(ctx, &item)...)
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

	var configurationTemplateID string
	configurationTemplateID = data.ConfigurationTemplateID.ValueString()

	request := operations.DeleteConfigurationTemplateRequest{
		ConfigurationTemplateID: configurationTemplateID,
	}
	res, err := r.client.ConfigurationTemplates.Delete(ctx, request)
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

}

func (r *ConfigurationTemplateResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.AddError("Not Implemented", "No available import state operation is available for resource configuration_template.")
}
