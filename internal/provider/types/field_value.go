// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type FieldValue struct {
	Str     types.String `queryParam:"inline" tfsdk:"str" tfPlanOnly:"true"`
	Boolean types.Bool   `queryParam:"inline" tfsdk:"boolean" tfPlanOnly:"true"`
}
