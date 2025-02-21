// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"net/http"
)

type GetGroupReviewersRequest struct {
	// The ID of the group.
	GroupID string `pathParam:"style=simple,explode=false,name=group_id"`
}

func (o *GetGroupReviewersRequest) GetGroupID() string {
	if o == nil {
		return ""
	}
	return o.GroupID
}

type GetGroupReviewersResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// The IDs of owners that are reviewers for this group.
	Strings []string
}

func (o *GetGroupReviewersResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetGroupReviewersResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetGroupReviewersResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetGroupReviewersResponse) GetStrings() []string {
	if o == nil {
		return nil
	}
	return o.Strings
}
