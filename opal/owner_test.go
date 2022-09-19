package opal

import (
	"context"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/opalsecurity/opal-go"
)

var knownUserID = os.Getenv("OPAL_TEST_KNOWN_USER_ID")

func TestAccExampleOwner_basic(t *testing.T) {
	baseName := "owner_" + acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	resourceName := "opal_owner." + baseName

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckOwnerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccOwnerResource(baseName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", baseName),
					resource.TestCheckResourceAttr(resourceName, "user.0.id", knownUserID),
					resource.TestCheckResourceAttr(resourceName, "access_request_escalation_period", "0"),
				),
			},
		},
	})
}

func testAccOwnerResource(name string) string {
	return fmt.Sprintf(`
resource "opal_owner" "%s" {
	name = "%s"
	user {
		id = "%s"
	}
}
`, name, name, knownUserID)
}

func testAccCheckOwnerDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*opal.APIClient)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "opal_owner" {
			continue
		}
		owner, _, err := client.OwnersApi.GetOwner(context.Background(), rs.Primary.ID).Execute()
		if err == nil {
			if owner != nil {
				return fmt.Errorf("Opal resource still exists: %s", rs.Primary.ID)
			}
			return nil
		}
		if !strings.Contains(err.Error(), "404 Not Found") {
			return err
		}
	}

	return nil
}
