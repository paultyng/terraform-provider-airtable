# Airtable Terraform Provider

## Data Sources

### `airtable_table`

```hcl
data "airtable_table" "test" {
	"workspace_id" = "appOYVvt71h5txnFZ"
	"table"        = "Table 1"
}
```

#### Arguments

* `workspace_id`
* `table`
* `view`

#### Attributes

* `records`:
  * `id`
  * `created_time`: RFC3339 time
  * `fields`: map of fields and values