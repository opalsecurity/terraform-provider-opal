package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
)

func testAccCustomResourceConfig(name string, description string, requireMFAToApprove bool) string {
	return testAccProviderConfig() + fmt.Sprintf(`
resource "opal_resource" "%s" {
	name                   = "%s"
	description            = "%s"
	app_id                 = "%s"
	resource_type          = "CUSTOM"
	admin_owner_id         = "%s"
	visibility             = "GLOBAL"
	require_mfa_to_approve = %t
	%s
}
`, name, name, description, knownCustomAppID, knownCustomAppAdminOwnerID, requireMFAToApprove,
		GenerateRequestConfigurations([]RequestConfigurationConfig{
			{
				IsRequestable: true,
				ReviewerStages: []ReviewerStageConfig{
					{
						OwnerIDs: []string{knownCustomAppAdminOwnerID},
						Operator: "AND",
					},
				},
			},
		}),
	)
}

func TestAccResource_CustomApp_CRUD(t *testing.T) {
	t.Parallel()
	name := testAccName("resource")
	resourceName := "opal_resource." + name

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProviderFactories,
		PreCheck:                 func() { testAccPreCheckCustomApp(t) },
		CheckDestroy:             testAccCheckDestroyedByGET("opal_resource", "id", "/resources/%s"),
		Steps: []resource.TestStep{
			{
				Config: testAccCustomResourceConfig(name, "first", false),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "app_id", knownCustomAppID),
					resource.TestCheckResourceAttr(resourceName, "resource_type", "CUSTOM"),
					resource.TestCheckResourceAttr(resourceName, "admin_owner_id", knownCustomAppAdminOwnerID),
					resource.TestCheckResourceAttr(resourceName, "require_mfa_to_approve", "false"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccCustomResourceConfig(name, "second", true),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionUpdate),
					},
				},
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "description", "second"),
					resource.TestCheckResourceAttr(resourceName, "require_mfa_to_approve", "true"),
				),
			},
		},
	})
}
