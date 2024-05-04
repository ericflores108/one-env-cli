/*
Copyright © 2024 Eric Flores <eflorty108@gmail.com>
*/
package cmd

import (
	"bytes"
	"log"
	"os"
	"path/filepath"

	"github.com/ericflores108/one-env-cli/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func RootCMD() *cobra.Command {
	return rootCmd
}

// rootCmd represents the base command when called without any subcommands
var (
	rootCmd = &cobra.Command{
		Use:   "one-env-cli",
		Short: "create environments with 1Password",
		Long: `one-env-cli is a command-line tool that streamlines environment creation
		 using your password manager as the provider.
		 
		 It provides a convenient way to manage and create environments, such as Postman environments, quickly and securely.
		`,
	}
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "verbose")
}

func Configure() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("Failed to get user home directory: %v", err)
	}

	configDir := filepath.Join(homeDir, ".one-env-cli")
	configFile := filepath.Join(configDir, ".one-env-cli")

	// Create the configuration file if it doesn't exist
	err = utils.CreateFilesIfNotExists([]string{configFile})
	if err != nil {
		log.Fatalf("Failed to create configuration file: %v", err)
	}

	viper.SetConfigType("json")
	viper.SetConfigFile(configFile)

	err = viper.ReadInConfig()
	if err != nil {
		// If the configuration file is empty, write the default configuration
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			defaultConfig := `{
                "plugin": {
                    "postman": {
                        "keyName": "Postman",
                        "keySecretName": "api-key"
                    }
                },
                "op": {
                    "vault": "Developer"
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
            }`

			err = viper.ReadConfig(bytes.NewBufferString(defaultConfig))
			if err != nil {
				log.Fatalf("Failed to read default configuration: %v", err)
			}

			err = viper.WriteConfig()
			if err != nil {
				log.Fatalf("Failed to write default configuration: %v", err)
			}
		} else {
			log.Fatalf("Failed to read configuration file: %v", err)
		}
	}
}
