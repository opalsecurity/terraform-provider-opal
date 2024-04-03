package provider

import (
	"fmt"
	"os"
	"strings"
)

var opalToken = os.Getenv("OPAL_TEST_TOKEN")
var opalBaseURL = os.Getenv("OPAL_TEST_BASE_URL")

type ReviewerStageConfig struct {
	OwnerIDs               []string
	Operator               string
	RequireManagerApproval bool
}

type RequestConfigurationConfig struct {
	IsRequestable  bool
	ReviewerStages []ReviewerStageConfig
	AutoApproval   bool
	Priority       int
}

type OpalGroupConfig struct {
	ResourceName          string
	Name                  string
	AppID                 string
	GroupType             string
	AdminOwnerID          string
	RequestConfigurations []RequestConfigurationConfig
	Additional            string
	OnCallScheduleIDs     []string
	VisibilityGroupIDs    []string
	Visibility            string
	MessageChannelIDs     []string
}

func GenerateReviewerStages(reviewerStages []ReviewerStageConfig) string {
	reviewerStagesStr := "reviewer_stages = ["
	for _, reviewerStage := range reviewerStages {
		reviewerStagesStr += fmt.Sprintf(
			`{
					operator = "%s"
					owner_ids = [
						"%s"
					]
					require_manager_approval = %t
				}`,
			reviewerStage.Operator, reviewerStage.OwnerIDs[0], reviewerStage.RequireManagerApproval,
		)
	}
	reviewerStagesStr += "]"
	return reviewerStagesStr
}

func GenerateRequestConfigurations(requestConfigurations []RequestConfigurationConfig) string {
	requestConfigurationsStr := "request_configurations = ["
	for _, requestConfiguration := range requestConfigurations {
		requestConfigurationsStr += fmt.Sprintf(
			`{
					allow_requests = %t
					auto_approval = %t
					max_duration = "%d"
					priority = %d
					recommended_duration = %d
					require_mfa_to_request = %t
					require_support_ticket = %t
					%s
				}`,
			requestConfiguration.IsRequestable, requestConfiguration.AutoApproval, 120,
			requestConfiguration.Priority, 120, false,
			false, GenerateReviewerStages(requestConfiguration.ReviewerStages),
		)
	}
	requestConfigurationsStr += "]"
	return requestConfigurationsStr
}

func GenerateGroupResource(ogc *OpalGroupConfig) string {
	requestConfigStr := GenerateRequestConfigurations(ogc.RequestConfigurations)
	resourceStr := fmt.Sprintf(
		`
		provider "opal" {
			server_url = "%s"
			bearer_auth = "%s"
		}
		
		resource "opal_group" "%s" {
			name = "%s"
			app_id = "%s"
			group_type = "%s"
			admin_owner_id = "%s"
			visibility = "%s"
			message_channel_ids = [%s]
			on_call_schedule_ids = [%s]
			%s
		}
		`,
		opalBaseURL, opalToken, ogc.ResourceName, ogc.Name, ogc.AppID, ogc.GroupType,
		ogc.AdminOwnerID, ogc.Visibility, strings.Join(ogc.MessageChannelIDs, ", "), strings.Join(ogc.OnCallScheduleIDs, ", "), requestConfigStr,
	)
	return resourceStr
}
