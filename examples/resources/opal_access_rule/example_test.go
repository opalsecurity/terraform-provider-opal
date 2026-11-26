package opal_access_rule_test

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

// Speakeasy fills unknown string examples with "...my_<name>...". After the
// EPRD-5367 overlay maps query to a string, that becomes invalid JSON and the
// access_rule_query hook rejects it. Guard the committed example.
func TestAccessRuleExampleQueryIsUsable(t *testing.T) {
	t.Parallel()

	_, thisFile, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("runtime.Caller failed")
	}
	body, err := os.ReadFile(filepath.Join(filepath.Dir(thisFile), "resource.tf"))
	if err != nil {
		t.Fatalf("read resource.tf: %v", err)
	}
	text := string(body)
	if strings.Contains(text, "...my_query...") {
		t.Fatal(
			`examples/resources/opal_access_rule/resource.tf still has Speakeasy placeholder "...my_query..."; set a real jsonencode(...) example (and an overlay example: so regen does not reinstate it)`,
		)
	}
	if !strings.Contains(text, "jsonencode(") {
		t.Fatal(`expected query = jsonencode({...}) in the access_rule example`)
	}
}
