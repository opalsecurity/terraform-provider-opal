package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/opalsecurity/terraform-provider-opal/v3/internal/provider/types"
)

func TestValidateRequestTemplateCustomField_calloutRequiredForbidden(t *testing.T) {
	var diags diag.Diagnostics
	seen := map[string]int{}
	validateRequestTemplateCustomField(0, tfTypes.RequestTemplateCustomField{
		Type:     types.StringValue("CALLOUT"),
		Required: types.BoolValue(true),
		Metadata: &tfTypes.RequestTemplateCustomFieldMetadata{
			CalloutData: &tfTypes.RequestTemplateCustomFieldCalloutMetadata{
				Severity: types.StringValue("WARNING"),
				Text:     types.StringValue("Production access"),
			},
		},
	}, seen, &diags)
	if !diags.HasError() {
		t.Fatal("expected error for required=true on CALLOUT")
	}
}

func TestValidateRequestTemplateCustomField_calloutRequiresText(t *testing.T) {
	var diags diag.Diagnostics
	seen := map[string]int{}
	validateRequestTemplateCustomField(0, tfTypes.RequestTemplateCustomField{
		Type: types.StringValue("CALLOUT"),
	}, seen, &diags)
	if !diags.HasError() {
		t.Fatal("expected error for CALLOUT without callout_data")
	}
}

func TestValidateRequestTemplateCustomField_answerableRequiresName(t *testing.T) {
	var diags diag.Diagnostics
	seen := map[string]int{}
	validateRequestTemplateCustomField(0, tfTypes.RequestTemplateCustomField{
		Type: types.StringValue("SHORT_TEXT"),
	}, seen, &diags)
	if !diags.HasError() {
		t.Fatal("expected error for SHORT_TEXT without name")
	}
}

func TestValidateRequestTemplateCustomField_calloutNameOptional(t *testing.T) {
	var diags diag.Diagnostics
	seen := map[string]int{}
	validateRequestTemplateCustomField(0, tfTypes.RequestTemplateCustomField{
		Type: types.StringValue("CALLOUT"),
		Metadata: &tfTypes.RequestTemplateCustomFieldMetadata{
			CalloutData: &tfTypes.RequestTemplateCustomFieldCalloutMetadata{
				Severity: types.StringValue("INFO"),
				Text:     types.StringValue("Heads up"),
			},
		},
	}, seen, &diags)
	if diags.HasError() {
		t.Fatalf("unexpected errors: %v", diags.Errors())
	}
}

func TestValidateRequestTemplateCustomField_dropdownRequiresOptions(t *testing.T) {
	var diags diag.Diagnostics
	seen := map[string]int{}
	validateRequestTemplateCustomField(0, tfTypes.RequestTemplateCustomField{
		Name: types.StringValue("Reason"),
		Type: types.StringValue("MULTI_CHOICE"),
	}, seen, &diags)
	if !diags.HasError() {
		t.Fatal("expected error for MULTI_CHOICE without options")
	}
}

func TestValidateRequestTemplateCustomField_duplicateNames(t *testing.T) {
	var diags diag.Diagnostics
	seen := map[string]int{}
	field := tfTypes.RequestTemplateCustomField{
		Name: types.StringValue("Reason"),
		Type: types.StringValue("SHORT_TEXT"),
	}
	validateRequestTemplateCustomField(0, field, seen, &diags)
	validateRequestTemplateCustomField(1, field, seen, &diags)
	if !diags.HasError() {
		t.Fatal("expected duplicate name error")
	}
}
