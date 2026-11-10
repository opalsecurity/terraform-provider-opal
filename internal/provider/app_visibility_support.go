package provider

// appTypesSupportingVisibilityEndpoint are the only app types for which
// GET/PUT /apps/{id}/visibility is allowed by the API. The managed opal_app
// resource is likewise limited to these types (create + import).
var appTypesSupportingVisibilityEndpoint = map[string]struct{}{
	"CUSTOM":           {},
	"CUSTOM_CONNECTOR": {},
}

// appTypeSupportsVisibilityEndpoint reports whether the visibility sub-endpoint
// (and therefore the managed opal_app resource) supports the given app type.
func appTypeSupportsVisibilityEndpoint(appType string) bool {
	_, ok := appTypesSupportingVisibilityEndpoint[appType]
	return ok
}
