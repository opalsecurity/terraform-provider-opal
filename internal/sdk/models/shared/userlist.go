// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// UserList - A list of users.
type UserList struct {
	Users []User `json:"users"`
}

func (o *UserList) GetUsers() []User {
	if o == nil {
		return []User{}
	}
	return o.Users
}
