// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type UARScope struct {
	Admins []types.String `tfsdk:"admins"`
	Names  []types.String `tfsdk:"names"`
	Tags   []TagFilter    `tfsdk:"tags"`
}