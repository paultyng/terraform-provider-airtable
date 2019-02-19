package provider

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccDataSourceTable_basic(t *testing.T) {
	//TODO: create the table via the API first? not sure if that is possible

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		// TODO: CheckDestroy:
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTableConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.airtable_table.test", "records.1.id", "recaiY04Qlfti0m9K"),
					resource.TestCheckResourceAttr("data.airtable_table.test", "records.1.fields.Name", "name-1"),
					resource.TestCheckResourceAttr("data.airtable_table.test", "records.1.fields.Notes", "foo"),
				),
			},
		},
	})
}

func testAccDataSourceTableConfig() string {
	return `
data "airtable_table" "test" {
	"workspace_id" = "appOYVvt71h5txnFZ"
	"table"        = "Table 1"
}
`
}
