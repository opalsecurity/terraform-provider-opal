package opal

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/hashicorp/go-multierror"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/opalsecurity/opal-go"
)

var allowedGroupTypes = enumSliceToStringSlice(opal.AllowedGroupTypeEnumEnumValues)

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
			"require_manager_approval": {
				Description: "Require the requester's manager's approval for requests to this group.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"auto_approval": {
				Description: "Automatically approve all requests for this group without review.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"require_mfa_to_approve": {
				Description: "Require that reviewers MFA to approve requests for this group.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"require_support_ticket": {
				Description: "Require that requesters attach a support ticket to requests for this group.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"max_duration": {
				Description: "The maximum duration for which this group can be requested (in minutes). By default, the max duration is indefinite access.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"request_template_id": {
				Description: "The ID of a request template for this group. You can get this ID from the URL in the Opal web app.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"remote_info": {
				Description: "Remote info that is required for the creation of remote groups.",
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
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
				Type:        schema.TypeList,
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
			"reviewer": {
				Description: "A required reviewer for this group. If none are specified.",
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Description: "The ID of the owner that must review requests to this group.",
							Type:        schema.TypeString,
							Required:    true,
						},
					},
				},
			},
			"audit_message_channel": {
				Description: "An audit message channel for this group.",
				Type:        schema.TypeList,
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
		},
	}
}

func resourceGroupCreate(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	client := m.(*opal.APIClient)

	name := d.Get("name").(string)
	groupType := opal.GroupTypeEnum(d.Get("group_type").(string))
	appID := d.Get("app_id").(string)

	createInfo := opal.NewCreateGroupInfo(name, groupType, appID)
	if descI, ok := d.GetOk("description"); ok {
		createInfo.SetDescription(descI.(string))
	}
	if remoteInfoI, ok := d.GetOk("remote_info"); ok {
		remoteInfo, err := parseGroupRemoteInfo(remoteInfoI)
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

	// Because group creation does not let us set some properties immediately,
	// we may have to update them in a follow up request.
	adminOwnerIDI, adminOwnerIDOk := d.GetOk("admin_owner_id")
	requireManagerApprovalI, requireManagerApprovalOk := d.GetOk("require_manager_approval")
	autoApprovalI, autoApprovalOk := d.GetOk("auto_approval")
	requireMfaToApproveI, requireMfaToApproveOk := d.GetOk("require_mfa_to_approve")
	requireSupportTicketI, requireSupportTicketOk := d.GetOk("require_support_ticket")
	maxDurationI, maxDurationOk := d.GetOk("max_duration")
	requestTemplateIDI, requestTemplateIDOk := d.GetOk("request_template_id")
	if adminOwnerIDOk || requireManagerApprovalOk || autoApprovalOk || requireMfaToApproveOk || requireSupportTicketOk || maxDurationOk || requestTemplateIDOk {
		updateInfo := opal.NewUpdateGroupInfo(group.GroupId)
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
		if requireSupportTicketOk {
			updateInfo.SetRequireSupportTicket(requireSupportTicketI.(bool))
		}
		if maxDurationOk {
			updateInfo.SetMaxDuration(int32(maxDurationI.(int)))
		}
		if requestTemplateIDOk {
			updateInfo.SetRequestTemplateId(requestTemplateIDI.(string))
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

	if reviewersI, ok := d.GetOk("reviewer"); ok {
		if diag := resourceGroupUpdateReviewers(ctx, d, client, reviewersI); diag != nil {
			return diag
		}
	} else if adminOwnerIDOk {
		// If the admin owner was set during creation, we should also set
		// the required reviewer to be the same so that it is consistent.
		//
		// Otherwise, if it's unset, the Opal API will automatically set it to
		// the app owner.
		if diag := resourceGroupUpdateReviewers(ctx, d, client, []any{map[string]any{"id": adminOwnerIDI}}); diag != nil {
			return diag
		}
	}

	if _, ok := d.GetOk("audit_message_channel"); ok {
		if diag := resourceGroupUpdateAuditMessageChannels(ctx, d, client); diag != nil {
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
		rawGroups := visibilityGroupI.([]any)
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
		rawChannels := auditMessageChannelsI.([]any)
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

func resourceGroupUpdateReviewers(ctx context.Context, d *schema.ResourceData, client *opal.APIClient, reviewersI any) diag.Diagnostics {
	rawReviewers := reviewersI.([]any)
	reviewerIds := make([]string, 0, len(rawReviewers))
	for _, rawReviewer := range rawReviewers {
		reviewer := rawReviewer.(map[string]any)
		reviewerIds = append(reviewerIds, reviewer["id"].(string))
	}
	tflog.Debug(ctx, "Setting group reviewers", map[string]any{
		"id":          d.Id(),
		"reviewerIds": reviewerIds,
	})

	if _, _, err := client.GroupsApi.SetGroupReviewers(ctx, d.Id()).ReviewerIDList(*opal.NewReviewerIDList(reviewerIds)).Execute(); err != nil {
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

	d.SetId(group.GroupId)
	if err := multierror.Append(
		d.Set("name", group.Name),
		d.Set("description", group.Description),
		d.Set("group_type", group.GroupType),
		d.Set("app_id", group.AppId),
		d.Set("admin_owner_id", group.AdminOwnerId),
		d.Set("require_manager_approval", group.RequireManagerApproval),
		d.Set("auto_approval", group.AutoApproval),
		d.Set("require_mfa_to_approve", group.RequireMfaToApprove),
		d.Set("require_support_ticket", group.RequireSupportTicket),
		d.Set("max_duration", group.MaxDuration),
		d.Set("request_template_id", group.RequestTemplateId),
		// XXX: We don't get the metadata back. Will terraform state be okay?
	); err.ErrorOrNil() != nil {
		return diagFromErr(ctx, err)
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

	auditChannels := make([]any, 0, len(auditChannelsResponse.Channels))
	for _, channel := range auditChannelsResponse.Channels {
		auditChannels = append(auditChannels, map[string]any{
			"id": channel.MessageChannelId,
		})
	}
	d.Set("audit_message_channel", auditChannels)

	reviewerIDs, _, err := client.GroupsApi.GetGroupReviewers(ctx, group.GroupId).Execute()
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

	// XXX: Read out message channels.

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
	if d.HasChange("require_manager_approval") {
		hasBasicUpdate = true
		updateInfo.SetRequireManagerApproval(d.Get("require_manager_approval").(bool))
	}
	if d.HasChange("auto_approval") {
		hasBasicUpdate = true
		updateInfo.SetAutoApproval(d.Get("auto_approval").(bool))
	}
	if d.HasChange("require_mfa_to_approve") {
		hasBasicUpdate = true
		updateInfo.SetRequireMfaToApprove(d.Get("require_mfa_to_approve").(bool))
	}
	if d.HasChange("require_support_ticket") {
		hasBasicUpdate = true
		updateInfo.SetRequireSupportTicket(d.Get("require_support_ticket").(bool))
	}
	if d.HasChange("max_duration") {
		hasBasicUpdate = true
		updateInfo.SetMaxDuration(int32(d.Get("max_duration").(int)))
	}
	if d.HasChange("request_template_id") {
		hasBasicUpdate = true
		updateInfo.SetRequestTemplateId(d.Get("request_template_id").(string))
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

	if d.HasChange("reviewer") {
		// If all reviewer blocks were unset, let's use the admin owner id. If we don't do this,
		// the group will be configured to an invalid state that the Opal API will still accept,
		// but the group will be unrequestable.
		reviewers := any([]any{map[string]any{"id": d.State().Attributes["admin_owner_id"]}})
		if reviewersBlock, ok := d.GetOk("reviewer"); ok {
			reviewers = reviewersBlock
		}
		if diag := resourceGroupUpdateReviewers(ctx, d, client, reviewers); diag != nil {
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
