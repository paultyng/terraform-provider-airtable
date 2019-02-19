package provider

import (
	"fmt"
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
				Config: testAccDataSourceTableConfig("BasicTest"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.airtable_table.test", "records.1.id", "recaiY04Qlfti0m9K"),
					resource.TestCheckResourceAttr("data.airtable_table.test", "records.1.fields.Name", "name-1"),
					resource.TestCheckResourceAttr("data.airtable_table.test", "records.1.fields.Notes", "foo"),

					resource.TestCheckOutput("records_length", "2"),
				),
			},
		},
	})
}

func TestAccDataSourceTable_pagination(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		// TODO: CheckDestroy:
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTableConfig("Pagination"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckOutput("records_length", "101"),
				),
			},
		},
	})
}

func testAccDataSourceTableConfig(view string) string {
	return fmt.Sprintf(`
data "airtable_table" "test" {
	"workspace_id" = "appOYVvt71h5txnFZ"
	"table"        = "Table 1"
	"view"         = "%s"
}

output "records_length" {
	value = "${length(data.airtable_table.test.records)}"
}
`, view)
}
