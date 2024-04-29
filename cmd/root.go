/*
Copyright © 2024 Eric Flores <eflorty108@gmail.com>
*/
package cmd

import (
	"os"

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
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.ReadInConfig()
}
