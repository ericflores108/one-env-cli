package postman

import (
	"fmt"
	"io"
	"net/http"
	"sync"
)

var BaseURL string = "https://api.getpostman.com"
var APIKey string
var initAPIKeyOnce sync.Once

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
		return nil, err
	}

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
