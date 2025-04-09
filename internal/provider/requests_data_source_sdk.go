// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/opalsecurity/terraform-provider-opal/internal/provider/typeconvert"
	tfTypes "github.com/opalsecurity/terraform-provider-opal/internal/provider/types"
	"github.com/opalsecurity/terraform-provider-opal/internal/sdk/models/shared"
)

func (r *RequestsDataSourceModel) RefreshFromSharedRequestList(ctx context.Context, resp *shared.RequestList) diag.Diagnostics {
	var diags diag.Diagnostics

	if resp != nil {
		r.Cursor = types.StringPointerValue(resp.Cursor)
		r.Requests = []tfTypes.Request{}
		if len(r.Requests) > len(resp.Requests) {
			r.Requests = r.Requests[:len(resp.Requests)]
		}
		for requestsCount, requestsItem := range resp.Requests {
			var requests tfTypes.Request
			requests.CreatedAt = types.StringValue(typeconvert.TimeToString(requestsItem.CreatedAt))
			requests.CustomFieldsResponses = []tfTypes.RequestCustomFieldResponse{}
			for customFieldsResponsesCount, customFieldsResponsesItem := range requestsItem.CustomFieldsResponses {
				var customFieldsResponses tfTypes.RequestCustomFieldResponse
				customFieldsResponses.FieldName = types.StringValue(customFieldsResponsesItem.FieldName)
				customFieldsResponses.FieldType = types.StringValue(string(customFieldsResponsesItem.FieldType))
				if customFieldsResponsesItem.FieldValue.Str != nil {
					customFieldsResponses.FieldValue.Str = types.StringPointerValue(customFieldsResponsesItem.FieldValue.Str)
				}
				if customFieldsResponsesItem.FieldValue.Boolean != nil {
					customFieldsResponses.FieldValue.Boolean = types.BoolPointerValue(customFieldsResponsesItem.FieldValue.Boolean)
				}
				if customFieldsResponsesCount+1 > len(requests.CustomFieldsResponses) {
					requests.CustomFieldsResponses = append(requests.CustomFieldsResponses, customFieldsResponses)
				} else {
					requests.CustomFieldsResponses[customFieldsResponsesCount].FieldName = customFieldsResponses.FieldName
					requests.CustomFieldsResponses[customFieldsResponsesCount].FieldType = customFieldsResponses.FieldType
					requests.CustomFieldsResponses[customFieldsResponsesCount].FieldValue = customFieldsResponses.FieldValue
				}
			}
			requests.DurationMinutes = types.Int64PointerValue(requestsItem.DurationMinutes)
			requests.ID = types.StringValue(requestsItem.ID)
			requests.Reason = types.StringValue(requestsItem.Reason)
			requests.RequestedItemsList = []tfTypes.RequestedItem{}
			for requestedItemsListCount, requestedItemsListItem := range requestsItem.RequestedItemsList {
				var requestedItemsList tfTypes.RequestedItem
				requestedItemsList.AccessLevelName = types.StringPointerValue(requestedItemsListItem.AccessLevelName)
				requestedItemsList.AccessLevelRemoteID = types.StringPointerValue(requestedItemsListItem.AccessLevelRemoteID)
				requestedItemsList.GroupID = types.StringPointerValue(requestedItemsListItem.GroupID)
				requestedItemsList.Name = types.StringPointerValue(requestedItemsListItem.Name)
				requestedItemsList.ResourceID = types.StringPointerValue(requestedItemsListItem.ResourceID)
				if requestedItemsListCount+1 > len(requests.RequestedItemsList) {
					requests.RequestedItemsList = append(requests.RequestedItemsList, requestedItemsList)
				} else {
					requests.RequestedItemsList[requestedItemsListCount].AccessLevelName = requestedItemsList.AccessLevelName
					requests.RequestedItemsList[requestedItemsListCount].AccessLevelRemoteID = requestedItemsList.AccessLevelRemoteID
					requests.RequestedItemsList[requestedItemsListCount].GroupID = requestedItemsList.GroupID
					requests.RequestedItemsList[requestedItemsListCount].Name = requestedItemsList.Name
					requests.RequestedItemsList[requestedItemsListCount].ResourceID = requestedItemsList.ResourceID
				}
			}
			requests.RequesterID = types.StringValue(requestsItem.RequesterID)
			requests.Status = types.StringValue(string(requestsItem.Status))
			requests.TargetGroupID = types.StringPointerValue(requestsItem.TargetGroupID)
			requests.TargetUserID = types.StringPointerValue(requestsItem.TargetUserID)
			requests.UpdatedAt = types.StringValue(typeconvert.TimeToString(requestsItem.UpdatedAt))
			if requestsCount+1 > len(r.Requests) {
				r.Requests = append(r.Requests, requests)
			} else {
				r.Requests[requestsCount].CreatedAt = requests.CreatedAt
				r.Requests[requestsCount].CustomFieldsResponses = requests.CustomFieldsResponses
				r.Requests[requestsCount].DurationMinutes = requests.DurationMinutes
				r.Requests[requestsCount].ID = requests.ID
				r.Requests[requestsCount].Reason = requests.Reason
				r.Requests[requestsCount].RequestedItemsList = requests.RequestedItemsList
				r.Requests[requestsCount].RequesterID = requests.RequesterID
				r.Requests[requestsCount].Status = requests.Status
				r.Requests[requestsCount].TargetGroupID = requests.TargetGroupID
				r.Requests[requestsCount].TargetUserID = requests.TargetUserID
				r.Requests[requestsCount].UpdatedAt = requests.UpdatedAt
			}
		}
	}

	return diags
}
