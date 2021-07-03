---
page_title: "airtable_table Data Source - terraform-provider-airtable"
subcategory: ""
description: |-
  The airtable_table resource allows you to read information from a table for use in other Terraform resources.
---

# Data Source `airtable_table`

The `airtable_table` resource allows you to read information from a table for use in other Terraform resources.

## Example Usage

```terraform
data "airtable_table" "test" {
	workspace_id      = "appOYVvt71h5txnFZ"
	table             = "Table 1"
	fields            = ["some_field", "another field", "One more (field)"]
	filter_by_formula = "AND({another field} > 9000, {One more (field)} != 'over nine thousand!')"
}
```

## Schema

### Required

- **table** (String, Required) Name of the table from which to query information.
- **workspace_id** (String, Required) ID of the workspace in Airtable, you can find this in the base URL on your [Bases page](https://airtable.com/) when logged in.

### Optional

- **id** (String, Optional) The ID of this resource.
- **timeouts** (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))
- **view** (String, Optional) Name of the view from which to query information.
- **fields** (List, Optional) Name of the fields (or columns) returned from the table
- **filter_by_formula** (String, Optional) A formula used to filter records. The formula will be evaluated for each record, and if the result is not `0`, `false`, `""`, `NaN`, `[]`, or `#Error!` the record will be included in the response. See the [formula documentation](https://support.airtable.com/hc/en-us/articles/203255215-Formula-Field-Reference) from Airtable.

### Read-only

- **records** (List of Object, Read-only) Records in the table / view. (see [below for nested schema](#nestedatt--records))

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- **read** (String, Optional)


<a id="nestedatt--records"></a>
### Nested Schema for `records`

- **created_time** (String)
- **fields** (Map of String)
- **id** (String)
