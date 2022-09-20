package opal

import (
	"context"
	"fmt"
	"os"
	"regexp"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/opalsecurity/opal-go"
)

var knownCustomAppID = os.Getenv("OPAL_TEST_KNOWN_CUSTOM_APP_ID")
var knownCustomAppAdminOwnerID = os.Getenv("OPAL_TEST_KNOWN_CUSTOM_APP_ADMIN_OWNER_ID")
var knownRequestTemplateID = os.Getenv("OPAL_TEST_KNOWN_REQUEST_TEMPLATE_ID")

func TestAccResource_Import(t *testing.T) {
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

func TestAccResource_CRUD(t *testing.T) {
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
					resource.TestCheckResourceAttr(resourceName, "name", baseName),                             // Verify that the name was set.
					resource.TestCheckResourceAttr(resourceName, "description", ""),                            // Verify that optional works.
					resource.TestCheckResourceAttr(resourceName, "max_duration", "0"),                          // Verify that optional works.
					resource.TestCheckResourceAttr(resourceName, "admin_owner_id", knownCustomAppAdminOwnerID), // Verify that optional works.
					resource.TestCheckResourceAttr(resourceName, "require_manager_approval", "false"),          // Verify that optional works.
					resource.TestCheckResourceAttr(resourceName, "require_support_ticket", "false"),            // Verify that optional works.
					resource.TestCheckResourceAttr(resourceName, "require_mfa_to_approve", "false"),            // Verify that optional works.
					resource.TestCheckResourceAttr(resourceName, "auto_approval", "false"),                     // Verify that optional works.
					resource.TestCheckResourceAttr(resourceName, "visibility", "GLOBAL"),                       // Verify that optional works.
					resource.TestCheckResourceAttr(resourceName, "reviewer.0.id", knownCustomAppAdminOwnerID),  // Verify that optional works.
				),
			},
			{
				Config: testAccResourceResource(baseName, baseName+"_changed", `
description = "test desc"
max_duration = 60
require_manager_approval = true
require_support_ticket = true
require_mfa_to_approve = true
`),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", baseName+"_changed"),        // Verify that updating the name works.
					resource.TestCheckResourceAttr(resourceName, "description", "test desc"),         // Verify that updating the description works.
					resource.TestCheckResourceAttr(resourceName, "max_duration", "60"),               // Verify that updating works.
					resource.TestCheckResourceAttr(resourceName, "require_manager_approval", "true"), // Verify that updating works.
					resource.TestCheckResourceAttr(resourceName, "require_support_ticket", "true"),   // Verify that updating works.
					resource.TestCheckResourceAttr(resourceName, "require_mfa_to_approve", "true"),   // Verify that updating works.

				),
			},
		},
	})
}

// TestAccResource_Visibility tests that setting visibility works.
func TestAccResource_Visibility(t *testing.T) {
	baseName := "tf_acc_test_resource_" + acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	resourceName := "opal_resource." + baseName

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckResourceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceResource(baseName, baseName, `visibility = "LIMITED"`),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "visibility", "LIMITED"),
				),
			},
			{
				Config: testAccResourceResource(baseName, baseName, `visibility = "GLOBAL"`),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "visibility", "GLOBAL"),
				),
			},
			{
				Config:      testAccResourceResource(baseName, baseName, `visibility_group { id = "whatever" }`),
				ExpectError: regexp.MustCompile("cannot be specified"),
			},
		},
	})
}

// TestAccResource_SetOnCreate tests that setting attributes on creation
// works.
func TestAccResource_SetOnCreate(t *testing.T) {
	baseName := "tf_acc_test_resource_" + acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	resourceName := "opal_resource." + baseName

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckResourceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceResource(baseName, baseName, fmt.Sprintf(`
description = "test desc"
require_manager_approval = true
require_support_ticket = true
max_duration = 30
request_template_id = "%s"
`, knownRequestTemplateID)),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", baseName),
					resource.TestCheckResourceAttr(resourceName, "description", "test desc"),
					resource.TestCheckResourceAttr(resourceName, "auto_approval", "false"),
				),
			},
		},
	})
}

// TestAccResource_SetOnCreate tests that setting auto approve on creation works.
func TestAccResource_SetOnCreate_AutoApproval(t *testing.T) {
	baseName := "tf_acc_test_resource_" + acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	resourceName := "opal_resource." + baseName

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckResourceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceResource(baseName, baseName, `
auto_approval = true
`),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "auto_approval", "true"),
				),
			},
		},
	})
}

// XXX: Test metadata / Remote ID

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

// TestAccResource_SetOnCreate_WithOwner tests that setting an admin_owner_id works.
func TestAccResource_SetOnCreate_WithOwner(t *testing.T) {
	resourceBaseName := "tf_acc_test_resource_" + acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	ownerBaseName := "tf_acc_test_owner_" + acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	resourceResourceName := "opal_resource." + resourceBaseName
	ownerResourceName := "opal_owner." + ownerBaseName

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: combineCheck(testAccCheckResourceDestroy, testAccCheckOwnerDestroy),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceResourceWithOwner(ownerBaseName, resourceBaseName, fmt.Sprintf(`admin_owner_id = "${%s.id}"`, ownerResourceName)),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceResourceName, "name", resourceBaseName),
					resource.TestCheckResourceAttrPair(resourceResourceName, "admin_owner_id", ownerResourceName, "id"),
					resource.TestCheckResourceAttrPair(resourceResourceName, "reviewer.0.id", ownerResourceName, "id"),
				),
			},
			{
				// Change the owner and verify that the owner is changed.
				Config: testAccResourceResourceWithOwner(ownerBaseName, resourceBaseName, fmt.Sprintf(`admin_owner_id = "%s"`, knownCustomAppAdminOwnerID)),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceResourceName, "name", resourceBaseName),
					resource.TestCheckResourceAttr(resourceResourceName, "admin_owner_id", knownCustomAppAdminOwnerID),
					resource.TestCheckResourceAttrPair(resourceResourceName, "reviewer.0.id", ownerResourceName, "id"),
				),
			},
			{
				// Change the reviewer as well and verify it's changed.
				Config: testAccResourceResourceWithOwner(ownerBaseName, resourceBaseName, fmt.Sprintf(`admin_owner_id = "%s"
reviewer {
	id = "%s"
}`, knownCustomAppAdminOwnerID, knownCustomAppAdminOwnerID)),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceResourceName, "name", resourceBaseName),
					resource.TestCheckResourceAttr(resourceResourceName, "admin_owner_id", knownCustomAppAdminOwnerID),
					resource.TestCheckResourceAttr(resourceResourceName, "reviewer.0.id", knownCustomAppAdminOwnerID),
				),
			},
		},
	})
}

func testAccResourceResourceWithOwner(ownerName, resourceName, additional string) string {
	return fmt.Sprintf(`
resource "opal_owner" "%s" {
	name = "%s"

	user {
		id = "%s"
	}
}

resource "opal_resource" "%s" {
	name = "%s"
	resource_type = "CUSTOM"
	app_id = "%s"
	%s
}
`, ownerName, ownerName, knownUserID1, resourceName, resourceName, knownCustomAppID, additional)
}

func combineCheck(fns ...resource.TestCheckFunc) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		for _, fn := range fns {
			if err := fn(s); err != nil {
				return err
			}
		}
		return nil
	}
}
