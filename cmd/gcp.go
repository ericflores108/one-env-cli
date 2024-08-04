/*
Copyright © 2024 Eric Flores <eflorty108@gmail.com>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// gcpCmd represents the gcp command
// https://cloud.google.com/secret-manager/docs/create-secret-quickstart#secretmanager-quickstart-go
var gcpCmd = &cobra.Command{
	Use:   "gcp",
	Short: "Create a Google Cloud secret from a password manager item",
	Long: `Create a secret in Google Cloud Secret Manager using an item from your password manager. 
This command requires Application Default Credentials to be set up for Google Cloud authentication.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("gcp called")
	},
}

func init() {
	gcpCmd.Flags().StringP("item", "i", "", "password manager item name")
	gcpCmd.MarkFlagRequired("item")
	gcpCmd.Flags().StringP("project", "p", "", "gcp project")
	gcpCmd.MarkFlagRequired("item")
	rootCmd.AddCommand(gcpCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// gcpCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// gcpCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
