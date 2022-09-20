package opal

import (
	"context"

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
var allowedResourceVisibilityTypes = enumSliceToStringSlice(opal.AllowedVisibilityTypeEnumEnumValues)

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
				Required:    true,
			},
			"resource_type": {
				Description:  "The type of the resource, i.e. AWS_EC2_INSTANCE.",
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice(allowedResourceTypes, false),
				Required:     true,
			},
			"app_id": {
				Description: "The ID of the app integration that provides the resource as a UUID.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"admin_owner_id": {
				Description: "The admin owner ID for this resource. By default, this is set to the application admin owner.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"require_manager_approval": {
				Description: "Require the requester's manager's approval for requests to this resource.",
				Type:        schema.TypeBool,
				Optional:    true,
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
			"require_support_ticket": {
				Description: "Require that requesters attach a support ticket to requests for this resource.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"max_duration": {
				Description: "The maximum duration for which this resource can be requested (in minutes). By default, the max duration is indefinite access.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"request_template_id": {
				Description: "The ID of a request template for this resource. You can get this ID from the URL in the Opal web app.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"metadata": {
				Description:  "The JSON metadata about the remote resource. Include only for items linked to remote systems. See [the guide](https://docs.opal.dev/reference/how-opal).",
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validateMetadata,
			},
			"visibility": {
				Description: "The visibility of this resource.",
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"group": {
							Description: "The groups that can see this resource.",
							Type:        schema.TypeList,
							Required:    true,
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
						"level": {
							Description:  "The visiblity level of the resource, i.e. LIMITED or GLOBAL.",
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice(allowedResourceVisibilityTypes, false),
							Required:     true,
						},
					},
				},
			},
			"reviewer": {
				Description: "A required reviewer for this resource.",
				Type:        schema.TypeList,
				Optional:    true,
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
			// XXX: Require mfa to connect to this resource.
		},
	}
}

func resourceResourceCreate(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	client := m.(*opal.APIClient)

	name := d.Get("name").(string)
	resourceType := opal.ResourceTypeEnum(d.Get("resource_type").(string))
	appID := d.Get("app_id").(string)

	createInfo := opal.NewCreateResourceInfo(name, resourceType, appID)
	// createInfo.SetMetadata()
	// createInfo.SetDescription()

	resource, _, err := client.ResourcesApi.CreateResource(ctx).CreateResourceInfo(*createInfo).Execute()
	if err != nil {
		return diag.FromErr(err)
	}

	tflog.Debug(ctx, "Created opal resource", map[string]any{
		"name": name,
		"id":   resource.AdminOwnerId,
	})

	// Because resource creation does not let us set an owner immediately (the owner is),
	// we need to update the resource after creation.
	if ownerIDI, ok := d.GetOk("admin_owner_id"); ok {
		updateInfo := opal.NewUpdateResourceInfo(resource.ResourceId)
		updateInfo.SetAdminOwnerId(ownerIDI.(string))
		if _, _, err := client.ResourcesApi.UpdateResources(ctx).UpdateResourceInfoList(*opal.NewUpdateResourceInfoList([]opal.UpdateResourceInfo{*updateInfo})).Execute(); err != nil {
			return diag.FromErr(err)
		}
	}

	if visibilityInfoI, ok := d.GetOk("visibility"); ok {
		info := (visibilityInfoI.([]any)[0]).(map[string]any)
		visibilityInfo := *opal.NewVisibilityInfo(opal.VisibilityTypeEnum(info["level"].(string)))

		rawGroups := d.Get("group").([]any)
		groupIds := make([]string, 0, len(rawGroups))
		for _, rawGroup := range rawGroups {
			group := rawGroup.(map[string]any)
			groupIds = append(groupIds, group["id"].(string))
		}
		visibilityInfo.SetVisibilityGroupIds(groupIds)

		if _, _, err := client.ResourcesApi.SetResourceVisibility(ctx, d.Id()).VisibilityInfo(visibilityInfo).Execute(); err != nil {
			return diag.FromErr(err)
		}
	}

	if reviewersI, ok := d.GetOk("reviewer"); ok {
		rawReviewers := reviewersI.([]any)
		reviewerIds := make([]string, 0, len(rawReviewers))
		for _, rawReviewer := range rawReviewers {
			group := rawReviewer.(map[string]any)
			reviewerIds = append(reviewerIds, group["id"].(string))
		}

		if _, _, err := client.ResourcesApi.SetResourceReviewers(ctx, d.Id()).ReviewerIDList(*opal.NewReviewerIDList(reviewerIds)).Execute(); err != nil {
			return diag.FromErr(err)
		}
	}

	// XXX: Update audit channel...
	// XXX: Update mfa required for connnect...

	d.SetId(resource.ResourceId)
	return resourceResourceRead(ctx, d, m)
}

func resourceResourceRead(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	client := m.(*opal.APIClient)

	resource, _, err := client.ResourcesApi.GetResource(ctx, d.Id()).Execute()
	if err != nil {
		return diag.FromErr(err)
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
		d.Set("require_suppot_ticket", resource.RequireSupportTicket),
		d.Set("max_duration", resource.MaxDuration),
		d.Set("request_template_id", resource.RequestTemplateId),
		// XXX: We don't get the metadata back. Will terraform state be okay?
	); err.ErrorOrNil() != nil {
		return diag.FromErr(err)
	}

	visibility, _, err := client.ResourcesApi.GetResourceVisibility(ctx, resource.ResourceId).Execute()
	if err != nil {
		return diag.FromErr(err)
	}

	visibilityGroups := make([]any, 0, len(visibility.VisibilityGroupIds))
	for _, groupID := range visibility.VisibilityGroupIds {
		visibilityGroups = append(visibilityGroups, map[string]any{
			"id": groupID,
		})
	}
	d.Set("visibility", map[string]any{
		"level": visibility.Visibility,
		"group": visibilityGroups,
	})

	reviewerIDs, _, err := client.ResourcesApi.GetResourceReviewers(ctx, resource.ResourceId).Execute()
	reviewers := make([]any, 0, len(reviewerIDs))
	for _, reviewerID := range reviewerIDs {
		reviewers = append(reviewers, map[string]any{
			"id": reviewerID,
		})
	}

	// XXX: Read out message channels, mfa required to connect.

	return nil
}

func resourceResourceUpdate(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	return resourceResourceRead(ctx, d, m)
}

func resourceResourceDelete(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	client := m.(*opal.APIClient)

	if _, err := client.ResourcesApi.DeleteResource(ctx, d.Id()).Execute(); err != nil {
		return diag.FromErr(err)
	}

	return nil
}
