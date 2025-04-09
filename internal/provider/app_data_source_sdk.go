// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"context"
	"encoding/json"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/opalsecurity/terraform-provider-opal/internal/provider/typeconvert"
	tfTypes "github.com/opalsecurity/terraform-provider-opal/internal/provider/types"
	"github.com/opalsecurity/terraform-provider-opal/internal/sdk/models/shared"
)

func (r *AppDataSourceModel) RefreshFromSharedApp(ctx context.Context, resp *shared.App) diag.Diagnostics {
	var diags diag.Diagnostics

	if resp != nil {
		r.AdminOwnerID = types.StringValue(resp.AdminOwnerID)
		r.Description = types.StringValue(resp.Description)
		r.ID = types.StringValue(resp.ID)
		r.Name = types.StringValue(resp.Name)
		r.Type = types.StringValue(resp.Type)
		r.Validations = []tfTypes.AppValidation{}
		if len(r.Validations) > len(resp.Validations) {
			r.Validations = r.Validations[:len(resp.Validations)]
		}
		for validationsCount, validationsItem := range resp.Validations {
			var validations tfTypes.AppValidation
			validations.Details = types.StringPointerValue(validationsItem.Details)
			validations.Key = types.StringValue(validationsItem.Key)
			nameResult, _ := json.Marshal(validationsItem.Name)
			validations.Name = types.StringValue(string(nameResult))
			validations.Severity = types.StringValue(string(validationsItem.Severity))
			validations.Status = types.StringValue(string(validationsItem.Status))
			validations.UpdatedAt = types.StringValue(typeconvert.TimeToString(validationsItem.UpdatedAt))
			validations.UsageReason = types.StringPointerValue(validationsItem.UsageReason)
			if validationsCount+1 > len(r.Validations) {
				r.Validations = append(r.Validations, validations)
			} else {
				r.Validations[validationsCount].Details = validations.Details
				r.Validations[validationsCount].Key = validations.Key
				r.Validations[validationsCount].Name = validations.Name
				r.Validations[validationsCount].Severity = validations.Severity
				r.Validations[validationsCount].Status = validations.Status
				r.Validations[validationsCount].UpdatedAt = validations.UpdatedAt
				r.Validations[validationsCount].UsageReason = validations.UsageReason
			}
		}
	}

	return diags
}
