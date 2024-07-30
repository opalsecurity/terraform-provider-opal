// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type User struct {
	Email       types.String `tfsdk:"email"`
	FirstName   types.String `tfsdk:"first_name"`
	HrIdpStatus types.String `tfsdk:"hr_idp_status"`
	ID          types.String `tfsdk:"id"`
	LastName    types.String `tfsdk:"last_name"`
	Name        types.String `tfsdk:"name"`
	Position    types.String `tfsdk:"position"`
}
