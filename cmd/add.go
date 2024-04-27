/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add [postman] [-i|--item]",
	Short: "add a 1Password item to an integrated application",
	Long:  `This command allows you to add an item to postman from 1Password.`,
}

var (
	EnvName = ""
)

func init() {
	addCmd.PersistentFlags().StringVarP(&EnvName, "item", "i", "", "item to add")
	addCmd.MarkPersistentFlagRequired("item")
	rootCmd.AddCommand(addCmd)
}
