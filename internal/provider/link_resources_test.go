package provider

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

// Tests for the resources that link two existing objects (bundle_group,
// group_tag, group_user, ...). Each test applies a shared fixture plus the
// link, checks the link through the parent's list endpoint, then removes only
// the link and checks it is gone from the API.

// testAccLinkFixtureConfig creates the objects the link resources point at:
// two Opal groups, a custom app resource, a tag and a bundle.
func testAccLinkFixtureConfig(prefix string) string {
	requestConfigurations := GenerateRequestConfigurations([]RequestConfigurationConfig{
		{
			IsRequestable: true,
			ReviewerStages: []ReviewerStageConfig{
				{
					OwnerIDs: []string{knownOpalAppAdminOwnerID},
					Operator: "AND",
				},
			},
		},
	})

	return testAccProviderConfig() + fmt.Sprintf(`
resource "opal_group" "group" {
	name           = "%[1]s_group"
	app_id         = "%[2]s"
	group_type     = "OPAL_GROUP"
	admin_owner_id = "%[3]s"
	visibility     = "GLOBAL"
	%[6]s
}

resource "opal_group" "containing_group" {
	name           = "%[1]s_containing_group"
	app_id         = "%[2]s"
	group_type     = "OPAL_GROUP"
	admin_owner_id = "%[3]s"
	visibility     = "GLOBAL"
	%[6]s
}

resource "opal_resource" "resource" {
	name           = "%[1]s_resource"
	app_id         = "%[4]s"
	resource_type  = "CUSTOM"
	admin_owner_id = "%[5]s"
	visibility     = "GLOBAL"
	%[6]s
}

resource "opal_tag" "tag" {
	key   = "%[1]s_tag"
	value = "v"
}

resource "opal_bundle" "bundle" {
	name           = "%[1]s_bundle"
	admin_owner_id = "%[3]s"
	visibility     = "GLOBAL"
}
`, prefix, knownOpalAppID, knownOpalAppAdminOwnerID, knownCustomAppID, knownCustomAppAdminOwnerID, requestConfigurations)
}

// testAccLinkCheckTimeout bounds how long testAccCheckLinkInAPI waits for some
// link changes, which the API applies asynchronously.
const testAccLinkCheckTimeout = 60 * time.Second

// testAccStateAttr returns attr of resourceName from state, or attr itself as
// a literal value when resourceName is empty.
func testAccStateAttr(s *terraform.State, resourceName string, attr string) (string, error) {
	if resourceName == "" {
		return attr, nil
	}
	rs, ok := s.RootModule().Resources[resourceName]
	if !ok {
		return "", fmt.Errorf("resource %s not found in state", resourceName)
	}
	return rs.Primary.Attributes[attr], nil
}

// testAccCheckLinkInAPI GETs pathFmt (formatted with parentAttr of
// parentResource) and checks whether childAttr of childResource appears in the
// response body, polling until it matches wantPresent or the timeout passes.
// Pass an empty resource name to use the attr as a literal ID.
func testAccCheckLinkInAPI(
	parentResource string,
	parentAttr string,
	pathFmt string,
	childResource string,
	childAttr string,
	wantPresent bool,
) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		parentID, err := testAccStateAttr(s, parentResource, parentAttr)
		if err != nil {
			return err
		}
		childID, err := testAccStateAttr(s, childResource, childAttr)
		if err != nil {
			return err
		}

		url := strings.TrimSuffix(opalBaseURL, "/") + fmt.Sprintf(pathFmt, parentID)
		deadline := time.Now().Add(testAccLinkCheckTimeout)
		for {
			present, err := testAccGETContains(url, childID)
			if err != nil {
				return fmt.Errorf("GET %s: %w", pathFmt, err)
			}
			if present == wantPresent {
				return nil
			}
			if time.Now().After(deadline) {
				return fmt.Errorf("expected %s in GET %s to be present=%t, still present=%t after %s",
					childID, pathFmt, wantPresent, present, testAccLinkCheckTimeout)
			}
			time.Sleep(2 * time.Second)
		}
	}
}

func testAccGETContains(url string, want string) (bool, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return false, err
	}
	req.Header.Set("Authorization", "Bearer "+opalToken)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return false, err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return false, err
	}
	if res.StatusCode != http.StatusOK {
		return false, fmt.Errorf("status %d", res.StatusCode)
	}
	return strings.Contains(string(body), want), nil
}

func testAccLinkPreCheck(t *testing.T) {
	testAccPreCheckCustomApp(t)
}

func testAccLinkCheckDestroy(s *terraform.State) error {
	for _, check := range []func(*terraform.State) error{
		testAccCheckGroupDestroy,
		testAccCheckDestroyedByGET("opal_resource", "id", "/resources/%s"),
		testAccCheckDestroyedByGET("opal_tag", "id", "/tag/%s"),
		testAccCheckDestroyedByGET("opal_bundle", "bundle_id", "/bundles/%s"),
	} {
		if err := check(s); err != nil {
			return err
		}
	}
	return nil
}

func TestAccBundleGroup_CRUD(t *testing.T) {
	t.Parallel()
	prefix := testAccName("bundle_group")
	fixture := testAccLinkFixtureConfig(prefix)
	linkName := "opal_bundle_group.link"

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProviderFactories,
		PreCheck:                 func() { testAccLinkPreCheck(t) },
		CheckDestroy:             testAccLinkCheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: fixture + `
resource "opal_bundle_group" "link" {
	bundle_id = opal_bundle.bundle.bundle_id
	group_id  = opal_group.group.id
}
`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(linkName, "bundle_id", "opal_bundle.bundle", "bundle_id"),
					resource.TestCheckResourceAttrPair(linkName, "group_id", "opal_group.group", "id"),
					testAccCheckLinkInAPI("opal_bundle.bundle", "bundle_id", "/bundles/%s/groups", "opal_group.group", "id", true),
				),
			},
			{
				Config: fixture,
				Check:  testAccCheckLinkInAPI("opal_bundle.bundle", "bundle_id", "/bundles/%s/groups", "opal_group.group", "id", false),
			},
		},
	})
}

func TestAccBundleResource_CRUD(t *testing.T) {
	t.Parallel()
	prefix := testAccName("bundle_resource")
	fixture := testAccLinkFixtureConfig(prefix)
	linkName := "opal_bundle_resource.link"

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProviderFactories,
		PreCheck:                 func() { testAccLinkPreCheck(t) },
		CheckDestroy:             testAccLinkCheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: fixture + `
resource "opal_bundle_resource" "link" {
	bundle_id   = opal_bundle.bundle.bundle_id
	resource_id = opal_resource.resource.id
}
`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(linkName, "bundle_id", "opal_bundle.bundle", "bundle_id"),
					resource.TestCheckResourceAttrPair(linkName, "resource_id", "opal_resource.resource", "id"),
					testAccCheckLinkInAPI("opal_bundle.bundle", "bundle_id", "/bundles/%s/resources", "opal_resource.resource", "id", true),
				),
			},
			{
				Config: fixture,
				Check:  testAccCheckLinkInAPI("opal_bundle.bundle", "bundle_id", "/bundles/%s/resources", "opal_resource.resource", "id", false),
			},
		},
	})
}

func TestAccGroupTag_CRUD(t *testing.T) {
	t.Parallel()
	prefix := testAccName("group_tag")
	fixture := testAccLinkFixtureConfig(prefix)
	linkName := "opal_group_tag.link"

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProviderFactories,
		PreCheck:                 func() { testAccLinkPreCheck(t) },
		CheckDestroy:             testAccLinkCheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: fixture + `
resource "opal_group_tag" "link" {
	group_id = opal_group.group.id
	tag_id   = opal_tag.tag.id
}
`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(linkName, "group_id", "opal_group.group", "id"),
					resource.TestCheckResourceAttrPair(linkName, "tag_id", "opal_tag.tag", "id"),
					testAccCheckLinkInAPI("opal_group.group", "id", "/groups/%s/tags", "opal_tag.tag", "id", true),
				),
			},
			{
				Config: fixture,
				Check:  testAccCheckLinkInAPI("opal_group.group", "id", "/groups/%s/tags", "opal_tag.tag", "id", false),
			},
		},
	})
}

func TestAccResourceTag_CRUD(t *testing.T) {
	t.Parallel()
	prefix := testAccName("resource_tag")
	fixture := testAccLinkFixtureConfig(prefix)
	linkName := "opal_resource_tag.link"

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProviderFactories,
		PreCheck:                 func() { testAccLinkPreCheck(t) },
		CheckDestroy:             testAccLinkCheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: fixture + `
resource "opal_resource_tag" "link" {
	resource_id = opal_resource.resource.id
	tag_id      = opal_tag.tag.id
}
`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(linkName, "resource_id", "opal_resource.resource", "id"),
					resource.TestCheckResourceAttrPair(linkName, "tag_id", "opal_tag.tag", "id"),
					testAccCheckLinkInAPI("opal_resource.resource", "id", "/resources/%s/tags", "opal_tag.tag", "id", true),
				),
			},
			{
				Config: fixture,
				Check:  testAccCheckLinkInAPI("opal_resource.resource", "id", "/resources/%s/tags", "opal_tag.tag", "id", false),
			},
		},
	})
}

func TestAccTagUser_CRUD(t *testing.T) {
	t.Parallel()
	prefix := testAccName("tag_user")
	fixture := testAccLinkFixtureConfig(prefix)
	linkName := "opal_tag_user.link"

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProviderFactories,
		PreCheck:                 func() { testAccLinkPreCheck(t) },
		CheckDestroy:             testAccLinkCheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: fixture + fmt.Sprintf(`
resource "opal_tag_user" "link" {
	tag_id  = opal_tag.tag.id
	user_id = "%s"
}
`, knownUserID2),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(linkName, "tag_id", "opal_tag.tag", "id"),
					resource.TestCheckResourceAttr(linkName, "user_id", knownUserID2),
					testAccCheckLinkInAPI("", knownUserID2, "/users/%s/tags", "opal_tag.tag", "id", true),
				),
			},
			{
				Config: fixture,
				Check:  testAccCheckLinkInAPI("", knownUserID2, "/users/%s/tags", "opal_tag.tag", "id", false),
			},
		},
	})
}

func TestAccGroupUser_CRUD(t *testing.T) {
	t.Parallel()
	prefix := testAccName("group_user")
	fixture := testAccLinkFixtureConfig(prefix)
	linkName := "opal_group_user.link"

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProviderFactories,
		PreCheck:                 func() { testAccLinkPreCheck(t) },
		CheckDestroy:             testAccLinkCheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: fixture + fmt.Sprintf(`
resource "opal_group_user" "link" {
	group_id = opal_group.group.id
	user_id  = "%s"
}
`, knownUserID2),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(linkName, "group_id", "opal_group.group", "id"),
					resource.TestCheckResourceAttr(linkName, "user_id", knownUserID2),
					testAccCheckLinkInAPI("opal_group.group", "id", "/groups/%s/users", "", knownUserID2, true),
				),
			},
			{
				Config: fixture,
				Check:  testAccCheckLinkInAPI("opal_group.group", "id", "/groups/%s/users", "", knownUserID2, false),
			},
		},
	})
}

func TestAccGroupContainingGroup_CRUD(t *testing.T) {
	t.Parallel()
	prefix := testAccName("group_containing_group")
	fixture := testAccLinkFixtureConfig(prefix)
	linkName := "opal_group_containing_group.link"

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProviderFactories,
		PreCheck:                 func() { testAccLinkPreCheck(t) },
		CheckDestroy:             testAccLinkCheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: fixture + `
resource "opal_group_containing_group" "link" {
	group_id            = opal_group.group.id
	containing_group_id = opal_group.containing_group.id
}
`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(linkName, "group_id", "opal_group.group", "id"),
					resource.TestCheckResourceAttrPair(linkName, "containing_group_id", "opal_group.containing_group", "id"),
					testAccCheckLinkInAPI("opal_group.group", "id", "/groups/%s/containing-groups", "opal_group.containing_group", "id", true),
				),
			},
			{
				ResourceName: linkName,
				ImportState:  true,
				ImportStateIdFunc: func(s *terraform.State) (string, error) {
					rs := s.RootModule().Resources[linkName]
					return fmt.Sprintf(`{"containing_group_id": %q, "group_id": %q}`,
						rs.Primary.Attributes["containing_group_id"],
						rs.Primary.Attributes["group_id"],
					), nil
				},
				ImportStateVerify:                    true,
				ImportStateVerifyIdentifierAttribute: "group_id",
			},
			{
				Config: fixture,
				Check:  testAccCheckLinkInAPI("opal_group.group", "id", "/groups/%s/containing-groups", "opal_group.containing_group", "id", false),
			},
		},
	})
}

func TestAccGroupResourceList_CRUD(t *testing.T) {
	t.Parallel()
	prefix := testAccName("group_resource_list")
	fixture := testAccLinkFixtureConfig(prefix)
	linkName := "opal_group_resource_list.link"

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProviderFactories,
		PreCheck:                 func() { testAccLinkPreCheck(t) },
		CheckDestroy:             testAccLinkCheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: fixture + `
resource "opal_group_resource_list" "link" {
	group_id  = opal_group.group.id
	resources = [
		{
			resource_id = opal_resource.resource.id
		},
	]
}
`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(linkName, "group_id", "opal_group.group", "id"),
					resource.TestCheckResourceAttr(linkName, "resources.#", "1"),
					resource.TestCheckResourceAttrPair(linkName, "resources.0.resource_id", "opal_resource.resource", "id"),
					testAccCheckLinkInAPI("opal_group.group", "id", "/groups/%s/resources", "opal_resource.resource", "id", true),
				),
			},
			{
				ResourceName:                         linkName,
				ImportState:                          true,
				ImportStateIdFunc:                    testAccImportStateIDFromAttr(linkName, "group_id"),
				ImportStateVerify:                    true,
				ImportStateVerifyIdentifierAttribute: "group_id",
				// Reads only populate the computed group_resources, so an import
				// can't recover the configured resources list
				ImportStateVerifyIgnore: []string{"resources"},
			},
			// No removal step: opal_group_resource_list has no DELETE operation, so
			// removing it from config leaves the resources on the group. Add a step
			// asserting they're gone once delete is implemented.
		},
	})
}
