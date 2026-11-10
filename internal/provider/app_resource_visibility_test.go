package provider

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"github.com/opalsecurity/terraform-provider-opal/v3/internal/sdk"
	"github.com/opalsecurity/terraform-provider-opal/v3/internal/sdk/models/shared"
)

const stubManagedAppID = "32acc112-21ff-4669-91c2-21e27683eaa1"

// TestAppResourceReadVisibilityTypeGuard drives AppResource.Read against a stub
// API to confirm non-writable app types are rejected before GET
// /apps/{id}/visibility (which would 400), while CUSTOM still calls it.
func TestAppResourceReadVisibilityTypeGuard(t *testing.T) {
	t.Parallel()

	const appBody = `{"app_id":"` + stubManagedAppID + `","name":"app","description":"d","admin_owner_id":"o1","app_type":"%s","visibility":"GLOBAL","import_visibility":"GLOBAL","visibility_group_ids":[]}`

	for _, testCase := range []struct {
		name              string
		appType           string
		wantError         bool
		wantErrorContains string
		wantVisibilityGET bool
		wantVisibility    string
	}{
		{
			name:              "okta directory rejected without visibility GET",
			appType:           "OKTA_DIRECTORY",
			wantError:         true,
			wantErrorContains: "only CUSTOM and CUSTOM_CONNECTOR are supported",
			wantVisibilityGET: false,
		},
		{
			name:              "custom app calls visibility endpoint",
			appType:           "CUSTOM",
			wantVisibilityGET: true,
			wantVisibility:    "LIMITED",
		},
		{
			name:              "custom connector calls visibility endpoint",
			appType:           "CUSTOM_CONNECTOR",
			wantVisibilityGET: true,
			wantVisibility:    "GLOBAL",
		},
	} {
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			visibilityGETs := 0
			server := httptest.NewServer(
				http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					w.Header().Set("Content-Type", "application/json")
					switch r.URL.Path {
					case "/apps/" + stubManagedAppID:
						_, _ = fmt.Fprintf(w, appBody, testCase.appType)
					case "/apps/" + stubManagedAppID + "/visibility":
						visibilityGETs++
						if testCase.appType == "OKTA_DIRECTORY" {
							w.WriteHeader(http.StatusBadRequest)
							_, _ = w.Write(
								[]byte(
									`{"message":"app type OKTA_DIRECTORY cannot be managed via this endpoint"}`,
								),
							)
							return
						}
						_, _ = fmt.Fprintf(
							w,
							`{"visibility":%q,"visibility_group_ids":[],"import_visibility":"GLOBAL"}`,
							testCase.wantVisibility,
						)
					default:
						t.Errorf("unexpected request path %q", r.URL.Path)
						w.WriteHeader(http.StatusNotFound)
					}
				}),
			)
			defer server.Close()

			res := &AppResource{client: sdk.New(
				sdk.WithServerURL(server.URL),
				sdk.WithSecurity(shared.Security{BearerAuth: "test-token"}),
				sdk.WithClient(server.Client()),
			)}

			ctx := context.Background()
			var schemaResp resource.SchemaResponse
			res.Schema(ctx, resource.SchemaRequest{}, &schemaResp)
			objType := schemaResp.Schema.Type().TerraformType(ctx).(tftypes.Object)

			attrs := map[string]tftypes.Value{}
			for name, typ := range objType.AttributeTypes {
				attrs[name] = tftypes.NewValue(typ, nil)
			}
			attrs["id"] = tftypes.NewValue(tftypes.String, stubManagedAppID)

			readReq := resource.ReadRequest{State: tfsdk.State{
				Schema: schemaResp.Schema,
				Raw:    tftypes.NewValue(objType, attrs),
			}}
			readResp := resource.ReadResponse{State: tfsdk.State{
				Schema: schemaResp.Schema,
				Raw:    tftypes.NewValue(objType, nil),
			}}
			res.Read(ctx, readReq, &readResp)

			if testCase.wantVisibilityGET && visibilityGETs != 1 {
				t.Errorf("visibility GET count = %d, want 1", visibilityGETs)
			}
			if !testCase.wantVisibilityGET && visibilityGETs != 0 {
				t.Errorf("visibility GET count = %d, want 0", visibilityGETs)
			}

			if testCase.wantError {
				if !readResp.Diagnostics.HasError() {
					t.Fatalf("expected an error diagnostic, got none")
				}
				errText := fmt.Sprintf("%v", readResp.Diagnostics)
				if !strings.Contains(errText, testCase.wantErrorContains) {
					t.Fatalf(
						"diagnostics %q do not contain %q",
						errText,
						testCase.wantErrorContains,
					)
				}
				return
			}
			if readResp.Diagnostics.HasError() {
				t.Fatalf("unexpected diagnostics: %v", readResp.Diagnostics)
			}

			var model AppResourceModel
			readResp.Diagnostics.Append(readResp.State.Get(ctx, &model)...)
			if readResp.Diagnostics.HasError() {
				t.Fatalf("reading state: %v", readResp.Diagnostics)
			}
			if model.Type.ValueString() != testCase.appType {
				t.Errorf("type = %q, want %q", model.Type.ValueString(), testCase.appType)
			}
			if model.AppType.ValueString() != testCase.appType {
				t.Errorf(
					"app_type = %q, want %q (copied from type on read/import)",
					model.AppType.ValueString(),
					testCase.appType,
				)
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
