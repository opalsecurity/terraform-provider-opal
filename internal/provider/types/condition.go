// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type Condition struct {
	GroupIds      []types.String `tfsdk:"group_ids"`
	RoleRemoteIds []types.String `tfsdk:"role_remote_ids"`
}
