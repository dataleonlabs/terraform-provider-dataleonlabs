// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package internal

import (
	"context"
	"os"

	"github.com/dataleonlabs/dataleonlabs-go"
	"github.com/dataleonlabs/dataleonlabs-go/option"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stainless-sdks/dataleonlabs-terraform/internal/services/company"
	"github.com/stainless-sdks/dataleonlabs-terraform/internal/services/individual"
)

var _ provider.ProviderWithConfigValidators = (*DataleonlabsProvider)(nil)

// DataleonlabsProvider defines the provider implementation.
type DataleonlabsProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// DataleonlabsProviderModel describes the provider data model.
type DataleonlabsProviderModel struct {
	BaseURL types.String `tfsdk:"base_url" json:"base_url,optional"`
	APIKey  types.String `tfsdk:"api_key" json:"api_key,optional"`
}

func (p *DataleonlabsProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "dataleonlabs"
	resp.Version = p.version
}

func ProviderSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"base_url": schema.StringAttribute{
				Description: "Set the base url that the provider connects to.",
				Optional:    true,
			},
			"api_key": schema.StringAttribute{
				Optional: true,
			},
		},
	}
}

func (p *DataleonlabsProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = ProviderSchema(ctx)
}

func (p *DataleonlabsProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {

	var data DataleonlabsProviderModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	opts := []option.RequestOption{}

	if !data.BaseURL.IsNull() && !data.BaseURL.IsUnknown() {
		opts = append(opts, option.WithBaseURL(data.BaseURL.ValueString()))
	} else if o, ok := os.LookupEnv("DATALEONLABS_BASE_URL"); ok {
		opts = append(opts, option.WithBaseURL(o))
	}

	if !data.APIKey.IsNull() && !data.APIKey.IsUnknown() {
		opts = append(opts, option.WithAPIKey(data.APIKey.ValueString()))
	} else if o, ok := os.LookupEnv("DATALEONLABS_API_KEY"); ok {
		opts = append(opts, option.WithAPIKey(o))
	} else {
		resp.Diagnostics.AddAttributeError(
			path.Root("api_key"),
			"Missing api_key value",
			"The api_key field is required. Set it in provider configuration or via the \"DATALEONLABS_API_KEY\" environment variable.",
		)
		return
	}

	client := dataleonlabs.NewClient(
		opts...,
	)

	resp.DataSourceData = &client
	resp.ResourceData = &client
}

func (p *DataleonlabsProvider) ConfigValidators(_ context.Context) []provider.ConfigValidator {
	return []provider.ConfigValidator{}
}

func (p *DataleonlabsProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		company.NewResource,
		individual.NewResource,
	}
}

func (p *DataleonlabsProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		company.NewCompanyDataSource,
		individual.NewIndividualDataSource,
	}
}

func NewProvider(version string) func() provider.Provider {
	return func() provider.Provider {
		return &DataleonlabsProvider{
			version: version,
		}
	}
}
