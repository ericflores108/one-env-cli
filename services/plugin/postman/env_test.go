package postman

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/spf13/viper"
)

func TestGetAllEnv(t *testing.T) {
	var resp EnvironmentsResponse
	var err error
	// Set up test configuration
	viper.AddConfigPath("../../../configs")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	err = viper.ReadInConfig()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Perform manual initialization
	err = initializeAPIKey()
	if err != nil {
		t.Errorf("Failed to initialize API key: %v", err)
	}

	resp, err = GetAllEnv()
	if err != nil {
		t.Errorf("Failed to get all postman environments: %v", err)
	}
	t.Log(resp)
}

func TestCreateEnv(t *testing.T) {
	// Create a mock HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Assert the request method and URL
		expectedMethod := "POST"
		expectedURL := "/environments"
		if r.Method != expectedMethod {
			t.Errorf("Expected request method %s, but got %s", expectedMethod, r.Method)
		}
		if r.URL.Path != expectedURL {
			t.Errorf("Expected request URL %s, but got %s", expectedURL, r.URL.Path)
		}

		// Read the request body
		body, err := io.ReadAll(r.Body)
		if err != nil {
			t.Errorf("Failed to read request body: %v", err)
		}
		defer r.Body.Close()

		// Unmarshal the request body
		var req CreateEnvironmentRequest
		err = json.Unmarshal(body, &req)
		if err != nil {
			t.Errorf("Failed to unmarshal request body: %v", err)
		}

		// Assert the environment data
		expectedEnvData := EnvironmentData{
			Name: "Test Environment",
			Values: []EnvironmentVariable{
				{Key: "key1", Value: "value1", Enabled: true, Type: "text"},
				{Key: "key2", Value: "value2", Enabled: false, Type: "secret"},
			},
		}
		if !reflect.DeepEqual(req.Environment, expectedEnvData) {
			t.Errorf("Expected environment data %+v, but got %+v", expectedEnvData, req.Environment)
		}

		// Send a response
		w.WriteHeader(http.StatusCreated)
		fmt.Fprint(w, `{"id": "env-id"}`)
	}))
	defer server.Close()

	// Set the base URL to the mock server URL
	originalBaseURL := BaseURL
	BaseURL = server.URL

	// Create environment variables
	variables := []EnvironmentVariable{
		{Key: "key1", Value: "value1", Enabled: true, Type: "text"},
		{Key: "key2", Value: "value2", Enabled: false, Type: "secret"},
	}

	envData := EnvironmentData{
		Name:   "Test Environment",
		Values: variables,
	}

	// Call the CreateEnv function
	resp, err := CreateEnv(envData)

	// Assert that no error occurred
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Assert the response status code
	expectedStatusCode := http.StatusCreated
	if resp.StatusCode != expectedStatusCode {
		t.Errorf("Expected status code %d, but got %d", expectedStatusCode, resp.StatusCode)
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		t.Errorf("Failed to read response body: %v", err)
	}

	// Assert the response body
	expectedBody := `{"id": "env-id"}`
	if string(body) != expectedBody {
		t.Errorf("Expected response body %s, but got %s", expectedBody, string(body))
	}

	// Restore the original base URL
	BaseURL = originalBaseURL
}
