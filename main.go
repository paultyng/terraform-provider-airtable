package main // import "github.com/paultyng/terraform-provider-airtable"

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/hashicorp/terraform/terraform"

	"github.com/paultyng/terraform-provider-airtable/provider"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() terraform.ResourceProvider {
			return provider.New()
		},
	})
}
