package listvalidators

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestRequestConfigurationsRejectsEmptyList(t *testing.T) {
	ctx := context.Background()
	elemType := types.ObjectType{AttrTypes: map[string]attr.Type{}}
	v := RequestConfigurations()

	cases := []struct {
		name      string
		value     types.List
		wantError bool
	}{
		{
			name:      "omitted",
			value:     types.ListNull(elemType),
			wantError: false,
		},
		{
			name:      "unknown",
			value:     types.ListUnknown(elemType),
			wantError: false,
		},
		{
			name:      "empty",
			value:     types.ListValueMust(elemType, []attr.Value{}),
			wantError: true,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			resp := &validator.ListResponse{}
			v.ValidateList(ctx, validator.ListRequest{ConfigValue: tc.value}, resp)
			gotError := resp.Diagnostics.HasError()
			if gotError != tc.wantError {
				t.Fatalf("HasError=%v, want %v; diags=%v", gotError, tc.wantError, resp.Diagnostics)
			}
			if tc.wantError && !resp.Diagnostics.Contains(resp.Diagnostics[0]) {
				t.Fatal("expected an Invalid Attribute Value diagnostic")
			}
			if tc.wantError {
				summary := resp.Diagnostics[0].Summary()
				if summary != "Invalid Attribute Value" {
					t.Fatalf("summary=%q, want Invalid Attribute Value", summary)
				}
			}
		})
	}
}
