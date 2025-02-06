// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"encoding/json"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/opalsecurity/terraform-provider-opal/internal/provider/types"
	"github.com/opalsecurity/terraform-provider-opal/internal/sdk/models/shared"
	"time"
)

func (r *AppDataSourceModel) RefreshFromSharedApp(resp *shared.App) {
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
			var validations1 tfTypes.AppValidation
			validations1.Details = types.StringPointerValue(validationsItem.Details)
			validations1.Key = types.StringValue(validationsItem.Key)
			nameResult, _ := json.Marshal(validationsItem.Name)
			validations1.Name = types.StringValue(string(nameResult))
			validations1.Severity = types.StringValue(string(validationsItem.Severity))
			validations1.Status = types.StringValue(string(validationsItem.Status))
			validations1.UpdatedAt = types.StringValue(validationsItem.UpdatedAt.Format(time.RFC3339Nano))
			validations1.UsageReason = types.StringPointerValue(validationsItem.UsageReason)
			if validationsCount+1 > len(r.Validations) {
				r.Validations = append(r.Validations, validations1)
			} else {
				r.Validations[validationsCount].Details = validations1.Details
				r.Validations[validationsCount].Key = validations1.Key
				r.Validations[validationsCount].Name = validations1.Name
				r.Validations[validationsCount].Severity = validations1.Severity
				r.Validations[validationsCount].Status = validations1.Status
				r.Validations[validationsCount].UpdatedAt = validations1.UpdatedAt
				r.Validations[validationsCount].UsageReason = validations1.UsageReason
			}
		}
	}
}
