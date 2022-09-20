package opal

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/opalsecurity/opal-go"
)

// allowedResourceTypes are the values from AllowedResourceTypeEnumEnumValues
// with type []string.
var allowedResourceTypes = (func() []string {
	rv := make([]string, 0, len(opal.AllowedResourceTypeEnumEnumValues))
	for _, v := range opal.AllowedResourceTypeEnumEnumValues {
		rv = append(rv, string(v))
	}
	return rv
})()

func resourceResource() *schema.Resource {
	return &schema.Resource{
		Description: "An Opal Resource resource.",
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
			"owner_id": {
				Description: "The owner ID for this resource.",
				Type:        schema.TypeString,
				Required:    true,
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
				Description: "The JSON metadata about the remote resource. Include only for items linked to remote systems. See [the guide](https://docs.opal.dev/reference/how-opal).",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"visibility_group": {
				Description: "The groups that can see this resource.",
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
