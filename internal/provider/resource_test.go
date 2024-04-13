package provider

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	tfversion "github.com/hashicorp/terraform-plugin-sdk/v2/version"
	"github.com/opalsecurity/terraform-provider-opal/opal/internal/sdk"
)

var knownOpalAppID = os.Getenv("OPAL_TEST_KNOWN_OPAL_APP_ID")
var knownOpalAppAdminOwnerID = os.Getenv("OPAL_TEST_KNOWN_OPAL_APP_ADMIN_OWNER_ID")

func generateBaseNameAndResourceName() (string, string) {
	baseName := "tf_acc_resource_test_" + acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	resourceName := "opal_resource." + baseName
	return baseName, resourceName
}

func generateSimpleOpalResourceConfig(baseName string, resourceName string) OpalResourceConfig {
	return OpalResourceConfig{
		ResourceName: resourceName,
		Name:         baseName,
		Visibility:   "GLOBAL",
	}
}

func TestAccResource_Import(t *testing.T) {
	t.Parallel()
	baseName, resourceName := generateBaseNameAndResourceName()
	config := generateSimpleOpalResourceConfig(baseName, resourceName)

	resource.Test(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_0_0),
		},
		ProtoV6ProviderFactories: testAccProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		CheckDestroy:             testAccCheckResourceDestroy,
		Steps: []resource.TestStep{
			{
				Config: GenerateSimpleResource(&config),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", config.Name),
					resource.TestCheckResourceAttr(resourceName, "visibility", "GLOBAL"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckResourceDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "opal_resource" {
			continue
		}

		security := shared.Security{
			BearerAuth: &opalToken,
		}
		opts := []sdk.SDKOption{
			sdk.WithServerURL(opalBaseURL),
			sdk.WithSecurity(security),
			sdk.WithClient(http.DefaultClient),
		}
		client := sdk.New(opts...)
		resource, err := client.Resources.GetResource(
			context.Background(),
			operations.GetResourceRequest{
				ID: rs.Primary.ID,
			},
		)
		if err == nil {
			if resource != nil {
				return fmt.Errorf("Resource %s still exists", rs.Primary.ID)
			}
		}

		if !strings.Contains(err.Error(), "404") {
			return err
		}
	}

	return nil
}
