package provider

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"strings"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/opalsecurity/terraform-provider-opal/v3/internal/sdk"
)

// These tests cover EPRD-3572: a 403 (forbidden) from a Read()-time sub-call
// (visibility, message channels, on-call schedules) must not be treated the
// same as a 404 (gone). It should surface as a warning and leave that
// specific attribute at its last-known state value, while the rest of the
// resource/group is read normally and stays in state. A genuine 404 on any
// of these sub-calls should still remove the resource/group from state, same
// as before opalsecurity/opal#29350.

func jsonHandler(t *testing.T, status int, body any) http.HandlerFunc {
	t.Helper()
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		if body != nil {
			if err := json.NewEncoder(w).Encode(body); err != nil {
				t.Fatalf("failed to encode response body: %v", err)
			}
		}
	}
}

func emptyHandler(status int) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(status)
	}
}

func newTestSDK(serverURL string) *sdk.OpalAPI {
	return sdk.New(sdk.WithServerURL(serverURL))
}

func TestResourceResourceRead_VisibilityForbidden(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/resources/resource-1", jsonHandler(t, 200, map[string]any{
		"resource_id": "resource-1",
		"name":        "my-resource",
	}))
	mux.HandleFunc("/resources/resource-1/visibility", emptyHandler(http.StatusForbidden))
	ts := httptest.NewServer(mux)
	defer ts.Close()

	ctx := context.Background()
	r := &ResourceResource{client: newTestSDK(ts.URL)}

	var schemaResp resource.SchemaResponse
	r.Schema(ctx, resource.SchemaRequest{}, &schemaResp)
	if schemaResp.Diagnostics.HasError() {
		t.Fatalf("unexpected diagnostics building schema: %v", schemaResp.Diagnostics)
	}

	initial := &ResourceResourceModel{
		ID:                 types.StringValue("resource-1"),
		Visibility:         types.StringValue("LIMITED"),
		VisibilityGroupIds: []types.String{types.StringValue("group-1")},
	}

	state := tfsdk.State{Schema: schemaResp.Schema}
	if diags := state.Set(ctx, initial); diags.HasError() {
		t.Fatalf("unexpected diagnostics seeding state: %v", diags)
	}

	req := resource.ReadRequest{State: state}
	resp := &resource.ReadResponse{State: state}

	r.Read(ctx, req, resp)

	if resp.Diagnostics.HasError() {
		t.Fatalf("Read() returned unexpected error diagnostics: %v", resp.Diagnostics)
	}
	if !hasWarningContaining(resp.Diagnostics, "Unable to read resource visibility") {
		t.Errorf("expected a warning about being unable to read resource visibility, got: %v", resp.Diagnostics)
	}
	if resp.State.Raw.IsNull() {
		t.Fatal("expected resource to remain in state on a 403, but it was removed")
	}

	var result ResourceResourceModel
	if diags := resp.State.Get(ctx, &result); diags.HasError() {
		t.Fatalf("unexpected diagnostics reading back result state: %v", diags)
	}

	if result.Name.ValueString() != "my-resource" {
		t.Errorf("expected name to be populated from the main GET, got %q", result.Name.ValueString())
	}
	if result.Visibility.ValueString() != "LIMITED" {
		t.Errorf("expected visibility to keep its last-known value %q, got %q", "LIMITED", result.Visibility.ValueString())
	}
	if len(result.VisibilityGroupIds) != 1 || result.VisibilityGroupIds[0].ValueString() != "group-1" {
		t.Errorf("expected visibility_group_ids to keep its last-known value, got %v", result.VisibilityGroupIds)
	}
}

func TestResourceResourceRead_VisibilityNotFoundStillRemovesResource(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/resources/resource-1", jsonHandler(t, 200, map[string]any{
		"resource_id": "resource-1",
		"name":        "my-resource",
	}))
	mux.HandleFunc("/resources/resource-1/visibility", emptyHandler(http.StatusNotFound))
	ts := httptest.NewServer(mux)
	defer ts.Close()

	ctx := context.Background()
	r := &ResourceResource{client: newTestSDK(ts.URL)}

	var schemaResp resource.SchemaResponse
	r.Schema(ctx, resource.SchemaRequest{}, &schemaResp)

	initial := &ResourceResourceModel{ID: types.StringValue("resource-1")}
	state := tfsdk.State{Schema: schemaResp.Schema}
	if diags := state.Set(ctx, initial); diags.HasError() {
		t.Fatalf("unexpected diagnostics seeding state: %v", diags)
	}

	req := resource.ReadRequest{State: state}
	resp := &resource.ReadResponse{State: state}

	r.Read(ctx, req, resp)

	if resp.Diagnostics.HasError() {
		t.Fatalf("Read() returned unexpected error diagnostics: %v", resp.Diagnostics)
	}
	if !resp.State.Raw.IsNull() {
		t.Fatal("expected resource to be removed from state on a genuine 404")
	}
}

func TestGroupResourceRead_SubResourcesForbidden(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/groups/group-1", jsonHandler(t, 200, map[string]any{
		"group_id": "group-1",
		"name":     "my-group",
	}))
	mux.HandleFunc("/groups/group-1/message-channels", emptyHandler(http.StatusForbidden))
	mux.HandleFunc("/groups/group-1/on-call-schedules", emptyHandler(http.StatusForbidden))
	mux.HandleFunc("/groups/group-1/visibility", emptyHandler(http.StatusForbidden))
	ts := httptest.NewServer(mux)
	defer ts.Close()

	ctx := context.Background()
	r := &GroupResource{client: newTestSDK(ts.URL)}

	var schemaResp resource.SchemaResponse
	r.Schema(ctx, resource.SchemaRequest{}, &schemaResp)
	if schemaResp.Diagnostics.HasError() {
		t.Fatalf("unexpected diagnostics building schema: %v", schemaResp.Diagnostics)
	}

	initial := &GroupResourceModel{
		ID:                 types.StringValue("group-1"),
		Visibility:         types.StringValue("LIMITED"),
		VisibilityGroupIds: []types.String{types.StringValue("group-2")},
	}

	state := tfsdk.State{Schema: schemaResp.Schema}
	if diags := state.Set(ctx, initial); diags.HasError() {
		t.Fatalf("unexpected diagnostics seeding state: %v", diags)
	}

	req := resource.ReadRequest{State: state}
	resp := &resource.ReadResponse{State: state}

	r.Read(ctx, req, resp)

	if resp.Diagnostics.HasError() {
		t.Fatalf("Read() returned unexpected error diagnostics: %v", resp.Diagnostics)
	}
	for _, substr := range []string{
		"Unable to read group message channels",
		"Unable to read group on-call schedules",
		"Unable to read group visibility",
	} {
		if !hasWarningContaining(resp.Diagnostics, substr) {
			t.Errorf("expected a warning containing %q, got: %v", substr, resp.Diagnostics)
		}
	}
	if resp.State.Raw.IsNull() {
		t.Fatal("expected group to remain in state when sub-resources are forbidden, but it was removed")
	}

	var result GroupResourceModel
	if diags := resp.State.Get(ctx, &result); diags.HasError() {
		t.Fatalf("unexpected diagnostics reading back result state: %v", diags)
	}

	if result.Name.ValueString() != "my-group" {
		t.Errorf("expected name to be populated from the main GET, got %q", result.Name.ValueString())
	}
	if result.Visibility.ValueString() != "LIMITED" {
		t.Errorf("expected visibility to keep its last-known value %q, got %q", "LIMITED", result.Visibility.ValueString())
	}
	if len(result.VisibilityGroupIds) != 1 || result.VisibilityGroupIds[0].ValueString() != "group-2" {
		t.Errorf("expected visibility_group_ids to keep its last-known value, got %v", result.VisibilityGroupIds)
	}
}

func TestGroupResourceRead_MessageChannelsNotFoundStillRemovesGroup(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/groups/group-1", jsonHandler(t, 200, map[string]any{
		"group_id": "group-1",
		"name":     "my-group",
	}))
	mux.HandleFunc("/groups/group-1/message-channels", emptyHandler(http.StatusNotFound))
	ts := httptest.NewServer(mux)
	defer ts.Close()

	ctx := context.Background()
	r := &GroupResource{client: newTestSDK(ts.URL)}

	var schemaResp resource.SchemaResponse
	r.Schema(ctx, resource.SchemaRequest{}, &schemaResp)

	initial := &GroupResourceModel{ID: types.StringValue("group-1")}
	state := tfsdk.State{Schema: schemaResp.Schema}
	if diags := state.Set(ctx, initial); diags.HasError() {
		t.Fatalf("unexpected diagnostics seeding state: %v", diags)
	}

	req := resource.ReadRequest{State: state}
	resp := &resource.ReadResponse{State: state}

	r.Read(ctx, req, resp)

	if resp.Diagnostics.HasError() {
		t.Fatalf("Read() returned unexpected error diagnostics: %v", resp.Diagnostics)
	}
	if !resp.State.Raw.IsNull() {
		t.Fatal("expected group to be removed from state on a genuine 404")
	}
}

func hasWarningContaining(diags diag.Diagnostics, substr string) bool {
	for _, w := range diags.Warnings() {
		if strings.Contains(w.Summary(), substr) || strings.Contains(w.Detail(), substr) {
			return true
		}
	}
	return false
}
