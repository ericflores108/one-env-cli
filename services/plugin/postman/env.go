package postman

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Environment struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
	Owner     string `json:"owner"`
	UID       string `json:"uid"`
	IsPublic  bool   `json:"isPublic"`
}

type EnvironmentsResponse struct {
	Environments []Environment `json:"environments"`
}

type EnvironmentType string

type Workspaces struct {
	Workspaces []PostmanWorkspace `json:"workspaces"`
}

type PostmanWorkspace struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Type       string `json:"type"`
	Visibility string `json:"visibility"`
	CreatedBy  string `json:"createdBy"`
}

type WorkspaceID string

const (
	SecretType  EnvironmentType = "secret"
	DefaultType EnvironmentType = "default"
)

type EnvironmentVariable struct {
	Key     string          `json:"key"`
	Value   string          `json:"value"`
	Enabled bool            `json:"enabled"`
	Type    EnvironmentType `json:"type"`
}

type EnvironmentData struct {
	Name   string                `json:"name"`
	Values []EnvironmentVariable `json:"values"`
}

type CreateEnvironmentRequest struct {
	Environment EnvironmentData `json:"environment"`
}

func Workspace(name string) WorkspaceID {
	if name == "" {
		return ""
	}
	resp, err := makeRequest("GET", "/workspaces", nil)
	if err != nil {
		fmt.Println("An error occurred. Using default workspace in Postman. ", err.Error())
		return ""
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		fmt.Println("Using default workspace in Postman. Status code: ", resp.StatusCode)
		return ""
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("An error occurred. Using default workspace in Postman. ", err.Error())
		return ""
	}
	var workspaces Workspaces
	err = json.Unmarshal(body, &workspaces)
	if err != nil {
		fmt.Println("An error occurred. Using default workspace in Postman. ", err.Error())
		return ""
	}
	for _, w := range workspaces.Workspaces {
		if w.Name == name {
			return WorkspaceID(w.ID)
		}
	}
	fmt.Println("No workspace found. Using default workspace in Postman.")
	return ""
}

func CreateEnv(envData EnvironmentData, workspaceName string) (*http.Response, error) {
	// Create the request payload
	payload := CreateEnvironmentRequest{
		Environment: envData,
	}

	// Convert the payload to JSON
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal JSON payload: %w", err)
	}

	// Create the request body
	body := bytes.NewBuffer(jsonPayload)

	// Make the POST request
	endpoint := "/environments"
	if workspaceID := Workspace(workspaceName); workspaceID != "" {
		endpoint += "?workspace=" + string(workspaceID)
	}
	resp, err := makeRequest("POST", endpoint, body)
	if err != nil {
		return nil, fmt.Errorf("failed to make POST request: %w", err)
	}

	// Check for the specific status code
	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		resp.Body.Close()
		return nil, fmt.Errorf("unexpected status code: %d. Expected: %d. Response: %s",
			resp.StatusCode, http.StatusOK, string(bodyBytes))
	}

	return resp, nil
}
