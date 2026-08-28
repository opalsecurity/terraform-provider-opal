package hooks

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

var configurationTemplateConflicts = map[string][]string{
	"updateGroups": {
		"admin_owner_id",
		"custom_request_notification",
		"request_configurations",
		"require_mfa_to_approve",
	},
	"updateResources": {
		"admin_owner_id",
		"custom_request_notification",
		"request_configurations",
		"require_mfa_to_approve",
		"require_mfa_to_connect",
		"ticket_propagation",
	},
}

type configurationTemplateUpdateHook struct{}

func (h *configurationTemplateUpdateHook) BeforeRequest(hookCtx BeforeRequestContext, req *http.Request) (*http.Request, error) {
	conflicts, ok := configurationTemplateConflicts[hookCtx.OperationID]
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

	itemsKey := "groups"
	if hookCtx.OperationID == "updateResources" {
		itemsKey = "resources"
	}

	changed := false
	for _, item := range payload[itemsKey] {
		templateID, present := item["configuration_template_id"]
		if !present || bytes.Equal(templateID, []byte("null")) || bytes.Equal(templateID, []byte(`""`)) {
			continue
		}
		for _, field := range conflicts {
			if _, present := item[field]; present {
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
