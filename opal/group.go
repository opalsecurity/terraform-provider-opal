package opal

import (
	"context"
	"fmt"
	"reflect"
	"sort"

	"github.com/opalsecurity/opal-go"
	"github.com/pkg/errors"

	"github.com/hashicorp/go-multierror"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

var allowedGroupTypes = enumSliceToStringSlice(opal.AllowedGroupTypeEnumEnumValues)
var allowedReviewerStageOperators = []string{"AND", "OR"}

func dataSourceGroup() *schema.Resource {
	return &schema.Resource{
		Description: "An Opal group data source.",
		ReadContext: dataSourceGroupRead,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"id": {
				Description: "The ID of the group.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"name": {
				Description: "The name of the group.",
				Type:        schema.TypeString,
				Computed:    true,
			},
		},
	}
}

func dataSourceGroupRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*opal.APIClient)

	id, idOk := d.GetOk("id")
	var group opal.Group
	if idOk {
		groups, _, err := client.GroupsApi.GetGroups(ctx).GroupIds([]string{id.(string)}).Execute()
		if err != nil {
			return diagFromErr(ctx, err)
		}
		if len(groups.Results) != 1 {
			return diagFromErr(ctx, fmt.Errorf("expected 1 group returned but got %d", len(groups.Results)))
		}
		group = groups.Results[0]
		if err != nil {
			return diagFromErr(ctx, err)
		}
	} else {
		return diagFromErr(ctx, errors.New("must provide id for resource data source"))
	}

	d.SetId(group.GroupId)
	if err := multierror.Append(
		d.Set("name", group.Name),
	); err.ErrorOrNil() != nil {
		return diagFromErr(ctx, err)
	}
	return nil
}

func resourceGroup() *schema.Resource {
	return &schema.Resource{
		Description:   "An Opal Group resource.",
		CreateContext: resourceGroupCreate,
		ReadContext:   resourceGroupRead,
		UpdateContext: resourceGroupUpdate,
		DeleteContext: resourceGroupDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"id": {
				Description: "The ID of the group.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"name": {
				Description: "The name of the group.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"description": {
				Description: "The description of the group.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"group_type": {
				Description:  "The type of the group, i.e. GIT_HUB_TEAM.",
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice(allowedGroupTypes, false),
				ForceNew:     true,
				Required:     true,
			},
			"app_id": {
				Description: "The ID of the app integration that provides the group. You can get this value from the URL of the app in the Opal web app. For an Opal group, use the ID from the Opal app in the apps view.",
				Type:        schema.TypeString,
				ForceNew:    true,
				Required:    true,
			},
			"admin_owner_id": {
				Description: "The admin owner ID for this group. By default, this is set to the application admin owner.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"require_mfa_to_approve": {
				Description: "Require that reviewers MFA to approve requests for this group.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"remote_info": {
				Description: "Remote info that is required for the creation of remote groups.",
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Computed:    true,
				MaxItems:    1,
				Elem:        groupRemoteInfoElem(),
			},
			"visibility": {
				Description:  "The visibility level of the group, i.e. LIMITED or GLOBAL.",
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice(allowedVisibilityTypes, false),
				Optional:     true,
				Default:      "GLOBAL",
			},
			"visibility_group": {
				Description: "The groups that can see this group when visibility is limited. If not specified, only users with direct access can see this resource when visibility is set to LIMITED.",
				Type:        schema.TypeSet,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Description: "The ID of the group that can see this group.",
							Type:        schema.TypeString,
							Required:    true,
						},
					},
				},
			},
			"audit_message_channel": {
				Description: "An audit message channel for this group.",
				Type:        schema.TypeSet,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Description: "The ID of the message channel for this group.",
							Type:        schema.TypeString,
							Required:    true,
						},
					},
				},
			},
			"manage_resources": {
				Description: "Boolean flag to indicate if you intend to manage group <-> resource relationships via terraform.",
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
			},
			"resource": {
				Description: "A resource that members of the group get access to.",
				Type:        schema.TypeSet,
				Optional:    true,
				DiffSuppressFunc: func(k, oldValue, newValue string, d *schema.ResourceData) bool {
					if manage, ok := d.GetOk("manage_resources"); ok {
						return !manage.(bool)
					}
					return true
				},
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Description: "The ID of the resource.",
							Type:        schema.TypeString,
							Required:    true,
						},
						"access_level_remote_id": {
							Description: "The access level remote id of the resource that this group gives access to.",
							Type:        schema.TypeString,
							Optional:    true,
						},
					},
				},
			},
			"on_call_schedule": {
				Description: "An on call schedule for this group.",
				Type:        schema.TypeSet,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Description: "The UUID of the on call schedule for this group.",
							Type:        schema.TypeString,
							Required:    true,
						},
					},
				},
			},
			"request_configuration": {
				Description: "A request configuration for this group.",
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"group_ids": {
							Description: "The group IDs satisfying this request configuration. For the default request configuration, this should be empty and priority should be 0, otherwise, this should contain one group ID.",
							Type:        schema.TypeList,
							Optional:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							MaxItems: 1,
						},
						"is_requestable": {
							Description: "For users satisfying the condition, allow them to create an access request for this group. By default, any group is requestable.",
							Type:        schema.TypeBool,
							Optional:    true,
							Default:     true,
						},
						"auto_approval": {
							Description: "For users satisfying the condition, automatically approve all requests for this group without review.",
							Type:        schema.TypeBool,
							Optional:    true,
						},
						"require_mfa_to_request": {
							Description: "For users satisfying the condition, require that users MFA to request this group.",
							Type:        schema.TypeBool,
							Optional:    true,
						},
						"require_support_ticket": {
							Description: "For users satisfying the condition, require that requesters attach a support ticket to request for this group.",
							Type:        schema.TypeBool,
							Optional:    true,
						},
						"max_duration": {
							Description: "For users satisfying the condition, the maximum duration for which this group can be requested (in minutes).",
							Type:        schema.TypeInt,
							Optional:    true,
							Default:     -1,
						},
						"recommended_duration": {
							Description: "For users satisfying the condition, the recommended duration for which the group should be requested (in minutes). Will be the default value in a request. Use -1 to set to indefinite.",
							Type:        schema.TypeInt,
							Optional:    true,
							Default:     -1,
						},
						"request_template_id": {
							Description: "For users satisfying the condition, the ID of a request template for this group. You can get this ID from the URL in the Opal web app.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"priority": {
							Description: "The priority of this request configuration. The higher the number, the higher the priority. Defaults to 0.",
							Type:        schema.TypeInt,
							Optional:    true,
							Default:     0,
						},
						"reviewer_stage": {
							Description: "A reviewer stage for this request configuration. You are allowed to provide up to 3.",
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
					},
				},
			},
		},
	}
}

func resourceGroupCreate(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	client := m.(*opal.APIClient)

	if err := validateReviewerConfigDuringCreate(d); err != nil {
		return diagFromErr(ctx, err)
	}

	if err := validateRequestConfigurationListDuringCreate(ctx, d); err != nil {
		return diagFromErr(ctx, err)
	}

	name := d.Get("name").(string)
	groupType := opal.GroupTypeEnum(d.Get("group_type").(string))
	appID := d.Get("app_id").(string)

	createInfo := opal.NewCreateGroupInfo(name, groupType, appID)
	if descI, ok := d.GetOk("description"); ok {
		createInfo.SetDescription(descI.(string))
	}
	if remoteInfoI, ok := d.GetOk("remote_info"); ok {
		remoteInfo, err := groupRemoteInfoTerraformToAPI(remoteInfoI)
		if err != nil {
			return diagFromErr(ctx, err)
		}
		createInfo.SetRemoteInfo(*remoteInfo)
	}

	group, _, err := client.GroupsApi.CreateGroup(ctx).CreateGroupInfo(*createInfo).Execute()
	if err != nil {
		return diagFromErr(ctx, err)
	}
	d.SetId(group.GroupId)

	tflog.Debug(ctx, "Created opal group", map[string]any{
		"name": name,
		"id":   d.Id(),
	})

	requestConfigurationsListI, requestConfigurationsListOk := d.GetOk("request_configuration")

	// Because group creation does not let us set some properties immediately,
	// we may have to update them in a follow up request.
	adminOwnerIDI, adminOwnerIDOk := d.GetOk("admin_owner_id")
	requireMfaToApproveI, requireMfaToApproveOk := d.GetOkExists("require_mfa_to_approve")

	if adminOwnerIDOk || requireMfaToApproveOk || requestConfigurationsListOk {
		updateInfo := opal.NewUpdateGroupInfo(group.GroupId)
		if adminOwnerIDOk {
			updateInfo.SetAdminOwnerId(adminOwnerIDI.(string))
		}
		if requireMfaToApproveOk {
			updateInfo.SetRequireMfaToApprove(requireMfaToApproveI.(bool))
		}
		if requestConfigurationsListOk {
			requestConfigurationsList, err := parseRequestConfigurationList(ctx, requestConfigurationsListI)
			if err != nil {
				return diagFromErr(ctx, err)
			}
			if err := validateRequestConfigurationListDuringCreate(ctx, d); err != nil {
				return diagFromErr(ctx, err)
			}
			updateInfo.SetRequestConfigurationList(*requestConfigurationsList)
		}

		tflog.Debug(ctx, "Immediately updating opal group", map[string]any{
			"name":       name,
			"updateInfo": updateInfo,
		})

		if _, _, err := client.GroupsApi.UpdateGroups(ctx).UpdateGroupInfoList(*opal.NewUpdateGroupInfoList([]opal.UpdateGroupInfo{*updateInfo})).Execute(); err != nil {
			return diagFromErr(ctx, err)
		}
	}

	if _, ok := d.GetOk("visibility"); ok {
		if diag := resourceGroupUpdateVisibility(ctx, d, client); diag != nil {
			return diag
		}
	}

	if _, ok := d.GetOk("resource"); ok {
		if diag := resourceGroupUpdateResources(ctx, d, client); diag != nil {
			return diag
		}
	}

	if _, ok := d.GetOk("audit_message_channel"); ok {
		if diag := resourceGroupUpdateAuditMessageChannels(ctx, d, client); diag != nil {
			return diag
		}
	}

	if _, ok := d.GetOk("on_call_schedule"); ok {
		if diag := resourceGroupUpdateOnCallSchedules(ctx, d, client); diag != nil {
			return diag
		}
	}

	return resourceGroupRead(ctx, d, m)
}

func resourceGroupUpdateVisibility(ctx context.Context, d *schema.ResourceData, client *opal.APIClient) diag.Diagnostics {
	visibilityInfo := *opal.NewVisibilityInfo(opal.VisibilityTypeEnum(opal.VISIBILITYTYPEENUM_GLOBAL))
	if visibilityI, ok := d.GetOk("visibility"); ok {
		visibilityInfo.SetVisibility(opal.VisibilityTypeEnum(visibilityI.(string)))
	}

	if visibilityGroupI, ok := d.GetOk("visibility_group"); ok {
		rawGroups := visibilityGroupI.(*schema.Set).List()
		groupIds := make([]string, 0, len(rawGroups))
		for _, rawGroup := range rawGroups {
			group := rawGroup.(map[string]any)
			groupIds = append(groupIds, group["id"].(string))
		}
		visibilityInfo.SetVisibilityGroupIds(groupIds)
	}

	if _, _, err := client.GroupsApi.SetGroupVisibility(ctx, d.Id()).VisibilityInfo(visibilityInfo).Execute(); err != nil {
		return diagFromErr(ctx, err)
	}
	return nil
}

func resourceGroupUpdateAuditMessageChannels(ctx context.Context, d *schema.ResourceData, client *opal.APIClient) diag.Diagnostics {
	var channelIDs []string
	if auditMessageChannelsI, ok := d.GetOk("audit_message_channel"); ok {
		rawChannels := auditMessageChannelsI.(*schema.Set).List()
		for _, rawChannel := range rawChannels {
			channel := rawChannel.(map[string]any)
			channelIDs = append(channelIDs, channel["id"].(string))
		}
	}

	channelList := *opal.NewMessageChannelIDList(channelIDs)
	if _, _, err := client.GroupsApi.SetGroupMessageChannels(ctx, d.Id()).MessageChannelIDList(channelList).Execute(); err != nil {
		return diagFromErr(ctx, err)
	}

	return nil
}

func resourceGroupUpdateOnCallSchedules(ctx context.Context, d *schema.ResourceData, client *opal.APIClient) diag.Diagnostics {
	var onCallScheduleIDs []string
	if onCallSchedulesI, ok := d.GetOk("on_call_schedule"); ok {
		rawOnCallSchedules := onCallSchedulesI.(*schema.Set).List()
		for _, rawOnCallSchedule := range rawOnCallSchedules {
			onCallSchedule := rawOnCallSchedule.(map[string]any)
			onCallScheduleIDs = append(onCallScheduleIDs, onCallSchedule["id"].(string))
		}
	}

	onCallScheduleList := *opal.NewOnCallScheduleIDList(onCallScheduleIDs)
	if _, _, err := client.GroupsApi.SetGroupOnCallSchedules(ctx, d.Id()).OnCallScheduleIDList(onCallScheduleList).Execute(); err != nil {
		return diagFromErr(ctx, err)
	}

	return nil
}

func resourceGroupUpdateResources(ctx context.Context, d *schema.ResourceData, client *opal.APIClient) diag.Diagnostics {
	var rawResources []any
	if resourceI, ok := d.GetOk("resource"); ok {
		rawResources = resourceI.(*schema.Set).List()
	}

	resourcesWithAccessLevel := make([]opal.ResourceWithAccessLevel, 0, len(rawResources))
	for _, rawResource := range rawResources {
		resource := rawResource.(map[string]any)
		var accessLevelRemoteIDPtr *string
		accessLevelRemoteID := resource["access_level_remote_id"].(string)
		if accessLevelRemoteID != "" {
			accessLevelRemoteIDPtr = &accessLevelRemoteID
		}

		resourcesWithAccessLevel = append(resourcesWithAccessLevel, opal.ResourceWithAccessLevel{
			ResourceId:          resource["id"].(string),
			AccessLevelRemoteId: accessLevelRemoteIDPtr,
		})
	}
	tflog.Debug(ctx, "Setting group resources", map[string]any{
		"id":        d.Id(),
		"resources": resourcesWithAccessLevel,
	})

	updateInfo := opal.UpdateGroupResourcesInfo{
		Resources: resourcesWithAccessLevel,
	}

	if _, err := client.GroupsApi.SetGroupResources(ctx, d.Id()).UpdateGroupResourcesInfo(updateInfo).Execute(); err != nil {
		return diagFromErr(ctx, err)
	}
	return nil
}

func resourceGroupUpdateReviewerStages(ctx context.Context, d *schema.ResourceData, client *opal.APIClient, reviewerStagesI any) diag.Diagnostics {
	reviewerStages, err := parseReviewerStages(reviewerStagesI)
	if err != nil {
		return diagFromErr(ctx, err)
	}

	for _, reviewerStage := range reviewerStages {
		tflog.Debug(ctx, "Updating reviewer stage", map[string]any{
			"id":                     d.Id(),
			"requireManagerApproval": reviewerStage.RequireManagerApproval,
			"operator":               reviewerStage.Operator,
			"reviewerIds":            reviewerStage.OwnerIds,
		})
	}

	if _, _, err := client.GroupsApi.SetGroupReviewerStages(ctx, d.Id()).ReviewerStageList(*opal.NewReviewerStageList(reviewerStages)).Execute(); err != nil {
		return diagFromErr(ctx, err)
	}
	return nil
}

func extractReviewerIDs(reviewersI any) ([]string, error) {
	var rawReviewers []any
	switch reviewersI := reviewersI.(type) {
	case []any:
		rawReviewers = reviewersI
	case *schema.Set:
		rawReviewers = reviewersI.List()
	default:
		return nil, errors.Errorf("bad type passed: %v", reflect.TypeOf(reviewersI))
	}
	reviewerIds := make([]string, 0, len(rawReviewers))
	for _, rawReviewer := range rawReviewers {
		reviewer := rawReviewer.(map[string]any)
		reviewerIds = append(reviewerIds, reviewer["id"].(string))
	}

	return reviewerIds, nil
}

func resourceGroupRead(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	client := m.(*opal.APIClient)

	groups, _, err := client.GroupsApi.GetGroups(ctx).GroupIds([]string{d.Id()}).Execute()
	if err != nil {
		return diagFromErr(ctx, err)
	}
	if len(groups.Results) != 1 {
		return diagFromErr(ctx, fmt.Errorf("expected 1 group returned but got %d", len(groups.Results)))
	}
	group := groups.Results[0]

	requestConfigurations := make([]map[string]interface{}, 0, len(group.RequestConfigurationList))
	sort.SliceStable(group.RequestConfigurationList, func(i, j int) bool {
		return group.RequestConfigurationList[i].Priority < group.RequestConfigurationList[j].Priority
	})

	for _, requestConfiguration := range group.RequestConfigurationList {
		requestConfigurationI, err := parseSDKRequestConfiguration(ctx, &requestConfiguration)
		if err != nil {
			return diagFromErr(ctx, err)
		}
		requestConfigurations = append(requestConfigurations, requestConfigurationI)
	}

	d.SetId(group.GroupId)
	if err := multierror.Append(
		d.Set("name", group.Name),
		d.Set("description", group.Description),
		d.Set("group_type", group.GroupType),
		d.Set("app_id", group.AppId),
		d.Set("admin_owner_id", group.AdminOwnerId),
		d.Set("require_mfa_to_approve", group.RequireMfaToApprove),
	); err.ErrorOrNil() != nil {
		return diagFromErr(ctx, err)
	}

	remoteInfoI, err := groupRemoteInfoAPIToTerraform(group.RemoteInfo)
	if err != nil {
		return diagFromErr(ctx, err)
	}
	if remoteInfoI != nil {
		d.Set("remote_info", remoteInfoI)
	}

	if len(requestConfigurations) != 0 {
		d.Set("request_configuration", requestConfigurations)
	}

	visibility, _, err := client.GroupsApi.GetGroupVisibility(ctx, group.GroupId).Execute()
	if err != nil {
		return diagFromErr(ctx, err)
	}

	d.Set("visibility", visibility.Visibility)
	flattenedGroups := make([]any, 0, len(visibility.VisibilityGroupIds))
	for _, groupID := range visibility.VisibilityGroupIds {
		flattenedGroups = append(flattenedGroups, map[string]any{"id": groupID})
	}
	d.Set("visibility_group", flattenedGroups)

	auditChannelsResponse, _, err := client.GroupsApi.GetGroupMessageChannels(ctx, group.GroupId).Execute()
	if err != nil {
		return diagFromErr(ctx, err)
	}

	auditChannels := make([]any, 0, len(auditChannelsResponse.Channels))
	for _, channel := range auditChannelsResponse.Channels {
		auditChannels = append(auditChannels, map[string]any{
			"id": channel.MessageChannelId,
		})
	}
	d.Set("audit_message_channel", auditChannels)

	onCallSchedulesResponse, _, err := client.GroupsApi.GetGroupOnCallSchedules(ctx, group.GroupId).Execute()
	if err != nil {
		return diagFromErr(ctx, err)
	}
	onCallSchedules := make([]any, 0, len(onCallSchedulesResponse.OnCallSchedules))
	for _, onCallSchedule := range onCallSchedulesResponse.OnCallSchedules {
		onCallSchedules = append(onCallSchedules, map[string]any{
			"id": onCallSchedule.OnCallScheduleId,
		})
	}
	d.Set("on_call_schedule", onCallSchedules)

	groupResources, _, err := client.GroupsApi.GetGroupResources(ctx, group.GroupId).Execute()
	if err != nil {
		return diagFromErr(ctx, err)
	}

	groupResourcesI := make([]any, 0, len(groupResources.GroupResources))
	for _, groupResource := range groupResources.GroupResources {
		groupResourceI := map[string]any{
			"id":                     groupResource.ResourceId,
			"access_level_remote_id": groupResource.AccessLevel.AccessLevelRemoteId,
		}
		groupResourcesI = append(groupResourcesI, groupResourceI)
	}
	d.Set("resource", groupResourcesI)

	return nil
}

func resourceGroupUpdate(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	client := m.(*opal.APIClient)

	// Note that metadata, app_id, and"group_type force a recreation, so we do not need to
	// worry about those values here.
	hasBasicUpdate := false
	updateInfo := opal.NewUpdateGroupInfo(d.Id())
	if d.HasChange("name") {
		hasBasicUpdate = true
		updateInfo.SetName(d.Get("name").(string))
	}
	if d.HasChange("description") {
		hasBasicUpdate = true
		updateInfo.SetDescription(d.Get("description").(string))
	}
	if d.HasChange("admin_owner_id") {
		hasBasicUpdate = true
		updateInfo.SetAdminOwnerId(d.Get("admin_owner_id").(string))
	}
	if d.HasChange("require_mfa_to_approve") {
		hasBasicUpdate = true
		updateInfo.SetRequireMfaToApprove(d.Get("require_mfa_to_approve").(bool))
	}
	if d.HasChange("request_configuration") {
		hasBasicUpdate = true
		requestConfigurationsListI, ok := d.GetOk("request_configuration")
		if ok {
			requestConfigurationsList, err := parseRequestConfigurationList(ctx, requestConfigurationsListI)
			if err != nil {
				return diagFromErr(ctx, err)
			}
			updateInfo.SetRequestConfigurationList(*requestConfigurationsList)
		}
	}

	if hasBasicUpdate {
		_, _, err := client.GroupsApi.UpdateGroups(ctx).UpdateGroupInfoList(*opal.NewUpdateGroupInfoList([]opal.UpdateGroupInfo{*updateInfo})).Execute()
		if err != nil {
			return diagFromErr(ctx, err)
		}

	}

	if d.HasChange("visibility") || d.HasChange("visibility_group") {
		if diag := resourceGroupUpdateVisibility(ctx, d, client); diag != nil {
			return diag
		}
	}

	if d.HasChange("audit_message_channel") {
		if diag := resourceGroupUpdateAuditMessageChannels(ctx, d, client); diag != nil {
			return diag
		}
	}

	if d.HasChange("on_call_schedule") {
		if diag := resourceGroupUpdateOnCallSchedules(ctx, d, client); diag != nil {
			return diag
		}
	}

	if d.HasChange("resource") {
		if diag := resourceGroupUpdateResources(ctx, d, client); diag != nil {
			return diag
		}
	}

	return resourceGroupRead(ctx, d, m)
}

func resourceGroupDelete(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	client := m.(*opal.APIClient)
	tflog.Debug(ctx, "Deleting group", map[string]any{
		"id": d.Id(),
	})

	if _, err := client.GroupsApi.DeleteGroup(ctx, d.Id()).Execute(); err != nil {
		return diagFromErr(ctx, err)
	}

	return nil
}
