package postman

import (
	"fmt"
	"io"
	"net/http"
)

var BaseURL string = "https://api.getpostman.com"
var APIKey string

func init() {
	err := initializeAPIKey()
	if err != nil {
		fmt.Println("Failed to initialize API secret:", err)
	}
}

func initializeAPIKey() error {
	var err error
	APIKey, err = GetPostmanAPISecret()
	if err != nil {
		return fmt.Errorf("failed to get Postman API secret: %v", err)
	}
	return nil
}

func makeRequest(method, endpoint string, body io.Reader) (*http.Response, error) {
	url := BaseURL + endpoint

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("X-Api-Key", APIKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
