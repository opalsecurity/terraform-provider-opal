package provider

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stretchr/testify/require"
)

func TestToSharedUpdateGroupInfoConfigurationTemplateIDNullOnWire(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	model := &GroupResourceModel{
		ID:                      types.StringValue("group-id"),
		ConfigurationTemplateID: types.StringNull(),
		Name:                    types.StringValue("Engineering"),
		RequireMfaToApprove:     types.BoolValue(false),
	}

	info, diags := model.ToSharedUpdateGroupInfo(ctx)
	require.False(t, diags.HasError(), diags.Errors())
	require.True(t, info.ConfigurationTemplateID.IsSet())
	require.True(t, info.ConfigurationTemplateID.IsNull())

	body, err := json.Marshal(info)
	require.NoError(t, err)

	var payload map[string]json.RawMessage
	require.NoError(t, json.Unmarshal(body, &payload))
	require.JSONEq(t, `null`, string(payload["configuration_template_id"]))
}

func TestToSharedUpdateGroupInfoConfigurationTemplateIDOmittedWhenUnknown(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	model := &GroupResourceModel{
		ID:                      types.StringValue("group-id"),
		ConfigurationTemplateID: types.StringUnknown(),
		Name:                    types.StringValue("Engineering"),
		RequireMfaToApprove:     types.BoolValue(false),
	}

	info, diags := model.ToSharedUpdateGroupInfo(ctx)
	require.False(t, diags.HasError(), diags.Errors())
	require.False(t, info.ConfigurationTemplateID.IsSet())

	body, err := json.Marshal(info)
	require.NoError(t, err)

	var payload map[string]json.RawMessage
	require.NoError(t, json.Unmarshal(body, &payload))
	_, present := payload["configuration_template_id"]
	require.False(t, present)
}

func TestToSharedUpdateGroupInfoConfigurationTemplateIDUUIDOnWire(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	templateID := "06851574-e50d-40ca-8c78-f72ae6ab4304"
	model := &GroupResourceModel{
		ID:                      types.StringValue("group-id"),
		ConfigurationTemplateID: types.StringValue(templateID),
		Name:                    types.StringValue("Engineering"),
		RequireMfaToApprove:     types.BoolValue(false),
	}

	info, diags := model.ToSharedUpdateGroupInfo(ctx)
	require.False(t, diags.HasError(), diags.Errors())

	body, err := json.Marshal(info)
	require.NoError(t, err)

	var payload map[string]json.RawMessage
	require.NoError(t, json.Unmarshal(body, &payload))
	require.JSONEq(t, `"`+templateID+`"`, string(payload["configuration_template_id"]))
}

func TestToSharedUpdateResourceInfoConfigurationTemplateIDNullOnWire(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	model := &ResourceResourceModel{
		ID:                      types.StringValue("resource-id"),
		ConfigurationTemplateID: types.StringNull(),
		Name:                    types.StringValue("Production"),
		RequireMfaToApprove:     types.BoolValue(false),
	}

	info, diags := model.ToSharedUpdateResourceInfo(ctx)
	require.False(t, diags.HasError(), diags.Errors())
	require.True(t, info.ConfigurationTemplateID.IsSet())
	require.True(t, info.ConfigurationTemplateID.IsNull())

	body, err := json.Marshal(info)
	require.NoError(t, err)

	var payload map[string]json.RawMessage
	require.NoError(t, json.Unmarshal(body, &payload))
	require.JSONEq(t, `null`, string(payload["configuration_template_id"]))
}
