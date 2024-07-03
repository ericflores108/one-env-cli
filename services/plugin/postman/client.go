package postman

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

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

func initializeAPIKey() error {
	var err error
	initAPIKeyOnce.Do(func() {
		APIKey, err = GetPostmanAPISecret()
		if err != nil {
			err = fmt.Errorf("failed to get Postman API secret: %v", err)
		}
	})
	return err
}

func makeRequest(method, endpoint string, body io.Reader) (*http.Response, error) {
	err := initializeAPIKey()
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
