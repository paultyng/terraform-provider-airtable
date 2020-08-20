---
layout: ""
page_title: "Provider: Airtable"
description: |-
  The Airtable provider provides resources to interact with a Airtable API.
---

# Airtable Provider

The Airtable provider provides resources to interact with the Airtable API.

## Example Usage

```terraform
data "airtable_table" "test" {
	"workspace_id" = "appOYVvt71h5txnFZ"
	"table"        = "Table 1"
}
```

## Schema

### Optional

- **api_key** (String, Optional) API key from your [account](https://airtable.com/account) page. You can set this via the `AIRTABLE_API_KEY` environment variable as well.
