package postman

import (
	"testing"

	"github.com/spf13/viper"
)

func TestGetPostmanAPISecret(t *testing.T) {
	var _ any
	var err error
	// Set up test configuration
	viper.AddConfigPath("..")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	err = viper.ReadInConfig()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	// Test successful retrieval of secret
	_, err = GetPostmanAPISecret()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Test error handling
	viper.Set("entity.postman.keySecretName", "")
	_, err = GetPostmanAPISecret()
	if err == nil {
		t.Error("Expected error for entity.postman.keySecretName, but got nil")
	}
}
