package cmd

import (
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add [postman] [-i|--item]",
	Short: "add an item to an integrated application",
	Long:  `This command allows you to add an item to a plugin from your password manager (provider).`,
}

var (
	EnvName = ""
)

func init() {
	addCmd.PersistentFlags().StringVarP(&EnvName, "item", "i", "", "item to add")
	addCmd.MarkPersistentFlagRequired("item")
	rootCmd.AddCommand(addCmd)
}
