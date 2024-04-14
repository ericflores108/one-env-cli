/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)

// postmanCmd represents the postman command
var postmanCmd = &cobra.Command{
	Use:   "postman",
	Short: "add an environment to postman",
	Long:  `1Password secrets will be used to create a Postman environment.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		envName, err := cmd.Flags().GetString("env")
		if err != nil {
			fmt.Printf("error retrieving environment: %s\n", err.Error())
			return err
		}
		if envName == "" {
			return errors.New("missing environment to upload")
		}
		fmt.Println("uploading environment to postman, ", envName)
		return nil
	},
}

func init() {
	postmanCmd.Flags().StringP("env", "e", "", "environment name")
	addCmd.AddCommand(postmanCmd)
}
