package hooks

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"
)

type configurationTemplateUpdateOperation struct {
	itemsKey      string
	allowedFields map[string]struct{}
}

var configurationTemplateUpdateOperations = map[string]configurationTemplateUpdateOperation{
	"updateGroups": {
		itemsKey: "groups",
		allowedFields: fieldSet(
			"group_id",
			"configuration_template_id",
			"name",
			"description",
			"match_remote_name",
			"match_remote_description",
			"group_leader_user_ids",
			"risk_sensitivity_override",
		),
	},
	"updateResources": {
		itemsKey: "resources",
		allowedFields: fieldSet(
			"resource_id",
			"configuration_template_id",
			"name",
			"description",
			"match_remote_name",
			"match_remote_description",
			"parent_resource_id",
			"risk_sensitivity_override",
		),
	},
}

func fieldSet(fields ...string) map[string]struct{} {
	set := make(map[string]struct{}, len(fields))
	for _, field := range fields {
		set[field] = struct{}{}
	}
	return set
}

type configurationTemplateUpdateHook struct {
	mu                sync.RWMutex
	linkedGroupIDs    map[string]struct{}
	linkedResourceIDs map[string]struct{}
}

type configurationTemplateTransport struct {
	next HTTPClient
	hook *configurationTemplateUpdateHook
}

type configurationTemplateEntity struct {
	kind string
	id   string
}

func (h *configurationTemplateUpdateHook) SDKInit(baseURL string, client HTTPClient) (string, HTTPClient) {
	h.linkedGroupIDs = make(map[string]struct{})
	h.linkedResourceIDs = make(map[string]struct{})
	return baseURL, &configurationTemplateTransport{next: client, hook: h}
}

func (t *configurationTemplateTransport) Do(req *http.Request) (*http.Response, error) {
	if t.hook.shouldSkipFollowUp(req) {
		return &http.Response{
			StatusCode:    http.StatusOK,
			Status:        "200 OK",
			Header:        make(http.Header),
			Body:          http.NoBody,
			ContentLength: 0,
			Request:       req,
		}, nil
	}

	entities := configurationTemplateEntities(req)
	res, err := t.next.Do(req)
	if err == nil && res != nil && res.StatusCode >= 200 && res.StatusCode < 300 {
		t.hook.recordLinkedEntities(entities)
	}
	return res, err
}

func configurationTemplateEntities(req *http.Request) []configurationTemplateEntity {
	if req.Method != http.MethodPut || req.Body == nil {
		return nil
	}

	segments := pathSegments(req.URL.Path)
	if len(segments) == 0 {
		return nil
	}

	operation, ok := map[string]configurationTemplateUpdateOperation{
		"groups":    configurationTemplateUpdateOperations["updateGroups"],
		"resources": configurationTemplateUpdateOperations["updateResources"],
	}[segments[len(segments)-1]]
	if !ok {
		return nil
	}

	body, err := io.ReadAll(req.Body)
	if err != nil {
		return nil
	}
	req.Body = io.NopCloser(bytes.NewReader(body))

	var payload map[string][]map[string]json.RawMessage
	if json.Unmarshal(body, &payload) != nil {
		return nil
	}

	idKey := strings.TrimSuffix(operation.itemsKey, "s") + "_id"
	entities := make([]configurationTemplateEntity, 0, len(payload[operation.itemsKey]))
	for _, item := range payload[operation.itemsKey] {
		var templateID, entityID string
		if json.Unmarshal(item["configuration_template_id"], &templateID) != nil || templateID == "" ||
			json.Unmarshal(item[idKey], &entityID) != nil || entityID == "" {
			continue
		}
		entities = append(entities, configurationTemplateEntity{kind: operation.itemsKey, id: entityID})
	}
	return entities
}

func (h *configurationTemplateUpdateHook) recordLinkedEntities(entities []configurationTemplateEntity) {
	h.mu.Lock()
	defer h.mu.Unlock()
	for _, entity := range entities {
		if entity.kind == "groups" {
			h.linkedGroupIDs[entity.id] = struct{}{}
		} else {
			h.linkedResourceIDs[entity.id] = struct{}{}
		}
	}
}

func (h *configurationTemplateUpdateHook) shouldSkipFollowUp(req *http.Request) bool {
	if req.Method != http.MethodPut {
		return false
	}
	segments := pathSegments(req.URL.Path)
	if len(segments) < 3 {
		return false
	}
	entityKind, entityID, operation := segments[len(segments)-3], segments[len(segments)-2], segments[len(segments)-1]

	h.mu.RLock()
	defer h.mu.RUnlock()
	switch entityKind {
	case "groups":
		if operation != "message-channels" && operation != "on-call-schedules" && operation != "visibility" {
			return false
		}
		_, linked := h.linkedGroupIDs[entityID]
		return linked
	case "resources":
		if operation != "visibility" {
			return false
		}
		_, linked := h.linkedResourceIDs[entityID]
		return linked
	default:
		return false
	}
}

func pathSegments(path string) []string {
	return strings.Split(strings.Trim(path, "/"), "/")
}

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
			if _, allowed := operation.allowedFields[field]; !allowed {
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
