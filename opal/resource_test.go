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

var knownCustomAppID = os.Getenv("OPAL_TEST_KNOWN_CUSTOM_APP_ID")

func TestAccExampleResource_CRUD(t *testing.T) {
	baseName := "tf_acc_test_resource_" + acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	resourceName := "opal_resource." + baseName

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckResourceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceResource(baseName, baseName, ""),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", baseName),  // Verify that the name was set.
					resource.TestCheckResourceAttr(resourceName, "description", ""), // Verify that optional works.
				),
			},
		},
	})
}

func testAccResourceResource(tfName, name, additional string) string {
	return fmt.Sprintf(`
resource "opal_resource" "%s" {
	name = "%s"
	app_id = "%s"
	resource_type = "CUSTOM"

	%s
}
`, tfName, name, knownCustomAppID, additional)
}

func testAccCheckResourceDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*opal.APIClient)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "opal_resource" {
			continue
		}
		resource, _, err := client.ResourcesApi.GetResource(context.Background(), rs.Primary.ID).Execute()
		if err == nil {
			if resource != nil {
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
