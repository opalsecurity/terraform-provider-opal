package opal

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/opalsecurity/opal-go"
	"github.com/pkg/errors"
)

// validateReviewerConfigDuringCreate validates that when an item is created one of the following is true:
// - resource_type is set to a parent resource type (e.g. AWS_ACCOUNT)
// - request_configuration is set to a non-empty list and is valid satisfying one of the following:
//   - a reviewer_stage is defined
//   - auto_approve is set to true
//   - is_requestable is set to false
//
// NOTE: We only care that one of these 4 is correct in order for the item to have a valid reviewer config
// without needing to fall back on the default creation behavior which would cause an immediate diff after
// creation
func validateReviewerConfigDuringCreate(d *schema.ResourceData) error {
	if resourceTypeI, ok := d.GetOkExists("resource_type"); ok {
		if opal.ResourceTypeEnum(resourceTypeI.(string)) == opal.RESOURCETYPEENUM_AWS_ACCOUNT {
			return nil
		}
	}
	if requestConfigurationListI, ok := d.GetOkExists("request_configuration"); ok {
		if len(requestConfigurationListI.([]interface{})) > 0 {
			return nil
		}
	}

	return errors.New("Invalid reviewer configuration. Please specify a request_configuration block if the resource is not an AWS_ACCOUNT.")
}

func validateRequestConfigurationListDuringCreate(ctx context.Context, d *schema.ResourceData) error {
	_, autoApprovalOk := d.GetOkExists("auto_approval")
	_, requireMfaToRequestOk := d.GetOkExists("require_mfa_to_request")
	_, requireSupportTicketOk := d.GetOkExists("require_support_ticket")
	_, maxDurationOk := d.GetOk("max_duration")
	_, recommendedDurationOk := d.GetOk("recommended_duration")
	_, requestTemplateIDOk := d.GetOk("request_template_id")
	requestConfigurationListI, requestConfigurationListOk := d.GetOk("request_configuration")
	oldRequestConfigurationFieldsChanged := autoApprovalOk || requireMfaToRequestOk || requireSupportTicketOk || maxDurationOk || recommendedDurationOk || requestTemplateIDOk

	if requestConfigurationListOk {
		if oldRequestConfigurationFieldsChanged {
			return errors.New("Cannot set both request_configuration and any of auto_approval, require_mfa_to_request, require_support_ticket, is_requestable, max_duration, recommended_duration, or request_template_id.")
		}
		if len(requestConfigurationListI.([]interface{})) < 1 {
			return errors.New("Invalid request configuration list. Please specify at least 1 request configuration")
		}

		requestConfigurationCreateInfoList, err := parseRequestConfigurationList(ctx, requestConfigurationListI)
		if err != nil {
			return err
		}

		for _, requestConfiguration := range requestConfigurationCreateInfoList.RequestConfigurations {
			// verify priority
			if requestConfiguration.Priority != 0 && requestConfiguration.Condition == nil {
				return errors.New("non-default request configurations must have a condition")
			} else if requestConfiguration.Priority == 0 {
				if requestConfiguration.Condition != nil {
					return errors.New("default request configurations cannot have a condition")
				}
			}

			// 	- one of these is true:
			//   - a reviewer_stage is defined
			//   - auto_approve is set to true
			//   - is_requestable is set to false
			hasReviewerStages, isAutoApprove, isNotRequestable := len(requestConfiguration.ReviewerStages) > 0, requestConfiguration.AutoApproval, !requestConfiguration.AllowRequests

			if !(hasReviewerStages || isAutoApprove || isNotRequestable) {
				return errors.New("invalid request configuration. Please specify a reviewer_stage, set auto_approve to true, or set is_requestable to false")
			}

			for _, reviewerStage := range requestConfiguration.ReviewerStages {
				// validate operator
				if reviewerStage.Operator != "AND" && reviewerStage.Operator != "OR" {
					return errors.New("invalid operator, must be \"AND\" or \"OR\"")
				}
			}
		}
	}

	return nil
}
