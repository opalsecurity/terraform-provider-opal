package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/opalsecurity/terraform-provider-opal/v3/internal/provider/types"
)

func TestCalloutText(t *testing.T) {
	got := calloutText(tfTypes.RequestTemplateCustomField{
		Metadata: &tfTypes.RequestTemplateCustomFieldMetadata{
			CalloutData: &tfTypes.RequestTemplateCustomFieldCalloutMetadata{
				Text: types.StringValue("Heads up"),
			},
		},
	})
	if got != "Heads up" {
		t.Fatalf("got %q", got)
	}
	if calloutText(tfTypes.RequestTemplateCustomField{}) != "" {
		t.Fatal("expected empty for missing metadata")
	}
}
