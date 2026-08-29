package sdk

import (
	"context"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/opalsecurity/terraform-provider-opal/v3/internal/sdk/models/shared"
	"github.com/stretchr/testify/require"
)

type captureRequestClient struct {
	body string
}

func (c *captureRequestClient) Do(req *http.Request) (*http.Response, error) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}
	c.body = string(body)

	return &http.Response{
		StatusCode: http.StatusBadRequest,
		Header:     make(http.Header),
		Body:       io.NopCloser(strings.NewReader(`{"message":"stop after capture"}`)),
		Request:    req,
	}, nil
}

func TestConfigurationTemplateHookIsAppliedBySDK(t *testing.T) {
	t.Parallel()

	client := &captureRequestClient{}
	sdk := New(
		WithClient(client),
		WithServerURL("https://example.com"),
	)
	templateID := "template-id"
	name := "Production"

	_, err := sdk.Resources.Update(context.Background(), shared.UpdateResourceInfoList{
		Resources: []shared.UpdateResourceInfo{{
			ID:                      "resource-id",
			ConfigurationTemplateID: &templateID,
			Name:                    &name,
		}},
	})
	require.Error(t, err)
	require.JSONEq(t,
		`{"resources":[{"resource_id":"resource-id","configuration_template_id":"template-id"}]}`,
		client.body,
	)
}
