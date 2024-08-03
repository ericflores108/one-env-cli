package bwmanager

import (
	"bytes"
	"encoding/json"
	"os/exec"

	"github.com/ericflores108/one-env-cli/providermanager"
)

type BWManager struct {
	ItemName   string
	Item       ItemResponse
	PluginKey  string
	PluginType string
}

func (bwm *BWManager) GetItem() error {
	cmd := exec.Command("bw", "get", "item", bwm.ItemName)

	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		return err
	}

	err = json.Unmarshal(out.Bytes(), &bwm.Item)
	if err != nil {
		return err
	}

	return nil
}

func (bwm *BWManager) GetSecret() (string, error) {
	cmd := exec.Command("bw", "get", "password", bwm.PluginKey)
	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		return "", err
	}

	return out.String(), nil
}

func (bwm *BWManager) PostmanEnv() (*providermanager.PostmanEnvironmentData, error) {
	var envVars []providermanager.PostmanEnvironmentVariable

	for _, field := range bwm.Item.Fields {
		// default and secret are terms to describe data in postman
		envType := providermanager.PostmanDefault
		if field.Type == 1 {
			envType = providermanager.PostmanSecret
		}

		envVar := providermanager.PostmanEnvironmentVariable{
			Key:     field.Name,
			Value:   field.Value,
			Enabled: true,
			Type:    envType,
		}
		envVars = append(envVars, envVar)
	}

	envData := &providermanager.PostmanEnvironmentData{
		Name:   bwm.Item.Name,
		Values: envVars,
	}

	return envData, nil
}

func New(itemName, pluginKey, pluginType string) *BWManager {
	return &BWManager{
		ItemName:   itemName,
		PluginKey:  pluginKey,
		PluginType: pluginType,
		Item:       ItemResponse{},
	}
}
