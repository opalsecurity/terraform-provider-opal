package provider

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"github.com/opalsecurity/terraform-provider-opal/v3/internal/sdk"
	"github.com/opalsecurity/terraform-provider-opal/v3/internal/sdk/models/shared"
)

const stubDataAppID = "32acc112-21ff-4669-91c2-21e27683eaa1"

// TestAppDataSourceReadVisibilityInline drives AppDataSource.Read against a stub
// API to confirm it never calls GET /apps/{id}/visibility and still populates
// visibility from the inline GET /apps/{id} fields for every app type.
func TestAppDataSourceReadVisibilityInline(t *testing.T) {
	t.Parallel()

	const appBody = `{"app_id":"` + stubDataAppID + `","name":"app","description":"d","admin_owner_id":"o1","app_type":"%s","visibility":"%s","import_visibility":"GLOBAL","visibility_group_ids":[]}`

	for _, testCase := range []struct {
		name           string
		appType        string
		visibility     string
		wantVisibility string
	}{
		{
			name:           "okta directory uses inline visibility",
			appType:        "OKTA_DIRECTORY",
			visibility:     "GLOBAL",
			wantVisibility: "GLOBAL",
		},
		{
			name:           "custom app uses inline visibility",
			appType:        "CUSTOM",
			visibility:     "LIMITED",
			wantVisibility: "LIMITED",
		},
		{
			name:           "aws sso uses inline visibility",
			appType:        "AWS_SSO",
			visibility:     "GLOBAL",
			wantVisibility: "GLOBAL",
		},
	} {
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			visibilityGETs := 0
			server := httptest.NewServer(
				http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					w.Header().Set("Content-Type", "application/json")
					switch r.URL.Path {
					case "/apps/" + stubDataAppID:
						_, _ = fmt.Fprintf(w, appBody, testCase.appType, testCase.visibility)
					case "/apps/" + stubDataAppID + "/visibility":
						visibilityGETs++
						w.WriteHeader(http.StatusBadRequest)
						_, _ = w.Write([]byte(`{"message":"should not be called"}`))
					default:
						t.Errorf("unexpected request path %q", r.URL.Path)
						w.WriteHeader(http.StatusNotFound)
					}
				}),
			)
			defer server.Close()

			ds := &AppDataSource{client: sdk.New(
				sdk.WithServerURL(server.URL),
				sdk.WithSecurity(shared.Security{BearerAuth: "test-token"}),
				sdk.WithClient(server.Client()),
			)}

			ctx := context.Background()
			var schemaResp datasource.SchemaResponse
			ds.Schema(ctx, datasource.SchemaRequest{}, &schemaResp)
			objType := schemaResp.Schema.Type().TerraformType(ctx).(tftypes.Object)

			attrs := map[string]tftypes.Value{}
			for name, typ := range objType.AttributeTypes {
				attrs[name] = tftypes.NewValue(typ, nil)
			}
			attrs["id"] = tftypes.NewValue(tftypes.String, stubDataAppID)

			readReq := datasource.ReadRequest{Config: tfsdk.Config{
				Schema: schemaResp.Schema,
				Raw:    tftypes.NewValue(objType, attrs),
			}}
			readResp := datasource.ReadResponse{State: tfsdk.State{
				Schema: schemaResp.Schema,
				Raw:    tftypes.NewValue(objType, nil),
			}}
			ds.Read(ctx, readReq, &readResp)

			if visibilityGETs != 0 {
				t.Fatalf("visibility GET count = %d, want 0", visibilityGETs)
			}
			if readResp.Diagnostics.HasError() {
				t.Fatalf("unexpected diagnostics: %v", readResp.Diagnostics)
			}

			var model AppDataSourceModel
			readResp.Diagnostics.Append(readResp.State.Get(ctx, &model)...)
			if readResp.Diagnostics.HasError() {
				t.Fatalf("reading state: %v", readResp.Diagnostics)
			}
			if model.Type.ValueString() != testCase.appType {
				t.Errorf("type = %q, want %q", model.Type.ValueString(), testCase.appType)
			}
			if model.Visibility.ValueString() != testCase.wantVisibility {
				t.Errorf(
					"visibility = %q, want %q",
					model.Visibility.ValueString(),
					testCase.wantVisibility,
				)
			}
		})
	}
}
