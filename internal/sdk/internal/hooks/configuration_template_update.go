package hooks

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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
