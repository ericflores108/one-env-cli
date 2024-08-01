package op

import (
	"bytes"
	"encoding/json"
	"os/exec"
)

type Item struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

type Items []Item

type Field struct {
	ID        string `json:"id"`
	Type      string `json:"type"`
	Purpose   string `json:"purpose"`
	Label     string `json:"label"`
	Value     string `json:"value"`
	Reference string `json:"reference"`
}

type ItemResponse struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Fields []Field
}

// Provider - Get item to send to Plugin
func GetItem(name string) (ItemResponse, error) {
	cmd := exec.Command("op", "item", "get", name, "--format", "json")
	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		return ItemResponse{}, err
	}
	var resp ItemResponse
	err = json.Unmarshal(out.Bytes(), &resp)
	if err != nil {
		return ItemResponse{}, err
	}

	return resp, nil
}

// func GetItems() (Items, error) {
// 	var items Items
// 	cmd := exec.Command("op", "item", "list", "--format", "json")
// 	var out bytes.Buffer
// 	cmd.Stdout = &out

// 	err := cmd.Run()
// 	if err != nil {
// 		return items, err
// 	}

// 	err = json.Unmarshal(out.Bytes(), &items)
// 	if err != nil {
// 		return items, err
// 	}

// 	return items, nil
// }
