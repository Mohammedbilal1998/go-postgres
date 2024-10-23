package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Port string `mapstructure:"PORT"`
	DBUrl string `mapstructure:"DB_URL"`
}

func GetConfig() (Config, error){
	viper.AddConfigPath("./pkg/common/envs")
	viper.SetConfigName("dev")
	viper.SetConfigType("env")
	err := viper.ReadInConfig()
	if err != nil {
		return Config{}, err
	}
	c := Config{}
	err = viper.Unmarshal(&c) 
	if err != nil {
		return Config{}, err
	}
	return c, nil
}