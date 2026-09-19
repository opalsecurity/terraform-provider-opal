package provider

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

var opalToken = os.Getenv("OPAL_TEST_TOKEN")
var opalBaseURL = os.Getenv("OPAL_TEST_BASE_URL")

type ReviewerStageConfig struct {
	OwnerIDs               []string
	Operator               string
	RequireManagerApproval bool
}

type ConditionConfig struct {
	GroupIDs      []string
	RemoteRoleIds []string
}

type RequestConfigurationConfig struct {
	Condition      *ConditionConfig
	IsRequestable  bool
	ReviewerStages []ReviewerStageConfig
	AutoApproval   bool
	ReasonOptional bool
	Priority       int
	// MaxDuration overrides the default max_duration; nil keeps
	// defaultMaxDuration so existing tests are unaffected.
	MaxDuration *int64
}

// defaultMaxDuration is the finite duration used by tests that do not care
// about max_duration.
const defaultMaxDuration int64 = 120

func Int64Ptr(value int64) *int64 {
	return &value
}

type OpalGroupConfig struct {
	ResourceName          string
	Description           string
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

func generateQuotedArrayString(arr []string) string {
	if len(arr) == 0 {
		return "[]"
	}
	return fmt.Sprintf(`["%s"]`, strings.Join(arr, `", "`))
}

func GenerateErrorMessageRegexp(message string) *regexp.Regexp {
	return regexp.MustCompile(fmt.Sprintf(".*%s.*", message))
}

func GenerateReviewerStages(reviewerStages []ReviewerStageConfig) string {
	if len(reviewerStages) == 0 {
		return ""
	}

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

func GenerateCondition(condition *ConditionConfig) string {
	if condition == nil {
		return ""
	}
	return fmt.Sprintf(
		`condition = {
			group_ids = %s
			remote_role_ids = %s
		}`,
		generateQuotedArrayString(
			condition.GroupIDs,
		),
		generateQuotedArrayString(condition.RemoteRoleIds),
	)
}

func GenerateRequestConfigurations(requestConfigurations []RequestConfigurationConfig) string {
	var configurations []string

	for _, requestConfiguration := range requestConfigurations {
		maxDuration := defaultMaxDuration
		if requestConfiguration.MaxDuration != nil {
			maxDuration = *requestConfiguration.MaxDuration
		}
		configuration := fmt.Sprintf(
			`{
				allow_requests = %t
				auto_approval = %t
				max_duration = %d
				priority = %d
				recommended_duration = %d
				require_mfa_to_request = %t
				require_support_ticket = %t
				reason_optional = %t
				%s
				%s
			}`,
			requestConfiguration.IsRequestable,
			requestConfiguration.AutoApproval,
			maxDuration,
			requestConfiguration.Priority,
			120,
			false,
			false,
			requestConfiguration.ReasonOptional,
			GenerateCondition(requestConfiguration.Condition),
			GenerateReviewerStages(requestConfiguration.ReviewerStages),
		)
		configurations = append(configurations, configuration)
	}

	return "request_configurations = [" + strings.Join(configurations, ",") + "]"
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
			description = "%s"
			app_id = "%s"
			group_type = "%s"
			admin_owner_id = "%s"
			visibility = "%s"
			visibility_group_ids = %s
			message_channel_ids = %s
			on_call_schedule_ids = %s
			%s
			%s
		}
		`,
		opalBaseURL,
		opalToken,
		ogc.ResourceName,
		ogc.Name,
		ogc.Description,
		ogc.AppID,
		ogc.GroupType,
		ogc.AdminOwnerID,
		ogc.Visibility,
		generateQuotedArrayString(ogc.VisibilityGroupIDs),
		generateQuotedArrayString(ogc.MessageChannelIDs),
		generateQuotedArrayString(ogc.OnCallScheduleIDs),
		requestConfigStr,
		ogc.Additional,
	)
	return resourceStr
}
