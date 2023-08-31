package opal

import (
	"context"
	"errors"
	"fmt"

	"github.com/hashicorp/go-multierror"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/opalsecurity/opal-go"
)

func dataSourceOwner() *schema.Resource {
	return &schema.Resource{
		Description: "An Opal owner data source.",
		ReadContext: dataSourceOwnerRead,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"id": {
				Description: "The ID of the owner.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"name": {
				Description: "The name of the owner.",
				Type:        schema.TypeString,
				Optional:    true,
			},
		},
	}
}

func dataSourceOwnerRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*opal.APIClient)

	id, idOk := d.GetOk("id")
	name, nameOk := d.GetOk("name")
	var owner *opal.Owner
	var err error
	if idOk {
		owner, _, err = client.OwnersApi.GetOwner(ctx, id.(string)).Execute()
		if err != nil {
			return diagFromErr(ctx, err)
		}
	} else if nameOk {
		owners, _, err := client.OwnersApi.GetOwners(ctx).Name(name.(string)).Execute()
		if err != nil {
			return diagFromErr(ctx, err)
		}
		if len(owners.Results) == 1 {
			owner = &owners.Results[0]
		} else if len(owners.Results) > 1 {
			return diagFromErr(ctx, fmt.Errorf("more than one owner found with name %s", name.(string)))
		} else {
			return diagFromErr(ctx, fmt.Errorf("no owners found with name %s", name.(string)))
		}
	} else {
		return diagFromErr(ctx, errors.New("must provide id or name for owner data source"))
	}

	d.SetId(owner.OwnerId)
	if err := multierror.Append(
		d.Set("name", owner.Name),
	); err.ErrorOrNil() != nil {
		return diagFromErr(ctx, err)
	}

	return nil
}

func resourceOwner() *schema.Resource {
	return &schema.Resource{
		Description:   "An Opal Owner resource.",
		CreateContext: resourceOwnerCreate,
		ReadContext:   resourceOwnerRead,
		UpdateContext: resourceOwnerUpdate,
		DeleteContext: resourceOwnerDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"id": {
				Description: "The ID of the owner.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"name": {
				Description: "The name of the owner.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"description": {
				Description: "A description of the owner.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"access_request_escalation_period": {
				Description: "The amount of time (in minutes) before the next reviewer is notified. By default, there is no escalation policy.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"user": {
				Description: "The users for this owner. If an escalation period is set, the order of the users will determine the escalation order.",
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Description: "The ID of the user.",
							Required:    true,
						},
					},
				},
			},
			"reviewer_message_channel_id": {
				Description: "The id of the message_channel that incoming reviews should be posted to.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"source_group_id": {
				Description: "The id of the group that owner users will be synced with. If set, adding or removing users will fail.",
				Type:        schema.TypeString,
				Optional:    true,
			},
		},
	}
}

func resourceOwnerCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*opal.APIClient)

	name := d.Get("name").(string)
	rawUsers := d.Get("user").([]interface{})
	userIds := make([]string, 0, len(rawUsers))
	for _, rawUser := range rawUsers {
		user := rawUser.(map[string]interface{})
		userIds = append(userIds, user["id"].(string))
	}

	createInfo := opal.NewCreateOwnerInfo(name, userIds)
	if descI, ok := d.GetOk("description"); ok {
		createInfo.SetDescription(descI.(string))
	}
	if reviewerMessageChannelIDI, ok := d.GetOk("reviewer_message_channel_id"); ok {
		createInfo.SetReviewerMessageChannelId(reviewerMessageChannelIDI.(string))
	}
	if accessRequestEscalationPeriodI, ok := d.GetOk("access_request_escalation_period"); ok {
		createInfo.SetAccessRequestEscalationPeriod(int32(accessRequestEscalationPeriodI.(int)))
	}
	if sourceGroupIDI, ok := d.GetOk("source_group_id"); ok {
		createInfo.SetSourceGroupId(sourceGroupIDI.(string))
	}

	owner, _, err := client.OwnersApi.CreateOwner(ctx).CreateOwnerInfo(*createInfo).Execute()
	if err != nil {
		return diagFromErr(ctx, err)
	}
	tflog.Debug(ctx, "Created owner", map[string]any{
		"name": name,
		"id":   owner.OwnerId,
	})

	d.SetId(owner.OwnerId)
	return resourceOwnerRead(ctx, d, m)
}

func resourceOwnerRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*opal.APIClient)

	id := d.Get("id").(string)
	owner, _, err := client.OwnersApi.GetOwner(ctx, id).Execute()
	if err != nil {
		return diagFromErr(ctx, err)
	}

	d.SetId(owner.OwnerId)
	if err := multierror.Append(
		d.Set("name", owner.Name),
		d.Set("description", owner.Description),
		d.Set("access_request_escalation_period", owner.AccessRequestEscalationPeriod),
	); err.ErrorOrNil() != nil {
		return diagFromErr(ctx, err)
	}

	if owner.ReviewerMessageChannelId.IsSet() {
		if err := d.Set("reviewer_message_channel_id", owner.ReviewerMessageChannelId.Get()); err != nil {
			return diagFromErr(ctx, err)
		}
	}

	ownerHasSourceGroup := false
	if owner.SourceGroupId.IsSet() {
		// NOTE: IsSet() is misleading, if the source_group_id value is nil,
		// IsSet will still return true, we need to check whether there's an
		// actual source group id in the value to ensure we need to import the
		// users as well or not
		ownerHasSourceGroup = owner.SourceGroupId.Get() != nil
		if err := d.Set("source_group_id", owner.SourceGroupId.Get()); err != nil {
			return diagFromErr(ctx, err)
		}
	}

	if !ownerHasSourceGroup {
		users, _, err := client.OwnersApi.GetOwnerUsers(ctx, id).Execute()
		if err != nil {
			return diagFromErr(ctx, err)
		}

		if err := d.Set("user", flattenOwnerUsers(users)); err != nil {
			return diagFromErr(ctx, err)
		}
	}

	return nil
}

func flattenOwnerUsers(userList *opal.UserList) []interface{} {
	if userList == nil {
		return nil
	}

	users := make([]interface{}, 0, len(userList.Users))

	for _, u := range userList.Users {
		users = append(users, map[string]interface{}{
			"id": u.UserId,
		})
	}

	return users
}

func resourceOwnerUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*opal.APIClient)

	updateInfo := opal.NewUpdateOwnerInfo(d.Id())
	updateInfo.SetName(d.Get("name").(string))
	if d.HasChange("description") {
		updateInfo.SetDescription(d.Get("description").(string))
	}

	if d.HasChange("access_request_escalation_period") {
		updateInfo.SetAccessRequestEscalationPeriod(int32(d.Get("access_request_escalation_period").(int)))
	}

	if d.HasChange("reviewer_message_channel_id") {
		updateInfo.SetReviewerMessageChannelId(d.Get("reviewer_message_channel_id").(string))
	}

	hasChangedSourceGroupID := false
	if d.HasChange("source_group_id") {
		updateInfo.SetSourceGroupId(d.Get("source_group_id").(string))
		hasChangedSourceGroupID = true
	}

	owner, _, err := client.OwnersApi.UpdateOwners(ctx).UpdateOwnerInfoList(*opal.NewUpdateOwnerInfoList([]opal.UpdateOwnerInfo{*updateInfo})).Execute()
	if err != nil {
		return diagFromErr(ctx, err)
	}

	d.SetId(owner.Owners[0].OwnerId)

	// We use HasChange here to prevent an extra API call if unchanged.
	if d.HasChange("user") || hasChangedSourceGroupID {
		rawUsers := d.Get("user").([]interface{})
		userIds := make([]string, 0, len(rawUsers))
		for _, rawUser := range rawUsers {
			user := rawUser.(map[string]interface{})
			userIds = append(userIds, user["id"].(string))
		}

		tflog.Debug(ctx, "Updating owner users", map[string]any{
			"id":    d.Id(),
			"users": userIds,
		})
		_, _, err := client.OwnersApi.SetOwnerUsers(ctx, d.Id()).UserIDList(*opal.NewUserIDList(userIds)).Execute()
		if err != nil {
			return diagFromErr(ctx, err)
		}
	}

	return resourceOwnerRead(ctx, d, m)
}

func resourceOwnerDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*opal.APIClient)

	_, err := client.OwnersApi.DeleteOwner(ctx, d.Id()).Execute()
	if err != nil {
		return diagFromErr(ctx, err)
	}

	return nil
}
