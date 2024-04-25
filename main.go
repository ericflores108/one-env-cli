/*
Copyright © 2024 Eric Flores <eflorty108@gmail.com>
*/
package main

import (
	"fmt"

	"github.com/ericflores108/one-env-cli/cmd"
	"github.com/spf13/viper"
)

func main() {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("error reading in config: ", err)
	}
	cmd.Execute()
}
