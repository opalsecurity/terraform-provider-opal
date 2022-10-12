package opal

import (
	"context"
	"errors"
	"log"

	"github.com/hashicorp/go-multierror"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
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
		CustomizeDiff: customdiff.All(
			func(ctx context.Context, diff *schema.ResourceDiff, v interface{}) error {
				if diff.Get("visibility").(string) == "GLOBAL" && len(diff.Get("visibility_group").([]any)) > 0 {
					return errors.New("`visibility_group` cannot be specified when `visibility` is set to GLOBAL")
				}
				return nil
			},
			// XXX: We could enforce that remote_resource_id/metadata must be passed for resource types that need it.
		),
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
				Description: "The admin owner ID for this resource. By default, this is set to the application admin owner.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"require_manager_approval": {
				Description: "Require the requester's manager's approval for requests to this resource.",
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
			},
			"auto_approval": {
				Description: "Automatically approve all requests for this resource without review.",
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
			},
			"require_mfa_to_approve": {
				Description: "Require that reviewers MFA to approve requests for this resource.",
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
			},
			"require_mfa_to_connect": {
				Description: "Require that users MFA to connect to this resource. Only applicable for resources where a session can be started from Opal (i.e. AWS RDS database)",
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
			},
			"require_support_ticket": {
				Description: "Require that requesters attach a support ticket to requests for this resource.",
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
			},
			"max_duration": {
				Description: "The maximum duration for which this resource can be requested (in minutes). By default, the max duration is indefinite access.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"request_template_id": {
				Description: "The ID of a request template for this resource. You can get this ID from the URL in the Opal web app.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
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
				Computed:     true,
			},
			"visibility_group": {
				Description: "The groups that can see this resource when visibility is limited. If not specified, only admins and users with direct access can see this resource when visibility is set to LIMITED.",
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
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
			"reviewer": {
				Description: "A required reviewer for this resource. If none are specified, then the admin owner will be used.",
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Description: "The ID of the owner that must review requests to this resource.",
							Type:        schema.TypeString,
							Required:    true,
						},
					},
				},
			},
			// XXX: Audit message channel...
		},
	}
}

func resourceResourceCreate(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	client := m.(*opal.APIClient)

	name := d.Get("name").(string)
	resourceType := opal.ResourceTypeEnum(d.Get("resource_type").(string))
	appID := d.Get("app_id").(string)

	createInfo := opal.NewCreateResourceInfo(name, resourceType, appID)
	if descI, ok := d.GetOk("description"); ok {
		createInfo.SetDescription(descI.(string))
	}

	if remoteInfoI, ok := d.GetOk("remote_info"); ok {
		remoteInfo, err := parseResourceRemoteInfo(remoteInfoI)
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

	// Because resource creation does not let us set some properties immediately,
	// we may have to update them in a follow up request.
	adminOwnerIDI, adminOwnerIDOk := d.GetOk("admin_owner_id")
	requireManagerApprovalI, requireManagerApprovalOk := d.GetOk("require_manager_approval")
	autoApprovalI, autoApprovalOk := d.GetOk("auto_approval")
	requireMfaToApproveI, requireMfaToApproveOk := d.GetOk("require_mfa_to_approve")
	requireMfaToConnectI, requireMfaToConnectOk := d.GetOk("require_mfa_to_connect")
	requireSupportTicketI, requireSupportTicketOk := d.GetOk("require_support_ticket")
	maxDurationI, maxDurationOk := d.GetOk("max_duration")
	requestTemplateIDI, requestTemplateIDOk := d.GetOk("request_template_id")
	if adminOwnerIDOk || requireManagerApprovalOk || autoApprovalOk || requireMfaToApproveOk || requireMfaToConnectOk || requireSupportTicketOk || maxDurationOk || requestTemplateIDOk {
		updateInfo := opal.NewUpdateResourceInfo(resource.ResourceId)
		if adminOwnerIDOk {
			updateInfo.SetAdminOwnerId(adminOwnerIDI.(string))
		}
		if requireManagerApprovalOk {
			updateInfo.SetRequireManagerApproval(requireManagerApprovalI.(bool))
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
		if requireSupportTicketOk {
			updateInfo.SetRequireSupportTicket(requireSupportTicketI.(bool))
		}
		if maxDurationOk {
			updateInfo.SetMaxDuration(int32(maxDurationI.(int)))
		}
		if requestTemplateIDOk {
			updateInfo.SetRequestTemplateId(requestTemplateIDI.(string))
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

	if reviewersI, ok := d.GetOk("reviewer"); ok {
		if diag := resourceResourceUpdateReviewers(ctx, d, client, reviewersI); diag != nil {
			return diag
		}
	} else if adminOwnerIDOk {
		// If the admin owner was set during creation, we should also set
		// the required reviewer to be the same so that it is consistent.
		//
		// Otherwise, if it's unset, the Opal API will automatically set it to
		// the app owner.
		if diag := resourceResourceUpdateReviewers(ctx, d, client, []any{map[string]any{"id": adminOwnerIDI}}); diag != nil {
			return diag
		}
	}

	// XXX: Update audit channel...
	// XXX: Update mfa required for connnect...

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

func resourceResourceUpdateReviewers(ctx context.Context, d *schema.ResourceData, client *opal.APIClient, reviewersI any) diag.Diagnostics {
	rawReviewers := reviewersI.([]any)
	reviewerIds := make([]string, 0, len(rawReviewers))
	for _, rawReviewer := range rawReviewers {
		reviewer := rawReviewer.(map[string]any)
		reviewerIds = append(reviewerIds, reviewer["id"].(string))
	}
	tflog.Debug(ctx, "Setting resource reviewers", map[string]any{
		"id":          d.Id(),
		"reviewerIds": reviewerIds,
	})

	if _, _, err := client.ResourcesApi.SetResourceReviewers(ctx, d.Id()).ReviewerIDList(*opal.NewReviewerIDList(reviewerIds)).Execute(); err != nil {
		var gErr *opal.GenericOpenAPIError
		if errors.As(err, &gErr) {
			log.Println("error string", string(gErr.Body()))
		} else {
			log.Println("not", err)
		}
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
		d.Set("require_manager_approval", resource.RequireManagerApproval),
		d.Set("auto_approval", resource.AutoApproval),
		d.Set("require_mfa_to_approve", resource.RequireMfaToApprove),
		d.Set("require_support_ticket", resource.RequireSupportTicket),
		d.Set("max_duration", resource.MaxDuration),
		d.Set("request_template_id", resource.RequestTemplateId),
		// XXX: We don't get the metadata back. Will terraform state be okay?
	); err.ErrorOrNil() != nil {
		return diagFromErr(ctx, err)
	}

	visibility, _, err := client.ResourcesApi.GetResourceVisibility(ctx, resource.ResourceId).Execute()
	if err != nil {
		return diagFromErr(ctx, err)
	}

	visibilityGroups := make([]any, 0, len(visibility.VisibilityGroupIds))
	for _, groupID := range visibility.VisibilityGroupIds {
		visibilityGroups = append(visibilityGroups, map[string]any{
			"id": groupID,
		})
	}
	d.Set("visibility", visibility.Visibility)
	flattenedGroups := make([]any, 0, len(visibility.VisibilityGroupIds))
	for _, groupID := range visibility.VisibilityGroupIds {
		flattenedGroups = append(flattenedGroups, map[string]any{"id": groupID})
	}
	d.Set("visibility_group", flattenedGroups)

	reviewerIDs, _, err := client.ResourcesApi.GetResourceReviewers(ctx, resource.ResourceId).Execute()
	if err != nil {
		return diagFromErr(ctx, err)
	}

	reviewers := make([]any, 0, len(reviewerIDs))
	for _, reviewerID := range reviewerIDs {
		reviewers = append(reviewers, map[string]any{
			"id": reviewerID,
		})
	}
	d.Set("reviewer", reviewers)

	// XXX: Read out message channels, mfa required to connect.

	return nil
}

func resourceResourceUpdate(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	client := m.(*opal.APIClient)

	// Note that metadata, app_id, and resource_type force a recreation, so we do not need to
	// worry about those values here.
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
	if d.HasChange("require_manager_approval") {
		hasBasicChange = true
		updateInfo.SetRequireManagerApproval(d.Get("require_manager_approval").(bool))
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
	if d.HasChange("require_support_ticket") {
		hasBasicChange = true
		updateInfo.SetRequireSupportTicket(d.Get("require_support_ticket").(bool))
	}
	if d.HasChange("max_duration") {
		hasBasicChange = true
		updateInfo.SetMaxDuration(int32(d.Get("max_duration").(int)))
	}
	if d.HasChange("request_template_id") {
		hasBasicChange = true
		updateInfo.SetRequestTemplateId(d.Get("request_template_id").(string))
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

	if d.HasChange("reviewer") {
		// If all reviewer blocks were unset, let's use the admin owner id. If we don't do this,
		// the resource will be configured to an invalid state that the Opal API will still accept,
		// but the resource will be unrequestable.
		reviewers := any([]any{map[string]any{"id": d.State().Attributes["admin_owner_id"]}})
		if reviewersBlock, ok := d.GetOk("reviewer"); ok {
			reviewers = reviewersBlock
		}
		if diag := resourceResourceUpdateReviewers(ctx, d, client, reviewers); diag != nil {
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
