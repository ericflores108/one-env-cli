package utils

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

var (
	C          Config
	dir        = ".one-env-cli"
	file       = ".one-env-cli"
	configFile string
)

type Config struct {
	Plugin struct {
		Postman struct {
			KeyName string `json:"keyName"`
			Type    string `json:"type"`
		} `json:"postman"`
		GCP struct {
			Type string `json:"type"`
		} `json:"gcp"`
	} `json:"plugin"`
	Provider struct {
		OP struct {
			Vault   string `json:"vault"`
			Enabled bool   `json:"enabled"`
		} `json:"op"`
		BW struct {
			Enabled bool `json:"enabled"`
		} `json:"bw"`
	} `json:"provider"`
	CLI struct {
		Logging struct {
			Level       string   `json:"level"`
			Encoding    string   `json:"encoding"`
			OutputPaths []string `json:"outputPaths"`
		} `json:"logging"`
	} `json:"cli"`
}

func init() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("Failed to get user home directory: %v", err)
	}

	configDir := filepath.Join(homeDir, dir)
	configFile = filepath.Join(configDir, file)

	viper.SetConfigType("json")
	viper.SetConfigFile(configFile)
}

func InitConfig() error {
	return viper.Unmarshal(&C)
}

func ReadConfigPath() error {
	return viper.ReadInConfig()
}

func CreateAndWriteConfigFile() error {
	err := CreateFilesIfNotExists([]string{configFile})
	if err != nil {
		return err
	}
	return nil
}

func DefaultConfig() string {
	return `
{
  "plugin": {
    "postman": {
      "keyName": "PostmanAPI",
      "type": "api-key"
    },
    "gcp": {
      "type": "cli"
    }
  },
  "provider": {
    "op": {
      "vault": "Developer",
      "enabled": true
    },
    "bw": {
      "enabled": false
    }
  },
  "cli": {
    "logging": {
      "level": "debug",
      "encoding": "json",
      "outputPaths": [
        "tmp/log/one-env-cli.json"
      ]
    }
  }
}
`
}

func EditProvider(providerName string) error {
	if providerName != "op" && providerName != "bw" {
		return fmt.Errorf("invalid provider name. Must be 'op' or 'bw'")
	}

	// Read the current configuration
	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("failed to read config: %w", err)
	}

	// Update the provider settings
	if providerName == "op" {
		viper.Set("provider.op.enabled", true)
		viper.Set("provider.bw.enabled", false)
	} else {
		viper.Set("provider.op.enabled", false)
		viper.Set("provider.bw.enabled", true)
	}

	// Save the updated configuration
	if err := viper.WriteConfig(); err != nil {
		return fmt.Errorf("failed to write config: %w", err)
	}

	// Update the global Config struct
	if err := viper.Unmarshal(&C); err != nil {
		return fmt.Errorf("failed to unmarshal config: %w", err)
	}

	return nil
}
