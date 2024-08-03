package cmd

import (
	"fmt"

	"github.com/ericflores108/one-env-cli/utils"
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage the application configuration",
	Long: `The config command allows you to manage various aspects of the application configuration.
Currently, it supports setting the active provider (1Password or BitWarden).`,
}

var setProviderCmd = &cobra.Command{
	Use:   "set-provider [provider]",
	Short: "Set the active provider",
	Long: `Set the active provider for the application. 
The provider can be either "op" for 1Password or "bw" for BitWarden.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		provider := args[0]
		err := utils.EditProvider(provider)
		if err != nil {
			fmt.Printf("Error setting provider: %v\n", err)
			return
		}
		fmt.Printf("Successfully set provider to %s\n", provider)
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
	configCmd.AddCommand(setProviderCmd)
}
