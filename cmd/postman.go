package cmd

import (
	"fmt"
	"io"

	"github.com/ericflores108/one-env-cli/services/plugin/postman"
	"github.com/ericflores108/one-env-cli/services/provider/op"
	"github.com/ericflores108/one-env-cli/utils"
	"github.com/spf13/cobra"
)

var postmanCmd = &cobra.Command{
	Use:   "postman",
	Short: "add a 1Password item to create postman environment",
	Long:  `1Password secrets will be used to create a Postman environment.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		verbose, _ := cmd.Flags().GetBool("verbose")
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

		// Transform the item to Postman environment data for request
		envData := *postman.TransformItemToEnv(item)

		// Check optional flag for workspace
		workspace, err := cmd.Flags().GetString("workspace")
		if err != nil {
			fmt.Printf("No workspace set. Using default workspace.")
		}

		// Create the environment in Postman
		resp, err := postman.CreateEnv(envData, workspace)
		if err != nil {
			fmt.Printf("error creating environment in Postman: %s\n", err.Error())
			return err
		}
		defer resp.Body.Close()
		b, err := io.ReadAll(resp.Body)
		if err != nil {
			return utils.Error("\n  reading response: %v\n  ", err, verbose)
		}
		utils.LogHTTPResponse(verbose, resp, b)
		fmt.Printf("Environment '%s' created successfully in Postman\n", envData.Name)
		return nil
	},
}

func init() {
	postmanCmd.Flags().StringP("item", "i", "", "op item name")
	postmanCmd.MarkFlagRequired("item")
	postmanCmd.Flags().StringP("workspace", "w", "", "postman workspace")
	addCmd.AddCommand(postmanCmd)
}
