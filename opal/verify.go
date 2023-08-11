package opal

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/opalsecurity/opal-go"
	"github.com/pkg/errors"
)

// validateReviewerConfigDuringCreate validates that when an item is created one of the following is true:
// - a reviewer_stage is defined
// - auto_approve is set to true
// - is_requestable is set to false
// - resource_type is set to a parent resource type (e.g. AWS_ACCOUNT)
// NOTE: We only care that one of these 4 is correct in order for the item to have a valid reviewer config
// without needing to fall back on the default creation behavior which would cause an immediate diff after
// creation
func validateReviewerConfigDuringCreate(d *schema.ResourceData) error {
	if reviewerStagesI, ok := d.GetOk("reviewer_stage"); ok {
		if len(reviewerStagesI.([]any)) > 0 {
			return nil
		}
	}
	if autoApprovalI, ok := d.GetOkExists("auto_approval"); ok {
		if autoApprovalI.(bool) {
			return nil
		}
	}
	if isRequestableI, ok := d.GetOkExists("is_requestable"); ok {
		if !isRequestableI.(bool) {
			return nil
		}
	}
	if resourceTypeI, ok := d.GetOkExists("resource_type"); ok {
		if opal.ResourceTypeEnum(resourceTypeI.(string)) == opal.RESOURCETYPEENUM_AWS_ACCOUNT {
			return nil
		}
	}

	return errors.New("Invalid reviewer configuration. Please specify at least 1 reviewer stage, or set auto_approval to true or set is_requestable to false")
}
