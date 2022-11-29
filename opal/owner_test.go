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

var knownUserID1 = os.Getenv("OPAL_TEST_KNOWN_USER_ID_1")
var knownUserID2 = os.Getenv("OPAL_TEST_KNOWN_USER_ID_2")

func TestAccOwner_Import(t *testing.T) {
	baseName := "tf_acc_test_owner_" + acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	resourceName := "opal_owner." + baseName

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckOwnerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccOwnerResource(baseName, baseName, ""),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", baseName), // Verify that the name was set.
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccOwner_CRUD(t *testing.T) {
	baseName := "tf_acc_test_owner_" + acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	resourceName := "opal_owner." + baseName

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckOwnerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccOwnerResource(baseName, baseName, ""),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", baseName),                        // Verify that the name was set.
					resource.TestCheckResourceAttr(resourceName, "user.0.id", knownUserID1),               // Verify that the user was set.
					resource.TestCheckResourceAttr(resourceName, "description", ""),                       // Verify that optional works.
					resource.TestCheckResourceAttr(resourceName, "access_request_escalation_period", "0"), // Verify that optional works.
				),
			},
			{
				Config: testAccOwnerResource(baseName, baseName+"_changed", `description = "some desc"
					access_request_escalation_period = 30`),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", baseName+"_changed"),              // Verify that changing the name works.
					resource.TestCheckResourceAttr(resourceName, "description", "some desc"),               // Verify that changing the description works.
					resource.TestCheckResourceAttr(resourceName, "user.0.id", knownUserID1),                // Verify that the existing user wasn't changed.
					resource.TestCheckResourceAttr(resourceName, "access_request_escalation_period", "30"), // Verify that changing the escalation period works.
				),
			},
			{
				Config: testAccOwnerResource(baseName, baseName+"_changed", fmt.Sprintf(`user { id = "%s" }`, knownUserID2)),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "user.0.id", knownUserID1), // Verify that the existing user wasn't changed.
					resource.TestCheckResourceAttr(resourceName, "user.1.id", knownUserID2), // Verify that adding a user works.
				),
			},
			{
				Config: testAccOwnerResourceNoUser(baseName, baseName+"_changed", ""),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckNoResourceAttr(resourceName, "user.0.id"),
					resource.TestCheckNoResourceAttr(resourceName, "user.1.id"),
				),
			},
		},
	})
}

func testAccOwnerResourceNoUser(tfName, name, additional string) string {
	return fmt.Sprintf(`
resource "opal_owner" "%s" {
	name = "%s"

	%s
}
`, tfName, name, additional)
}

func testAccOwnerResource(tfName, name, additional string) string {
	return fmt.Sprintf(`
resource "opal_owner" "%s" {
	name = "%s"
	user {
		id = "%s"
	}

	%s
}
`, tfName, name, knownUserID1, additional)
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
