package postman

import (
	"bytes"
	"fmt"
	"os/exec"

	"github.com/ericflores108/one-env-cli/services/provider/op"
)

func GetPostmanAPISecret() (string, error) {
	key, err := op.GetKeyConfig("postman")
	if err != nil {
		return "", err
	}

	cmd := exec.Command("op", "read", fmt.Sprintf("op://%s/%s/%s", key.Vault, key.PluginKeyName, key.PluginKeySecret), "--no-newline")
	var out bytes.Buffer
	cmd.Stdout = &out

	err = cmd.Run()
	if err != nil {
		return "", err
	}

	return out.String(), nil
}
