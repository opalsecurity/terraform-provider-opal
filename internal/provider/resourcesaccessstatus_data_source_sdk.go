// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/opal-dev/terraform-provider-opal/internal/provider/types"
	"github.com/opal-dev/terraform-provider-opal/internal/sdk/models/shared"
	"time"
)

func (r *ResourcesAccessStatusDataSourceModel) RefreshFromSharedResourceUserAccessStatus(resp *shared.ResourceUserAccessStatus) {
	if resp != nil {
		if resp.AccessLevel == nil {
			r.AccessLevel = nil
		} else {
			r.AccessLevel = &tfTypes.ResourceAccessLevel{}
			r.AccessLevel.AccessLevelName = types.StringValue(resp.AccessLevel.AccessLevelName)
			r.AccessLevel.AccessLevelRemoteID = types.StringValue(resp.AccessLevel.AccessLevelRemoteID)
		}
		if resp.ExpirationDate != nil {
			r.ExpirationDate = types.StringValue(resp.ExpirationDate.Format(time.RFC3339Nano))
		} else {
			r.ExpirationDate = types.StringNull()
		}
		r.ResourceID = types.StringValue(resp.ResourceID)
		r.Status = types.StringValue(string(resp.Status))
		r.UserID = types.StringValue(resp.UserID)
	}
}
