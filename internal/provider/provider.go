// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/opalsecurity/terraform-provider-opal/internal/sdk"
	"github.com/opalsecurity/terraform-provider-opal/internal/sdk/models/shared"
	"net/http"
	"os"
)

var _ provider.Provider = &OpalProvider{}

type OpalProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// OpalProviderModel describes the provider data model.
type OpalProviderModel struct {
	ServerURL  types.String `tfsdk:"server_url"`
	BearerAuth types.String `tfsdk:"bearer_auth"`
}

func (p *OpalProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "opal"
	resp.Version = p.version
}

func (p *OpalProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: `Opal API: Your Home For Developer Resources.`,
		Attributes: map[string]schema.Attribute{
			"server_url": schema.StringAttribute{
				MarkdownDescription: "Server URL (defaults to https://api.opal.dev/v1)",
				Optional:            true,
				Required:            false,
			},
			"bearer_auth": schema.StringAttribute{
				Optional:  true,
				Sensitive: true,
			},
		},
	}
}

func (p *OpalProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var data OpalProviderModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	ServerURL := data.ServerURL.ValueString()

	if ServerURL == "" {
		ServerURL = "https://api.opal.dev/v1"
	}

	bearerAuth := new(string)
	if !data.BearerAuth.IsUnknown() && !data.BearerAuth.IsNull() {
		*bearerAuth = data.BearerAuth.ValueString()
	} else {
		if len(os.Getenv("OPAL_AUTH_TOKEN")) > 0 {
			*bearerAuth = os.Getenv("OPAL_AUTH_TOKEN")
		} else {
			bearerAuth = nil
		}
	}
	security := shared.Security{
		BearerAuth: bearerAuth,
	}

	opts := []sdk.SDKOption{
		sdk.WithServerURL(ServerURL),
		sdk.WithSecurity(security),
		sdk.WithClient(http.DefaultClient),
	}
	client := sdk.New(opts...)

	resp.DataSourceData = client
	resp.ResourceData = client
}

func (p *OpalProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewConfigurationTemplateResource,
		NewGroupResource,
		NewGroupResourceListResource,
		NewGroupTagResource,
		NewGroupUserResource,
		NewMessageChannelResource,
		NewOnCallScheduleResource,
		NewOwnerResource,
		NewResourceResource,
		NewResourceTagResource,
		NewTagResource,
		NewTagUserResource,
	}
}

func (p *OpalProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewAppDataSource,
		NewAppsDataSource,
		NewConfigurationTemplateListDataSource,
		NewEventsDataSource,
		NewGroupDataSource,
		NewGroupListDataSource,
		NewGroupResourceListDataSource,
		NewGroupReviewersStagesListDataSource,
		NewGroupTagsDataSource,
		NewGroupUsersDataSource,
		NewMessageChannelDataSource,
		NewMessageChannelListDataSource,
		NewOnCallScheduleDataSource,
		NewOnCallScheduleListDataSource,
		NewOwnerDataSource,
		NewOwnerFromNameDataSource,
		NewOwnersDataSource,
		NewRequestsDataSource,
		NewResourceDataSource,
		NewResourceMessageChannelListDataSource,
		NewResourceReviewersListDataSource,
		NewResourcesListDataSource,
		NewResourcesAccessStatusDataSource,
		NewResourcesUsersListDataSource,
		NewResourceTagsDataSource,
		NewResourceVisibilityDataSource,
		NewSessionsDataSource,
		NewTagDataSource,
		NewTagsListDataSource,
		NewUarDataSource,
		NewUARSListDataSource,
		NewUserDataSource,
		NewUsersDataSource,
		NewUserTagsDataSource,
	}
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &OpalProvider{
			version: version,
		}
	}
}
