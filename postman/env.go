package postman

import (
	"encoding/json"
	"fmt"
	"io"
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
