package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

var (
	_ resource.ResourceWithModifyPlan = &GroupResource{}
	_ resource.ResourceWithModifyPlan = &ResourceResource{}
)

// message_channel_ids and on_call_schedule_ids live in schemas shared with
// other entities, so they cannot be added to GroupConfigurationTemplateID
// without leaking the conflict. Visibility and visibility_group_ids are
// enforced there instead.
var groupConfigurationTemplateLinkedOnlyUpdates = []string{
	"message_channel_ids",
	"on_call_schedule_ids",
}

func (r *GroupResource) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	validateConfigurationTemplatePlan(ctx, req, resp, groupConfigurationTemplateLinkedOnlyUpdates)
}

func (r *ResourceResource) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	validateConfigurationTemplatePlan(ctx, req, resp, nil)
}

func validateConfigurationTemplatePlan(
	ctx context.Context,
	req resource.ModifyPlanRequest,
	resp *resource.ModifyPlanResponse,
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

	stateTemplateID, stateHasTemplate := state["configuration_template_id"]
	if stateHasTemplate && stateTemplateID.IsKnown() && !stateTemplateID.IsNull() {
		validateConfiguredChanges(config, state, linkedOnlyUpdates, resp)
	}

	preserveTemplateGovernedVisibility(ctx, config, state, resp)
}

// preserveTemplateGovernedVisibility keeps the prior visibility in the plan when
// the template governs it. The group's visibility is populated by its read
// operation, so leaving it unconfigured would otherwise plan as "known after
// apply" on every run and produce a permanent diff.
func preserveTemplateGovernedVisibility(
	ctx context.Context,
	config map[string]tftypes.Value,
	state map[string]tftypes.Value,
	resp *resource.ModifyPlanResponse,
) {
	if configured, ok := config["visibility"]; !ok || !configured.IsNull() {
		return
	}

	stateValue, ok := state["visibility"]
	if !ok || !stateValue.IsKnown() || stateValue.IsNull() {
		return
	}

	var visibility string
	if err := stateValue.As(&visibility); err != nil {
		return
	}

	resp.Diagnostics.Append(
		resp.Plan.SetAttribute(ctx, path.Root("visibility"), types.StringValue(visibility))...,
	)
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
