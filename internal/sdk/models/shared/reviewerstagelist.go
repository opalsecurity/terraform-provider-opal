// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type ReviewerStageList struct {
	// A list of reviewer stages.
	Stages []ReviewerStage `json:"stages"`
}

func (o *ReviewerStageList) GetStages() []ReviewerStage {
	if o == nil {
		return []ReviewerStage{}
	}
	return o.Stages
}
