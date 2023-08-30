package opal

import "fmt"

func testReviewerStage(op string, requireManagerApproval bool, reviewerIDs ...string) string {
	reviewerString := ""
	for _, reviewerID := range reviewerIDs {
		reviewerString += fmt.Sprintf(`
reviewer {
	id = "%s"
}
`, reviewerID)
	}

	return fmt.Sprintf(`
	request_configuration {
reviewer_stage {
	operator = "%s"
	require_manager_approval = "%t"
	
	%s
}
}`, op, requireManagerApproval, reviewerString)
}
