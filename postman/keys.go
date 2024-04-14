package postman

import (
	"bytes"
	"errors"
	"fmt"
	"os/exec"

	"github.com/spf13/viper"
)

type Key struct {
	Vault  string
	Name   string
	Secret string
}

func getPostmanKeyConfig() (Key, error) {
	vault := viper.GetString("op.vault")
	if vault == "" {
		return Key{}, errors.New("op.vault is empty")
	}

	keyName := viper.GetString("keys.postmanKeyName")
	if keyName == "" {
		return Key{}, errors.New("keys.postmanKeyName is empty")
	}

	keySecret := viper.GetString("keys.postmanSecretName")
	if keySecret == "" {
		return Key{}, errors.New("keys.postmanSecretName is empty")
	}

	key := Key{Vault: vault, Name: keyName, Secret: keySecret}
	return key, nil
}

func GetPostmanAPISecret() (string, error) {
	key, err := getPostmanKeyConfig()
	if err != nil {
		return "", err
	}

	cmd := exec.Command("op", "read", fmt.Sprintf("op://%s/%s/%s", key.Vault, key.Name, key.Secret), "--no-newline")
	var out bytes.Buffer
	cmd.Stdout = &out

	err = cmd.Run()
	if err != nil {
		return "", err
	}

	return out.String(), nil
}
