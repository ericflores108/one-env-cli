package op

import (
	"fmt"
	"os/exec"
	"testing"
)

func TestItems(t *testing.T) {
	// Capture the JSON output from the command
	cmd := exec.Command("op", "item", "list")
	output, err := cmd.Output()
	if err != nil {
		t.Errorf("Failed to run command: %v", err)
		return
	}

	// Print the JSON output
	fmt.Println("JSON Output:")
	fmt.Println(string(output))

	items, err := GetItems()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	} else {
		t.Logf("List: %+v", items)
	}
}
