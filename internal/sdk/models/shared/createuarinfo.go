// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/opalsecurity/terraform-provider-opal/internal/sdk/internal/utils"
	"time"
)

// CreateUARInfo - Information needed to start a user access review.
type CreateUARInfo struct {
	// The last day for reviewers to complete their access reviews.
	Deadline time.Time `json:"deadline"`
	// The name of the UAR.
	Name                   string  `json:"name"`
	ReminderIncludeManager *bool   `json:"reminder_include_manager,omitempty"`
	ReminderSchedule       []int64 `json:"reminder_schedule,omitempty"`
	// A policy for auto-assigning reviewers. If auto-assignment is on, specific assignments can still be manually adjusted after the access review is started. Default is Manually.
	ReviewerAssignmentPolicy UARReviewerAssignmentPolicyEnum `json:"reviewer_assignment_policy"`
	// A bool representing whether to present a warning when a user is the only reviewer for themself. Default is False.
	SelfReviewAllowed bool `json:"self_review_allowed"`
	// A bool representing whether to send a notification to reviewers when they're assigned a new review. Default is False.
	SendReviewerAssignmentNotification bool `json:"send_reviewer_assignment_notification"`
	// The time zone name (as defined by the IANA Time Zone database) used in the access review deadline and exported audit report. Default is America/Los_Angeles.
	TimeZone string `json:"time_zone"`
	// If set, the access review will only contain resources and groups that match at least one of the filters in scope.
	UarScope *UARScope `json:"uar_scope,omitempty"`
}

func (c CreateUARInfo) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CreateUARInfo) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CreateUARInfo) GetDeadline() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.Deadline
}

func (o *CreateUARInfo) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *CreateUARInfo) GetReminderIncludeManager() *bool {
	if o == nil {
		return nil
	}
	return o.ReminderIncludeManager
}

func (o *CreateUARInfo) GetReminderSchedule() []int64 {
	if o == nil {
		return nil
	}
	return o.ReminderSchedule
}

func (o *CreateUARInfo) GetReviewerAssignmentPolicy() UARReviewerAssignmentPolicyEnum {
	if o == nil {
		return UARReviewerAssignmentPolicyEnum("")
	}
	return o.ReviewerAssignmentPolicy
}

func (o *CreateUARInfo) GetSelfReviewAllowed() bool {
	if o == nil {
		return false
	}
	return o.SelfReviewAllowed
}

func (o *CreateUARInfo) GetSendReviewerAssignmentNotification() bool {
	if o == nil {
		return false
	}
	return o.SendReviewerAssignmentNotification
}

func (o *CreateUARInfo) GetTimeZone() string {
	if o == nil {
		return ""
	}
	return o.TimeZone
}

func (o *CreateUARInfo) GetUarScope() *UARScope {
	if o == nil {
		return nil
	}
	return o.UarScope
}
