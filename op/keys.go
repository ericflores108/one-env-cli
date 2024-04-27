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

func GetKeyConfig(integratedParty string) (Key, error) {
	var key Key
	if integratedParty == "" {
		return key, errors.New("entity is empty")
	}

	vault := viper.GetString("op.vault")
	if vault == "" {
		return key, errors.New("op.vault is empty")
	}

	viperStringPrefix := "integration." + integratedParty
	viperKeyName := viper.GetString(viperStringPrefix + ".keyName")
	viperSecretName := viper.GetString(viperStringPrefix + ".keySecretName")

	if viperKeyName == "" {
		return key, errors.New("viperKeyName is empty")
	}
	if viperSecretName == "" {
		return key, errors.New("viperKeyName is empty")
	}

	return Key{Vault: vault, KeyName: viperKeyName, SecretName: viperSecretName}, nil
}
