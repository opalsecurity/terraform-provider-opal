package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"github.com/opalsecurity/terraform-provider-opal/v3/internal/sdk/models/operations"
)

var (
	_ resource.ResourceWithModifyPlan = &GroupResource{}
	_ resource.ResourceWithModifyPlan = &ResourceResource{}
)

// These attributes are sent through the bulk group/resource update endpoints.
// The public API rejects changes to them while a configuration template is
// attached. Attributes managed by dedicated endpoints, such as visibility and
// group message channels, are intentionally not included.
var groupConfigurationTemplateRestrictedUpdates = []string{
	"description",
	"extensions_duration_in_minutes",
	"group_leader_user_ids",
	"match_remote_description",
	"match_remote_name",
	"name",
	"risk_sensitivity_override",
}

var resourceConfigurationTemplateRestrictedUpdates = []string{
	"description",
	"extensions_duration_in_minutes",
	"match_remote_description",
	"match_remote_name",
	"name",
	"parent_resource_id",
	"risk_sensitivity_override",
}

var groupConfigurationTemplateLinkedOnlyUpdates = []string{
	"message_channel_ids",
	"on_call_schedule_ids",
	"visibility",
	"visibility_group_ids",
}

var resourceConfigurationTemplateLinkedOnlyUpdates = []string{
	"visibility",
	"visibility_group_ids",
}

func (r *GroupResource) ModifyPlan(_ context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	validateConfigurationTemplatePlan(
		req,
		resp,
		groupConfigurationTemplateRestrictedUpdates,
		groupConfigurationTemplateLinkedOnlyUpdates,
	)
}

func (r *ResourceResource) ModifyPlan(_ context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	validateConfigurationTemplatePlan(
		req,
		resp,
		resourceConfigurationTemplateRestrictedUpdates,
		resourceConfigurationTemplateLinkedOnlyUpdates,
	)
}

func validateConfigurationTemplatePlan(
	req resource.ModifyPlanRequest,
	resp *resource.ModifyPlanResponse,
	restrictedUpdates []string,
	linkedOnlyUpdates []string,
) {
	// Creation is supported as two REST calls: create the entity, then attach
	// the template with a minimal update. A null plan represents destruction.
	if req.State.Raw.IsNull() || req.Plan.Raw.IsNull() {
		return
	}

	config, err := terraformObjectValues(req.Config.Raw)
	if err != nil {
		resp.Diagnostics.AddError("Unable to inspect configuration", err.Error())
		return
	}
	plan, err := terraformObjectValues(req.Plan.Raw)
	if err != nil {
		resp.Diagnostics.AddError("Unable to inspect planned state", err.Error())
		return
	}
	state, err := terraformObjectValues(req.State.Raw)
	if err != nil {
		resp.Diagnostics.AddError("Unable to inspect prior state", err.Error())
		return
	}

	templateID, ok := plan["configuration_template_id"]
	if !ok || !templateID.IsKnown() || templateID.IsNull() {
		return
	}
	var templateIDString string
	if err := templateID.As(&templateIDString); err != nil || templateIDString == "" {
		return
	}

	validateConfiguredChanges(config, state, restrictedUpdates, resp)

	stateTemplateID, stateHasTemplate := state["configuration_template_id"]
	if stateHasTemplate && stateTemplateID.IsKnown() && !stateTemplateID.IsNull() {
		validateConfiguredChanges(config, state, linkedOnlyUpdates, resp)
	}
}

func validateConfiguredChanges(
	config map[string]tftypes.Value,
	state map[string]tftypes.Value,
	attributes []string,
	resp *resource.ModifyPlanResponse,
) {
	for _, attribute := range attributes {
		configValue, configured := config[attribute]
		if !configured || !configValue.IsKnown() || configValue.IsNull() {
			continue
		}

		stateValue, presentInState := state[attribute]
		if presentInState && stateValue.IsKnown() && configValue.Equal(stateValue) {
			continue
		}

		resp.Diagnostics.AddAttributeError(
			path.Root(attribute),
			"Cannot update an entity linked to a configuration template",
			fmt.Sprintf(
				"%q cannot be changed while configuration_template_id is set. "+
					"The public REST API only accepts the entity ID and configuration_template_id when attaching or changing a configuration template. "+
					"Unlink the template in the Opal UI before changing this attribute.",
				attribute,
			),
		)
	}
}

func terraformObjectValues(value tftypes.Value) (map[string]tftypes.Value, error) {
	values := make(map[string]tftypes.Value)
	if err := value.As(&values); err != nil {
		return nil, err
	}
	return values, nil
}

// refreshResourceVisibility keeps the resource state aligned with the
// dedicated public REST endpoint. The base resource response does not include
// visibility, which otherwise makes an imported linked resource appear to
// change visibility on every plan.
func (r *ResourceResource) refreshResourceVisibility(ctx context.Context, data *ResourceResourceModel) diag.Diagnostics {
	var diags diag.Diagnostics

	res, err := r.client.Resources.GetVisibility(ctx, operations.GetResourceVisibilityRequest{
		ID: data.ID.ValueString(),
	})
	if err != nil {
		diags.AddError("failure to invoke API", err.Error())
		if res != nil && res.RawResponse != nil {
			diags.AddError("unexpected http request/response", debugResponse(res.RawResponse))
		}
		return diags
	}
	if res == nil || res.StatusCode != 200 || res.VisibilityInfo == nil {
		diags.AddError("unexpected response from API", "Unable to read resource visibility")
		return diags
	}

	data.Visibility = types.StringValue(string(res.VisibilityInfo.Visibility))
	data.VisibilityGroupIds = make([]types.String, 0, len(res.VisibilityInfo.VisibilityGroupIds))
	for _, groupID := range res.VisibilityInfo.VisibilityGroupIds {
		data.VisibilityGroupIds = append(data.VisibilityGroupIds, types.StringValue(groupID))
	}

	return diags
}
