package config

import (
	"fmt"

	"github.com/spf13/viper"
)

const (
	envPrefix          = "opengo"
	envBindOpenaiToken = "openai_token"
	viperConfigName    = "chatgpt"
	viperConfigType    = "yaml"
	viperConfigPath    = "."
)

type configProvider struct{}

func NewConfigProvider() *configProvider {
	viper.SetEnvPrefix(envPrefix)
	viper.BindEnv(envBindOpenaiToken)
	viper.SetConfigName(viperConfigName)
	viper.SetConfigType(viperConfigType)
	viper.AddConfigPath(viperConfigPath)
	// Don't fail if the yaml file is not found. The key could be in the env var
	viper.ReadInConfig()

	return &configProvider{}
}

func (p *configProvider) ReadString(key string) string {
	value, ok := viper.Get(key).(string)
	if !ok {
		panic(fmt.Sprintf("Error while reader config key `%s`.\n", key))
	}

	return value
}
