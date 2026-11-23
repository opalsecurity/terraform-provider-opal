package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
)

func testAccRequestTemplateConfig(name string, reasonRequired bool) string {
	return testAccProviderConfig() + fmt.Sprintf(`
resource "opal_request_template" "%s" {
  name = "%s"
  custom_fields = [
    {
      name     = "Why do you need this?"
      type     = "SHORT_TEXT"
      required = %t
    },
    {
      name = "Reason category"
      type = "MULTI_CHOICE"
      metadata = {
        multi_choice_data = {
          options = ["Incident response", "Scheduled maintenance"]
        }
      }
    },
    {
      type = "CALLOUT"
      metadata = {
        callout_data = {
          severity = "WARNING"
          text     = "This role grants access to production customer data."
        }
      }
    },
  ]
}
`, name, name, reasonRequired)
}

func TestAccRequestTemplate_CRUD(t *testing.T) {
	t.Parallel()
	name := testAccName("reqtmpl")
	resourceName := "opal_request_template." + name

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProviderFactories,
		PreCheck:                 func() { testAccPreCheckRequestTemplate(t) },
		CheckDestroy: testAccCheckDestroyedByGET(
			"opal_request_template",
			"request_template_id",
			"/request-templates/%s",
		),
		Steps: []resource.TestStep{
			{
				Config: testAccRequestTemplateConfig(name, true),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "request_template_id"),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "custom_fields.#", "3"),
					resource.TestCheckResourceAttr(resourceName, "custom_fields.0.name", "Why do you need this?"),
					resource.TestCheckResourceAttr(resourceName, "custom_fields.0.type", "SHORT_TEXT"),
					resource.TestCheckResourceAttr(resourceName, "custom_fields.0.required", "true"),
					resource.TestCheckResourceAttr(resourceName, "custom_fields.1.type", "MULTI_CHOICE"),
					resource.TestCheckResourceAttr(resourceName, "custom_fields.2.type", "CALLOUT"),
					resource.TestCheckResourceAttrSet(resourceName, "custom_fields.2.name"),
					resource.TestCheckResourceAttr(
						resourceName,
						"custom_fields.2.metadata.callout_data.text",
						"This role grants access to production customer data.",
					),
				),
			},
			{
				ResourceName:                         resourceName,
				ImportState:                          true,
				ImportStateVerify:                    true,
				ImportStateIdFunc:                    testAccImportStateIDFromAttr(resourceName, "request_template_id"),
				ImportStateVerifyIdentifierAttribute: "request_template_id",
			},
			{
				Config: testAccRequestTemplateConfig(name, false),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionUpdate),
					},
				},
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "custom_fields.0.required", "false"),
					resource.TestCheckResourceAttrSet(resourceName, "custom_fields.2.name"),
				),
			},
		},
	})
}
