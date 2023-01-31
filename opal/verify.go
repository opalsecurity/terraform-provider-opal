package opal

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"
)

// ignoreReviewerDefaultValue is a DiffSuppressFunc that lets us correctly
// handle a missing, optional reviewer block.
func ignoreReviewerDefaultValue(k, oldValue, newValue string, d *schema.ResourceData) bool {
	// If the list of reviewers went from 0 to 1, we don't want to force a diff iff
	// the oldValue is the same as the admin_owner_id. This covers the case where
	// the user omits the reviewer block and lets the API compute the value. Without
	// this code, there will be a permanent (no-op) diff where terraform constantly
	// tries to delete the reviewer.
	if oldValue == "1" && newValue == "0" {
		// The case where we go from 1 -> 0 is where the length of the set goes from
		// 1 to 0. To actually check if we're in the case outlined above, we need to
		// query for the actual values (instead of the lengths).
		oldI, _ := d.GetChange("reviewer")
		old := oldI.(*schema.Set)
		// This should be case to index into because the schema validation would have complained
		// before we get here.
		if old.List()[0].(map[string]any)["id"] == d.Get("admin_owner_id") {
			return true
		}
	}

	return false
}

// validateReviewerConfigDuringCreate validates that when an item is created one of the following is true:
// - a reviewer_stage is defined
// - auto_approve is set to true
// - is_requestable is set to false
// NOTE: We only care that one of these 3 is correct in order for the item to have a valid reviewer config
// without needing to fall back on the default creation behavior which would cause an immediate diff after
// creation
func validateReviewerConfigDuringCreate(d *schema.ResourceData) error {
	if reviewerStagesI, ok := d.GetOk("reviewer_stage"); ok {
		if reviewerStagesI.(*schema.Set).Len() > 0 {
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

	return errors.New("Invalid reviewer configuration. Please specify at least 1 reviewer stage, or set auto_approval to true or set is_requestable to false")
}
