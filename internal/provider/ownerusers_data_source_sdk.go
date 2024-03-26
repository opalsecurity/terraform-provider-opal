// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/opal-dev/terraform-provider-opal/internal/provider/types"
	"github.com/opal-dev/terraform-provider-opal/internal/sdk/models/shared"
)

func (r *OwnerUsersDataSourceModel) RefreshFromSharedUserList(resp *shared.UserList) {
	if resp != nil {
		if len(r.Users) > len(resp.Users) {
			r.Users = r.Users[:len(resp.Users)]
		}
		for usersCount, usersItem := range resp.Users {
			var users1 tfTypes.User
			users1.Email = types.StringValue(usersItem.Email)
			users1.FirstName = types.StringValue(usersItem.FirstName)
			if usersItem.HrIdpStatus != nil {
				users1.HrIdpStatus = types.StringValue(string(*usersItem.HrIdpStatus))
			} else {
				users1.HrIdpStatus = types.StringNull()
			}
			users1.ID = types.StringPointerValue(usersItem.ID)
			users1.LastName = types.StringValue(usersItem.LastName)
			users1.Name = types.StringPointerValue(usersItem.Name)
			users1.Position = types.StringValue(usersItem.Position)
			if usersCount+1 > len(r.Users) {
				r.Users = append(r.Users, users1)
			} else {
				r.Users[usersCount].Email = users1.Email
				r.Users[usersCount].FirstName = users1.FirstName
				r.Users[usersCount].HrIdpStatus = users1.HrIdpStatus
				r.Users[usersCount].ID = users1.ID
				r.Users[usersCount].LastName = users1.LastName
				r.Users[usersCount].Name = users1.Name
				r.Users[usersCount].Position = users1.Position
			}
		}
	}
}