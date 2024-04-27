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

func CreateEnv(name string, variables []EnvironmentVariable) (*http.Response, error) {
	// Create the environment data
	envData := EnvironmentData{
		Name:   name,
		Values: variables,
	}

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
	resp, err := makeRequest("POST", "/environments", body)
	if err != nil {
		return nil, fmt.Errorf("failed to make POST request: %v", err)
	}

	return resp, nil
}
