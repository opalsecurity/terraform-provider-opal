// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// # App Object
// ### Description
// The `App` object is used to represent an app to an application.
//
// ### Usage Example
// List from the `GET Apps` endpoint.
type App struct {
	// The ID of the owner of the app.
	AdminOwnerID string `json:"admin_owner_id"`
	// A description of the app.
	Description string `json:"description"`
	// The ID of the app.
	ID string `json:"app_id"`
	// The name of the app.
	Name string `json:"name"`
	// The type of an app.
	Type string `json:"app_type"`
	// Validation checks of an apps' configuration and permissions.
	Validations []AppValidation `json:"validations,omitempty"`
}

func (o *App) GetAdminOwnerID() string {
	if o == nil {
		return ""
	}
	return o.AdminOwnerID
}

func (o *App) GetDescription() string {
	if o == nil {
		return ""
	}
	return o.Description
}

func (o *App) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *App) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *App) GetType() string {
	if o == nil {
		return ""
	}
	return o.Type
}

func (o *App) GetValidations() []AppValidation {
	if o == nil {
		return nil
	}
	return o.Validations
}
