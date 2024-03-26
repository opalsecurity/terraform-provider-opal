// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/opal-dev/terraform-provider-opal/internal/sdk/models/shared"
)

func (r *AppDataSourceModel) RefreshFromSharedApp(resp *shared.App) {
	if resp != nil {
		r.AdminOwnerID = types.StringValue(resp.AdminOwnerID)
		r.Description = types.StringValue(resp.Description)
		r.ID = types.StringPointerValue(resp.ID)
		r.Name = types.StringValue(resp.Name)
		r.Type = types.StringPointerValue(resp.Type)
	}
}
