package provider

import "testing"

func TestAppTypeSupportsVisibilityEndpoint(t *testing.T) {
	t.Parallel()

	tests := []struct {
		appType string
		want    bool
	}{
		{appType: "CUSTOM", want: true},
		{appType: "CUSTOM_CONNECTOR", want: true},
		{appType: "OKTA_DIRECTORY", want: false},
		{appType: "AWS_SSO", want: false},
		{appType: "", want: false},
	}

	for _, testCase := range tests {
		t.Run(testCase.appType, func(t *testing.T) {
			t.Parallel()
			got := appTypeSupportsVisibilityEndpoint(testCase.appType)
			if got != testCase.want {
				t.Errorf(
					"appTypeSupportsVisibilityEndpoint(%q) = %v, want %v",
					testCase.appType,
					got,
					testCase.want,
				)
			}
		})
	}
}
