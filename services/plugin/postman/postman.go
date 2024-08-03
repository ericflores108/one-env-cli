package postman

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"

	"github.com/ericflores108/one-env-cli/providermanager"
)

type PluginJob struct {
	Provider providermanager.ProviderManager
}

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

type CreateEnvironmentRequest struct {
	Environment providermanager.PostmanEnvironmentData `json:"environment"`
}

var (
	BaseURL        = "https://api.getpostman.com"
	APIKey         string
	initAPIKeyOnce sync.Once
	httpClient     *http.Client
)

func init() {
	httpClient = &http.Client{
		Timeout: 30 * time.Second,
	}
}

func (p *PluginJob) initializeAPIKey() error {
	var err error
	initAPIKeyOnce.Do(func() {
		APIKey, err = p.Provider.GetSecret()
		if err != nil {
			err = fmt.Errorf("failed to get Postman API secret: %v", err)
		}
	})
	return err
}

func (p *PluginJob) makeRequest(method, endpoint string, body io.Reader) (*http.Response, error) {
	err := p.initializeAPIKey()
	if err != nil {
		return nil, fmt.Errorf("failed to initialize API key: %v", err)
	}

	url := BaseURL + endpoint

	bodyBytes, err := io.ReadAll(body)
	if err != nil {
		return nil, fmt.Errorf("failed to read request body: %v", err)
	}

	bodyReader := bytes.NewReader(bodyBytes)

	req, err := http.NewRequest(method, url, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP request: %v", err)
	}

	req.Header.Set("X-Api-Key", APIKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send HTTP request: %v", err)
	}

	return resp, nil
}

func (p *PluginJob) Workspace(name string) WorkspaceID {
	if name == "" {
		return ""
	}
	resp, err := p.makeRequest("GET", "/workspaces", nil)
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

func (p *PluginJob) CreateEnv(envData providermanager.PostmanEnvironmentData, workspaceName string) (*http.Response, error) {
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
	if workspaceID := p.Workspace(workspaceName); workspaceID != "" {
		endpoint += "?workspace=" + string(workspaceID)
	}
	resp, err := p.makeRequest("POST", endpoint, body)
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

func NewPluginJob(providerManager providermanager.ProviderManager) *PluginJob {
	return &PluginJob{
		Provider: providerManager,
	}
}
