package provider

import (
	"os"
	"testing"
)

func testAccPreCheck(t *testing.T) {
	if v := os.Getenv("AIRTABLE_API_KEY"); v == "" {
		t.Fatal("AIRTABLE_API_KEY environment variable must be set")
	}
}
