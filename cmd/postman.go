package cmd

import (
	"fmt"
	"io"

	"github.com/ericflores108/one-env-cli/services/plugin/postman"
	"github.com/ericflores108/one-env-cli/utils"
	"github.com/spf13/cobra"
)

var postmanCmd = &cobra.Command{
	Use:   "postman",
	Short: "add an item to create postman environment",
	Long:  `Your password manager secrets will be used to create a Postman environment for the item identified.`,
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

		// register password manager
		provider, err := utils.GetEnabledProvider(itemName, utils.C.Plugin.Postman.KeyName, utils.C.Plugin.Postman.Type)
		if err != nil {
			fmt.Printf("error enabling provider: %s\n", err.Error())
			return err
		}

		// Get item
		err = provider.GetItem()
		if err != nil {
			fmt.Printf("error retrieving item from provider: %s\n", err.Error())
			return err
		}

		// Get Postman Env
		envData, err := provider.PostmanEnv()
		if err != nil {
			fmt.Printf("error retrieving item from provider: %s\n", err.Error())
			return err
		}

		// Check optional flag for workspace
		workspace, err := cmd.Flags().GetString("workspace")
		if err != nil {
			fmt.Printf("No workspace set. Using default workspace.")
		}

		job := postman.NewPluginJob(provider)

		// Create the environment in Postman
		resp, err := job.CreateEnv(*envData, workspace)
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
	postmanCmd.Flags().StringP("item", "i", "", "provider item name")
	postmanCmd.MarkFlagRequired("item")
	postmanCmd.Flags().StringP("workspace", "w", "", "postman workspace")
	addCmd.AddCommand(postmanCmd)
}
