package utils

import (
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
