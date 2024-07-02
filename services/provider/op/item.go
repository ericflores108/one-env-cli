package op

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os/exec"
)

const (
	opCommand  = "op"
	jsonFormat = "json"
)

// Item represents a basic 1Password item
type Item struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

// Items is a slice of Item pointers
type Items []*Item

// Field represents a field in a 1Password item
type Field struct {
	ID        string `json:"id"`
	Type      string `json:"type"`
	Purpose   string `json:"purpose"`
	Label     string `json:"label"`
	Value     string `json:"value"`
	Reference string `json:"reference"`
}

// ItemResponse represents a detailed 1Password item
type ItemResponse struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Fields []Field `json:"fields"`
}

// GetItems retrieves all items from 1Password
func GetItems() (Items, error) {
	cmd := exec.Command(opCommand, "item", "list", "--format", jsonFormat)
	var out bytes.Buffer
	cmd.Stdout = &out

	if err := cmd.Run(); err != nil {
		return nil, fmt.Errorf("failed to run op command: %w", err)
	}

	var items Items
	if err := json.Unmarshal(out.Bytes(), &items); err != nil {
		return nil, fmt.Errorf("failed to unmarshal items: %w", err)
	}

	return items, nil
}

// GetItem retrieves a specific item from 1Password by name
func GetItem(name string) (*ItemResponse, error) {
	cmd := exec.Command(opCommand, "item", "get", name, "--format", jsonFormat)
	var out bytes.Buffer
	cmd.Stdout = &out

	if err := cmd.Run(); err != nil {
		return nil, fmt.Errorf("failed to run op command: %w", err)
	}

	var resp ItemResponse
	if err := json.Unmarshal(out.Bytes(), &resp); err != nil {
		return nil, fmt.Errorf("failed to unmarshal item response: %w", err)
	}

	return &resp, nil
}
