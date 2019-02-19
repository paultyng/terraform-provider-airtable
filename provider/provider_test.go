package provider

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

var testAccProviders map[string]terraform.ResourceProvider
var testAccProvider *schema.Provider

func init() {
	testAccProvider = New()
	testAccProviders = map[string]terraform.ResourceProvider{
		"airtable": testAccProvider,
	}
}

func TestProvider(t *testing.T) {
	if err := New().InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestProvider_impl(t *testing.T) {
	var _ terraform.ResourceProvider = New()
}

func testAccPreCheck(t *testing.T) {
	if v := os.Getenv("AIRTABLE_API_KEY"); v == "" {
		t.Fatal("AIRTABLE_API_KEY environment variable must be set")
	}
}
