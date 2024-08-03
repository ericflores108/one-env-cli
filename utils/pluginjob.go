package utils

import (
	"errors"

	"github.com/ericflores108/one-env-cli/providermanager"
	"github.com/ericflores108/one-env-cli/services/provider/bwmanager"
	"github.com/ericflores108/one-env-cli/services/provider/opmanager"
)

// GetEnabledProvider checks the configuration and returns the enabled ProviderManager
func GetEnabledProvider(itemName, pluginKeyName, pluginType string) (providermanager.ProviderManager, error) {
	if C.Provider.OP.Enabled {
		return opmanager.New(itemName, C.Provider.OP.Vault, pluginKeyName, pluginType), nil
	}

	if C.Provider.BW.Enabled {
		return bwmanager.New(itemName, pluginKeyName, pluginType), nil
	}

	return nil, errors.New("no enabled provider found in configuration")
}
