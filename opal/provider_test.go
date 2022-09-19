package opal

import (
	"errors"
	"fmt"
	"net/url"
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

	if os.Getenv("OPAL_TEST_KNOWN_USER_ID_1") == "" {
		t.Fatal("OPAL_TEST_KNOWN_USER_ID_1 must be set for acceptance tests. You should get this value from any user in the test organization.")
	}

	if os.Getenv("OPAL_TEST_KNOWN_USER_ID_2") == "" {
		t.Fatal("OPAL_TEST_KNOWN_USER_ID_2 must be set for acceptance tests. You should get this value from any user in the test organization.")
	}

	if os.Getenv("OPAL_TEST_KNOWN_USER_ID_1") == os.Getenv("OPAL_TEST_KNOWN_USER_ID_2") {
		t.Fatal("OPAL_TEST_KNOWN_USER_ID_1 should not be the same as OPAL_TEST_KNOWN_USER_ID_2")
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
		return nil, err
	}
	conf.Host = u.Host
	conf.Scheme = u.Scheme

	return opal.NewAPIClient(conf), nil
}
