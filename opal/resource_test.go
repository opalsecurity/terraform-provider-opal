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
var knownCustomAppAdminOwnerID = os.Getenv("OPAL_TEST_KNOWN_CUSTOM_APP_ADMIN_OWNER_ID")
var knownRequestTemplateID = os.Getenv("OPAL_TEST_KNOWN_REQUEST_TEMPLATE_ID")

// Commenting this out for now since while supporting the deprecated request configuration
// fields, is_requestable has a default value and gets populated, which causes a diff from
// parsing just the request_configuration fields. Will add this back once we remove support
// for the deprecated fields.
// func TestAccResource_Import(t *testing.T) {
// 	baseName := "tf_acc_test_resource_" + acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
// 	resourceName := "opal_resource." + baseName

// 	resource.Test(t, resource.TestCase{
// 		PreCheck:  func() { testAccPreCheck(t) },
// 		Providers: testAccProviders,
// 		Steps: []resource.TestStep{
// 			{
// 				Config: testAccResourceResourceWithRequestConfigAndReviewers(baseName, baseName, "", ""),
// 				Check: resource.ComposeTestCheckFunc(
// 					resource.TestCheckResourceAttr(resourceName, "name", baseName), // Verify that the name was set.
// 				),
// 			},
// 			{
// 				ResourceName:      resourceName,
// 				ImportState:       true,
// 				ImportStateVerify: true,
// 			},
// 		},
// 	})
// }

func TestAccResource_CRUD(t *testing.T) {
	baseName := "tf_acc_test_resource_" + acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	resourceName := "opal_resource." + baseName

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckResourceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceResourceWithRequestConfigAndReviewers(baseName, baseName, "", ""),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", baseName),                                                                     // Verify that the name was set.
					resource.TestCheckResourceAttr(resourceName, "description", ""),                                                                    // Verify that optional works.
					resource.TestCheckResourceAttr(resourceName, "request_configuration.0.max_duration", "-1"),                                         // Verify that optional works.
					resource.TestCheckResourceAttr(resourceName, "admin_owner_id", knownCustomAppAdminOwnerID),                                         // Verify that optional works.
					resource.TestCheckResourceAttr(resourceName, "request_configuration.0.require_support_ticket", "false"),                            // Verify that optional works.
					resource.TestCheckResourceAttr(resourceName, "require_mfa_to_approve", "false"),                                                    // Verify that optional works.
					resource.TestCheckResourceAttr(resourceName, "request_configuration.0.require_mfa_to_request", "false"),                            // Verify that optional works.
					resource.TestCheckResourceAttr(resourceName, "require_mfa_to_connect", "false"),                                                    // Verify that optional works.
					resource.TestCheckResourceAttr(resourceName, "request_configuration.0.is_requestable", "true"),                                     // Verify that optional works.
					resource.TestCheckResourceAttr(resourceName, "request_configuration.0.auto_approval", "false"),                                     // Verify that optional works.
					resource.TestCheckResourceAttr(resourceName, "visibility", "GLOBAL"),                                                               // Verify that optional works.
					resource.TestCheckResourceAttr(resourceName, "request_configuration.0.reviewer_stage.0.reviewer.0.id", knownCustomAppAdminOwnerID), // Verify that optional works.
				),
			},
			{
				Config: testAccResourceResourceWithRequestConfigAndReviewers(baseName, baseName+"_changed", `
max_duration = 60
require_support_ticket = true
`, `
description = "description"
`),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", baseName+"_changed"),                              // Verify that updating the name works.
					resource.TestCheckResourceAttr(resourceName, "request_configuration.0.max_duration", "60"),             // Verify that updating works.
					resource.TestCheckResourceAttr(resourceName, "request_configuration.0.require_support_ticket", "true"), // Verify that updating works.
					resource.TestCheckResourceAttr(resourceName, "request_configuration.0.is_requestable", "true"),         // Verify that updating works.
					resource.TestCheckResourceAttr(resourceName, "description", "description"),                             // Verify that updating works.
				),
			},
		},
	})
}

// TestAccResource_Visibility tests that setting visibility works.
func TestAccResource_Visibility(t *testing.T) {
	resourceBaseName := "tf_acc_test_resource_" + acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	resourceResourceName := "opal_resource." + resourceBaseName
	groupBaseName := "tf_acc_test_group_" + acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	groupResourceName := "opal_group." + groupBaseName

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckResourceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceResourceWithGroup(resourceBaseName, groupBaseName, `visibility = "LIMITED"`),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceResourceName, "visibility", "LIMITED"),
				),
			},
			{
				Config: testAccResourceResourceWithGroup(resourceBaseName, groupBaseName, `visibility = "GLOBAL"`),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceResourceName, "visibility", "GLOBAL"),
				),
			},
			{
				Config: testAccResourceResourceWithGroup(resourceBaseName, groupBaseName, fmt.Sprintf(`
visibility = "LIMITED"
visibility_group { id = "${%s.id}" }
`, groupResourceName)),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceResourceName, "visibility", "LIMITED"),
					resource.TestCheckResourceAttrPair(resourceResourceName, "visibility_group.0.id", groupResourceName, "id"),
				),
			},
		},
	})
}

func testAccResourceResourceWithGroup(resourceName, groupName, additional string) string {
	return fmt.Sprintf(`
resource "opal_resource" "%s" {
	name = "%s"
	app_id = "%s"
	resource_type = "CUSTOM"
	admin_owner_id = "%s"
	
	request_configuration {
		reviewer_stage {
			reviewer {
				id = "%s"
			}
		}		
	}
	%s
}

resource "opal_group" "%s" {
	name = "%s"
	app_id = "%s"
	group_type = "OPAL_GROUP"
	admin_owner_id = "%s"
	
	request_configuration {
	reviewer_stage {
		reviewer {
			id = "%s"
		}
	}		
	}
}
`, resourceName, resourceName, knownCustomAppID, knownCustomAppAdminOwnerID, knownCustomAppAdminOwnerID, additional, groupName, groupName, knownOpalAppID, knownOpalAppAdminOwnerID, knownOpalAppAdminOwnerID)
}

func TestAccResource_Reviewer(t *testing.T) {
	baseName := "tf_acc_resource_test_" + acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	resourceName := "opal_resource." + baseName

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceResource(baseName, baseName, testReviewerStage("AND", false, knownOpalAppAdminOwnerID)),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", baseName),
					resource.TestCheckResourceAttr(resourceName, "request_configuration.0.reviewer_stage.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "request_configuration.0.reviewer_stage.0.reviewer.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "request_configuration.0.reviewer_stage.0.operator", "AND"),
					resource.TestCheckResourceAttr(resourceName, "request_configuration.0.reviewer_stage.0.require_manager_approval", "false"),
					resource.TestCheckTypeSetElemNestedAttrs(resourceName, "request_configuration.0.reviewer_stage.0.reviewer.*", map[string]string{"id": knownOpalAppAdminOwnerID}),
				),
			},
			{
				Config: testAccResourceResource(baseName, baseName, testReviewerStage("OR", true, knownOpalAppAdminOwnerID, knownCustomAppAdminOwnerID)),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", baseName),
					resource.TestCheckResourceAttr(resourceName, "request_configuration.0.reviewer_stage.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "request_configuration.0.reviewer_stage.0.reviewer.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "request_configuration.0.reviewer_stage.0.operator", "OR"),
					resource.TestCheckResourceAttr(resourceName, "request_configuration.0.reviewer_stage.0.require_manager_approval", "true"),
					resource.TestCheckTypeSetElemNestedAttrs(resourceName, "request_configuration.0.reviewer_stage.0.reviewer.*", map[string]string{"id": knownOpalAppAdminOwnerID}),
					resource.TestCheckTypeSetElemNestedAttrs(resourceName, "request_configuration.0.reviewer_stage.0.reviewer.*", map[string]string{"id": knownCustomAppAdminOwnerID}),
				),
			},
			{
				Config: testAccResourceResource(baseName, baseName, testReviewerStage("AND", false, knownCustomAppAdminOwnerID)),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", baseName),
					resource.TestCheckResourceAttr(resourceName, "request_configuration.0.reviewer_stage.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "request_configuration.0.reviewer_stage.0.reviewer.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "request_configuration.0.reviewer_stage.0.operator", "AND"),
					resource.TestCheckResourceAttr(resourceName, "request_configuration.0.reviewer_stage.0.require_manager_approval", "false"),
					resource.TestCheckTypeSetElemNestedAttrs(resourceName, "request_configuration.0.reviewer_stage.0.reviewer.*", map[string]string{"id": knownCustomAppAdminOwnerID}),
				),
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
				Config: testAccResourceResourceWithRequestConfigAndReviewers(baseName, baseName, "", ""),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", baseName),
				),
			},
		},
	})
}

// TestAccResource_SetOnCreate tests that setting auto approve on creation works.
func TestAccResource_SetOnCreate_AutoApproval(t *testing.T) {
	t.Skip("Skip until API behavior is fixed")

	baseName := "tf_acc_test_resource_" + acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	resourceName := "opal_resource." + baseName

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckResourceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceResource(baseName, baseName, `
is_requestable = true
auto_approval = true
`),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "request_configuration.0.auto_approval", "true"),
				),
			},
		},
	})
}

var knownGithubAppID = os.Getenv("OPAL_TEST_KNOWN_GITHUB_APP_ID")
var knownGithubRepoName = os.Getenv("OPAL_TEST_KNOWN_GITHUB_APP_REPO_NAME")

// TestAccResource_Remote tests creating a resource with a remote system.
func TestAccResource_Remote(t *testing.T) {
	baseName := "tf_acc_test_resource_" + acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	resourceName := "opal_resource." + baseName

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckResourceDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "opal_resource" "%s" {
	name = "%s"
	app_id = "%s"
	admin_owner_id = "%s"
	request_configuration {
		reviewer_stage {
			reviewer {
				id = "%s"
			}
		}
	}
	resource_type = "GIT_HUB_REPO"
	remote_info {
		github_repo {
			repo_name = "%s"
		}
    }
}
`, baseName, baseName, knownGithubAppID, knownOpalAppAdminOwnerID, knownOpalAppAdminOwnerID, knownGithubRepoName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", baseName),
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
	admin_owner_id = "%s"
	%s
}
`, tfName, name, knownCustomAppID, knownCustomAppAdminOwnerID, additional)
}

func testAccResourceResourceWithRequestConfigAndReviewers(tfName, name, additionalRequestConfig, additional string) string {
	return testAccResourceResource(tfName, name, fmt.Sprintf(`
request_configuration {
	is_requestable = true
	reviewer_stage {
		reviewer {
			id = "%s"
		}
	}
	%s
}
%s
`, knownCustomAppAdminOwnerID, additionalRequestConfig, additional))
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
					resource.TestCheckResourceAttr(resourceResourceName, "request_configuration.0.reviewer_stage.0.reviewer.0.id", knownOpalAppAdminOwnerID),
				),
			},
			{
				// Change the owner and verify that the owner is changed.
				Config: testAccResourceResourceWithOwner(ownerBaseName, resourceBaseName, fmt.Sprintf(`admin_owner_id = "%s"`, knownCustomAppAdminOwnerID)),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceResourceName, "name", resourceBaseName),
					resource.TestCheckResourceAttr(resourceResourceName, "admin_owner_id", knownCustomAppAdminOwnerID),
					resource.TestCheckResourceAttr(resourceResourceName, "request_configuration.0.reviewer_stage.0.reviewer.0.id", knownOpalAppAdminOwnerID),
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

	request_configuration {
		reviewer_stage {
			reviewer {
				id = "%s"
			}
		}
	}
	%s
}
`, ownerName, ownerName, knownUserID1, resourceName, resourceName, knownCustomAppID, knownOpalAppAdminOwnerID, additional)
}

// TestAccResource_RequestConfiguration tests that setting a request configuration works.
func TestAccResource_RequestConfiguration(t *testing.T) {
	resourceBaseName := "tf_acc_test_resource_" + acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	resourceResourceName := "opal_resource." + resourceBaseName

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: combineCheck(testAccCheckResourceDestroy, testAccCheckOwnerDestroy),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceResourceWithRequestConfiguration(resourceBaseName, ""),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceResourceName, "name", resourceBaseName),
					resource.TestCheckResourceAttr(resourceResourceName, "request_configuration.0.max_duration", "60"),
					resource.TestCheckResourceAttr(resourceResourceName, "request_configuration.0.recommended_duration", "30"),
					resource.TestCheckResourceAttr(resourceResourceName, "request_configuration.0.require_support_ticket", "true"),
					resource.TestCheckResourceAttr(resourceResourceName, "request_configuration.0.require_mfa_to_request", "true"),
					resource.TestCheckResourceAttr(resourceResourceName, "request_configuration.0.is_requestable", "true"),
					resource.TestCheckResourceAttr(resourceResourceName, "request_configuration.0.auto_approval", "false"),
					resource.TestCheckResourceAttr(resourceResourceName, "request_configuration.0.reviewer_stage.0.reviewer.0.id", knownOpalAppAdminOwnerID),
				),
			},
			{
				Config: testAccResourceResourceWithRequestConfiguration(resourceBaseName, fmt.Sprintf(`
request_configuration {
	group_ids = ["%s"]
	priority = 1
	max_duration = 30
	recommended_duration = 15
	require_support_ticket = false
	require_mfa_to_request = false
	is_requestable = true
	auto_approval = true
	reviewer_stage {
		reviewer {
			id = "%s"
		}
	}
}
`, knownGroupID, knownOpalAppAdminOwnerID)),

				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceResourceName, "name", resourceBaseName),
					resource.TestCheckResourceAttr(resourceResourceName, "request_configuration.1.max_duration", "30"),
					resource.TestCheckResourceAttr(resourceResourceName, "request_configuration.1.recommended_duration", "15"),
					resource.TestCheckResourceAttr(resourceResourceName, "request_configuration.1.require_support_ticket", "false"),
					resource.TestCheckResourceAttr(resourceResourceName, "request_configuration.1.require_mfa_to_request", "false"),
					resource.TestCheckResourceAttr(resourceResourceName, "request_configuration.1.is_requestable", "true"),
					resource.TestCheckResourceAttr(resourceResourceName, "request_configuration.1.auto_approval", "true"),
					resource.TestCheckResourceAttr(resourceResourceName, "request_configuration.1.reviewer_stage.0.reviewer.0.id", knownOpalAppAdminOwnerID),
				),
			},
		},
	})
}

func testAccResourceResourceWithRequestConfiguration(resourceName, additional string) string {
	return fmt.Sprintf(`
resource "opal_resource" "%s" {
	name = "%s"
	resource_type = "CUSTOM"
	app_id = "%s"
	admin_owner_id = "%s"
	request_configuration {
		max_duration = 60
		recommended_duration = 30
		require_support_ticket = true
		require_mfa_to_request = true
 		is_requestable = true
		auto_approval = false
		reviewer_stage {
			reviewer {
				id = "%s"
			}
		}
	}
	%s
}
`, resourceName, resourceName, knownCustomAppID, knownOpalAppAdminOwnerID, knownOpalAppAdminOwnerID, additional)
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
