// Package customdefaults supplies Terraform plan default values referenced from
// the Speakeasy overlay via `x-speakeasy-terraform-custom-default`. Each
// exported constructor returns a `defaults.*` implementation whose only job is
// to mirror a backend-side default so Terraform's plan phase can resolve it
// before apply (otherwise Computed+Optional fields that the resource generator
// never writes back from the API response stay Unknown through apply and
// Terraform errors with "provider returned invalid result object after apply").
package customdefaults

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/defaults"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ImportVisibilityGlobal defaults `opal_app.import_visibility` to `"GLOBAL"`
// when the user omits it. The backend's createApp handler already uses the
// same default (see opal.git `web/backend/router/handlers/public/handlers_apps.go`:
// `importVisibility := model.VisibilityGlobal`) and the OpenAPI docstring on
// `CreateAppInfo.import_visibility` documents it, so this preserves the
// existing API contract — the only change is where the default is resolved
// (plan-time in the provider, instead of request-time on the backend).
func ImportVisibilityGlobal() defaults.String {
	return importVisibilityGlobal{}
}

type importVisibilityGlobal struct{}

func (importVisibilityGlobal) Description(context.Context) string {
	return "Defaults import_visibility to GLOBAL, matching the API default."
}

func (importVisibilityGlobal) MarkdownDescription(context.Context) string {
	return "Defaults `import_visibility` to `GLOBAL`, matching the API default."
}

func (importVisibilityGlobal) DefaultString(
	_ context.Context,
	_ defaults.StringRequest,
	resp *defaults.StringResponse,
) {
	resp.PlanValue = types.StringValue("GLOBAL")
}
