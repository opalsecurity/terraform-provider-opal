package provider

import (
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

func TestConfigurationTemplateIDConfiguredInConfig(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		value tftypes.Value
		want  bool
	}{
		{
			name:  "missing",
			value: tftypes.Value{},
			want:  false,
		},
		{
			name:  "null",
			value: tftypes.NewValue(tftypes.String, nil),
			want:  false,
		},
		{
			name:  "empty string",
			value: tftypes.NewValue(tftypes.String, ""),
			want:  false,
		},
		{
			name:  "literal",
			value: tftypes.NewValue(tftypes.String, "00000000-0000-0000-0000-000000000003"),
			want:  true,
		},
		{
			name:  "unknown",
			value: tftypes.NewValue(tftypes.String, tftypes.UnknownValue),
			want:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			config := map[string]tftypes.Value{}
			if tt.value.Type() != nil {
				config["configuration_template_id"] = tt.value
			}
			if got := configurationTemplateIDConfiguredInConfig(config); got != tt.want {
				t.Fatalf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidateConfigurationTemplateConflicts(t *testing.T) {
	t.Parallel()

	objectType := tftypes.Object{
		AttributeTypes: map[string]tftypes.Type{
			"configuration_template_id": tftypes.String,
			"admin_owner_id":            tftypes.String,
			"visibility":                tftypes.String,
		},
	}

	tests := []struct {
		name       string
		attrs      map[string]tftypes.Value
		wantErrors int
		wantSubstr string
	}{
		{
			name: "literal template and literal admin_owner_id",
			attrs: map[string]tftypes.Value{
				"configuration_template_id": tftypes.NewValue(
					tftypes.String,
					"00000000-0000-0000-0000-000000000003",
				),
				"admin_owner_id": tftypes.NewValue(
					tftypes.String,
					"00000000-0000-0000-0000-000000000002",
				),
				"visibility": tftypes.NewValue(tftypes.String, nil),
			},
			wantErrors: 1,
			wantSubstr: `Attribute "admin_owner_id" cannot be specified`,
		},
		{
			name: "unknown template and literal admin_owner_id",
			attrs: map[string]tftypes.Value{
				"configuration_template_id": tftypes.NewValue(tftypes.String, tftypes.UnknownValue),
				"admin_owner_id": tftypes.NewValue(
					tftypes.String,
					"00000000-0000-0000-0000-000000000002",
				),
				"visibility": tftypes.NewValue(tftypes.String, nil),
			},
			wantErrors: 1,
			wantSubstr: `Attribute "admin_owner_id" cannot be specified`,
		},
		{
			name: "literal template and unknown admin_owner_id",
			attrs: map[string]tftypes.Value{
				"configuration_template_id": tftypes.NewValue(
					tftypes.String,
					"00000000-0000-0000-0000-000000000003",
				),
				"admin_owner_id": tftypes.NewValue(tftypes.String, tftypes.UnknownValue),
				"visibility":     tftypes.NewValue(tftypes.String, nil),
			},
			wantErrors: 1,
			wantSubstr: `Attribute "admin_owner_id" cannot be specified`,
		},
		{
			name: "unknown template and unknown admin_owner_id",
			attrs: map[string]tftypes.Value{
				"configuration_template_id": tftypes.NewValue(tftypes.String, tftypes.UnknownValue),
				"admin_owner_id":            tftypes.NewValue(tftypes.String, tftypes.UnknownValue),
				"visibility":                tftypes.NewValue(tftypes.String, nil),
			},
			wantErrors: 1,
			wantSubstr: `Attribute "admin_owner_id" cannot be specified`,
		},
		{
			name: "template only — no conflict",
			attrs: map[string]tftypes.Value{
				"configuration_template_id": tftypes.NewValue(
					tftypes.String,
					"00000000-0000-0000-0000-000000000003",
				),
				"admin_owner_id": tftypes.NewValue(tftypes.String, nil),
				"visibility":     tftypes.NewValue(tftypes.String, nil),
			},
			wantErrors: 0,
		},
		{
			name: "unknown template only — no conflict",
			attrs: map[string]tftypes.Value{
				"configuration_template_id": tftypes.NewValue(tftypes.String, tftypes.UnknownValue),
				"admin_owner_id":            tftypes.NewValue(tftypes.String, nil),
				"visibility":                tftypes.NewValue(tftypes.String, nil),
			},
			wantErrors: 0,
		},
		{
			name: "no template — admin_owner_id allowed",
			attrs: map[string]tftypes.Value{
				"configuration_template_id": tftypes.NewValue(tftypes.String, nil),
				"admin_owner_id": tftypes.NewValue(
					tftypes.String,
					"00000000-0000-0000-0000-000000000002",
				),
				"visibility": tftypes.NewValue(tftypes.String, "GLOBAL"),
			},
			wantErrors: 0,
		},
		{
			name: "unknown template and visibility",
			attrs: map[string]tftypes.Value{
				"configuration_template_id": tftypes.NewValue(tftypes.String, tftypes.UnknownValue),
				"admin_owner_id":            tftypes.NewValue(tftypes.String, nil),
				"visibility":                tftypes.NewValue(tftypes.String, "GLOBAL"),
			},
			wantErrors: 1,
			wantSubstr: `Attribute "visibility" cannot be specified`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			var diagnostics diag.Diagnostics
			validateConfigurationTemplateConflicts(
				tftypes.NewValue(objectType, tt.attrs),
				[]string{"admin_owner_id", "visibility"},
				&diagnostics,
			)

			if diagnostics.HasError() != (tt.wantErrors > 0) {
				t.Fatalf("HasError=%v diagnostics=%v", diagnostics.HasError(), diagnostics)
			}
			errorCount := 0
			for _, d := range diagnostics {
				if d.Severity() == diag.SeverityError {
					errorCount++
				}
			}
			if errorCount != tt.wantErrors {
				t.Fatalf("got %d errors, want %d: %v", errorCount, tt.wantErrors, diagnostics)
			}
			if tt.wantSubstr != "" {
				joined := diagnostics.Errors()[0].Detail()
				if !strings.Contains(joined, tt.wantSubstr) {
					t.Fatalf("detail %q does not contain %q", joined, tt.wantSubstr)
				}
			}
		})
	}
}
