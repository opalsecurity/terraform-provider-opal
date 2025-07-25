// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package sdk

// Generated from OpenAPI doc version 1.0 and generator version 2.663.0

import (
	"context"
	"fmt"
	"github.com/opalsecurity/terraform-provider-opal/v3/internal/sdk/internal/config"
	"github.com/opalsecurity/terraform-provider-opal/v3/internal/sdk/internal/hooks"
	"github.com/opalsecurity/terraform-provider-opal/v3/internal/sdk/internal/utils"
	"github.com/opalsecurity/terraform-provider-opal/v3/internal/sdk/models/shared"
	"github.com/opalsecurity/terraform-provider-opal/v3/internal/sdk/retry"
	"net/http"
	"time"
)

const (
	// Production
	ServerProd string = "prod"
)

// ServerList contains the list of servers available to the SDK
var ServerList = map[string]string{
	ServerProd: "https://api.opal.dev/v1",
}

// HTTPClient provides an interface for supplying the SDK with a custom HTTP client
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// String provides a helper function to return a pointer to a string
func String(s string) *string { return &s }

// Bool provides a helper function to return a pointer to a bool
func Bool(b bool) *bool { return &b }

// Int provides a helper function to return a pointer to an int
func Int(i int) *int { return &i }

// Int64 provides a helper function to return a pointer to an int64
func Int64(i int64) *int64 { return &i }

// Float32 provides a helper function to return a pointer to a float32
func Float32(f float32) *float32 { return &f }

// Float64 provides a helper function to return a pointer to a float64
func Float64(f float64) *float64 { return &f }

// Pointer provides a helper function to return a pointer to a type
func Pointer[T any](v T) *T { return &v }

// OpalAPI - Opal API: The Opal API is a RESTful API that allows you to interact with the Opal Security platform programmatically.
type OpalAPI struct {
	SDKVersion string
	// Operations related to access rules
	AccessRules *AccessRules
	// Operations related to apps
	Apps *Apps
	// Operations related to bundles
	Bundles *Bundles
	// Operations related to configuration templates
	ConfigurationTemplates *ConfigurationTemplates
	// Operations related to events
	Events *Events
	// Operations related to group bindings
	GroupBindings *GroupBindings
	// Operations related to groups
	Groups *Groups
	// Operations related to IDP group mappings
	IdpGroupMappings *IdpGroupMappings
	// Operations related to message channels
	MessageChannels *MessageChannels
	// Operations related to non-human identities
	NonHumanIdentities *NonHumanIdentities
	// Operations related to on-call schedules
	OnCallSchedules *OnCallSchedules
	// Operations related to owners
	Owners *Owners
	// Operations related to requests
	Requests *Requests
	// Operations related to resources
	Resources *Resources
	// Operations related to scoped role permissions
	ScopedRolePermissions *ScopedRolePermissions
	// Operations related to sessions
	Sessions *Sessions
	// Operations related to tags
	Tags *Tags
	// Operations related to UARs
	Uars *Uars
	// Operations related to users
	Users *Users

	sdkConfiguration config.SDKConfiguration
	hooks            *hooks.Hooks
}

type SDKOption func(*OpalAPI)

// WithServerURL allows the overriding of the default server URL
func WithServerURL(serverURL string) SDKOption {
	return func(sdk *OpalAPI) {
		sdk.sdkConfiguration.ServerURL = serverURL
	}
}

// WithTemplatedServerURL allows the overriding of the default server URL with a templated URL populated with the provided parameters
func WithTemplatedServerURL(serverURL string, params map[string]string) SDKOption {
	return func(sdk *OpalAPI) {
		if params != nil {
			serverURL = utils.ReplaceParameters(serverURL, params)
		}

		sdk.sdkConfiguration.ServerURL = serverURL
	}
}

// WithServer allows the overriding of the default server by name
func WithServer(server string) SDKOption {
	return func(sdk *OpalAPI) {
		_, ok := ServerList[server]
		if !ok {
			panic(fmt.Errorf("server %s not found", server))
		}

		sdk.sdkConfiguration.Server = server
	}
}

// WithClient allows the overriding of the default HTTP client used by the SDK
func WithClient(client HTTPClient) SDKOption {
	return func(sdk *OpalAPI) {
		sdk.sdkConfiguration.Client = client
	}
}

// WithSecurity configures the SDK to use the provided security details
func WithSecurity(security shared.Security) SDKOption {
	return func(sdk *OpalAPI) {
		sdk.sdkConfiguration.Security = utils.AsSecuritySource(security)
	}
}

// WithSecuritySource configures the SDK to invoke the Security Source function on each method call to determine authentication
func WithSecuritySource(security func(context.Context) (shared.Security, error)) SDKOption {
	return func(sdk *OpalAPI) {
		sdk.sdkConfiguration.Security = func(ctx context.Context) (interface{}, error) {
			return security(ctx)
		}
	}
}

func WithRetryConfig(retryConfig retry.Config) SDKOption {
	return func(sdk *OpalAPI) {
		sdk.sdkConfiguration.RetryConfig = &retryConfig
	}
}

// WithTimeout Optional request timeout applied to each operation
func WithTimeout(timeout time.Duration) SDKOption {
	return func(sdk *OpalAPI) {
		sdk.sdkConfiguration.Timeout = &timeout
	}
}

// New creates a new instance of the SDK with the provided options
func New(opts ...SDKOption) *OpalAPI {
	sdk := &OpalAPI{
		SDKVersion: "3.1.0",
		sdkConfiguration: config.SDKConfiguration{
			UserAgent:  "speakeasy-sdk/terraform 3.1.0 2.663.0 1.0 github.com/opalsecurity/terraform-provider-opal/v3/internal/sdk",
			ServerList: ServerList,
		},
		hooks: hooks.New(),
	}
	for _, opt := range opts {
		opt(sdk)
	}

	// Use WithClient to override the default client if you would like to customize the timeout
	if sdk.sdkConfiguration.Client == nil {
		sdk.sdkConfiguration.Client = &http.Client{Timeout: 60 * time.Second}
	}

	currentServerURL, _ := sdk.sdkConfiguration.GetServerDetails()
	serverURL := currentServerURL
	serverURL, sdk.sdkConfiguration.Client = sdk.hooks.SDKInit(currentServerURL, sdk.sdkConfiguration.Client)
	if currentServerURL != serverURL {
		sdk.sdkConfiguration.ServerURL = serverURL
	}

	sdk.AccessRules = newAccessRules(sdk, sdk.sdkConfiguration, sdk.hooks)
	sdk.Apps = newApps(sdk, sdk.sdkConfiguration, sdk.hooks)
	sdk.Bundles = newBundles(sdk, sdk.sdkConfiguration, sdk.hooks)
	sdk.ConfigurationTemplates = newConfigurationTemplates(sdk, sdk.sdkConfiguration, sdk.hooks)
	sdk.Events = newEvents(sdk, sdk.sdkConfiguration, sdk.hooks)
	sdk.GroupBindings = newGroupBindings(sdk, sdk.sdkConfiguration, sdk.hooks)
	sdk.Groups = newGroups(sdk, sdk.sdkConfiguration, sdk.hooks)
	sdk.IdpGroupMappings = newIdpGroupMappings(sdk, sdk.sdkConfiguration, sdk.hooks)
	sdk.MessageChannels = newMessageChannels(sdk, sdk.sdkConfiguration, sdk.hooks)
	sdk.NonHumanIdentities = newNonHumanIdentities(sdk, sdk.sdkConfiguration, sdk.hooks)
	sdk.OnCallSchedules = newOnCallSchedules(sdk, sdk.sdkConfiguration, sdk.hooks)
	sdk.Owners = newOwners(sdk, sdk.sdkConfiguration, sdk.hooks)
	sdk.Requests = newRequests(sdk, sdk.sdkConfiguration, sdk.hooks)
	sdk.Resources = newResources(sdk, sdk.sdkConfiguration, sdk.hooks)
	sdk.ScopedRolePermissions = newScopedRolePermissions(sdk, sdk.sdkConfiguration, sdk.hooks)
	sdk.Sessions = newSessions(sdk, sdk.sdkConfiguration, sdk.hooks)
	sdk.Tags = newTags(sdk, sdk.sdkConfiguration, sdk.hooks)
	sdk.Uars = newUars(sdk, sdk.sdkConfiguration, sdk.hooks)
	sdk.Users = newUsers(sdk, sdk.sdkConfiguration, sdk.hooks)

	return sdk
}
