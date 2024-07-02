package postman

import (
	"github.com/ericflores108/one-env-cli/services/provider/op"
)

func TransformItemToEnv(item *op.ItemResponse) *EnvironmentData {
	var envVars []EnvironmentVariable

	for _, field := range item.Fields {
		envType := DefaultType
		if field.Type == "CONCEALED" {
			envType = SecretType
		}

		envVar := EnvironmentVariable{
			Key:     field.Label,
			Value:   field.Value,
			Enabled: true,
			Type:    envType,
		}
		envVars = append(envVars, envVar)
	}

	envData := &EnvironmentData{
		Name:   item.Title,
		Values: envVars,
	}

	return envData
}
