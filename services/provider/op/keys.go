package op

import (
	"errors"

	"github.com/spf13/viper"
)

type Key struct {
	Vault      string
	KeyName    string
	SecretName string
}

func GetKeyConfig(plugin string) (Key, error) {
	var key Key
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

	return Key{Vault: vault, KeyName: viperKeyName, SecretName: viperSecretName}, nil
}
