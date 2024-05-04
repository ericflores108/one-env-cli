package postman

import (
	"testing"

	"github.com/ericflores108/one-env-cli/services/provider/op"
)

func TestTransformItemToEnv(t *testing.T) {
	// Test case 1: Item with fields of different types
	item := op.ItemResponse{
		ID:    "1",
		Title: "Test Item",
		Fields: []op.Field{
			{
				ID:    "username",
				Type:  "STRING",
				Value: "testuser",
			},
			{
				ID:    "password",
				Type:  "CONCEALED",
				Value: "testpassword",
			},
			{
				ID:    "notes",
				Type:  "STRING",
				Value: "Test notes",
			},
		},
	}

	expectedEnvData := EnvironmentData{
		Name: "Test Item",
		Values: []EnvironmentVariable{
			{
				Key:     "username",
				Value:   "testuser",
				Enabled: true,
				Type:    DefaultType,
			},
			{
				Key:     "password",
				Value:   "testpassword",
				Enabled: true,
				Type:    SecretType,
			},
			{
				Key:     "notes",
				Value:   "Test notes",
				Enabled: true,
				Type:    DefaultType,
			},
		},
	}

	envData := TransformItemToEnv(item)

	if envData.Name != expectedEnvData.Name {
		t.Errorf("Expected environment name: %s, got: %s", expectedEnvData.Name, envData.Name)
	}

	if len(envData.Values) != len(expectedEnvData.Values) {
		t.Errorf("Expected number of environment variables: %d, got: %d", len(expectedEnvData.Values), len(envData.Values))
	}

	for i, expectedVar := range expectedEnvData.Values {
		if envData.Values[i] != expectedVar {
			t.Errorf("Expected environment variable: %+v, got: %+v", expectedVar, envData.Values[i])
		}
	}

	// Test case 2: Item with no fields
	item = op.ItemResponse{
		ID:     "2",
		Title:  "Empty Item",
		Fields: []op.Field{},
	}

	expectedEnvData = EnvironmentData{
		Name:   "Empty Item",
		Values: []EnvironmentVariable{},
	}

	envData = TransformItemToEnv(item)

	if envData.Name != expectedEnvData.Name {
		t.Errorf("Expected environment name: %s, got: %s", expectedEnvData.Name, envData.Name)
	}

	if len(envData.Values) != len(expectedEnvData.Values) {
		t.Errorf("Expected number of environment variables: %d, got: %d", len(expectedEnvData.Values), len(envData.Values))
	}
}
