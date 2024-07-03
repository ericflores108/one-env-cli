package postman

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestWorkspace(t *testing.T) {
	// Mock server to simulate Postman API
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/workspaces" {
			if r.Method != "GET" {
				t.Errorf("Expected GET request, got %s", r.Method)
			}
			workspaces := Workspaces{
				Workspaces: []PostmanWorkspace{
					{ID: "123", Name: "Test Workspace"},
					{ID: "456", Name: "Another Workspace"},
				},
			}
			json.NewEncoder(w).Encode(workspaces)
		} else {
			t.Errorf("Unexpected request to %s", r.URL.Path)
		}
	}))
	defer server.Close()

	// Override the BaseURL for testing
	originalBaseURL := BaseURL
	BaseURL = server.URL
	defer func() { BaseURL = originalBaseURL }()

	workspaceID := Workspace("Non-existent Workspace")
	if workspaceID != "" {
		t.Errorf("Expected empty workspace ID for non-existent workspace, got '%s'", workspaceID)
	}
}

func TestWorkspaceEmptyName(t *testing.T) {
	workspaceID := Workspace("")
	if workspaceID != "" {
		t.Errorf("Expected empty workspace ID for empty name, got '%s'", workspaceID)
	}
}
