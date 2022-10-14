package opal

import (
	"context"

	"github.com/hashicorp/go-multierror"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/opalsecurity/opal-go"
)

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
				Computed:    true,
			},
			"access_request_escalation_period": {
				Description: "The amount of time (in minutes) before the next reviewer is notified. By default, there is no escalation policy.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"user": {
				Description: "The users for this owner. If an escalation period is set, the order of the users will determine the escalation order.",
				Type:        schema.TypeList,
				Required:    true,
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
			// XXX: Linked reviewer message channel.
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
	if accessRequestEscalationPeriodI, ok := d.GetOk("access_request_escalation_period"); ok {
		createInfo.SetAccessRequestEscalationPeriod(int32(accessRequestEscalationPeriodI.(int)))
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

	users, _, err := client.OwnersApi.GetOwnerUsers(ctx, id).Execute()
	if err != nil {
		return diagFromErr(ctx, err)
	}

	if err := d.Set("user", flattenOwnerUsers(users)); err != nil {
		return diagFromErr(ctx, err)
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
			"id":         u.UserId,
			"email":      u.Email,
			"full_name":  u.FullName,
			"first_name": u.FirstName,
			"last_name":  u.LastName,
			"position":   u.Position,
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

	owner, _, err := client.OwnersApi.UpdateOwners(ctx).UpdateOwnerInfoList(*opal.NewUpdateOwnerInfoList([]opal.UpdateOwnerInfo{*updateInfo})).Execute()
	if err != nil {
		return diagFromErr(ctx, err)
	}

	d.SetId(owner.Owners[0].OwnerId)

	// We use HasChange here to prevent an extra API call if unchanged.
	if d.HasChange("user") {
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
