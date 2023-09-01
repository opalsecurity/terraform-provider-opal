package opal

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/hashicorp/go-multierror"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/opalsecurity/opal-go"
)

// enumSliceToSTringSlice converts the values from an SDK-provided enum slice
// to type []string.
func enumSliceToStringSlice[T ~string](input []T) []string {
	rv := make([]string, 0, len(input))
	for _, v := range input {
		rv = append(rv, string(v))
	}
	return rv
}

var allowedResourceTypes = enumSliceToStringSlice(opal.AllowedResourceTypeEnumEnumValues)
var allowedVisibilityTypes = enumSliceToStringSlice(opal.AllowedVisibilityTypeEnumEnumValues)

func dataSourceResource() *schema.Resource {
	return &schema.Resource{
		Description: "An Opal resource data source.",
		ReadContext: dataSourceResourceRead,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"id": {
				Description: "The ID of the resource.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"name": {
				Description: "The name of the resource.",
				Type:        schema.TypeString,
				Computed:    true,
			},
		},
	}
}

func dataSourceResourceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*opal.APIClient)

	id, idOk := d.GetOk("id")
	var resource *opal.Resource
	var err error
	if idOk {
		resource, _, err = client.ResourcesApi.GetResource(ctx, id.(string)).Execute()
		if err != nil {
			return diagFromErr(ctx, err)
		}
	} else {
		return diagFromErr(ctx, errors.New("must provide id for resource data source"))
	}

	d.SetId(resource.ResourceId)
	if err := multierror.Append(
		d.Set("name", resource.Name),
	); err.ErrorOrNil() != nil {
		return diagFromErr(ctx, err)
	}

	return nil
}

func resourceResource() *schema.Resource {
	return &schema.Resource{
		Description:   "An Opal Resource resource.",
		CreateContext: resourceResourceCreate,
		ReadContext:   resourceResourceRead,
		UpdateContext: resourceResourceUpdate,
		DeleteContext: resourceResourceDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"id": {
				Description: "The ID of the resource.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"name": {
				Description: "The name of the resource.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"description": {
				Description: "The description of the resource.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"resource_type": {
				Description:  "The type of the resource, i.e. AWS_EC2_INSTANCE.",
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice(allowedResourceTypes, false),
				ForceNew:     true,
				Required:     true,
			},
			"app_id": {
				Description: "The ID of the app integration that provides the resource. You can get this value from the URL of the app in the Opal web app.",
				Type:        schema.TypeString,
				ForceNew:    true,
				Required:    true,
			},
			"admin_owner_id": {
				Description: "The admin owner ID for this resource.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"auto_approval": {
				Description: "Automatically approve all requests for this resource without review.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"require_mfa_to_approve": {
				Description: "Require that reviewers MFA to approve requests for this resource.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"require_mfa_to_connect": {
				Description: "Require that users MFA to connect to this resource. Only applicable for resources where a session can be started from Opal (i.e. AWS RDS database)",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"require_mfa_to_request": {
				Description: "Require that users MFA to request this resource.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"require_support_ticket": {
				Description: "Require that requesters attach a support ticket to requests for this resource.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"max_duration": {
				Description: "The maximum duration for which this resource can be requested (in minutes).",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"recommended_duration": {
				Description: "The recommended duration for which the resource should be requested (in minutes). Will be the default value in a request. Use -1 to set to indefinite.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"request_template_id": {
				Description: "The ID of a request template for this resource. You can get this ID from the URL in the Opal web app.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"remote_info": {
				Description: "Remote info that is required for the creation of remote resources.",
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
				MaxItems:    1,
				Elem:        resourceRemoteInfoElem(),
			},
			"visibility": {
				Description:  "The visibility level of the resource, i.e. LIMITED or GLOBAL.",
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice(allowedVisibilityTypes, false),
				Optional:     true,
				Default:      "GLOBAL",
			},
			"visibility_group": {
				Description: "The groups that can see this resource when visibility is limited. If not specified, only admins and users with direct access can see this resource when visibility is set to LIMITED.",
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Description: "The ID of the group that can see this resource.",
							Type:        schema.TypeString,
							Required:    true,
						},
					},
				},
			},
			"reviewer_stage": {
				Description: "A reviewer stage for this resource. You are allowed to provide up to 3.",
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"operator": {
							Description:  "The operator of the stage. Operator is either \"AND\" or \"OR\".",
							Type:         schema.TypeString,
							Optional:     true,
							Default:      "AND",
							ValidateFunc: validation.StringInSlice(allowedReviewerStageOperators, false),
						},
						"require_manager_approval": {
							Description: "Whether this reviewer stage should require manager approval.",
							Type:        schema.TypeBool,
							Optional:    true,
							Default:     false,
						},
						"reviewer": {
							Description: "A reviewer for this stage.",
							Type:        schema.TypeSet,
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Description: "The ID of the owner.",
										Type:        schema.TypeString,
										Required:    true,
									},
								},
							},
						},
					},
				},
			},
			"is_requestable": {
				Description: "Allow users to create an access request for this resource. By default, any resource is requestable.",
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
			},
			// XXX: Audit message channel...
		},
	}
}

func resourceResourceCreate(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	client := m.(*opal.APIClient)

	if err := validateReviewerConfigDuringCreate(d); err != nil {
		return diagFromErr(ctx, err)
	}

	name := d.Get("name").(string)
	resourceType := opal.ResourceTypeEnum(d.Get("resource_type").(string))
	appID := d.Get("app_id").(string)

	createInfo := opal.NewCreateResourceInfo(name, resourceType, appID)
	if descI, ok := d.GetOk("description"); ok {
		createInfo.SetDescription(descI.(string))
	}

	if remoteInfoI, ok := d.GetOk("remote_info"); ok {
		remoteInfo, err := resourceRemoteInfoTerraformToAPI(remoteInfoI)
		if err != nil {
			return diagFromErr(ctx, err)
		}
		createInfo.SetRemoteInfo(*remoteInfo)
	}

	resource, _, err := client.ResourcesApi.CreateResource(ctx).CreateResourceInfo(*createInfo).Execute()
	if err != nil {
		return diagFromErr(ctx, err)
	}
	d.SetId(resource.ResourceId)

	tflog.Debug(ctx, "Created opal resource", map[string]any{
		"name": name,
		"id":   d.Id(),
	})

	// In the case that auto_approval is true or is_requestable is false, we still want to
	// update the reviewer stages to be empty to avoid the immediate diff from the default
	// reviewer configuration.
	// NOTE: This call should come before updating is_requestable and auto_approval as it otherwise
	// overrides those values
	var reviewerStages any = make([]any, 0)
	if reviewerStagesI, ok := d.GetOk("reviewer_stage"); ok {
		reviewerStages = reviewerStagesI
	}
	if diag := resourceResourceUpdateReviewerStages(ctx, d, client, reviewerStages); diag != nil {
		return diag
	}

	// Because resource creation does not let us set some properties immediately,
	// we may have to update them in a follow up request.
	adminOwnerIDI, adminOwnerIDOk := d.GetOk("admin_owner_id")
	autoApprovalI, autoApprovalOk := d.GetOkExists("auto_approval")
	requireMfaToApproveI, requireMfaToApproveOk := d.GetOkExists("require_mfa_to_approve")
	requireMfaToConnectI, requireMfaToConnectOk := d.GetOkExists("require_mfa_to_connect")
	requireMfaToRequestI, requireMfaToRequestOk := d.GetOkExists("require_mfa_to_request")
	requireSupportTicketI, requireSupportTicketOk := d.GetOkExists("require_support_ticket")
	isRequestableI, isRequestableOk := d.GetOkExists("is_requestable")
	maxDurationI, maxDurationOk := d.GetOk("max_duration")
	recommendedDurationI, recommendedDurationOk := d.GetOk("recommended_duration")
	requestTemplateIDI, requestTemplateIDOk := d.GetOk("request_template_id")
	if adminOwnerIDOk || autoApprovalOk || requireMfaToApproveOk || requireMfaToConnectOk || requireMfaToRequestOk || requireSupportTicketOk || isRequestableOk || maxDurationOk || recommendedDurationOk || requestTemplateIDOk {
		updateInfo := opal.NewUpdateResourceInfo(resource.ResourceId)
		if adminOwnerIDOk {
			updateInfo.SetAdminOwnerId(adminOwnerIDI.(string))
		}
		if autoApprovalOk {
			updateInfo.SetAutoApproval(autoApprovalI.(bool))
		}
		if requireMfaToApproveOk {
			updateInfo.SetRequireMfaToApprove(requireMfaToApproveI.(bool))
		}
		if requireMfaToConnectOk {
			updateInfo.SetRequireMfaToConnect(requireMfaToConnectI.(bool))
		}
		if requireMfaToRequestOk {
			updateInfo.SetRequireMfaToRequest(requireMfaToRequestI.(bool))
		}
		if requireSupportTicketOk {
			updateInfo.SetRequireSupportTicket(requireSupportTicketI.(bool))
		}
		if maxDurationOk {
			updateInfo.SetMaxDuration(int32(maxDurationI.(int)))
		}
		if recommendedDurationOk {
			updateInfo.SetRecommendedDuration(int32(recommendedDurationI.(int)))
		}
		if requestTemplateIDOk {
			updateInfo.SetRequestTemplateId(requestTemplateIDI.(string))
		}
		if isRequestableOk {
			updateInfo.SetIsRequestable(isRequestableI.(bool))
		}

		tflog.Debug(ctx, "Immediately updating opal resource", map[string]any{
			"name":       name,
			"updateInfo": updateInfo,
		})

		if _, _, err := client.ResourcesApi.UpdateResources(ctx).UpdateResourceInfoList(*opal.NewUpdateResourceInfoList([]opal.UpdateResourceInfo{*updateInfo})).Execute(); err != nil {
			return diagFromErr(ctx, err)
		}
	}

	if _, ok := d.GetOk("visibility"); ok {
		if diag := resourceResourceUpdateVisibility(ctx, d, client); diag != nil {
			return diag
		}
	}

	return resourceResourceRead(ctx, d, m)
}

func resourceResourceUpdateVisibility(ctx context.Context, d *schema.ResourceData, client *opal.APIClient) diag.Diagnostics {
	visibilityInfo := *opal.NewVisibilityInfo(opal.VisibilityTypeEnum(opal.VISIBILITYTYPEENUM_GLOBAL))
	if visibilityI, ok := d.GetOk("visibility"); ok {
		visibilityInfo.SetVisibility(opal.VisibilityTypeEnum(visibilityI.(string)))
	}

	if visibilityGroupI, ok := d.GetOk("visibility_group"); ok {
		rawGroups := visibilityGroupI.([]any)
		groupIds := make([]string, 0, len(rawGroups))
		for _, rawGroup := range rawGroups {
			group := rawGroup.(map[string]any)
			groupIds = append(groupIds, group["id"].(string))
		}
		visibilityInfo.SetVisibilityGroupIds(groupIds)
	}

	if _, _, err := client.ResourcesApi.SetResourceVisibility(ctx, d.Id()).VisibilityInfo(visibilityInfo).Execute(); err != nil {
		return diagFromErr(ctx, err)
	}
	return nil
}

func resourceResourceUpdateReviewerStages(ctx context.Context, d *schema.ResourceData, client *opal.APIClient, reviewerStagesI any) diag.Diagnostics {
	rawReviewerStages := reviewerStagesI.([]any)
	reviewerStages := make([]opal.ReviewerStage, 0, len(rawReviewerStages))
	for _, rawReviewerStage := range rawReviewerStages {
		reviewerStage := rawReviewerStage.(map[string]any)
		requireManagerApproval := reviewerStage["require_manager_approval"].(bool)
		operator := reviewerStage["operator"].(string)
		reviewersI := reviewerStage["reviewer"]
		reviewerIds, err := extractReviewerIDs(reviewersI)
		if err != nil {
			return diagFromErr(ctx, err)
		}

		reviewerStages = append(reviewerStages, *opal.NewReviewerStage(requireManagerApproval, operator, reviewerIds))
		tflog.Debug(ctx, "Setting resource reviewer stage", map[string]any{
			"id":                     d.Id(),
			"requireManagerApproval": requireManagerApproval,
			"operator":               operator,
			"reviewerIds":            reviewerIds,
		})
	}

	if _, _, err := client.ResourcesApi.SetResourceReviewerStages(ctx, d.Id()).ReviewerStageList(*opal.NewReviewerStageList(reviewerStages)).Execute(); err != nil {
		return diagFromErr(ctx, err)
	}
	return nil
}

func resourceResourceRead(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	client := m.(*opal.APIClient)

	resource, _, err := client.ResourcesApi.GetResource(ctx, d.Id()).Execute()
	if err != nil {
		return diagFromErr(ctx, err)
	}

	d.SetId(resource.ResourceId)
	if err := multierror.Append(
		d.Set("name", resource.Name),
		d.Set("description", resource.Description),
		d.Set("resource_type", resource.ResourceType),
		d.Set("app_id", resource.AppId),
		d.Set("admin_owner_id", resource.AdminOwnerId),
		d.Set("auto_approval", resource.AutoApproval),
		d.Set("require_mfa_to_approve", resource.RequireMfaToApprove),
		d.Set("require_mfa_to_connect", resource.RequireMfaToConnect),
		d.Set("require_mfa_to_request", resource.RequireMfaToRequest),
		d.Set("require_support_ticket", resource.RequireSupportTicket),
		d.Set("max_duration", resource.MaxDuration),
		d.Set("recommended_duration", resource.RecommendedDuration),
		d.Set("request_template_id", resource.RequestTemplateId),
		d.Set("is_requestable", resource.IsRequestable),
	); err.ErrorOrNil() != nil {
		return diagFromErr(ctx, err)
	}

	remoteInfoI, err := resourceRemoteInfoAPIToTerraform(resource.RemoteInfo)
	if err != nil {
		return diagFromErr(ctx, err)
	}
	if remoteInfoI != nil {
		d.Set("remote_info", remoteInfoI)
	}

	visibility, _, err := client.ResourcesApi.GetResourceVisibility(ctx, resource.ResourceId).Execute()
	if err != nil {
		return diagFromErr(ctx, err)
	}
	d.Set("visibility", visibility.Visibility)

	flattenedGroups := make([]any, 0, len(visibility.VisibilityGroupIds))
	for _, groupID := range visibility.VisibilityGroupIds {
		flattenedGroups = append(flattenedGroups, map[string]any{"id": groupID})
	}
	d.Set("visibility_group", flattenedGroups)

	reviewerStages, _, err := client.ResourcesApi.GetResourceReviewerStages(ctx, resource.ResourceId).Execute()
	if err != nil {
		return diagFromErr(ctx, err)
	}

	reviewerStagesI := make([]any, 0, len(reviewerStages))
	for _, reviewerStage := range reviewerStages {
		reviewersI := make([]any, 0, len(reviewerStage.OwnerIds))
		for _, reviewerID := range reviewerStage.OwnerIds {
			reviewersI = append(reviewersI, map[string]any{
				"id": reviewerID,
			})
		}

		reviewerStagesI = append(reviewerStagesI, map[string]any{
			"reviewer":                 reviewersI,
			"operator":                 reviewerStage.Operator,
			"require_manager_approval": reviewerStage.RequireManagerApproval,
		})
	}
	d.Set("reviewer_stage", reviewerStagesI)

	if resource.Metadata != nil {
		remoteInfoIList := make([]any, 0, 1)
		switch *resource.ResourceType {
		case opal.RESOURCETYPEENUM_AWS_SSO_PERMISSION_SET:
			// TODO: Handle other AWS Orgs resource types
			var metadata opal.AwsPermissionSetMetadata
			if err := json.Unmarshal([]byte(*resource.Metadata), &metadata); err != nil {
				return diagFromErr(ctx, err)
			}
			permissionSetIList := make([]any, 0, 1)
			permissionSetIList = append(permissionSetIList, map[string]any{
				"arn":        metadata.AwsPermissionSet.Arn,
				"account_id": metadata.AwsPermissionSet.AccountId,
			})
			remoteInfoIList = append(remoteInfoIList, map[string]any{
				"aws_permission_set": permissionSetIList,
			})
		}

		if len(remoteInfoIList) == 1 {
			d.Set("remote_info", remoteInfoIList)
		}
	}

	return nil
}

func resourceResourceUpdate(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	client := m.(*opal.APIClient)

	// Note that fields like metadata, app_id, resource_type, and remote_info
	// force a recreation, so we do not need to worry about those values here.
	hasBasicChange := false
	updateInfo := opal.NewUpdateResourceInfo(d.Id())
	if d.HasChange("name") {
		hasBasicChange = true
		updateInfo.SetName(d.Get("name").(string))
	}
	if d.HasChange("description") {
		hasBasicChange = true
		updateInfo.SetDescription(d.Get("description").(string))
	}
	if d.HasChange("admin_owner_id") {
		hasBasicChange = true
		updateInfo.SetAdminOwnerId(d.Get("admin_owner_id").(string))
	}
	if d.HasChange("auto_approval") {
		hasBasicChange = true
		updateInfo.SetAutoApproval(d.Get("auto_approval").(bool))
	}
	if d.HasChange("require_mfa_to_approve") {
		hasBasicChange = true
		updateInfo.SetRequireMfaToApprove(d.Get("require_mfa_to_approve").(bool))
	}
	if d.HasChange("require_mfa_to_connect") {
		hasBasicChange = true
		updateInfo.SetRequireMfaToConnect(d.Get("require_mfa_to_connect").(bool))
	}
	if d.HasChange("require_mfa_to_request") {
		hasBasicChange = true
		updateInfo.SetRequireMfaToRequest(d.Get("require_mfa_to_request").(bool))
	}
	if d.HasChange("require_support_ticket") {
		hasBasicChange = true
		updateInfo.SetRequireSupportTicket(d.Get("require_support_ticket").(bool))
	}
	if d.HasChange("max_duration") {
		hasBasicChange = true
		updateInfo.SetMaxDuration(int32(d.Get("max_duration").(int)))
	}
	if d.HasChange("recommended_duration") {
		hasBasicChange = true
		updateInfo.SetRecommendedDuration(int32(d.Get("recommended_duration").(int)))
	}
	if d.HasChange("request_template_id") {
		hasBasicChange = true
		updateInfo.SetRequestTemplateId(d.Get("request_template_id").(string))
	}
	if d.HasChange("is_requestable") {
		hasBasicChange = true
		updateInfo.SetIsRequestable(d.Get("is_requestable").(bool))
	}

	if hasBasicChange {
		_, _, err := client.ResourcesApi.UpdateResources(ctx).UpdateResourceInfoList(*opal.NewUpdateResourceInfoList([]opal.UpdateResourceInfo{*updateInfo})).Execute()
		if err != nil {
			return diagFromErr(ctx, err)
		}
	}

	if d.HasChange("visibility") || d.HasChange("visibility_group") {
		if diag := resourceResourceUpdateVisibility(ctx, d, client); diag != nil {
			return diag
		}
	}

	if d.HasChange("reviewer_stage") {
		reviewerStages := any([]any{})
		if reviewersStagesBlock, ok := d.GetOk("reviewer_stage"); ok {
			reviewerStages = reviewersStagesBlock
		}
		if diag := resourceResourceUpdateReviewerStages(ctx, d, client, reviewerStages); diag != nil {
			return diag
		}
	}

	return resourceResourceRead(ctx, d, m)
}

func resourceResourceDelete(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	client := m.(*opal.APIClient)
	tflog.Debug(ctx, "Deleting resource", map[string]any{
		"id": d.Id(),
	})

	if _, err := client.ResourcesApi.DeleteResource(ctx, d.Id()).Execute(); err != nil {
		return diagFromErr(ctx, err)
	}

	return nil
}
