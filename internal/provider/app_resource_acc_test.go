package provider

import (
	"fmt"
	"os"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
)

// Non-writable fixture uses OPAL_TEST_KNOWN_OPAL_APP_ID (already wired in CI).
// OPAL is not CUSTOM/CUSTOM_CONNECTOR, so GET /apps/{id}/visibility returns 400 —
// same class of failure as OKTA_DIRECTORY in #233.
func testAccAppManagedPreCheck(t *testing.T) {
	t.Helper()
	if os.Getenv("OPAL_TEST_TOKEN") == "" {
		t.Fatal("OPAL_TEST_TOKEN must be set for acceptance tests")
	}
	if os.Getenv("OPAL_TEST_BASE_URL") == "" {
		t.Fatal("OPAL_TEST_BASE_URL must be set for acceptance tests")
	}
	if os.Getenv("OPAL_TEST_KNOWN_CUSTOM_APP_ADMIN_OWNER_ID") == "" {
		t.Fatal("OPAL_TEST_KNOWN_CUSTOM_APP_ADMIN_OWNER_ID must be set for acceptance tests")
	}
}

func testAccAppNonWritablePreCheck(t *testing.T) {
	t.Helper()
	testAccAppManagedPreCheck(t)
	if os.Getenv("OPAL_TEST_KNOWN_OPAL_APP_ID") == "" {
		t.Fatal("OPAL_TEST_KNOWN_OPAL_APP_ID must be set for acceptance tests")
	}
}

func testAccNonWritableAppDataSourceConfig(id string) string {
	return testAccProviderConfig() + fmt.Sprintf(`
data "opal_app" "non_writable" {
  id = "%s"
}
`, id)
}

func testAccAppCustomConfig(name, adminOwnerID, description, visibility string) string {
	return testAccProviderConfig() + fmt.Sprintf(`
resource "opal_app" "%s" {
  name              = "%s"
  description       = "%s"
  admin_owner_id    = "%s"
  app_type          = "CUSTOM"
  visibility        = "%s"
  import_visibility = "GLOBAL"
}
`, name, name, description, adminOwnerID, visibility)
}

func testAccAppCustomWithVisibilityDataSourceConfig(name, adminOwnerID string) string {
	return testAccAppCustomConfig(
		name,
		adminOwnerID,
		"visibility ds fixture",
		"GLOBAL",
	) + fmt.Sprintf(
		`
data "opal_app_visibility" "custom" {
  id = opal_app.%s.id
}
`,
		name,
	)
}

func testAccAppVisibilityNonWritableConfig(id string) string {
	return testAccProviderConfig() + fmt.Sprintf(`
data "opal_app_visibility" "non_writable" {
  id = "%s"
}
`, id)
}

// TestAccApp_DataSourceNonWritableVisibility confirms data.opal_app reads a
// non-CUSTOM app (including visibility) without calling the visibility
// sub-endpoint that would 400.
func TestAccApp_DataSourceNonWritableVisibility(t *testing.T) {
	t.Parallel()
	appID := os.Getenv("OPAL_TEST_KNOWN_OPAL_APP_ID")

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProviderFactories,
		PreCheck:                 func() { testAccAppNonWritablePreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: testAccNonWritableAppDataSourceConfig(appID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.opal_app.non_writable", "id", appID),
					resource.TestCheckResourceAttr("data.opal_app.non_writable", "type", "OPAL"),
					resource.TestCheckResourceAttrSet("data.opal_app.non_writable", "visibility"),
					resource.TestCheckResourceAttrSet(
						"data.opal_app.non_writable",
						"import_visibility",
					),
				),
			},
		},
	})
}

// TestAccApp_CRUD creates a CUSTOM app, verifies import does not force
// replacement (app_type is recovered from type on Read), updates description
// in place, and destroys the app.
func TestAccApp_CRUD(t *testing.T) {
	t.Parallel()
	adminOwnerID := os.Getenv("OPAL_TEST_KNOWN_CUSTOM_APP_ADMIN_OWNER_ID")
	name := testAccName("app")
	resourceName := "opal_app." + name

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProviderFactories,
		PreCheck:                 func() { testAccAppManagedPreCheck(t) },
		CheckDestroy:             testAccCheckDestroyedByGET("opal_app", "id", "/apps/%s"),
		Steps: []resource.TestStep{
			{
				Config: testAccAppCustomConfig(name, adminOwnerID, "first", "GLOBAL"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "app_type", "CUSTOM"),
					resource.TestCheckResourceAttr(resourceName, "type", "CUSTOM"),
					resource.TestCheckResourceAttr(resourceName, "visibility", "GLOBAL"),
					resource.TestCheckResourceAttr(resourceName, "import_visibility", "GLOBAL"),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"custom_connector"},
			},
			{
				Config: testAccAppCustomConfig(name, adminOwnerID, "second", "GLOBAL"),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(
							resourceName,
							plancheck.ResourceActionUpdate,
						),
					},
				},
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "description", "second"),
					resource.TestCheckResourceAttr(resourceName, "app_type", "CUSTOM"),
					resource.TestCheckResourceAttr(resourceName, "visibility", "GLOBAL"),
				),
			},
		},
	})
}

// TestAccApp_ImportNonWritableRejected confirms importing a non-CUSTOM app
// into the managed opal_app resource fails with a clear unsupported-type error
// instead of a 400 from GET /apps/{id}/visibility.
func TestAccApp_ImportNonWritableRejected(t *testing.T) {
	t.Parallel()
	nonWritableID := os.Getenv("OPAL_TEST_KNOWN_OPAL_APP_ID")
	adminOwnerID := os.Getenv("OPAL_TEST_KNOWN_CUSTOM_APP_ADMIN_OWNER_ID")
	name := testAccName("app_import_reject")
	resourceName := "opal_app." + name

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProviderFactories,
		PreCheck:                 func() { testAccAppNonWritablePreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: testAccAppCustomConfig(
					name,
					adminOwnerID,
					"import reject fixture",
					"GLOBAL",
				),
				ResourceName:  resourceName,
				ImportState:   true,
				ImportStateId: nonWritableID,
				ExpectError: regexp.MustCompile(
					`(?s)only\s+CUSTOM\s+and\s+CUSTOM_CONNECTOR\s+are\s+supported`,
				),
			},
		},
	})
}

// TestAccAppVisibility_DataSourceCustom confirms opal_app_visibility reads a
// CUSTOM app via GET /apps/{id}/visibility.
func TestAccAppVisibility_DataSourceCustom(t *testing.T) {
	t.Parallel()
	adminOwnerID := os.Getenv("OPAL_TEST_KNOWN_CUSTOM_APP_ADMIN_OWNER_ID")
	name := testAccName("app_vis_ds")

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProviderFactories,
		PreCheck:                 func() { testAccAppManagedPreCheck(t) },
		CheckDestroy:             testAccCheckDestroyedByGET("opal_app", "id", "/apps/%s"),
		Steps: []resource.TestStep{
			{
				Config: testAccAppCustomWithVisibilityDataSourceConfig(name, adminOwnerID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(
						"data.opal_app_visibility.custom", "id",
						"opal_app."+name, "id",
					),
					resource.TestCheckResourceAttr(
						"data.opal_app_visibility.custom",
						"visibility",
						"GLOBAL",
					),
					resource.TestCheckResourceAttrSet(
						"data.opal_app_visibility.custom",
						"import_visibility",
					),
				),
			},
		},
	})
}

// TestAccAppVisibility_DataSourceNonWritableRejected confirms
// opal_app_visibility still errors for app types the visibility sub-endpoint
// rejects (use data.opal_app for those instead).
func TestAccAppVisibility_DataSourceNonWritableRejected(t *testing.T) {
	t.Parallel()
	nonWritableID := os.Getenv("OPAL_TEST_KNOWN_OPAL_APP_ID")

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProviderFactories,
		PreCheck:                 func() { testAccAppNonWritablePreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: testAccAppVisibilityNonWritableConfig(nonWritableID),
				ExpectError: regexp.MustCompile(
					`(?s)unexpected response code 400|cannot be managed via this endpoint`,
				),
			},
		},
	})
}
