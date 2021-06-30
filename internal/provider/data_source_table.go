package provider

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/paultyng/terraform-provider-airtable/sdk"
)

func dataSourceTable() *schema.Resource {
	return &schema.Resource{
		Description: "The `airtable_table` resource allows you to read information from a table for use " +
			"in other Terraform resources.",

		ReadContext: dataSourceTableRead,

		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(30 * time.Second),
		},

		Schema: map[string]*schema.Schema{
			"workspace_id": {
				Description: "ID of the workspace in Airtable, you can find this in the base URL on " +
					"your [Bases page](https://airtable.com/) when logged in.",
				Type:     schema.TypeString,
				Required: true,
			},
			"table": {
				Description: "Name of the table from which to query information.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"view": {
				Description: "Name of the view from which to query information.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"fields": {
				Description: "Only data for fields whose names are in this list will be included in the result." +
					"If you don't need every field, you can use this parameter to reduce the amount of data transferred.",
				Type:     schema.TypeString,
				Optional: true,
			},
			"filterByFormula": {
				Description: "A formula used to filter records. If combined with the `view` parameter, " +
					"only records in that new which satisfy the formula will be returned",
				Optional: true,
			},

			"records": {
				Description: "Records in the table / view.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Description: "ID of the record.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"created_time": {
							Description: "Record created time, in RFC 3339 format.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"fields": {
							Description: "Map of key/value pairs, each key is a column name, each value " +
								"is the rows value.",
							Type:     schema.TypeMap,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceTableRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*config)
	workspaceID := d.Get("workspace_id").(string)
	table := d.Get("table").(string)
	view := d.Get("view").(string)

	options := &sdk.ListRecordsOptions{
		View: view,
	}

	sdkRecords, err := c.client.ListRecords(workspaceID, table, options)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(fmt.Sprintf("%s/%s", workspaceID, table))
	d.Set("records", flattenRecords(sdkRecords))
	return nil
}

func flattenRecords(sdkRecords []sdk.Record) []interface{} {
	records := make([]interface{}, 0, len(sdkRecords))
	for _, sdkR := range sdkRecords {
		fields := map[string]string{}

		for name, value := range sdkR.Fields {
			fields[name] = value
		}

		records = append(records, map[string]interface{}{
			"id":           sdkR.ID,
			"created_time": sdkR.CreatedTime.Format(time.RFC3339),
			"fields":       fields,
		})
	}
	return records
}
