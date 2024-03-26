// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/opal-dev/terraform-provider-opal/internal/sdk/internal/utils"
)

// Operator - The operator of the reviewer stage.
type Operator string

const (
	OperatorAnd Operator = "AND"
	OperatorOr  Operator = "OR"
)

func (e Operator) ToPointer() *Operator {
	return &e
}

func (e *Operator) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "AND":
		fallthrough
	case "OR":
		*e = Operator(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Operator: %v", v)
	}
}

// ReviewerStage - A reviewer stage.
type ReviewerStage struct {
	// The operator of the reviewer stage.
	Operator *Operator `default:"AND" json:"operator"`
	OwnerIds []string  `json:"owner_ids"`
	// Whether this reviewer stage should require manager approval.
	RequireManagerApproval bool `json:"require_manager_approval"`
}

func (r ReviewerStage) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(r, "", false)
}

func (r *ReviewerStage) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &r, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ReviewerStage) GetOperator() *Operator {
	if o == nil {
		return nil
	}
	return o.Operator
}

func (o *ReviewerStage) GetOwnerIds() []string {
	if o == nil {
		return []string{}
	}
	return o.OwnerIds
}

func (o *ReviewerStage) GetRequireManagerApproval() bool {
	if o == nil {
		return false
	}
	return o.RequireManagerApproval
}
