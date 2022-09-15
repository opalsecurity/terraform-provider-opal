package opal

import (
	"context"

	"github.com/hashicorp/go-multierror"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/opalsecurity/opal-go"
)

func resourceOwner() *schema.Resource {
	return &schema.Resource{
		ReadContext: resourceOwnerRead,
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
				Description: "The amount of time (in minutes) before the next reviewer is notified. Use 0 to remove escalation policy.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"user": {
				Description: "The users for this owner.",
				Type:        schema.TypeList,
				Required:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Description: "The ID of the user.",
							Required:    true,
						},
						"email": {
							Type:        schema.TypeString,
							Description: "The email of the user.",
							Computed:    true,
						},
						"full_name": {
							Type:        schema.TypeString,
							Description: "The full name of the user.",
							Computed:    true,
						},
						"last_name": {
							Type:        schema.TypeString,
							Description: "The last name of the user.",
							Computed:    true,
						},
						"first_name": {
							Type:        schema.TypeString,
							Description: "The first name of the user.",
							Computed:    true,
						},
						"position": {
							Type:        schema.TypeString,
							Description: "The position of the user.",
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func resourceOwnerRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*opal.APIClient)

	id := d.Get("id").(string)
	owner, _, err := client.OwnersApi.GetOwner(ctx, id).Execute()
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(owner.OwnerId)
	if err := multierror.Append(
		d.Set("name", owner.Name),
		d.Set("description", owner.Description),
		d.Set("access_request_escalation_period", owner.AccessRequestEscalationPeriod),
	); err.ErrorOrNil() != nil {
		return diag.FromErr(err)
	}

	users, _, err := client.OwnersApi.GetOwnerUsers(ctx, id).Execute()
	if err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("user", flattenOwnerUsers(users)); err != nil {
		return diag.FromErr(err)
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
