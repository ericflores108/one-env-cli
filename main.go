/*
Copyright © 2024 Eric Flores <eflorty108@gmail.com>
*/
package main

import (
	"github.com/ericflores108/one-env-cli/cmd"
	"github.com/ericflores108/one-env-cli/utils"
)

func main() {
	cmd.Configure()
	utils.InitCLILogger()
	cmd.Execute()
}
