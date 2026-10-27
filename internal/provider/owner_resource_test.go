package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
)

func testAccOwnerConfig(name string, description string, userIDs []string, escalationPeriod int) string {
	return testAccProviderConfig() + fmt.Sprintf(`
resource "opal_owner" "%s" {
	name                             = "%s"
	description                      = "%s"
	user_ids                         = %s
	access_request_escalation_period = %d
}
`, name, name, description, generateQuotedArrayString(userIDs), escalationPeriod)
}

func TestAccOwner_CRUD(t *testing.T) {
	t.Parallel()
	name := testAccName("owner")
	resourceName := "opal_owner." + name

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		CheckDestroy:             testAccCheckDestroyedByGET("opal_owner", "id", "/owners/%s"),
		Steps: []resource.TestStep{
			{
				Config: testAccOwnerConfig(name, "first", []string{knownUserID1}, 0),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "first"),
					resource.TestCheckResourceAttr(resourceName, "user_ids.#", "1"),
					resource.TestCheckTypeSetElemAttr(resourceName, "user_ids.*", knownUserID1),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				// Owner reads don't populate user_ids, so an import can't recover them
				ImportStateVerifyIgnore: []string{"user_ids"},
			},
			{
				Config: testAccOwnerConfig(name, "second", []string{knownUserID1, knownUserID2}, 60),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionUpdate),
					},
				},
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "description", "second"),
					resource.TestCheckResourceAttr(resourceName, "access_request_escalation_period", "60"),
					resource.TestCheckResourceAttr(resourceName, "user_ids.#", "2"),
					resource.TestCheckTypeSetElemAttr(resourceName, "user_ids.*", knownUserID2),
				),
			},
			{
				// removing a user should update in place too
				Config: testAccOwnerConfig(name, "second", []string{knownUserID2}, 60),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionUpdate),
					},
				},
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "user_ids.#", "1"),
					resource.TestCheckTypeSetElemAttr(resourceName, "user_ids.*", knownUserID2),
				),
			},
		},
	})
}
