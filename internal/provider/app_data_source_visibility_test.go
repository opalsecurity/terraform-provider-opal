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

const stubAppID = "32acc112-21ff-4669-91c2-21e27683eaa1"

// TestAppDataSourceReadVisibility drives AppDataSource.Read against a stub API
// to confirm apps whose type is rejected by GET /apps/{id}/visibility (400)
// still read successfully with null visibility attributes, while a supported
// app keeps its visibility attributes.
func TestAppDataSourceReadVisibility(t *testing.T) {
	t.Parallel()

	const appBody = `{"app_id":"` + stubAppID + `","name":"app","description":"d","admin_owner_id":"o1","app_type":"%s"}`

	for _, testCase := range []struct {
		name           string
		appType        string
		visibilityCode int
		visibilityBody string
		wantError      bool
		wantVisibility string // empty means null
	}{
		{
			name:           "okta directory visibility 400 leaves visibility null",
			appType:        "OKTA_DIRECTORY",
			visibilityCode: http.StatusBadRequest,
			visibilityBody: `{"message":"app type OKTA_DIRECTORY cannot be managed via this endpoint"}`,
		},
		{
			name:           "custom app populates visibility",
			appType:        "CUSTOM",
			visibilityCode: http.StatusOK,
			visibilityBody: `{"visibility":"GLOBAL","visibility_group_ids":[]}`,
			wantVisibility: "GLOBAL",
		},
		{
			name:           "other visibility errors still fail",
			appType:        "CUSTOM",
			visibilityCode: http.StatusNotFound,
			wantError:      true,
		},
	} {
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				switch r.URL.Path {
				case "/apps/" + stubAppID:
					_, _ = fmt.Fprintf(w, appBody, testCase.appType)
				case "/apps/" + stubAppID + "/visibility":
					w.WriteHeader(testCase.visibilityCode)
					_, _ = w.Write([]byte(testCase.visibilityBody))
				default:
					t.Errorf("unexpected request path %q", r.URL.Path)
					w.WriteHeader(http.StatusNotFound)
				}
			}))
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
			attrs["id"] = tftypes.NewValue(tftypes.String, stubAppID)

			readReq := datasource.ReadRequest{Config: tfsdk.Config{
				Schema: schemaResp.Schema,
				Raw:    tftypes.NewValue(objType, attrs),
			}}
			readResp := datasource.ReadResponse{State: tfsdk.State{
				Schema: schemaResp.Schema,
				Raw:    tftypes.NewValue(objType, nil),
			}}
			ds.Read(ctx, readReq, &readResp)

			if testCase.wantError {
				if !readResp.Diagnostics.HasError() {
					t.Fatalf("expected an error diagnostic, got none")
				}
				return
			}
			if readResp.Diagnostics.HasError() {
				t.Fatalf("unexpected diagnostics: %v", readResp.Diagnostics)
			}

			var gotType, gotVisibility *string
			var model AppDataSourceModel
			readResp.Diagnostics.Append(readResp.State.Get(ctx, &model)...)
			if readResp.Diagnostics.HasError() {
				t.Fatalf("reading state: %v", readResp.Diagnostics)
			}
			if !model.Type.IsNull() {
				v := model.Type.ValueString()
				gotType = &v
			}
			if !model.Visibility.IsNull() {
				v := model.Visibility.ValueString()
				gotVisibility = &v
			}
			if gotType == nil || *gotType != testCase.appType {
				t.Errorf("type = %v, want %q", gotType, testCase.appType)
			}
			switch {
			case testCase.wantVisibility == "" && gotVisibility != nil:
				t.Errorf("visibility = %q, want null", *gotVisibility)
			case testCase.wantVisibility != "" && (gotVisibility == nil || *gotVisibility != testCase.wantVisibility):
				t.Errorf("visibility = %v, want %q", gotVisibility, testCase.wantVisibility)
			}
		})
	}
}
