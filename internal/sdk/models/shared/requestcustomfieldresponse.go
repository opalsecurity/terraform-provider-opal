// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"errors"
	"github.com/opal-dev/terraform-provider-opal/internal/sdk/internal/utils"
)

type FieldValueType string

const (
	FieldValueTypeStr     FieldValueType = "str"
	FieldValueTypeBoolean FieldValueType = "boolean"
)

type FieldValue struct {
	Str     *string
	Boolean *bool

	Type FieldValueType
}

func CreateFieldValueStr(str string) FieldValue {
	typ := FieldValueTypeStr

	return FieldValue{
		Str:  &str,
		Type: typ,
	}
}

func CreateFieldValueBoolean(boolean bool) FieldValue {
	typ := FieldValueTypeBoolean

	return FieldValue{
		Boolean: &boolean,
		Type:    typ,
	}
}

func (u *FieldValue) UnmarshalJSON(data []byte) error {

	str := ""
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = &str
		u.Type = FieldValueTypeStr
		return nil
	}

	boolean := false
	if err := utils.UnmarshalJSON(data, &boolean, "", true, true); err == nil {
		u.Boolean = &boolean
		u.Type = FieldValueTypeBoolean
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u FieldValue) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.Boolean != nil {
		return utils.MarshalJSON(u.Boolean, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type RequestCustomFieldResponse struct {
	FieldName string `json:"field_name"`
	// The type of the custom request field.
	FieldType  RequestTemplateCustomFieldTypeEnum `json:"field_type"`
	FieldValue FieldValue                         `json:"field_value"`
}

func (o *RequestCustomFieldResponse) GetFieldName() string {
	if o == nil {
		return ""
	}
	return o.FieldName
}

func (o *RequestCustomFieldResponse) GetFieldType() RequestTemplateCustomFieldTypeEnum {
	if o == nil {
		return RequestTemplateCustomFieldTypeEnum("")
	}
	return o.FieldType
}

func (o *RequestCustomFieldResponse) GetFieldValue() FieldValue {
	if o == nil {
		return FieldValue{}
	}
	return o.FieldValue
}
