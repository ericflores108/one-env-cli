/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add [postman] [-e|--env]",
	Short: "add an environment to postman",
	Long:  `This command allows you to add an environment to postman from 1Password.`,
}

var (
	EnvName = ""
)

func init() {
	addCmd.PersistentFlags().StringVarP(&EnvName, "env", "e", "", "environment to add")
	addCmd.MarkPersistentFlagRequired("env")
	rootCmd.AddCommand(addCmd)
}
