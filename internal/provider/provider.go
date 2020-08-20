package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/paultyng/terraform-provider-airtable/sdk"
)

type config struct {
	client *sdk.Client
}

// New returns the provider instance.
func New() *schema.Provider {
	var p *schema.Provider
	p = &schema.Provider{
		Schema: map[string]*schema.Schema{
			"api_key": {
				Description: "API key from your [account](https://airtable.com/account) page. " +
					"You can set this via the `AIRTABLE_API_KEY` environment variable as well.",
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("AIRTABLE_API_KEY", nil),
				Sensitive:   true,
			},
		},

		DataSourcesMap: map[string]*schema.Resource{
			"airtable_table": dataSourceTable(),
		},
	}
	p.ConfigureContextFunc = providerConfigure(p)
	return p
}

func providerConfigure(p *schema.Provider) schema.ConfigureContextFunc {
	return func(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
		c := &config{
			client: sdk.NewClient(d.Get("api_key").(string), "terraform-provider-airtable/1.0", nil),
		}

		return c, nil
	}
}
