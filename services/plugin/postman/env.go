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
	Workspaces []Workspace `json:"workspaces"`
}

type Workspace struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Type       string `json:"type"`
	Visibility string `json:"visibility"`
	CreatedBy  string `json:"createdBy"`
}

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

func GetAllEnv() (EnvironmentsResponse, error) {
	resp, err := makeRequest("GET", "/environments", nil)
	if err != nil {
		return EnvironmentsResponse{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return EnvironmentsResponse{}, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return EnvironmentsResponse{}, err
	}

	var envResponse EnvironmentsResponse
	err = json.Unmarshal(body, &envResponse)
	if err != nil {
		return EnvironmentsResponse{}, err
	}

	return envResponse, nil
}

func getWorkspace(workspace string) string {
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
		if w.Name == workspace {
			return w.ID
		}
	}
	fmt.Println("No workspace found. Using default workspace in Postman.")
	return ""
}

func CreateEnv(envData EnvironmentData, workspace string) (*http.Response, error) {
	// Create the request payload
	payload := CreateEnvironmentRequest{
		Environment: envData,
	}

	// Convert the payload to JSON
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal JSON payload: %v", err)
	}

	// Create the request body
	body := bytes.NewBuffer(jsonPayload)

	// Make the POST request
	endpoint := "/environments"
	if workspace != "" {
		if w := getWorkspace(workspace); w != "" {
			endpoint += "?workspace=" + w
		}
	}
	resp, err := makeRequest("POST", endpoint, body)
	if err != nil {
		return nil, fmt.Errorf("failed to make POST request: %v", err)
	}

	return resp, nil
}
