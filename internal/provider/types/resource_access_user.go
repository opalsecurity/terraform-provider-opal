// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type ResourceAccessUser struct {
	AccessLevel     ResourceAccessLevel `tfsdk:"access_level"`
	Email           types.String        `tfsdk:"email"`
	ExpirationDate  types.String        `tfsdk:"expiration_date"`
	FullName        types.String        `tfsdk:"full_name"`
	HasDirectAccess types.Bool          `tfsdk:"has_direct_access"`
	NumAccessPaths  types.Int64         `tfsdk:"num_access_paths"`
	ResourceID      types.String        `tfsdk:"resource_id"`
	UserID          types.String        `tfsdk:"user_id"`
}
