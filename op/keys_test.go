package op

import (
	"testing"

	"github.com/spf13/viper"
)

func TestGetKeyConfig(t *testing.T) {
	viper.AddConfigPath("..")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		t.Errorf("error: %v", err)
	}
	key, err := GetKeyConfig("postman")
	if err != nil {
		t.Errorf("error: %v", err)
	}
	t.Log(key)
}
