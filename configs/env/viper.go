package env

import (
	"log"

	"github.com/spf13/viper"
)

type EnvVars struct {
	PORT                string `mapstructure:"PORT"`
	MONGODBURI          string `mapstructure:"MONGODB_URI"`
	MONGODBNAME         string `mapstructure:"MONGODB_NAME"`
	PUBLICKEY           string `mapstructure:"PUBLIC_KEY"`
	PRIVATE_KEY         string `mapstructure:"PRIVATE_KEY"`
}

func LoadConfig() (config EnvVars, err error) {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatal(err)
	}

	return
}
