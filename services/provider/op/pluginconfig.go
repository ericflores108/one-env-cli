package op

import (
	"errors"

	"github.com/spf13/viper"
)

type PluginConfig struct {
	Vault      string
	KeyName    string
	SecretName string
}

// Provider - Give credentials to consume Plugin
func GetPluginConfig(plugin string) (PluginConfig, error) {
	var key PluginConfig
	if plugin == "" {
		return key, errors.New("plugin is not supported")
	}

	vault := viper.GetString("provider.op.vault")
	if vault == "" {
		return key, errors.New("provider.op.vault does not exist")
	}

	viperStringPrefix := "plugin." + plugin
	viperKeyName := viper.GetString(viperStringPrefix + ".keyName")
	viperSecretName := viper.GetString(viperStringPrefix + ".keySecretName")

	if viperKeyName == "" {
		return key, errors.New("key name does not exist")
	}
	if viperSecretName == "" {
		return key, errors.New("secret name does not exist")
	}

	return PluginConfig{Vault: vault, KeyName: viperKeyName, SecretName: viperSecretName}, nil
}
