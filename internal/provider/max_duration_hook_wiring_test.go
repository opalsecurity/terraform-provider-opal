package provider

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/opalsecurity/terraform-provider-opal/v3/internal/sdk"
	"github.com/opalsecurity/terraform-provider-opal/v3/internal/sdk/models/operations"
	"github.com/opalsecurity/terraform-provider-opal/v3/internal/sdk/models/shared"
)

// TestMaxDurationIndefiniteHookIsWired drives a real SDK client against a stub
// API to confirm the response hook is registered and reaches the decoded model.
// The hook's own behavior is covered in internal/sdk/internal/hooks; this test
// exists so a lost registration cannot pass silently (PLAT-654).
func TestMaxDurationIndefiniteHookIsWired(t *testing.T) {
	t.Parallel()

	for _, testCase := range []struct {
		name     string
		response string
		want     int64
	}{
		{
			name:     "omitted max duration reads back as indefinite",
			response: `{"group_id":"g1","name":"g","request_configurations":[{"priority":0}]}`,
			want:     -1,
		},
		{
			name:     "a real duration is untouched",
			response: `{"group_id":"g1","name":"g","request_configurations":[{"priority":0,"max_duration_minutes":120}]}`,
			want:     120,
		},
	} {
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			server := httptest.NewServer(
				http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
					w.Header().Set("Content-Type", "application/json")
					if _, err := w.Write([]byte(testCase.response)); err != nil {
						t.Errorf("writing stub response: %v", err)
					}
				}),
			)
			defer server.Close()

			client := sdk.New(
				sdk.WithServerURL(server.URL),
				sdk.WithSecurity(shared.Security{BearerAuth: "test-token"}),
				sdk.WithClient(server.Client()),
			)

			res, err := client.Groups.GetGroup(
				context.Background(),
				operations.GetGroupRequest{ID: "g1"},
			)
			if err != nil {
				t.Fatalf("GetGroup returned error: %v", err)
			}
			if res.Group == nil {
				t.Fatalf("GetGroup returned no group (status %d)", res.StatusCode)
			}
			if len(res.Group.RequestConfigurations) != 1 {
				t.Fatalf("got %d request configurations, want 1", len(res.Group.RequestConfigurations))
			}

			maxDuration := res.Group.RequestConfigurations[0].MaxDuration
			if maxDuration == nil {
				t.Fatalf("max duration is nil, want %d", testCase.want)
			}
			if *maxDuration != testCase.want {
				t.Errorf("max duration = %d, want %d", *maxDuration, testCase.want)
			}
		})
	}
}
