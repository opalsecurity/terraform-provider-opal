package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

func TestStringIDsFromTerraformCollection(t *testing.T) {
	idType := tftypes.Set{ElementType: tftypes.String}
	ids := []string{
		"7870617d-e72a-47f5-a84c-693817ab4567",
		"1520617d-e72a-47f5-a84c-693817ab48ad",
	}

	elems := make([]tftypes.Value, 0, len(ids))
	for _, id := range ids {
		elems = append(elems, tftypes.NewValue(tftypes.String, id))
	}

	got, ok := stringIDsFromTerraformCollection(tftypes.NewValue(idType, elems))
	if !ok {
		t.Fatal("expected to decode visibility_group_ids")
	}
	if len(got) != len(ids) {
		t.Fatalf("got %d ids, want %d", len(got), len(ids))
	}
	for i, id := range ids {
		gotID, ok := got[i].(types.String)
		if !ok || gotID.ValueString() != id {
			t.Fatalf("id[%d]=%v, want %s", i, got[i], id)
		}
	}

	empty, ok := stringIDsFromTerraformCollection(tftypes.NewValue(idType, []tftypes.Value{}))
	if !ok {
		t.Fatal("expected empty set to decode")
	}
	if len(empty) != 0 {
		t.Fatalf("got %d ids for empty set", len(empty))
	}
}

func TestAttributeOmitted(t *testing.T) {
	config := map[string]tftypes.Value{
		"visibility_group_ids": tftypes.NewValue(tftypes.Set{ElementType: tftypes.String}, nil),
		"visibility":           tftypes.NewValue(tftypes.String, "GLOBAL"),
	}
	if !attributeOmitted(config, "visibility_group_ids") {
		t.Fatal("null visibility_group_ids should be treated as omitted")
	}
	if attributeOmitted(config, "visibility") {
		t.Fatal("configured visibility should not be treated as omitted")
	}
	if !attributeOmitted(config, "missing") {
		t.Fatal("missing attribute should be treated as omitted")
	}
}
