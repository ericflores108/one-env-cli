package postman

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"

	"github.com/spf13/viper"
)

func TestInitializeAPIKey(t *testing.T) {
	var err error
	// Set up test configuration
	viper.AddConfigPath("../../../configs")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	err = viper.ReadInConfig()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Reset the initAPIKeyOnce variable before each test
	initAPIKeyOnce = sync.Once{}

	// Call the initializeAPIKey function multiple times concurrently
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			err := initializeAPIKey()
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
		}()
	}
	wg.Wait()

	// Assert that the API key is initialized
	if APIKey == "" {
		t.Error("API key is not initialized")
	}
}

func TestMakeRequest(t *testing.T) {
	var err error
	// Set up test configuration
	viper.AddConfigPath("../../../configs")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	err = viper.ReadInConfig()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	// Create a mock HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Assert the request method and URL
		expectedMethod := "GET"
		expectedURL := "/test"
		if r.Method != expectedMethod {
			t.Errorf("Expected request method %s, but got %s", expectedMethod, r.Method)
		}
		if r.URL.Path != expectedURL {
			t.Errorf("Expected request URL %s, but got %s", expectedURL, r.URL.Path)
		}

		// Send a response
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "Response body")
	}))
	defer server.Close()

	// Set the base URL to the mock server URL
	originalBaseURL := BaseURL
	BaseURL = server.URL

	// Set a test API key
	originalAPIKey := APIKey
	APIKey = "test_api_key"

	// Make a request
	resp, err := makeRequest("GET", "/test", nil)

	// Assert that no error occurred
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Assert the response status code
	expectedStatusCode := http.StatusOK
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
	expectedBody := "Response body"
	if string(body) != expectedBody {
		t.Errorf("Expected response body %s, but got %s", expectedBody, string(body))
	}

	// Restore the original base URL and API key
	BaseURL = originalBaseURL
	APIKey = originalAPIKey
}
