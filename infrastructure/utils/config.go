package util

import (
	"github.com/spf13/viper"
)

type Config struct {
	SqlServerHost     string `mapstructure:"SQL_SERVER_HOST"`
	SqlServerPort     string `mapstructure:"SQL_SERVER_PORT"`
	SqlServerUser     string `mapstructure:"SQL_SERVER_USER"`
	SqlServerPassword string `mapstructure:"SQL_SERVER_PASSWORD"`
	Database          string `mapstructure:"SQL_SERVER_DATABASE"`
	ServerAddress     string `mapstructure:"SERVER_ADDRESS"`
	TokenKey          string `mapstructure:"API_TOKEN"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
