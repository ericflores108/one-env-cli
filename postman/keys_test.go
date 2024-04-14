package postman

import (
	"errors"
	"testing"

	"github.com/spf13/viper"
)

func TestGetPostmanKeyConfig(t *testing.T) {
	tests := []struct {
		name          string
		vault         string
		keyName       string
		keySecret     string
		expectedKey   Key
		expectedError error
	}{
		{
			name:        "Valid config",
			vault:       "my-vault",
			keyName:     "my-key",
			keySecret:   "my-secret",
			expectedKey: Key{Vault: "my-vault", Name: "my-key", Secret: "my-secret"},
		},
		{
			name:          "Empty vault",
			expectedError: errors.New("op.vault is empty"),
		},
		{
			name:          "Empty keyName",
			vault:         "my-vault",
			expectedError: errors.New("keys.postmanKeyName is empty"),
		},
		{
			name:          "Empty keySecret",
			vault:         "my-vault",
			keyName:       "my-key",
			expectedError: errors.New("keys.postmanSecretName is empty"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			viper.Reset()
			viper.Set("op.vault", test.vault)
			viper.Set("keys.postmanKeyName", test.keyName)
			viper.Set("keys.postmanSecretName", test.keySecret)

			key, err := getPostmanKeyConfig()

			if test.expectedError != nil {
				if err == nil {
					t.Errorf("Expected error: %v, but got nil", test.expectedError)
				} else if err.Error() != test.expectedError.Error() {
					t.Errorf("Expected error: %v, but got: %v", test.expectedError, err)
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
				if key != test.expectedKey {
					t.Errorf("Expected key: %+v, but got: %+v", test.expectedKey, key)
				}
			}
		})
	}
}

func TestGetPostmanAPISecret(t *testing.T) {
	// Set up test configuration
	viper.Set("op.vault", "personal")
	viper.Set("keys.postmanKeyName", "POSTMAN_API_KEY")
	viper.Set("keys.postmanSecretName", "secret")

	// Test successful retrieval of secret
	_, err := GetPostmanAPISecret()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Test error handling
	viper.Set("keys.postmanKeyName", "")
	_, err = GetPostmanAPISecret()
	if err == nil {
		t.Error("Expected error for empty keyName, but got nil")
	}
}
