package postman

import (
	"bytes"
	"fmt"
	"os/exec"

	"github.com/ericflores108/one-env-cli/utils"
)

func GetPostmanAPISecret() (string, error) {
	cmd := exec.Command("op", "read", fmt.Sprintf("op://%s/%s/%s", utils.C.Provider.OP.Vault, utils.C.Plugin.Postman.KeyName, utils.C.Plugin.Postman.KeySecretName), "--no-newline")
	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		return "", err
	}

	return out.String(), nil
}
