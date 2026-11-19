package provider

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	tfTypes "github.com/opalsecurity/terraform-provider-opal/v3/internal/provider/types"
)

var _ resource.ResourceWithValidateConfig = &RequestTemplateResource{}

// Limits mirror web/backend/models/requesttemplates/constants.go.
const (
	requestTemplateNameMaxLen           = 100
	requestTemplateCustomFieldMaxNum    = 20
	requestTemplateCustomFieldNameMax   = 75
	requestTemplateCustomFieldDescMax   = 150
	requestTemplateDropdownOptionsMax   = 100
	requestTemplateDropdownOptionMaxLen = 75
)

func (r *RequestTemplateResource) ValidateConfig(
	ctx context.Context,
	req resource.ValidateConfigRequest,
	resp *resource.ValidateConfigResponse,
) {
	var data RequestTemplateResourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if !data.Name.IsNull() && !data.Name.IsUnknown() &&
		len(data.Name.ValueString()) > requestTemplateNameMaxLen {
		resp.Diagnostics.AddAttributeError(
			path.Root("name"),
			"Request template name is too long",
			fmt.Sprintf("name must be at most %d characters", requestTemplateNameMaxLen),
		)
	}

	if data.CustomFields == nil {
		return
	}
	if len(data.CustomFields) > requestTemplateCustomFieldMaxNum {
		resp.Diagnostics.AddAttributeError(
			path.Root("custom_fields"),
			"Too many custom fields",
			fmt.Sprintf("at most %d custom fields are allowed", requestTemplateCustomFieldMaxNum),
		)
	}

	seenNames := map[string]int{}
	for i, field := range data.CustomFields {
		validateRequestTemplateCustomField(i, field, seenNames, &resp.Diagnostics)
	}
}

func validateRequestTemplateCustomField(
	index int,
	field tfTypes.RequestTemplateCustomField,
	seenNames map[string]int,
	diagnostics *diag.Diagnostics,
) {
	fieldPath := path.Root("custom_fields").AtListIndex(index)

	if field.Type.IsNull() || field.Type.IsUnknown() {
		return
	}
	fieldType := field.Type.ValueString()

	if !field.Name.IsNull() && !field.Name.IsUnknown() {
		name := field.Name.ValueString()
		if len(name) > requestTemplateCustomFieldNameMax {
			diagnostics.AddAttributeError(
				fieldPath.AtName("name"),
				"Custom field name is too long",
				fmt.Sprintf("name must be at most %d characters", requestTemplateCustomFieldNameMax),
			)
		}
		if name != "" {
			if prev, ok := seenNames[name]; ok {
				diagnostics.AddAttributeError(
					fieldPath.AtName("name"),
					"Duplicate custom field name",
					fmt.Sprintf("custom field name %q is also used at custom_fields[%d]", name, prev),
				)
			} else {
				seenNames[name] = index
			}
		}
	}

	if !field.Description.IsNull() && !field.Description.IsUnknown() &&
		len(field.Description.ValueString()) > requestTemplateCustomFieldDescMax {
		diagnostics.AddAttributeError(
			fieldPath.AtName("description"),
			"Custom field description is too long",
			fmt.Sprintf("description must be at most %d characters", requestTemplateCustomFieldDescMax),
		)
	}

	switch fieldType {
	case "CALLOUT":
		if !field.Required.IsNull() && !field.Required.IsUnknown() && field.Required.ValueBool() {
			diagnostics.AddAttributeError(
				fieldPath.AtName("required"),
				"Callout fields cannot be required",
				"CALLOUT fields are display-only and collect no answer; omit required or set it to false.",
			)
		}
		if field.Metadata == nil || field.Metadata.CalloutData == nil {
			diagnostics.AddAttributeError(
				fieldPath.AtName("metadata"),
				"Callout fields require callout_data",
				"Set metadata.callout_data with severity and text for CALLOUT fields.",
			)
			return
		}
		if field.Metadata.CalloutData.Text.IsNull() ||
			(!field.Metadata.CalloutData.Text.IsUnknown() &&
				strings.TrimSpace(field.Metadata.CalloutData.Text.ValueString()) == "") {
			diagnostics.AddAttributeError(
				fieldPath.AtName("metadata").AtName("callout_data").AtName("text"),
				"Callout text is required",
				"metadata.callout_data.text must be a non-empty string.",
			)
		}
	case "MULTI_CHOICE", "MULTI_SELECT":
		if field.Name.IsNull() ||
			(!field.Name.IsUnknown() && strings.TrimSpace(field.Name.ValueString()) == "") {
			diagnostics.AddAttributeError(
				fieldPath.AtName("name"),
				"Custom field name is required",
				fmt.Sprintf("%s fields require a name", fieldType),
			)
		}
		if field.Metadata == nil || field.Metadata.MultiChoiceData == nil ||
			field.Metadata.MultiChoiceData.Options == nil {
			diagnostics.AddAttributeError(
				fieldPath.AtName("metadata"),
				"Dropdown fields require options",
				fmt.Sprintf("%s fields require metadata.multi_choice_data.options with at least one non-blank option", fieldType),
			)
			return
		}
		options := field.Metadata.MultiChoiceData.Options
		if len(options) == 0 {
			diagnostics.AddAttributeError(
				fieldPath.AtName("metadata").AtName("multi_choice_data").AtName("options"),
				"Dropdown fields require options",
				"provide at least one non-blank option",
			)
			return
		}
		if len(options) > requestTemplateDropdownOptionsMax {
			diagnostics.AddAttributeError(
				fieldPath.AtName("metadata").AtName("multi_choice_data").AtName("options"),
				"Too many dropdown options",
				fmt.Sprintf("at most %d options are allowed", requestTemplateDropdownOptionsMax),
			)
		}
		for oi, opt := range options {
			if opt.IsUnknown() {
				continue
			}
			if opt.IsNull() || strings.TrimSpace(opt.ValueString()) == "" {
				diagnostics.AddAttributeError(
					fieldPath.AtName("metadata").AtName("multi_choice_data").AtName("options").AtListIndex(oi),
					"Dropdown option cannot be blank",
					"each option must be a non-empty string",
				)
				continue
			}
			if len(opt.ValueString()) > requestTemplateDropdownOptionMaxLen {
				diagnostics.AddAttributeError(
					fieldPath.AtName("metadata").AtName("multi_choice_data").AtName("options").AtListIndex(oi),
					"Dropdown option is too long",
					fmt.Sprintf("each option must be at most %d characters", requestTemplateDropdownOptionMaxLen),
				)
			}
		}
	default:
		// SHORT_TEXT, LONG_TEXT, BOOLEAN, and any future answerable types.
		if field.Name.IsNull() ||
			(!field.Name.IsUnknown() && strings.TrimSpace(field.Name.ValueString()) == "") {
			diagnostics.AddAttributeError(
				fieldPath.AtName("name"),
				"Custom field name is required",
				fmt.Sprintf("%s fields require a name", fieldType),
			)
		}
	}
}
