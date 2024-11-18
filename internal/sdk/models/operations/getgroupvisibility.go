// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/opalsecurity/terraform-provider-opal/internal/sdk/internal/utils"
	"github.com/opalsecurity/terraform-provider-opal/internal/sdk/models/shared"
	"net/http"
)

type GetGroupVisibilityRequest struct {
	// The ID of the group.
	ID string `pathParam:"style=simple,explode=false,name=group_id"`
}

func (o *GetGroupVisibilityRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

// GetGroupVisibilityResponseBody - Visibility infomation of an entity.
type GetGroupVisibilityResponseBody struct {
	// The visibility level of the entity.
	Visibility         shared.VisibilityTypeEnum `json:"visibility"`
	VisibilityGroupIds []string                  `json:"visibility_group_ids"`
}

func (g GetGroupVisibilityResponseBody) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GetGroupVisibilityResponseBody) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *GetGroupVisibilityResponseBody) GetVisibility() shared.VisibilityTypeEnum {
	if o == nil {
		return shared.VisibilityTypeEnum("")
	}
	return o.Visibility
}

func (o *GetGroupVisibilityResponseBody) GetVisibilityGroupIds() []string {
	if o == nil {
		return nil
	}
	return o.VisibilityGroupIds
}

type GetGroupVisibilityResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// The visibility info of this group.
	Object *GetGroupVisibilityResponseBody
}

func (o *GetGroupVisibilityResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetGroupVisibilityResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetGroupVisibilityResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetGroupVisibilityResponse) GetObject() *GetGroupVisibilityResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
