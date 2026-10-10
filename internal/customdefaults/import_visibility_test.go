package customdefaults_test

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/defaults"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/opalsecurity/terraform-provider-opal/v3/internal/customdefaults"
	"github.com/stretchr/testify/assert"
)

// Test that the exported constructor returns the documented default — "GLOBAL"
// — so a Terraform plan on an `opal_app` resource that omits
// `import_visibility` resolves to the same value the backend's createApp
// handler would pick (`model.VisibilityGlobal` in
// opal.git `handlers_apps.go`). Guards against accidental edits that would
// desynchronize the TF-side default from the API-side default.
func TestImportVisibilityGlobal_DefaultsToGlobal(t *testing.T) {
	ctx := context.Background()

	var resp defaults.StringResponse
	customdefaults.ImportVisibilityGlobal().DefaultString(ctx, defaults.StringRequest{}, &resp)

	assert.Equal(t, types.StringValue("GLOBAL"), resp.PlanValue)
	assert.Empty(t, resp.Diagnostics, "the default must not emit diagnostics")
}

// Plain and markdown descriptions are what the Terraform plugin framework
// surfaces in `terraform plan` output and in IDE hover tooltips over the
// attribute. They need to be non-empty and mention `GLOBAL` so a user who sees
// the default value in a plan can trace it back to this helper.
func TestImportVisibilityGlobal_Descriptions(t *testing.T) {
	ctx := context.Background()
	d := customdefaults.ImportVisibilityGlobal()

	assert.Contains(t, d.Description(ctx), "GLOBAL")
	assert.Contains(t, d.MarkdownDescription(ctx), "GLOBAL")
}
