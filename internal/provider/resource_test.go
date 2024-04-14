package provider

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	tfprotov6 "github.com/hashicorp/terraform-plugin-go/tfprotov6"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// testAccCheckResourceDestroy verifies the resource has been destroyed.
func testAccCheckResourceDestroy(s *terraform.State) error {
	// Add logic to verify the resource has been destroyed
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "opal_resource" {
			continue
		}

		// Implement API call to check for the existence of the resource
		exists, err := checkResourceExists(rs.Primary.ID)
		if err != nil {
			return fmt.Errorf("Error checking if resource %s exists: %s", rs.Primary.ID, err)
		}
		if exists {
			return fmt.Errorf("Resource %s still exists", rs.Primary.ID)
		}
	}
	return nil
}

// GenerateSimpleResource generates a Terraform configuration for a simple resource.
func GenerateSimpleResource(resourceName string) string {
	return fmt.Sprintf(`
resource "opal_resource" "%s" {
	name = "%s"
	visibility = "GLOBAL"
}
`, resourceName, resourceName)
}

// testAccProviders is a map of Terraform providers for the test cases
var testAccProviders = map[string]func() provider.Provider{
	"opal": ProviderFactory,
}

func ProviderFactory() provider.Provider {
	return New("") // Assuming New returns a provider.Provider
}

// checkResourceExists simulates checking if a resource exists in the backend
// Placeholder for actual API call
func checkResourceExists(id string) (bool, error) {
	// Simulate API call to check if the resource exists
	// This should be replaced with actual API call logic
	return false, nil
}

// TestAccResource_Read tests the reading of an Opal resource.
func TestAccResource_Read(t *testing.T) {
	resourceName := "opal_resource.test"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckResourceConfig(resourceName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "test-resource"),
					resource.TestCheckResourceAttr(resourceName, "visibility", "GLOBAL"),
				),
			},
		},
	})
}

// testAccCheckResourceConfig returns a Terraform configuration for an Opal resource with a given name.
func testAccCheckResourceConfig(name string) string {
	return fmt.Sprintf(`
resource "opal_resource" "%s" {
	name       = "test-resource"
	visibility = "GLOBAL"
}
`, name)
}

// TestAccResource_Update tests the updating of an Opal resource.
func TestAccResource_Update(t *testing.T) {
	t.Parallel()
	baseName, resourceName := generateBaseNameAndResourceName()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProviders,
		PreCheck:                 func() { testAccPreCheck(t) },
		CheckDestroy:             testAccCheckResourceDestroy,
		Steps: []resource.TestStep{
			{
				Config: GenerateSimpleResource(resourceName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", baseName),
					resource.TestCheckResourceAttr(resourceName, "visibility", "GLOBAL"),
				),
			},
			{
				Config: GenerateSimpleResource(resourceName + "_updated"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", baseName+"_updated"),
					resource.TestCheckResourceAttr(resourceName, "visibility", "GLOBAL"),
				),
			},
		},
	})
}

// TestAccResource_Delete tests the deletion of an Opal resource.
func TestAccResource_Delete(t *testing.T) {
	t.Parallel()
	baseName, resourceName := generateBaseNameAndResourceName()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProviders,
		PreCheck:                 func() { testAccPreCheck(t) },
		CheckDestroy:             testAccCheckResourceDestroy,
		Steps: []resource.TestStep{
			{
				Config: GenerateSimpleResource(resourceName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", baseName),
					resource.TestCheckResourceAttr(resourceName, "visibility", "GLOBAL"),
				),
			},
			{
				Config: "",
				Check: resource.ComposeTestCheckFunc(
					testAccCheckResourceDestroy,
				),
			},
		},
	})
}
