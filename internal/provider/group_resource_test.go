package provider

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-plugin-testing/tfversion"
	"github.com/opalsecurity/terraform-provider-opal/internal/sdk"
	"github.com/opalsecurity/terraform-provider-opal/internal/sdk/models/operations"
	"github.com/opalsecurity/terraform-provider-opal/internal/sdk/models/shared"
	"github.com/pkg/errors"
)

var knownOpalAppID = os.Getenv("OPAL_TEST_KNOWN_OPAL_APP_ID")
var knownOpalAppAdminOwnerID = os.Getenv("OPAL_TEST_KNOWN_OPAL_APP_ADMIN_OWNER_ID")
var knownOpalGroupID = os.Getenv("OPAL_TEST_KNOWN_OPAL_GROUP_ID")
var knownOpalGithubTeamSlug = os.Getenv("OPAL_TEST_KNOWN_GITHUB_APP_TEAM_SLUG")
var knownGithubAppID = os.Getenv("OPAL_TEST_KNOWN_GITHUB_APP_ID")

func generateBaseNameAndResourceName() (string, string) {
	baseName := "tf_acc_group_test_" + acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	resourceName := "opal_group." + baseName
	return baseName, resourceName
}

func generateSimpleOpalGroupConfig(baseName string, resourceName string) OpalGroupConfig {
	return OpalGroupConfig{
		ResourceName:       resourceName,
		Name:               baseName,
		Description:        "Test description",
		AppID:              knownOpalAppID,
		GroupType:          "OPAL_GROUP",
		AdminOwnerID:       knownOpalAppAdminOwnerID,
		Visibility:         "GLOBAL",
		VisibilityGroupIDs: []string{},
		MessageChannelIDs:  []string{},
		OnCallScheduleIDs:  []string{},
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
}

func TestAccGroup_Import(t *testing.T) {
	t.Parallel()
	baseName, resourceName := generateBaseNameAndResourceName()
	config := generateSimpleOpalGroupConfig(baseName, baseName)

	resource.Test(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_0_0),
		},
		ProtoV6ProviderFactories: testAccProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		CheckDestroy:             testAccCheckGroupDestroy,
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

func TestAccGroup_CRUD(t *testing.T) {
	t.Parallel()
	baseName, resourceName := generateBaseNameAndResourceName()
	config := generateSimpleOpalGroupConfig(baseName, baseName)

	configString := GenerateGroupResource(&config)
	OLD_DESCRIPTION := config.Description
	NEW_DESCRIPTION := "New description"
	config.Description = NEW_DESCRIPTION
	updatedConfigString := GenerateGroupResource(&config)

	resource.Test(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_0_0),
		},
		ProtoV6ProviderFactories: testAccProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: configString,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", config.Name),
					resource.TestCheckResourceAttr(resourceName, "app_id", config.AppID),
					resource.TestCheckResourceAttr(resourceName, "description", OLD_DESCRIPTION),
					resource.TestCheckResourceAttr(resourceName, "group_type", "OPAL_GROUP"),
					resource.TestCheckResourceAttr(resourceName, "admin_owner_id", knownOpalAppAdminOwnerID),
					resource.TestCheckResourceAttr(resourceName, "visibility", "GLOBAL"),
					resource.TestCheckResourceAttr(resourceName, "on_call_schedule_ids.#", "0"),
					resource.TestCheckResourceAttr(resourceName, "message_channel_ids.#", "0"),
					resource.TestCheckResourceAttr(resourceName, "request_configurations.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "request_configurations.0.allow_requests", "true"),
					resource.TestCheckResourceAttr(resourceName, "request_configurations.0.auto_approval", "false"),
					resource.TestCheckResourceAttr(resourceName, "request_configurations.0.max_duration", "120"),
					resource.TestCheckResourceAttr(resourceName, "request_configurations.0.priority", "0"),
					resource.TestCheckResourceAttr(resourceName, "request_configurations.0.recommended_duration", "120"),
					resource.TestCheckResourceAttr(resourceName, "request_configurations.0.require_mfa_to_request", "false"),
					resource.TestCheckResourceAttr(resourceName, "request_configurations.0.require_support_ticket", "false"),
					resource.TestCheckResourceAttr(resourceName, "request_configurations.0.reviewer_stages.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "request_configurations.0.reviewer_stages.0.operator", "AND"),
					resource.TestCheckResourceAttr(resourceName, "request_configurations.0.reviewer_stages.0.owner_ids.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "request_configurations.0.reviewer_stages.0.owner_ids.0", knownOpalAppAdminOwnerID),
					resource.TestCheckResourceAttr(resourceName, "request_configurations.0.reviewer_stages.0.require_manager_approval", "false"),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"visibility", "on_call_schedule_ids", "message_channel_ids"},
			},
			{
				Config: updatedConfigString,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", config.Name),
					resource.TestCheckResourceAttr(resourceName, "app_id", config.AppID),
					resource.TestCheckResourceAttr(resourceName, "description", NEW_DESCRIPTION),
					resource.TestCheckResourceAttr(resourceName, "group_type", "OPAL_GROUP"),
					resource.TestCheckResourceAttr(resourceName, "admin_owner_id", knownOpalAppAdminOwnerID),
					resource.TestCheckResourceAttr(resourceName, "visibility", "GLOBAL"),
					resource.TestCheckResourceAttr(resourceName, "on_call_schedule_ids.#", "0"),
					resource.TestCheckResourceAttr(resourceName, "message_channel_ids.#", "0"),
					resource.TestCheckResourceAttr(resourceName, "request_configurations.#", "1"),
				),
			},
		},
	})
}

func TestAccGroup_Visibility(t *testing.T) {
	t.Parallel()
	baseName, resourceName := generateBaseNameAndResourceName()
	config := generateSimpleOpalGroupConfig(baseName, baseName)

	configString := GenerateGroupResource(&config)
	config.Visibility = "INVALID_VISIBILITY"
	invalidVisibilityTypeConfigString := GenerateGroupResource(&config)
	config.VisibilityGroupIDs = []string{knownOpalGroupID}
	config.Visibility = "LIMITED"
	teamVisibilityConfigString := GenerateGroupResource(&config)

	resource.Test(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_0_0),
		},
		ProtoV6ProviderFactories: testAccProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		CheckDestroy:             testAccCheckGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: configString,
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
			{
				Config:      invalidVisibilityTypeConfigString,
				PlanOnly:    true,
				ExpectError: GenerateErrorMessageRegexp("Invalid Attribute Value Match"),
			},
			{
				Config: teamVisibilityConfigString,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", config.Name),
					resource.TestCheckResourceAttr(resourceName, "app_id", config.AppID),
					resource.TestCheckResourceAttr(resourceName, "visibility", "LIMITED"),
					resource.TestCheckResourceAttr(resourceName, "visibility_group_ids.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "visibility_group_ids.0", knownOpalGroupID),
				),
			},
		},
	})
}

func TestAccGroup_RequestConfigurations(t *testing.T) {
	t.Parallel()
	baseName, resourceName := generateBaseNameAndResourceName()
	config := generateSimpleOpalGroupConfig(baseName, baseName)

	configString := GenerateGroupResource(&config)
	config.RequestConfigurations = []RequestConfigurationConfig{
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
			Priority:     1,
			Condition: &ConditionConfig{
				GroupIDs: []string{knownOpalGroupID},
			},
		},
	}
	sequentialPriorityConfigString := GenerateGroupResource(&config)

	config.RequestConfigurations = []RequestConfigurationConfig{
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
			Condition: &ConditionConfig{
				GroupIDs: []string{knownOpalGroupID},
			},
		},
	}
	invalidDefaultConditionConfigString := GenerateGroupResource(&config)

	config.RequestConfigurations = []RequestConfigurationConfig{
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
			Priority:     2,
		},
		{
			IsRequestable: true,
			AutoApproval:  true,
			Priority:      3,
		},
	}
	invalidSequentialPriorityConfigString := GenerateGroupResource(&config)

	emptyRequestConfigurationsConfig := []RequestConfigurationConfig{}
	config.RequestConfigurations = emptyRequestConfigurationsConfig
	emptyRequestConfigurationsConfigString := GenerateGroupResource(&config)

	resource.Test(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_0_0),
		},
		ProtoV6ProviderFactories: testAccProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		CheckDestroy:             testAccCheckGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: configString,
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
			{
				Config:      invalidDefaultConditionConfigString,
				PlanOnly:    true,
				ExpectError: GenerateErrorMessageRegexp("Invalid Attribute Type"),
			},
			{
				Config:      invalidSequentialPriorityConfigString,
				PlanOnly:    true,
				ExpectError: GenerateErrorMessageRegexp("Invalid Attribute Type"),
			},
			{
				Config:      emptyRequestConfigurationsConfigString,
				PlanOnly:    true,
				ExpectError: GenerateErrorMessageRegexp("Invalid Attribute Value"),
			},
			{
				Config: sequentialPriorityConfigString,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectNonEmptyPlan(),
					},
				},
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "request_configurations.#", "2"),
				),
			},
		},
	})
}

func TestAccGroup_RemoteInfo(t *testing.T) {
	t.Parallel()
	baseName, resourceName := generateBaseNameAndResourceName()
	config := generateSimpleOpalGroupConfig(baseName, baseName)
	// Manually adding remote_info since there's too many options for remote_info
	config.GroupType = "GIT_HUB_TEAM"
	config.AppID = knownGithubAppID
	config.Additional = fmt.Sprintf(`
	remote_info = {
		github_team = {
			team_slug = "%s"
		}
	}`, knownOpalGithubTeamSlug)

	resource.Test(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_0_0),
		},
		ProtoV6ProviderFactories: testAccProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		CheckDestroy:             testAccCheckGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: GenerateGroupResource(&config),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", config.Name),
					resource.TestCheckResourceAttr(resourceName, "app_id", config.AppID),
					resource.TestCheckResourceAttr(resourceName, "visibility", "GLOBAL"),
					resource.TestCheckResourceAttr(resourceName, "group_type", "GIT_HUB_TEAM"),
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

func testAccCheckGroupDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "opal_group" {
			continue
		}

		security := shared.Security{
			BearerAuth: opalToken,
		}
		opts := []sdk.SDKOption{
			sdk.WithServerURL(opalBaseURL),
			sdk.WithSecurity(security),
			sdk.WithClient(http.DefaultClient),
		}
		client := sdk.New(opts...)
		group, err := client.Groups.GetGroup(
			context.Background(),
			operations.GetGroupRequest{
				ID: rs.Primary.ID,
			},
		)
		if err != nil {
			return errors.Wrapf(err, "unexpected error while checking status of Terraform resource %v.", rs.Primary.ID)
		}

		if group.StatusCode != 404 {
			return fmt.Errorf("Expected 404 after destorying the group but got %d", group.StatusCode)
		}
	}

	return nil
}
