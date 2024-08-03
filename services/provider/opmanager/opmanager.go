package opmanager

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os/exec"

	"github.com/ericflores108/one-env-cli/providermanager"
)

type OPManager struct {
	ItemName   string
	Item       ItemResponse
	Vault      string
	PluginKey  string
	PluginType string
}

// Provider - Get item to send to Plugin
func (opm *OPManager) GetItem() error {
	cmd := exec.Command("op", "item", "get", opm.ItemName, "--format", "json")
	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		return err
	}

	err = json.Unmarshal(out.Bytes(), &opm.Item)
	if err != nil {
		return err
	}

	return nil
}

func (opm *OPManager) GetSecret() (string, error) {
	cmd := exec.Command("op", "read", fmt.Sprintf("op://%s/%s/%s", opm.Vault, opm.PluginKey, opm.PluginType), "--no-newline")
	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		return "", err
	}

	return out.String(), nil
}

func (opm *OPManager) PostmanEnv() (*providermanager.PostmanEnvironmentData, error) {
	var envVars []providermanager.PostmanEnvironmentVariable

	for _, field := range opm.Item.Fields {
		// default and secret are terms to describe data in postman
		envType := providermanager.PostmanDefault
		if field.Type == "CONCEALED" {
			envType = providermanager.PostmanSecret
		}

		envVar := providermanager.PostmanEnvironmentVariable{
			Key:     field.Label,
			Value:   field.Value,
			Enabled: true,
			Type:    envType,
		}
		envVars = append(envVars, envVar)
	}

	envData := &providermanager.PostmanEnvironmentData{
		Name:   opm.Item.Title,
		Values: envVars,
	}

	return envData, nil
}

func New(itemName, vault, pluginKey, pluginType string) *OPManager {
	return &OPManager{
		ItemName:   itemName,
		Vault:      vault,
		PluginKey:  pluginKey,
		PluginType: pluginType,
		Item:       ItemResponse{},
	}
}
