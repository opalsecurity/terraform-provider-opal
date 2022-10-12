package opal

import (
	"errors"
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/opalsecurity/opal-go"
)

var testAccProviders map[string]*schema.Provider
var testAccProvider *schema.Provider

func init() {
	os.Setenv("OPAL_AUTH_TOKEN", os.Getenv("OPAL_TEST_TOKEN"))
	os.Setenv("OPAL_BASE_URL", os.Getenv("OPAL_TEST_BASE_URL"))
	testAccProvider = NewProvider()
	testAccProviders = map[string]*schema.Provider{
		"opal": testAccProvider,
	}
}

// testAccPreCheck validates the necessary test API keys exist
// in the testing environment
func testAccPreCheck(t *testing.T) {
	if v := os.Getenv("OPAL_TEST_TOKEN"); v == "" {
		t.Fatal("OPAL_TEST_TOKEN must be set for acceptance tests")
	}

	if v := os.Getenv("OPAL_TEST_KNOWN_OPAL_APP_ID"); v == "" {
		t.Fatal("OPAL_TEST_KNOWN_OPAL_APP_ID must be set for acceptance tests. You should get this value from the Opal product connection in the test organization.")
	}

	if v := os.Getenv("OPAL_TEST_KNOWN_OPAL_APP_ADMIN_OWNER_ID"); v == "" {
		t.Fatal("OPAL_TEST_KNOWN_OPAL_APP_ADMIN_OWNER_ID must be set for acceptance tests. You should get this value from the owner of the Opal product connection in the test organization.")
	}

	if v := os.Getenv("OPAL_TEST_KNOWN_GITHUB_APP_REPO_ID"); v == "" {
		t.Fatal("OPAL_TEST_KNOWN_GITHUB_APP_RESOURCE_REMOTE_ID must be set for acceptance tests. This value is the id of a test repo and must match what's in the test github org.")
	}

	if v := os.Getenv("OPAL_TEST_KNOWN_GITHUB_APP_REPO_NAME"); v == "" {
		t.Fatal(`OPAL_TEST_KNOWN_GITHUB_APP_REPO_NAME must be set for acceptance tests. This value is the name of the repo you linked and must match what's in the test github org.`)
	}

	if v := os.Getenv("OPAL_TEST_KNOWN_GITHUB_APP_GROUP_REMOTE_ID"); v == "" {
		t.Fatal("OPAL_TEST_KNOWN_GITHUB_APP_GROUP_REMOTE_ID must be set for acceptance tests. This value is in the form known-org/known-team and must match what's in the test github org.")
	}

	if v := os.Getenv("OPAL_TEST_KNOWN_GITHUB_APP_GROUP_METADATA"); v == "" {
		t.Fatal(`OPAL_TEST_KNOWN_GITHUB_APP_GROUP_METADATA must be set for acceptance tests. This value is in the form {"git_hub_team"={"org_name"="example-org", "team_slug"="example-team"}} and must match what's in the test github org.`)
	}

	if os.Getenv("OPAL_TEST_KNOWN_USER_ID_1") == "" {
		t.Fatal("OPAL_TEST_KNOWN_USER_ID_1 must be set for acceptance tests. You should get this value from any user in the test organization.")
	}

	if os.Getenv("OPAL_TEST_KNOWN_USER_ID_2") == "" {
		t.Fatal("OPAL_TEST_KNOWN_USER_ID_2 must be set for acceptance tests. You should get this value from any user in the test organization.")
	}

	if os.Getenv("OPAL_TEST_KNOWN_USER_ID_1") == os.Getenv("OPAL_TEST_KNOWN_USER_ID_2") {
		t.Fatal("OPAL_TEST_KNOWN_USER_ID_1 should not be the same as OPAL_TEST_KNOWN_USER_ID_2")
	}

	if os.Getenv("OPAL_TEST_KNOWN_CUSTOM_APP_ID") == "" {
		t.Fatal("OPAL_TEST_KNOWN_CUSTOM_APP_ID must be set for acceptance tests. You should get this value from a custom app integration in the test organization.")
	}

	if os.Getenv("OPAL_TEST_KNOWN_CUSTOM_APP_ADMIN_OWNER_ID") == "" {
		t.Fatal("OPAL_TEST_KNOWN_CUSTOM_APP_ADMIN_OWNER_ID must be set for acceptance tests. You should get this value from a custom app integration in the test organization.")
	}

	if os.Getenv("OPAL_TEST_KNOWN_REQUEST_TEMPLATE_ID") == "" {
		t.Fatal("OPAL_TEST_KNOWN_REQUEST_TEMPLATE_ID must be set for acceptance tests. You should get this value from a custom app integration in the test organization.")
	}

	if os.Getenv("OPAL_TEST_BASE_URL") == "" {
		t.Fatal("OPAL_TEST_BASE_URL must be set for acceptance tests")
	}
}

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

func sweeperClient() (*opal.APIClient, error) {
	conf := opal.NewConfiguration()

	token := os.Getenv("OPAL_TEST_TOKEN")
	if token == "" {
		return nil, errors.New("OPAL_TEST_TOKEN must be set")
	}
	conf.DefaultHeader["Authorization"] = fmt.Sprintf("Bearer %s", token)

	baseUrl := os.Getenv("OPAL_TEST_BASE_URL")
	if baseUrl == "" {
		return nil, errors.New("OPAL_TEST_BASE_URL must be set")
	}
	conf.Servers = opal.ServerConfigurations{{
		URL: baseUrl,
	}}

	return opal.NewAPIClient(conf), nil
}
