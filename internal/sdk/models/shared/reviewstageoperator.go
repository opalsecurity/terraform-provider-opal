// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// ReviewStageOperator - The operator to apply to reviewers in a stage
type ReviewStageOperator string

const (
	ReviewStageOperatorAnd ReviewStageOperator = "AND"
	ReviewStageOperatorOr  ReviewStageOperator = "OR"
)

func (e ReviewStageOperator) ToPointer() *ReviewStageOperator {
	return &e
}
func (e *ReviewStageOperator) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "AND":
		fallthrough
	case "OR":
		*e = ReviewStageOperator(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ReviewStageOperator: %v", v)
	}
}
