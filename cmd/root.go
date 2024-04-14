/*
Copyright © 2024 Eric Flores <eflorty108@gmail.com>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "one-env-cli",
	Short: "Environment manager with 1Password",
	Long: `one-env-cli is a command-line tool that streamlines environment creation
	 using 1Password secrets as the source.
	 
	 It provides a convenient way to manage and create environments, such as Postman environments, quickly and securely.
	`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
