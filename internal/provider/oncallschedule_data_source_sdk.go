// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/opal-dev/terraform-provider-opal/internal/sdk/models/shared"
)

func (r *OnCallScheduleDataSourceModel) RefreshFromSharedOnCallSchedule(resp *shared.OnCallSchedule) {
	if resp != nil {
		r.ID = types.StringPointerValue(resp.ID)
		r.Name = types.StringPointerValue(resp.Name)
		r.RemoteID = types.StringPointerValue(resp.RemoteID)
		if resp.ThirdPartyProvider != nil {
			r.ThirdPartyProvider = types.StringValue(string(*resp.ThirdPartyProvider))
		} else {
			r.ThirdPartyProvider = types.StringNull()
		}
	}
}
