// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/opalsecurity/terraform-provider-opal/internal/sdk/models/shared"
	"net/http"
)

type GetEventsRequest struct {
	// An actor filter for the events. Supply the ID of the actor.
	ActorFilter *string `queryParam:"style=form,explode=true,name=actor_filter"`
	// An API filter for the events. Supply the name and preview of the API token.
	APITokenFilter *string `queryParam:"style=form,explode=true,name=api_token_filter"`
	// The pagination cursor value.
	Cursor *string `queryParam:"style=form,explode=true,name=cursor"`
	// An end date filter for the events.
	EndDateFilter *string `queryParam:"style=form,explode=true,name=end_date_filter"`
	// An event type filter for the events.
	EventTypeFilter *string `queryParam:"style=form,explode=true,name=event_type_filter"`
	// An object filter for the events. Supply the ID of the object.
	ObjectFilter *string `queryParam:"style=form,explode=true,name=object_filter"`
	// Number of results to return per page. Default is 200.
	PageSize *int64 `queryParam:"style=form,explode=true,name=page_size"`
	// A start date filter for the events.
	StartDateFilter *string `queryParam:"style=form,explode=true,name=start_date_filter"`
}

func (o *GetEventsRequest) GetActorFilter() *string {
	if o == nil {
		return nil
	}
	return o.ActorFilter
}

func (o *GetEventsRequest) GetAPITokenFilter() *string {
	if o == nil {
		return nil
	}
	return o.APITokenFilter
}

func (o *GetEventsRequest) GetCursor() *string {
	if o == nil {
		return nil
	}
	return o.Cursor
}

func (o *GetEventsRequest) GetEndDateFilter() *string {
	if o == nil {
		return nil
	}
	return o.EndDateFilter
}

func (o *GetEventsRequest) GetEventTypeFilter() *string {
	if o == nil {
		return nil
	}
	return o.EventTypeFilter
}

func (o *GetEventsRequest) GetObjectFilter() *string {
	if o == nil {
		return nil
	}
	return o.ObjectFilter
}

func (o *GetEventsRequest) GetPageSize() *int64 {
	if o == nil {
		return nil
	}
	return o.PageSize
}

func (o *GetEventsRequest) GetStartDateFilter() *string {
	if o == nil {
		return nil
	}
	return o.StartDateFilter
}

type GetEventsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// One page worth of events with the appropriate filters applied.
	PaginatedEventList *shared.PaginatedEventList
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetEventsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetEventsResponse) GetPaginatedEventList() *shared.PaginatedEventList {
	if o == nil {
		return nil
	}
	return o.PaginatedEventList
}

func (o *GetEventsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetEventsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
