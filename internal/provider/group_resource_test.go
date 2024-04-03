package provider

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/tfversion"
)

var knownOpalAppID = os.Getenv("OPAL_TEST_KNOWN_OPAL_APP_ID")
var knownOpalAppAdminOwnerID = os.Getenv("OPAL_TEST_KNOWN_OPAL_APP_ADMIN_OWNER_ID")

func TestAccGroup_Import(t *testing.T) {
	t.Parallel()
	baseName := "tf_acc_group_test_" + acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	resourceName := "opal_group." + baseName

	config := OpalGroupConfig{
		ResourceName:      baseName,
		Name:              baseName,
		AppID:             knownOpalAppID,
		GroupType:         "OPAL_GROUP",
		AdminOwnerID:      knownOpalAppAdminOwnerID,
		Visibility:        "GLOBAL",
		OnCallScheduleIDs: []string{},
		RequestConfigurations: []RequestConfigurationConfig{
			{
				IsRequestable: true,
				ReviewerStages: []ReviewerStageConfig{
					{
						OwnerIDs:               []string{knownOpalAppAdminOwnerID},
						Operator:               "AND",
						RequireManagerApproval: false,
					},
				},
				AutoApproval: false,
				Priority:     0,
			},
		},
	}

	resource.Test(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_0_0),
		},
		ProtoV6ProviderFactories: testAccProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: GenerateGroupResource(&config),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", config.Name),
					resource.TestCheckResourceAttr(resourceName, "app_id", config.AppID),
					resource.TestCheckResourceAttr(resourceName, "visibility", "GLOBAL"),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"visibility", "on_call_schedule_ids", "message_channel_ids"},
			},
		},
	})
}
