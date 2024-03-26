// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	speakeasy_stringplanmodifier "github.com/opal-dev/terraform-provider-opal/internal/planmodifiers/stringplanmodifier"
	"github.com/opal-dev/terraform-provider-opal/internal/sdk"
	"github.com/opal-dev/terraform-provider-opal/internal/sdk/models/operations"
	"github.com/opal-dev/terraform-provider-opal/internal/sdk/models/shared"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &OwnerResource{}
var _ resource.ResourceWithImportState = &OwnerResource{}

func NewOwnerResource() resource.Resource {
	return &OwnerResource{}
}

// OwnerResource defines the resource implementation.
type OwnerResource struct {
	client *sdk.OpalAPI
}

// OwnerResourceModel describes the resource data model.
type OwnerResourceModel struct {
	AccessRequestEscalationPeriod types.Int64    `tfsdk:"access_request_escalation_period"`
	Description                   types.String   `tfsdk:"description"`
	ID                            types.String   `tfsdk:"id"`
	Name                          types.String   `tfsdk:"name"`
	ReviewerMessageChannelID      types.String   `tfsdk:"reviewer_message_channel_id"`
	SourceGroupID                 types.String   `tfsdk:"source_group_id"`
	UserIds                       []types.String `tfsdk:"user_ids"`
}

func (r *OwnerResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_owner"
}

func (r *OwnerResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Owner Resource",

		Attributes: map[string]schema.Attribute{
			"access_request_escalation_period": schema.Int64Attribute{
				Computed:    true,
				Optional:    true,
				Description: `The amount of time (in minutes) before the next reviewer is notified. Use 0 to remove escalation policy.`,
			},
			"description": schema.StringAttribute{
				Computed:    true,
				Optional:    true,
				Description: `A description of the owner.`,
			},
			"id": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
				},
				Description: `The ID of the owner.`,
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: `The name of the owner.`,
			},
			"reviewer_message_channel_id": schema.StringAttribute{
				Computed:    true,
				Optional:    true,
				Description: `The message channel id for the reviewer channel.`,
			},
			"source_group_id": schema.StringAttribute{
				Computed:    true,
				Optional:    true,
				Description: `Sync this owner's user list with a source group.`,
			},
			"user_ids": schema.ListAttribute{
				PlanModifiers: []planmodifier.List{
					listplanmodifier.RequiresReplaceIfConfigured(),
				},
				Required:    true,
				ElementType: types.StringType,
				Description: `Users to add to the created owner. If setting a source_group_id this list must be empty. Requires replacement if changed. `,
			},
		},
	}
}

func (r *OwnerResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *OwnerResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *OwnerResourceModel
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

	request := *data.ToSharedCreateOwnerInfo()
	res, err := r.client.Owners.Create(ctx, request)
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
	if res.Owner == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedOwner(res.Owner)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)
	id := data.ID.ValueString()
	request1 := operations.GetOwnerIDRequest{
		ID: id,
	}
	res1, err := r.client.Owners.GetID(ctx, request1)
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
	if res1.Owner == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res1.RawResponse))
		return
	}
	data.RefreshFromSharedOwner(res1.Owner)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *OwnerResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *OwnerResourceModel
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

	id := data.ID.ValueString()
	request := operations.GetOwnerIDRequest{
		ID: id,
	}
	res, err := r.client.Owners.GetID(ctx, request)
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
	if res.Owner == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedOwner(res.Owner)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *OwnerResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *OwnerResourceModel
	var plan types.Object

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	merge(ctx, req, resp, &data)
	if resp.Diagnostics.HasError() {
		return
	}

	singleton := *data.ToSharedUpdateOwnerInfo()
	owners := []shared.UpdateOwnerInfo{singleton}
	request := shared.UpdateOwnerInfoList{
		Owners: owners,
	}
	res, err := r.client.Owners.Update(ctx, request)
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
	if res.UpdateOwnerInfoList == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedUpdateOwnerInfo(res.UpdateOwnerInfoList.Owners[0])
	refreshPlan(ctx, plan, &data, resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *OwnerResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *OwnerResourceModel
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

	id := data.ID.ValueString()
	request := operations.DeleteOwnerRequest{
		ID: id,
	}
	res, err := r.client.Owners.Delete(ctx, request)
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

func (r *OwnerResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), req.ID)...)
}