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

var knownOpalAppID = os.Getenv("OPAL_TEST_KNOWN_OPAL_APP_ID")
var knownOpalAppAdminOwnerID = os.Getenv("OPAL_TEST_KNOWN_OPAL_APP_ADMIN_OWNER_ID")
var knownGithubRepoResourceID = os.Getenv("OPAL_TEST_KNOWN_GITHUB_TEST_REPO_2_RESOURCE_ID")

func TestAccGroup_Import(t *testing.T) {
	baseName := "tf_acc_group_test_" + acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	resourceName := "opal_group." + baseName

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccGroupResourceWithReviewer(baseName, baseName, ""),
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
				Config: testAccGroupResourceWithReviewer(baseName, baseName, ""),
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
				Config: testAccGroupResourceWithReviewer(baseName, baseName+"_changed", `
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
		},
	})
}

func testAccResourceGroupWithGroup(resourceName, groupName, additional string) string {
	return fmt.Sprintf(`
resource "opal_group" "%s" {
	name = "%s"
	app_id = "%s"
	group_type = "OPAL_GROUP"
	admin_owner_id = "%s"

	reviewer {
		id = "%s"
	}

	%s
}

resource "opal_group" "%s" {
	name = "%s"
	app_id = "%s"
	group_type = "OPAL_GROUP"
	admin_owner_id = "%s"

	reviewer {
		id = "%s"
	}
}
`, resourceName, resourceName, knownOpalAppID, knownOpalAppAdminOwnerID, knownOpalAppAdminOwnerID, additional, groupName, groupName, knownOpalAppID, knownOpalAppAdminOwnerID, knownOpalAppAdminOwnerID)
}

func TestAccGroup_Reviewer(t *testing.T) {
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
					resource.TestCheckResourceAttr(resourceName, "name", baseName),
					// The reviewer isn't saved since it is not marked as Computed.
					resource.TestCheckResourceAttr(resourceName, "reviewer.#", "0"),
				),
			},
			{
				Config: testAccGroupResource(baseName, baseName, fmt.Sprintf(`reviewer { id = "%s" }`, knownOpalAppAdminOwnerID)),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", baseName),
					resource.TestCheckResourceAttr(resourceName, "reviewer.#", "1"),
					resource.TestCheckTypeSetElemNestedAttrs(resourceName, "reviewer.*", map[string]string{"id": knownOpalAppAdminOwnerID}),
				),
			},
			{
				Config: testAccGroupResource(baseName, baseName, fmt.Sprintf(`reviewer { id = "%s" }
reviewer { id = "%s" }`, knownOpalAppAdminOwnerID, knownCustomAppAdminOwnerID)),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", baseName),
					resource.TestCheckResourceAttr(resourceName, "reviewer.#", "2"),
					resource.TestCheckTypeSetElemNestedAttrs(resourceName, "reviewer.*", map[string]string{"id": knownOpalAppAdminOwnerID}),
					resource.TestCheckTypeSetElemNestedAttrs(resourceName, "reviewer.*", map[string]string{"id": knownCustomAppAdminOwnerID}),
				),
			},
			{
				Config: testAccGroupResource(baseName, baseName, fmt.Sprintf(`reviewer { id = "%s" }`, knownCustomAppAdminOwnerID)),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", baseName),
					resource.TestCheckResourceAttr(resourceName, "reviewer.#", "1"),
					resource.TestCheckTypeSetElemNestedAttrs(resourceName, "reviewer.*", map[string]string{"id": knownCustomAppAdminOwnerID}),
				),
			},
			{
				Config: testAccGroupResource(baseName, baseName, ""),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", baseName),
					resource.TestCheckResourceAttr(resourceName, "reviewer.#", "0"),
				),
			},
			{
				Config: testAccGroupResource(baseName, baseName, fmt.Sprintf(`reviewer { id = "%s" }`, knownOpalAppAdminOwnerID)),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", baseName),
					resource.TestCheckResourceAttr(resourceName, "reviewer.#", "1"),
					resource.TestCheckTypeSetElemNestedAttrs(resourceName, "reviewer.*", map[string]string{"id": knownOpalAppAdminOwnerID}),
				),
			},
			{
				Config: testAccGroupResource(baseName, baseName, ""),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", baseName),
					resource.TestCheckResourceAttr(resourceName, "reviewer.#", "1"),
					// This is a bit weird. If we delete the reviewer list and the only reviewer is the app owner, then
					// we don't make the diff and the reviewer stays in the state. This is fine but a bit inconsistent.
					resource.TestCheckTypeSetElemNestedAttrs(resourceName, "reviewer.*", map[string]string{"id": knownOpalAppAdminOwnerID}),
				),
			},
		},
	})
}

func TestAccGroup_Resource(t *testing.T) {
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
					resource.TestCheckResourceAttr(resourceName, "name", baseName),
					resource.TestCheckResourceAttr(resourceName, "resource.#", "0"),
				),
			},
			{
				Config: testAccGroupResource(baseName, baseName, testAccGroupResourceWithAccessLevel(knownGithubRepoResourceID, "pull")),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", baseName),
					resource.TestCheckResourceAttr(resourceName, "resource.#", "1"),
					resource.TestCheckTypeSetElemNestedAttrs(resourceName, "resource.*", map[string]string{"id": knownGithubRepoResourceID}),
				),
			},
			{
				Config: testAccGroupResource(baseName, baseName, ""),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", baseName),
					resource.TestCheckResourceAttr(resourceName, "resource.#", "0"),
				),
			},
		},
	})
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
				Config: testAccGroupResourceWithReviewer(baseName, baseName, fmt.Sprintf(`
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

var knownGithubTeamName = os.Getenv("OPAL_TEST_KNOWN_GITHUB_APP_TEAM_SLUG")

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
	admin_owner_id = "%s"
	reviewer {
		id = "%s"
	}
	group_type = "GIT_HUB_TEAM"
	remote_info {
		github_team {
			team_slug = "%s"
		}
	}
}
`, baseName, baseName, knownGithubAppID, knownOpalAppAdminOwnerID, knownOpalAppAdminOwnerID, knownGithubTeamName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", baseName),
				),
			},
		},
	})
}

func testAccGroupResourceWithAccessLevel(resourceID, accessLevelRemoteID string) string {
	return fmt.Sprintf(`
resource {
	id = "%s"
	access_level_remote_id = "%s"
}
`, resourceID, accessLevelRemoteID)

}

func testAccGroupResource(tfName, name, additional string) string {
	return fmt.Sprintf(`
resource "opal_group" "%s" {
	name = "%s"
	group_type = "OPAL_GROUP"
	app_id = "%s"
	admin_owner_id = "%s"

	%s
}
`, tfName, name, knownOpalAppID, knownOpalAppAdminOwnerID, additional)
}

func testAccGroupResourceWithReviewer(tfName, name, additional string) string {
	return testAccGroupResource(tfName, name, fmt.Sprintf(`
	reviewer {
		id = "%s"
	}

	%s
`, knownOpalAppAdminOwnerID, additional))
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
