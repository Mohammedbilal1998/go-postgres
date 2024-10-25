package config

import(
	"github.com/spf13/viper"
)

type Config struct {
	DBUrl string `mapstructure:"DBURL"`
	Port string `mapstructure:"PORT"`
}

func ReadConfig() Config {
	var config Config
	viper.AddConfigPath("./config")
	viper.SetConfigName("dev")
	viper.SetConfigType("env")

	viper.ReadInConfig()

	viper.Unmarshal(&config)
	return config
}