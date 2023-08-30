package opal

import (
	"context"

	"github.com/opalsecurity/opal-go"
)

// Ptr returns a pointer from a generic value
func Ptr[T any](val T) *T {
	return &val
}

type Condition struct {
	GroupIds []string `json:"group_ids"`
}

const NIL_MINUTES = -1

// parseRequestConfigurationList parses a request_configuration_list from a terraform resource data object
func parseRequestConfigurationList(
	ctx context.Context,
	requestConfigurationListI interface{},
) (*opal.CreateRequestConfigurationInfoList, error) {
	requestConfigurationsArrayI := requestConfigurationListI.([]interface{})
	requestConfigurations := make([]opal.RequestConfiguration, len(requestConfigurationsArrayI))

	for idx, requestConfigurationI := range requestConfigurationsArrayI {
		requestConfiguration, err := parseRequestConfiguration(ctx, requestConfigurationI)
		if err != nil {
			return nil, err
		}
		requestConfigurations[idx] = *requestConfiguration
	}
	r := &opal.CreateRequestConfigurationInfoList{
		RequestConfigurations: requestConfigurations,
	}
	return r, nil
}

// parseRequestConfiguration parses a single request_configuration from a terraform resource data object
func parseRequestConfiguration(
	ctx context.Context,
	requestConfigurationI interface{},
) (*opal.RequestConfiguration, error) {
	requestConfiguration := opal.RequestConfiguration{}
	requestConfigurationMap := requestConfigurationI.(map[string]interface{})

	if prioriryI, ok := requestConfigurationMap["priority"]; ok {
		priorityVal := int32(prioriryI.(int))
		requestConfiguration.Priority = priorityVal
	}

	if allowRequestsI, ok := requestConfigurationMap["is_requestable"]; ok {
		requestConfiguration.AllowRequests = allowRequestsI.(bool)
	}

	if autoApprovalI, ok := requestConfigurationMap["auto_approval"]; ok {
		requestConfiguration.AutoApproval = autoApprovalI.(bool)
	}

	if requireMfaToRequestI, ok := requestConfigurationMap["require_mfa_to_request"]; ok {
		requestConfiguration.RequireMfaToRequest = requireMfaToRequestI.(bool)
	}

	requestConfiguration.MaxDurationMinutes = Ptr(int32(NIL_MINUTES))
	if maxDurationI, ok := requestConfigurationMap["max_duration"]; ok {
		maxDurationVal := int32(maxDurationI.(int))
		if maxDurationVal > 0 {
			requestConfiguration.MaxDurationMinutes = &maxDurationVal
		}
	}

	requestConfiguration.RecommendedDurationMinutes = Ptr(int32(NIL_MINUTES))
	if recommendedDurationI, ok := requestConfigurationMap["recommended_duration"]; ok {
		recommendedDurationVal := int32(recommendedDurationI.(int))
		if recommendedDurationVal > 0 {
			requestConfiguration.RecommendedDurationMinutes = &recommendedDurationVal
		}
	}

	if requireSupportTicketI, ok := requestConfigurationMap["require_support_ticket"]; ok {
		requestConfiguration.RequireSupportTicket = requireSupportTicketI.(bool)
	}

	if requestTemplateIDI, ok := requestConfigurationMap["request_template_id"]; ok {
		requestTemplateVal := requestTemplateIDI.(string)
		if requestTemplateVal != "" {
			requestConfiguration.RequestTemplateId = &requestTemplateVal
		}
	}

	if groupIDsI, ok := requestConfigurationMap["group_ids"]; ok {
		groupIDsArrayI := groupIDsI.([]interface{})
		groupIDs := make([]string, len(groupIDsArrayI))
		for idx, groupIDI := range groupIDsArrayI {
			groupIDs[idx] = groupIDI.(string)
		}
		if len(groupIDs) > 0 {
			requestConfiguration.Condition = &opal.Condition{
				GroupIds: groupIDs,
			}
		}
	}

	if reviewerStagesI, ok := requestConfigurationMap["reviewer_stage"]; ok {
		reviewerStages, err := parseReviewerStages(reviewerStagesI)
		if err != nil {
			return nil, err
		}
		requestConfiguration.ReviewerStages = reviewerStages
	}

	return &requestConfiguration, nil
}

// parseReviewerStages parses a reviewer_stage from a terraform resource data object
func parseReviewerStages(reviewerStagesI any) ([]opal.ReviewerStage, error) {
	rawReviewerStages := reviewerStagesI.([]any)
	reviewerStages := make([]opal.ReviewerStage, 0, len(rawReviewerStages))
	for _, rawReviewerStage := range rawReviewerStages {
		reviewerStage := rawReviewerStage.(map[string]any)
		requireManagerApproval := reviewerStage["require_manager_approval"].(bool)
		// if reviewerStage["require_manager_approval"] == nil {
		// 	requireManagerApproval = false
		// }
		operator := reviewerStage["operator"].(string)
		// if reviewerStage["operator"] == nil {
		// 	operator = "AND"
		// }
		reviewersI := reviewerStage["reviewer"]
		reviewerIds, err := extractReviewerIDs(reviewersI)
		if err != nil {
			return nil, err
		}

		reviewerStages = append(reviewerStages, *opal.NewReviewerStage(requireManagerApproval, operator, reviewerIds))
	}
	return reviewerStages, nil
}

// parseSDKRequestConfiguration parses a request_configuration from the opal-go SDK
func parseSDKRequestConfiguration(
	ctx context.Context,
	requestConfiguration *opal.RequestConfiguration,
) (map[string]interface{}, error) {
	requestConfigurationMap := make(map[string]interface{})

	requestConfigurationMap["priority"] = requestConfiguration.Priority
	requestConfigurationMap["is_requestable"] = requestConfiguration.AllowRequests
	requestConfigurationMap["auto_approval"] = requestConfiguration.AutoApproval
	requestConfigurationMap["require_mfa_to_request"] = requestConfiguration.RequireMfaToRequest
	requestConfigurationMap["require_support_ticket"] = requestConfiguration.RequireSupportTicket

	if requestConfiguration.MaxDurationMinutes != nil && int(*requestConfiguration.MaxDurationMinutes) > 0 {
		requestConfigurationMap["max_duration"] = requestConfiguration.MaxDurationMinutes
	} else {
		requestConfigurationMap["max_duration"] = NIL_MINUTES
	}

	if requestConfiguration.RecommendedDurationMinutes != nil && int(*requestConfiguration.RecommendedDurationMinutes) > 0 {
		requestConfigurationMap["recommended_duration"] = requestConfiguration.RecommendedDurationMinutes
	} else {
		requestConfigurationMap["recommended_duration"] = NIL_MINUTES
	}

	if requestConfiguration.RequestTemplateId != nil {
		requestConfigurationMap["request_template_id"] = *requestConfiguration.RequestTemplateId
	}

	if requestConfiguration.Condition != nil && requestConfiguration.Condition.GroupIds != nil && len(requestConfiguration.Condition.GroupIds) > 0 {
		requestConfigurationMap["group_ids"] = requestConfiguration.Condition.GroupIds
	}

	if requestConfiguration.ReviewerStages != nil {
		reviewerStages, err := parseSDKReviewerStages(requestConfiguration.ReviewerStages)
		if err != nil {
			return nil, err
		}
		requestConfigurationMap["reviewer_stage"] = reviewerStages
	}

	return requestConfigurationMap, nil
}

// parseSDKReviewerStages parses a reviewer_stage from the opal-go SDK
func parseSDKReviewerStages(reviewerStages []opal.ReviewerStage) ([]map[string]interface{}, error) {
	rawReviewerStages := make([]map[string]interface{}, 0, len(reviewerStages))
	for _, reviewerStage := range reviewerStages {
		reviewerStageMap := make(map[string]interface{})
		reviewerStageMap["require_manager_approval"] = reviewerStage.RequireManagerApproval
		reviewerStageMap["operator"] = reviewerStage.Operator
		reviewers := make([]map[string]string, 0, len(reviewerStage.OwnerIds))
		for _, ownerID := range reviewerStage.OwnerIds {
			reviewers = append(reviewers, map[string]string{"id": ownerID})
		}
		reviewerStageMap["reviewer"] = reviewers
		rawReviewerStages = append(rawReviewerStages, reviewerStageMap)
	}
	return rawReviewerStages, nil
}
