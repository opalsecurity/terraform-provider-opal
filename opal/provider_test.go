package opal

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
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

	if os.Getenv("OPAL_TEST_KNOWN_USER_ID") == "" {
		t.Fatal("OPAL_TEST_KNOWN_USER_ID must be set for acceptance tests. You should get this value from any user in the test organization.")
	}

	if os.Getenv("OPAL_TEST_BASE_URL") == "" {
		t.Fatal("OPAL_TEST_BASE_URL must be set for acceptance tests")
	}
}
