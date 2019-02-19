package provider

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform/helper/schema"

	"github.com/paultyng/terraform-provider-airtable/sdk"
)

func dataSourceTable() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceTableRead,

		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(30 * time.Second),
		},

		Schema: map[string]*schema.Schema{
			"workspace_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"table": {
				Type:     schema.TypeString,
				Required: true,
			},
			"view": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"records": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"created_time": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"fields": {
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

func dataSourceTableRead(d *schema.ResourceData, meta interface{}) error {
	c := meta.(*config)
	workspaceID := d.Get("workspace_id").(string)
	table := d.Get("table").(string)
	view := d.Get("view").(string)

	options := &sdk.ListRecordsOptions{
		View: view,
	}

	sdkRecords, err := c.client.ListRecords(workspaceID, table, options)
	if err != nil {
		return err
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
