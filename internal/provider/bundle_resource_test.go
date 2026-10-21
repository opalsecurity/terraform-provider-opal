package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
)

func testAccBundleConfig(name string, description string) string {
	return testAccProviderConfig() + fmt.Sprintf(`
resource "opal_bundle" "%s" {
	name           = "%s"
	description    = "%s"
	admin_owner_id = "%s"
	visibility     = "GLOBAL"
}
`, name, name, description, knownOpalAppAdminOwnerID)
}

func TestAccBundle_CRUD(t *testing.T) {
	t.Parallel()
	name := testAccName("bundle")
	resourceName := "opal_bundle." + name

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		CheckDestroy:             testAccCheckDestroyedByGET("opal_bundle", "bundle_id", "/bundles/%s"),
		Steps: []resource.TestStep{
			{
				Config: testAccBundleConfig(name, "first"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "bundle_id"),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "first"),
					resource.TestCheckResourceAttr(resourceName, "admin_owner_id", knownOpalAppAdminOwnerID),
					resource.TestCheckResourceAttr(resourceName, "visibility", "GLOBAL"),
				),
			},
			{
				ResourceName:                         resourceName,
				ImportState:                          true,
				ImportStateVerify:                    true,
				ImportStateIdFunc:                    testAccImportStateIDFromAttr(resourceName, "bundle_id"),
				ImportStateVerifyIdentifierAttribute: "bundle_id",
			},
			{
				Config: testAccBundleConfig(name, "second"),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionUpdate),
					},
				},
				Check: resource.TestCheckResourceAttr(resourceName, "description", "second"),
			},
		},
	})
}
