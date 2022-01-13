package config

import (
	"github.com/spf13/viper"
)

var Config *viper.Viper

func Init() {

	Config = viper.New()
	Config.AddConfigPath("config")
	Config.SetConfigName("config")
	Config.SetConfigType("json")

	if err := Config.ReadInConfig(); err != nil {
		panic(err)
	}
}

func Get(path string) interface{} {
	return Config.Get(path)
}

func GetInt(path string) int {
	return Config.GetInt(path)
}

func GetString(path string) string {
	return Config.GetString(path)
}

func GetBool(path string) bool {
	return Config.GetBool(path)
}
