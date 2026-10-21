package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
)

func testAccTagConfig(name string, key string, value string) string {
	return testAccProviderConfig() + fmt.Sprintf(`
resource "opal_tag" "%s" {
	key   = "%s"
	value = "%s"
}
`, name, key, value)
}

func TestAccTag_CRUD(t *testing.T) {
	t.Parallel()
	name := testAccName("tag")
	resourceName := "opal_tag." + name

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		CheckDestroy:             testAccCheckDestroyedByGET("opal_tag", "id", "/tag/%s"),
		Steps: []resource.TestStep{
			{
				Config: testAccTagConfig(name, name, "v1"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "key", name),
					resource.TestCheckResourceAttr(resourceName, "value", "v1"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				// key and value both require replacement
				Config: testAccTagConfig(name, name, "v2"),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionDestroyBeforeCreate),
					},
				},
				Check: resource.TestCheckResourceAttr(resourceName, "value", "v2"),
			},
		},
	})
}
