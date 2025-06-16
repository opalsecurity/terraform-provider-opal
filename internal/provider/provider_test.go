package provider

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
)

var testAccProviderFactories map[string]func() (tfprotov6.ProviderServer, error)

func init() {
	// Validate required environment variables first
	opalTestToken := os.Getenv("OPAL_TEST_TOKEN")
	if opalTestToken == "" {
		panic("OPAL_TEST_TOKEN must be set for acceptance tests")
	}

	opalTestBaseURL := os.Getenv("OPAL_TEST_BASE_URL")
	if opalTestBaseURL == "" {
		panic("OPAL_TEST_BASE_URL must be set for acceptance tests")
	}

	// Now set the provider environment variables
	os.Setenv("OPAL_AUTH_TOKEN", opalTestToken)
	os.Setenv("OPAL_BASE_URL", opalTestBaseURL)

	testAccProviderFactories = map[string]func() (tfprotov6.ProviderServer, error){
		"opal": providerserver.NewProtocol6WithError(New("test")()),
	}
}

func testAccPreCheck(t *testing.T) {
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

	if os.Getenv("OPAL_TEST_KNOWN_OPAL_GROUP_ID") == "" {
		t.Fatal("OPAL_TEST_KNOWN_OPAL_GROUP_ID must be set for acceptance tests. You should get this value from an Opal group in the test organization.")
	}

	if os.Getenv("OPAL_TEST_KNOWN_ON_CALL_SCHEDULE_ID") == "" {
		t.Fatal("OPAL_TEST_KNOWN_ON_CALL_SCHEDULE_ID must be set for acceptance tests. You should get this value from an imported Opal on call schedule in the test organization.")
	}
}
