// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type AwsEc2Instance struct {
	AccountID  types.String `tfsdk:"account_id"`
	InstanceID types.String `tfsdk:"instance_id"`
	Region     types.String `tfsdk:"region"`
}