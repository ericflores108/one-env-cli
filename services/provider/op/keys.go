package op

import (
	"errors"
	"fmt"

	"github.com/spf13/viper"
)

// Key represents the configuration for a key in a vault
type Key struct {
	Vault           string
	PluginKeyName   string
	PluginKeySecret string
}

var (
	ErrEmptyPlugin      = errors.New("plugin is not supported")
	ErrMissingVault     = errors.New("op.vault does not exist")
	ErrMissingKeyName   = errors.New("plugin key name does not exist")
	ErrMissingKeySecret = errors.New("plugin key secret does not exist")
)

// GetKeyConfig retrieves the key configuration for a given plugin
// It returns a pointer to a Key struct and an error if any
func GetKeyConfig(plugin string) (*Key, error) {
	if plugin == "" {
		return nil, ErrEmptyPlugin
	}

	vault := viper.GetString("op.vault")
	if vault == "" {
		return nil, ErrMissingVault
	}

	configPrefix := fmt.Sprintf("plugin.%s", plugin)
	pluginKeyName := viper.GetString(fmt.Sprintf("%s.keyName", configPrefix))
	pluginKeySecret := viper.GetString(fmt.Sprintf("%s.keySecretName", configPrefix))

	if pluginKeyName == "" {
		return nil, ErrMissingKeyName
	}
	if pluginKeySecret == "" {
		return nil, ErrMissingKeySecret
	}

	return &Key{
		Vault:           vault,
		PluginKeyName:   pluginKeyName,
		PluginKeySecret: pluginKeySecret,
	}, nil
}
