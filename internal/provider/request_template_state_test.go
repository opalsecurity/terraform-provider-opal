package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/opalsecurity/terraform-provider-opal/v3/internal/provider/types"
)

func TestRestoreAPIComputedCustomFieldAttrs(t *testing.T) {
	data := &RequestTemplateResourceModel{
		CustomFields: []tfTypes.RequestTemplateCustomField{
			{Name: types.StringValue("Why?"), Required: types.BoolValue(true), Type: types.StringValue("SHORT_TEXT")},
			{Name: types.StringNull(), Required: types.BoolNull(), Type: types.StringValue("CALLOUT")},
		},
	}
	fromAPI := []tfTypes.RequestTemplateCustomField{
		{Name: types.StringValue("Why?"), Required: types.BoolValue(true), Type: types.StringValue("SHORT_TEXT")},
		{Name: types.StringValue("callout-1"), Required: types.BoolValue(false), Type: types.StringValue("CALLOUT")},
	}
	restoreAPIComputedCustomFieldAttrs(fromAPI, data)
	if got := data.CustomFields[1].Name.ValueString(); got != "callout-1" {
		t.Fatalf("name: got %q", got)
	}
	if data.CustomFields[1].Required.IsNull() || data.CustomFields[1].Required.ValueBool() {
		t.Fatalf("required: got %v", data.CustomFields[1].Required)
	}
	if data.CustomFields[0].Name.ValueString() != "Why?" {
		t.Fatal("should not change explicit names")
	}
}
