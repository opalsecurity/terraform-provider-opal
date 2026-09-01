package hooks

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type configurationTemplateUpdateOperation struct {
	itemsKey string
	idKey    string
}

var configurationTemplateUpdateOperations = map[string]configurationTemplateUpdateOperation{
	"updateGroups": {
		itemsKey: "groups",
		idKey:    "group_id",
	},
	"updateResources": {
		itemsKey: "resources",
		idKey:    "resource_id",
	},
}

var configurationTemplateBlockedFollowUpOperations = map[string]struct{}{
	"updateGroupMessageChannels": {},
	"updateGroupOnCallSchedules": {},
	"updateGroupVisibility":      {},
	"updateResourceVisibility":   {},
}

var configurationTemplateVisibilityOperations = map[string]struct{}{
	"updateGroupVisibility":    {},
	"updateResourceVisibility": {},
}

type configurationTemplateUpdateHook struct{}

func (h *configurationTemplateUpdateHook) BeforeRequest(hookCtx BeforeRequestContext, req *http.Request) (*http.Request, error) {
	operation, ok := configurationTemplateUpdateOperations[hookCtx.OperationID]
	if !ok || req.Body == nil {
		return req, nil
	}

	body, err := io.ReadAll(req.Body)
	if err != nil {
		return req, fmt.Errorf("read configuration template update request: %w", err)
	}
	if err := req.Body.Close(); err != nil {
		return req, fmt.Errorf("close configuration template update request: %w", err)
	}

	var payload map[string][]map[string]json.RawMessage
	if err := json.Unmarshal(body, &payload); err != nil {
		return req, fmt.Errorf("decode configuration template update request: %w", err)
	}

	changed := false
	for _, item := range payload[operation.itemsKey] {
		templateID, present := item["configuration_template_id"]
		if !present || bytes.Equal(templateID, []byte("null")) || bytes.Equal(templateID, []byte(`""`)) {
			continue
		}

		for field := range item {
			if field != operation.idKey && field != "configuration_template_id" {
				delete(item, field)
				changed = true
			}
		}
	}

	if !changed {
		req.Body = io.NopCloser(bytes.NewReader(body))
		return req, nil
	}

	body, err = json.Marshal(payload)
	if err != nil {
		return req, fmt.Errorf("encode configuration template update request: %w", err)
	}
	req.Body = io.NopCloser(bytes.NewReader(body))
	req.ContentLength = int64(len(body))
	req.GetBody = func() (io.ReadCloser, error) {
		return io.NopCloser(bytes.NewReader(body)), nil
	}

	return req, nil
}

func (h *configurationTemplateUpdateHook) AfterError(
	hookCtx AfterErrorContext,
	res *http.Response,
	err error,
) (*http.Response, error) {
	if _, ok := configurationTemplateBlockedFollowUpOperations[hookCtx.OperationID]; !ok ||
		res == nil ||
		res.StatusCode != http.StatusBadRequest ||
		res.Body == nil {
		return res, err
	}

	body, readErr := io.ReadAll(res.Body)
	if readErr != nil {
		return res, err
	}
	_ = res.Body.Close()

	if !suppressConfigurationTemplateFollowUpError(hookCtx.OperationID, string(body)) {
		res.Body = io.NopCloser(bytes.NewReader(body))
		return res, err
	}

	res.StatusCode = http.StatusOK
	res.Status = "200 OK"
	res.Body = http.NoBody
	res.ContentLength = 0
	res.Header.Del("Content-Length")
	res.Header.Del("Content-Type")

	return res, nil
}

// suppressConfigurationTemplateFollowUpError reports whether a 400 from one of
// the follow-up endpoints should be treated as success because the entity's
// configuration is governed by an attached configuration template.
func suppressConfigurationTemplateFollowUpError(operationID, body string) bool {
	if strings.Contains(body, "linked to configuration templates") {
		return true
	}

	// When a configuration template is attached, the Terraform configuration
	// omits `visibility` (the template governs it), so the generated
	// provider still calls the visibility endpoint with an empty visibility
	// level. The API parses the enum before checking for an attached
	// template, so it responds with `Unrecognized visibility level: ` (with
	// an empty value) rather than the linked-template message. Match only
	// the empty-value form: a non-empty invalid level is a genuine error and
	// is caught by the plan-time enum validator anyway.
	if _, ok := configurationTemplateVisibilityOperations[operationID]; ok {
		return strings.Contains(body, `Unrecognized visibility level: "`)
	}

	return false
}
