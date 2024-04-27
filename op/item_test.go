package op

import (
	"encoding/json"
	"testing"
)

func TestGetItems(t *testing.T) {
	items, err := GetItems()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	} else {
		jsonData, err := json.MarshalIndent(items, "", "  ")
		if err != nil {
			t.Errorf("Error marshaling JSON: %v", err)
		} else {
			t.Logf("List:\n%s", string(jsonData))
		}
	}
}

func TestGetItem(t *testing.T) {
	item, err := GetItem("POSTMAN_API_KEY")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	} else {
		jsonData, err := json.MarshalIndent(item, "", "  ")
		if err != nil {
			t.Errorf("Error marshaling JSON: %v", err)
		} else {
			t.Logf("Item:\n%s", string(jsonData))
		}
	}
}
