package postman

import (
	"testing"

	"github.com/spf13/viper"
)

func TestGetAllEnv(t *testing.T) {
	var resp EnvironmentsResponse
	var err error
	// Set up test configuration
	viper.AddConfigPath("..")
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
