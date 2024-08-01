package op

import (
	"testing"

	"github.com/spf13/viper"
)

func TestGetPluginConfig(t *testing.T) {
	viper.AddConfigPath("../../../configs")
	viper.SetConfigName(".example-one-env-cli")
	viper.SetConfigType("json")
	err := viper.ReadInConfig()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	key, err := GetPluginConfig("postman")
	if err != nil {
		t.Errorf("error: %v", err)
	}
	t.Log(key)
}
