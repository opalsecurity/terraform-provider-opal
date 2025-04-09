// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// RequestStatusEnum - # Request Status
// ### Description
// The `RequestStatus` enum is used to represent the status of a request.
//
// ### Usage Example
// Returned from the `GET Requests` endpoint.
type RequestStatusEnum string

const (
	RequestStatusEnumPending  RequestStatusEnum = "PENDING"
	RequestStatusEnumApproved RequestStatusEnum = "APPROVED"
	RequestStatusEnumDenied   RequestStatusEnum = "DENIED"
	RequestStatusEnumCanceled RequestStatusEnum = "CANCELED"
)

func (e RequestStatusEnum) ToPointer() *RequestStatusEnum {
	return &e
}
func (e *RequestStatusEnum) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "PENDING":
		fallthrough
	case "APPROVED":
		fallthrough
	case "DENIED":
		fallthrough
	case "CANCELED":
		*e = RequestStatusEnum(v)
		return nil
	default:
		return fmt.Errorf("invalid value for RequestStatusEnum: %v", v)
	}
}
