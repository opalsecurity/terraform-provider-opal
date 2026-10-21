package provider

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

var knownUserID1 = os.Getenv("OPAL_TEST_KNOWN_USER_ID_1")
var knownUserID2 = os.Getenv("OPAL_TEST_KNOWN_USER_ID_2")
var knownCustomAppID = os.Getenv("OPAL_TEST_KNOWN_CUSTOM_APP_ID")
var knownCustomAppAdminOwnerID = os.Getenv("OPAL_TEST_KNOWN_CUSTOM_APP_ADMIN_OWNER_ID")

// testAccProviderConfig is the provider block every test config starts with.
func testAccProviderConfig() string {
	return fmt.Sprintf(`
provider "opal" {
	server_url  = "%s"
	bearer_auth = "%s"
}
`, opalBaseURL, opalToken)
}

// testAccName returns a unique name for objects created by acceptance tests so
// leftovers from failed runs are easy to spot in the test org.
func testAccName(prefix string) string {
	return "tf_acc_" + prefix + "_" + acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
}

// testAccCheckDestroyedByGET returns a CheckDestroy that asserts a GET of
// pathFmt (formatted with the resource's idAttr) returns 404 for every
// resourceType in state.
func testAccCheckDestroyedByGET(resourceType string, idAttr string, pathFmt string) func(*terraform.State) error {
	return func(s *terraform.State) error {
		for _, rs := range s.RootModule().Resources {
			if rs.Type != resourceType {
				continue
			}

			id := rs.Primary.Attributes[idAttr]
			url := strings.TrimSuffix(opalBaseURL, "/") + fmt.Sprintf(pathFmt, id)
			req, err := http.NewRequest(http.MethodGet, url, nil)
			if err != nil {
				return err
			}
			req.Header.Set("Authorization", "Bearer "+opalToken)
			res, err := http.DefaultClient.Do(req)
			if err != nil {
				return fmt.Errorf("checking %s %s was destroyed: %w", resourceType, id, err)
			}
			res.Body.Close()

			if res.StatusCode != http.StatusNotFound {
				return fmt.Errorf("expected 404 after destroying %s %s but got %d", resourceType, id, res.StatusCode)
			}
		}
		return nil
	}
}

// testAccImportStateIDFromAttr returns an ImportStateIdFunc that reads attr
// from resourceName, for resources whose import ID is not the "id" attribute.
func testAccImportStateIDFromAttr(resourceName string, attr string) func(*terraform.State) (string, error) {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return "", fmt.Errorf("resource %s not found in state", resourceName)
		}
		return rs.Primary.Attributes[attr], nil
	}
}

func testAccPreCheckCustomApp(t *testing.T) {
	testAccPreCheck(t)
	if knownCustomAppID == "" || knownCustomAppAdminOwnerID == "" {
		t.Fatal("OPAL_TEST_KNOWN_CUSTOM_APP_ID and OPAL_TEST_KNOWN_CUSTOM_APP_ADMIN_OWNER_ID must be set for this test")
	}
}
