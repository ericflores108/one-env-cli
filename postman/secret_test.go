package postman

import (
	"testing"

	"github.com/spf13/viper"
)

func TestGetPostmanAPISecret(t *testing.T) {
	// Set up test configuration
	viper.Set("op.vault", "personal")
	viper.Set("keys.postman.KeyName", "POSTMAN_API_KEY")
	viper.Set("keys.postman.SecretName", "secret")

	// Test successful retrieval of secret
	_, err := GetPostmanAPISecret()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Test error handling
	viper.Set("keys.postman.SecretName", "")
	_, err = GetPostmanAPISecret()
	if err == nil {
		t.Error("Expected error for keys.postman.SecretName, but got nil")
	}
}
