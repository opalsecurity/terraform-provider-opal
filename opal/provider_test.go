package opal

import (
	"errors"
	"fmt"
	"net/url"
	"os"
	"path"
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

	if v := os.Getenv("OPAL_TEST_KNOWN_GITHUB_APP_REPO_NAME"); v == "" {
		t.Fatal(`OPAL_TEST_KNOWN_GITHUB_APP_REPO_NAME must be set for acceptance tests. This value is the name of the repo you linked and must match what's in the test github org.`)
	}

	if v := os.Getenv("OPAL_TEST_KNOWN_GITHUB_APP_TEAM_SLUG"); v == "" {
		t.Fatal(`OPAL_TEST_KNOWN_GITHUB_APP_TEAM_SLUG must be set for acceptance tests. This value is the name of the team you linked and must match what's in the test github org.`)
	}

	if v := os.Getenv("OPAL_TEST_KNOWN_GITHUB_TEST_REPO_2_RESOURCE_ID"); v == "" {
		t.Fatal(`OPAL_TEST_KNOWN_GITHUB_TEST_REPO_2_RESOURCE_ID must be set for acceptance tests. This value is the Opal id of the test-repo-2 GitHub repo.`)
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
	u, err := url.Parse(baseUrl)
	if err != nil {
		return nil, fmt.Errorf("cannot parse base url: %s", baseUrl)
	}
	u.Path = path.Join(u.Path, "/v1")
	conf.Servers = opal.ServerConfigurations{{
		URL: u.String(),
	}}

	return opal.NewAPIClient(conf), nil
}
