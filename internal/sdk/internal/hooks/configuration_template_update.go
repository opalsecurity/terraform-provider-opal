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

// configurationTemplateUpdateHook keeps the parent PUT /groups and PUT
// /resources calls (and their sibling PUTs) from sending template-owned
// settings to an entity that is linked to a configuration template.
//
// The link itself is managed through PUT /{groups|resources}/{id}/configuration-template.
// The provider runs that call first on Update, so the hook learns the entity's
// link state from it (SDKInit wraps the HTTP client to observe the calls) and
// then:
//
//   - BeforeRequest strips every field outside the allowlist from the parent
//     PUT of a linked entity. The backend rejects those fields while an entity
//     is linked, and Terraform state still carries the pre-attach values.
//   - the transport short-circuits the visibility / message-channel /
//     on-call-schedule sibling PUTs for a linked entity, because the template
//     owns those settings.
//   - the transport also short-circuits a visibility PUT that carries no
//     visibility value at all. That is what Create sends for an entity that is
//     about to be linked (the link is the last call of Create).
//
// An unlinked entity (including one just unlinked by a null template PUT) is
// left untouched so its own configuration can be written.
type configurationTemplateUpdateOperation struct {
	itemsKey      string
	allowedFields map[string]struct{}
}

var configurationTemplateUpdateOperations = map[string]configurationTemplateUpdateOperation{
	"updateGroups": {
		itemsKey: "groups",
		allowedFields: fieldSet(
			"group_id",
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
			"name",
			"description",
			"match_remote_name",
			"match_remote_description",
			"parent_resource_id",
			"risk_sensitivity_override",
		),
	},
}

// siblingOperations are the sibling PUTs a template owns, per entity kind.
var configurationTemplateSiblingOperations = map[string]map[string]struct{}{
	"groups":    fieldSet("message-channels", "on-call-schedules", "visibility"),
	"resources": fieldSet("visibility"),
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

func (h *configurationTemplateUpdateHook) SDKInit(
	baseURL string,
	client HTTPClient,
) (string, HTTPClient) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.linkedGroupIDs = make(map[string]struct{})
	h.linkedResourceIDs = make(map[string]struct{})
	return baseURL, &configurationTemplateTransport{next: client, hook: h}
}

func (t *configurationTemplateTransport) Do(req *http.Request) (*http.Response, error) {
	if t.hook.shouldSkipSibling(req) {
		return &http.Response{
			StatusCode:    http.StatusOK,
			Status:        "200 OK",
			Header:        make(http.Header),
			Body:          http.NoBody,
			ContentLength: 0,
			Request:       req,
		}, nil
	}

	link, observed := templateLinkFromRequest(req)
	res, err := t.next.Do(req)
	if observed && err == nil && res != nil && res.StatusCode >= 200 && res.StatusCode < 300 {
		t.hook.recordLink(link)
	}
	return res, err
}

type configurationTemplateLink struct {
	kind   string
	id     string
	linked bool
}

// templateLinkFromRequest recognizes PUT /{groups|resources}/{id}/configuration-template
// and reports whether the body links (non-empty id) or unlinks (null / empty).
func templateLinkFromRequest(req *http.Request) (configurationTemplateLink, bool) {
	if req.Method != http.MethodPut || req.Body == nil {
		return configurationTemplateLink{}, false
	}
	kind, id, operation, ok := entityOperation(req)
	if !ok || operation != "configuration-template" {
		return configurationTemplateLink{}, false
	}

	body, err := io.ReadAll(req.Body)
	if err != nil {
		return configurationTemplateLink{}, false
	}
	req.Body = io.NopCloser(bytes.NewReader(body))

	var payload struct {
		ConfigurationTemplateID *string `json:"configuration_template_id"`
	}
	if json.Unmarshal(body, &payload) != nil {
		return configurationTemplateLink{}, false
	}
	return configurationTemplateLink{
		kind:   kind,
		id:     id,
		linked: payload.ConfigurationTemplateID != nil && *payload.ConfigurationTemplateID != "",
	}, true
}

func (h *configurationTemplateUpdateHook) recordLink(link configurationTemplateLink) {
	h.mu.Lock()
	defer h.mu.Unlock()
	set := h.linkedGroupIDs
	if link.kind == "resources" {
		set = h.linkedResourceIDs
	}
	if link.linked {
		set[link.id] = struct{}{}
	} else {
		delete(set, link.id)
	}
}

func (h *configurationTemplateUpdateHook) isLinked(kind, id string) bool {
	h.mu.RLock()
	defer h.mu.RUnlock()
	switch kind {
	case "groups":
		_, ok := h.linkedGroupIDs[id]
		return ok
	case "resources":
		_, ok := h.linkedResourceIDs[id]
		return ok
	default:
		return false
	}
}

func (h *configurationTemplateUpdateHook) shouldSkipSibling(req *http.Request) bool {
	if req.Method != http.MethodPut {
		return false
	}
	kind, id, operation, ok := entityOperation(req)
	if !ok {
		return false
	}
	if _, owned := configurationTemplateSiblingOperations[kind][operation]; !owned {
		return false
	}
	if h.isLinked(kind, id) {
		return true
	}
	return operation == "visibility" && visibilityBodyIsEmpty(req)
}

// visibilityBodyIsEmpty is true when a visibility PUT has no visibility value.
// The backend rejects that with "Unrecognized visibility level", and it only
// happens for an entity whose configuration comes from a template.
func visibilityBodyIsEmpty(req *http.Request) bool {
	if req.Body == nil || req.Body == http.NoBody {
		return true
	}
	body, err := io.ReadAll(req.Body)
	if err != nil {
		return false
	}
	req.Body = io.NopCloser(bytes.NewReader(body))

	var payload struct {
		Visibility *string `json:"visibility"`
	}
	if json.Unmarshal(body, &payload) != nil {
		return false
	}
	return payload.Visibility == nil || *payload.Visibility == ""
}

// entityOperation splits .../{groups|resources}/{id}/{operation}.
func entityOperation(req *http.Request) (kind, id, operation string, ok bool) {
	segments := pathSegments(req.URL.Path)
	if len(segments) < 3 {
		return "", "", "", false
	}
	kind, id, operation = segments[len(segments)-3], segments[len(segments)-2], segments[len(segments)-1]
	if kind != "groups" && kind != "resources" {
		return "", "", "", false
	}
	return kind, id, operation, true
}

func pathSegments(path string) []string {
	return strings.Split(strings.Trim(path, "/"), "/")
}

func (h *configurationTemplateUpdateHook) BeforeRequest(
	hookCtx BeforeRequestContext,
	req *http.Request,
) (*http.Request, error) {
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

	idKey := strings.TrimSuffix(operation.itemsKey, "s") + "_id"
	changed := false
	for _, item := range payload[operation.itemsKey] {
		var entityID string
		if json.Unmarshal(item[idKey], &entityID) != nil ||
			!h.isLinked(operation.itemsKey, entityID) {
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
