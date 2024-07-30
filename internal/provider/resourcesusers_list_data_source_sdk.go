// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/opalsecurity/terraform-provider-opal/internal/provider/types"
	"github.com/opalsecurity/terraform-provider-opal/internal/sdk/models/shared"
	"time"
)

func (r *ResourcesUsersListDataSourceModel) RefreshFromSharedResourceAccessUserList(resp *shared.ResourceAccessUserList) {
	if resp != nil {
		r.Results = []tfTypes.ResourceAccessUser{}
		if len(r.Results) > len(resp.Results) {
			r.Results = r.Results[:len(resp.Results)]
		}
		for resultsCount, resultsItem := range resp.Results {
			var results1 tfTypes.ResourceAccessUser
			results1.AccessLevel.AccessLevelName = types.StringValue(resultsItem.AccessLevel.AccessLevelName)
			results1.AccessLevel.AccessLevelRemoteID = types.StringValue(resultsItem.AccessLevel.AccessLevelRemoteID)
			results1.Email = types.StringValue(resultsItem.Email)
			if resultsItem.ExpirationDate != nil {
				results1.ExpirationDate = types.StringValue(resultsItem.ExpirationDate.Format(time.RFC3339Nano))
			} else {
				results1.ExpirationDate = types.StringNull()
			}
			results1.FullName = types.StringValue(resultsItem.FullName)
			results1.HasDirectAccess = types.BoolValue(resultsItem.HasDirectAccess)
			results1.NumAccessPaths = types.Int64Value(int64(resultsItem.NumAccessPaths))
			results1.ResourceID = types.StringValue(resultsItem.ResourceID)
			results1.UserID = types.StringValue(resultsItem.UserID)
			if resultsCount+1 > len(r.Results) {
				r.Results = append(r.Results, results1)
			} else {
				r.Results[resultsCount].AccessLevel = results1.AccessLevel
				r.Results[resultsCount].Email = results1.Email
				r.Results[resultsCount].ExpirationDate = results1.ExpirationDate
				r.Results[resultsCount].FullName = results1.FullName
				r.Results[resultsCount].HasDirectAccess = results1.HasDirectAccess
				r.Results[resultsCount].NumAccessPaths = results1.NumAccessPaths
				r.Results[resultsCount].ResourceID = results1.ResourceID
				r.Results[resultsCount].UserID = results1.UserID
			}
		}
	}
}
