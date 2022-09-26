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

var knownOpalAppID = os.Getenv("OPAL_TEST_KNOWN_OPAL_APP_ID")
var knownOpalAppAdminOwnerID = os.Getenv("OPAL_TEST_KNOWN_OPAL_APP_ADMIN_OWNER_ID")

func TestAccGroup_Import(t *testing.T) {
	baseName := "tf_acc_group_test_" + acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	resourceName := "opal_group." + baseName

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccGroupResource(baseName, baseName, ""),
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

func TestAccGroup_CRUD(t *testing.T) {
	baseName := "tf_acc_group_test_" + acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	resourceName := "opal_group." + baseName

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccGroupResource(baseName, baseName, ""),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", baseName),                           // Verify that the name was set.
					resource.TestCheckResourceAttr(resourceName, "description", ""),                          // Verify that optional works.
					resource.TestCheckResourceAttr(resourceName, "max_duration", "0"),                        // Verify that optional works.
					resource.TestCheckResourceAttr(resourceName, "admin_owner_id", knownOpalAppAdminOwnerID), // Verify that optional works.
					resource.TestCheckResourceAttr(resourceName, "require_manager_approval", "false"),        // Verify that optional works.
					resource.TestCheckResourceAttr(resourceName, "require_support_ticket", "false"),          // Verify that optional works.
					resource.TestCheckResourceAttr(resourceName, "require_mfa_to_approve", "false"),          // Verify that optional works.
					resource.TestCheckResourceAttr(resourceName, "auto_approval", "false"),                   // Verify that optional works.
					resource.TestCheckResourceAttr(resourceName, "visibility", "GLOBAL"),                     // Verify that optional works.
					resource.TestCheckResourceAttr(resourceName, "reviewer.0.id", knownOpalAppAdminOwnerID),  // Verify that optional works.
				),
			},
			{
				Config: testAccGroupResource(baseName, baseName+"_changed", `
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
func TestAccGroup_Visibility(t *testing.T) {
	limitedGroupBaseName := "tf_acc_test_group_" + acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	limitedGroupResourceName := "opal_group." + limitedGroupBaseName
	groupBaseName := "tf_acc_test_group_" + acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	groupResourceName := "opal_group." + groupBaseName

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceGroupWithGroup(limitedGroupBaseName, groupBaseName, `visibility = "LIMITED"`),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(limitedGroupResourceName, "visibility", "LIMITED"),
				),
			},
			{
				Config: testAccResourceGroupWithGroup(limitedGroupBaseName, groupBaseName, `visibility = "GLOBAL"`),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(limitedGroupResourceName, "visibility", "GLOBAL"),
				),
			},
			{
				Config: testAccResourceGroupWithGroup(limitedGroupBaseName, groupBaseName, fmt.Sprintf(`
visibility = "LIMITED"
visibility_group { id = "${%s.id}" }
`, groupResourceName)),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(limitedGroupResourceName, "visibility", "LIMITED"),
					resource.TestCheckResourceAttrPair(limitedGroupResourceName, "visibility_group.0.id", groupResourceName, "id"),
				),
			},
			{
				Config:      testAccResourceGroupWithGroup(limitedGroupBaseName, groupBaseName, `visibility_group { id = "whatever" }`),
				ExpectError: regexp.MustCompile("cannot be specified"),
			},
		},
	})
}

func testAccResourceGroupWithGroup(resourceName, groupName, additional string) string {
	return fmt.Sprintf(`
resource "opal_group" "%s" {
	name = "%s"
	app_id = "%s"
	group_type = "OPAL_GROUP"

	%s
}

resource "opal_group" "%s" {
	name = "%s"
	app_id = "%s"
	group_type = "OPAL_GROUP"
}
`, resourceName, resourceName, knownOpalAppID, additional, groupName, groupName, knownOpalAppID)
}

// TestAccGroup_SetOnCreate tests that setting attributes on creation
// works.
func TestAccGroup_SetOnCreate(t *testing.T) {
	baseName := "tf_acc_group_test_" + acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	resourceName := "opal_group." + baseName

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccGroupResource(baseName, baseName, fmt.Sprintf(`
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

// TestAccGroup_SetOnCreate tests that setting auto approve on creation works.
func TestAccGroup_SetOnCreate_AutoApproval(t *testing.T) {
	baseName := "tf_acc_group_test_" + acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	resourceName := "opal_group." + baseName

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccGroupResource(baseName, baseName, `
auto_approval = true
`),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "auto_approval", "true"),
				),
			},
		},
	})
}

var knownGithubAppGroupMetadata = os.Getenv("OPAL_TEST_KNOWN_GITHUB_APP_GROUP_METADATA")
var knownGithubAppGroupRemoteID = os.Getenv("OPAL_TEST_KNOWN_GITHUB_APP_GROUP_REMOTE_ID")

// TestAccGroup_Remote tests creating a resource with a remote system.
func TestAccGroup_Remote(t *testing.T) {
	baseName := "tf_acc_group_test_" + acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	resourceName := "opal_group." + baseName

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`resource "opal_group" "%s" {
	name = "%s"
	app_id = "%s"
	group_type = "GIT_HUB_TEAM"
	metadata = jsonencode(%s)
	remote_group_id = "%s"
}
`, baseName, baseName, knownGithubAppID, knownGithubAppGroupMetadata, knownGithubAppGroupRemoteID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", baseName),
				),
			},
		},
	})
}

func testAccGroupResource(tfName, name, additional string) string {
	return fmt.Sprintf(`
resource "opal_group" "%s" {
	name = "%s"
	group_type = "OPAL_GROUP"
	app_id = "%s"

	%s
}
`, tfName, name, knownOpalAppID, additional)
}

func testAccCheckGroupDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*opal.APIClient)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "opal_group" {
			continue
		}
		groups, _, err := client.GroupsApi.GetGroups(context.Background()).GroupIds([]string{rs.Primary.ID}).Execute()
		if err == nil {
			if len(groups.Results) > 0 {
				return fmt.Errorf("Opal group still exists: %s", rs.Primary.ID)
			}
			return nil
		}
		if !strings.Contains(err.Error(), "404 Not Found") {
			return err
		}
	}

	return nil
}