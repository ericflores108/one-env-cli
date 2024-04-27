package cmd

import (
	"fmt"

	"github.com/ericflores108/one-env-cli/op"
	"github.com/ericflores108/one-env-cli/postman"
	"github.com/spf13/cobra"
)

var postmanCmd = &cobra.Command{
	Use:   "postman",
	Short: "add a 1Password item to create postman environment",
	Long:  `1Password secrets will be used to create a Postman environment.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		itemName, err := cmd.Flags().GetString("item")
		if err != nil {
			fmt.Printf("error retrieving item: %s\n", err.Error())
			return err
		}
		if itemName == "" {
			return fmt.Errorf("missing item to upload")
		}

		// Get the item from 1Password
		item, err := op.GetItem(itemName)
		if err != nil {
			fmt.Printf("error retrieving item from 1Password: %s\n", err.Error())
			return err
		}

		// Transform the item to environment data
		envData := postman.TransformItemToEnv(item)

		// Create the environment in Postman
		resp, err := postman.CreateEnv(envData)
		if err != nil {
			fmt.Printf("error creating environment in Postman: %s\n", err.Error())
			return err
		}
		defer resp.Body.Close()

		fmt.Printf("Environment '%s' created successfully in Postman\n", envData.Name)
		return nil
	},
}

func init() {
	postmanCmd.Flags().StringP("item", "i", "", "item name")
	addCmd.AddCommand(postmanCmd)
}
