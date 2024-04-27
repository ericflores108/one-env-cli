package postman

import (
	op "github.com/ericflores108/one-env-cli/op"
)

func TransformItemToEnv(item op.ItemResponse) EnvironmentData {
	var envVars []EnvironmentVariable

	for _, field := range item.Fields {
		envType := DefaultType
		if field.Type == "CONCEALED" {
			envType = SecretType
		}

		envVar := EnvironmentVariable{
			Key:     field.ID,
			Value:   field.Value,
			Enabled: true,
			Type:    envType,
		}
		envVars = append(envVars, envVar)
	}

	envData := EnvironmentData{
		Name:   item.Title,
		Values: envVars,
	}

	return envData
}
