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
var _ datasource.DataSource = &GroupReviewersStagesListDataSource{}
var _ datasource.DataSourceWithConfigure = &GroupReviewersStagesListDataSource{}

func NewGroupReviewersStagesListDataSource() datasource.DataSource {
	return &GroupReviewersStagesListDataSource{}
}

// GroupReviewersStagesListDataSource is the data source implementation.
type GroupReviewersStagesListDataSource struct {
	client *sdk.OpalAPI
}

// GroupReviewersStagesListDataSourceModel describes the data model.
type GroupReviewersStagesListDataSourceModel struct {
	Data    []tfTypes.ReviewerStage `tfsdk:"data"`
	GroupID types.String            `tfsdk:"group_id"`
}

// Metadata returns the data source type name.
func (r *GroupReviewersStagesListDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_group_reviewers_stages_list"
}

// Schema defines the schema for the data source.
func (r *GroupReviewersStagesListDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "GroupReviewersStagesList DataSource",

		Attributes: map[string]schema.Attribute{
			"data": schema.ListNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"operator": schema.StringAttribute{
							Computed:    true,
							Description: `The operator of the reviewer stage. must be one of ["AND", "OR"]`,
						},
						"owner_ids": schema.ListAttribute{
							Computed:    true,
							ElementType: types.StringType,
						},
						"require_manager_approval": schema.BoolAttribute{
							Computed:    true,
							Description: `Whether this reviewer stage should require manager approval.`,
						},
					},
				},
				Description: `The reviewer stages for this group.`,
			},
			"group_id": schema.StringAttribute{
				Required:    true,
				Description: `The ID of the group.`,
			},
		},
	}
}

func (r *GroupReviewersStagesListDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *GroupReviewersStagesListDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *GroupReviewersStagesListDataSourceModel
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

	groupID := data.GroupID.ValueString()
	request := operations.GetGroupReviewersStagesRequest{
		GroupID: groupID,
	}
	res, err := r.client.Groups.GetReviewersStages(ctx, request)
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
	if res.Classes == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedReviewerStage(res.Classes)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
