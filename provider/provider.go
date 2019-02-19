package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"

	"github.com/paultyng/terraform-provider-airtable/sdk"
)

type config struct {
	stopContext context.Context
	client      *sdk.Client
}

// New returns the provider instance.
func New() *schema.Provider {
	var p *schema.Provider
	p = &schema.Provider{
		Schema: map[string]*schema.Schema{
			"api_key": {
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
	p.ConfigureFunc = providerConfigure(p)
	return p
}

func providerConfigure(p *schema.Provider) schema.ConfigureFunc {
	return func(d *schema.ResourceData) (interface{}, error) {
		c := &config{
			stopContext: p.StopContext(),
			client:      sdk.NewClient(d.Get("api_key").(string), "terraform-provider-airtable/1.0", nil),
		}

		return c, nil
	}
}
