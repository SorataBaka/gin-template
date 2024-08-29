package bootstrap

import (
	"log"
	"github.com/spf13/viper"
)

type EnvObject struct {
	APP_ENV string `mapstructure:"APP_ENV"`
	PORT uint16 `mapstructure:"PORT"`
	GIN_MODE string `mapstructure:"GIN_MODE"`
	DURATION uint8 `mapstructure:"DURATION"`
}

func InitializeEnv() *EnvObject {
	env := EnvObject{}
	viper.SetConfigFile(".env")
	viper.SetDefault("APP_ENV", "production")
	viper.SetDefault("PORT", 3000)
	viper.SetDefault("GIN_MODE", env.GIN_MODE)
	viper.SetDefault("DURATION", 2)

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Could not find .env file")
	}
	err = viper.Unmarshal(&env)
	if err != nil {
		log.Fatal("Could not load environmental variables")
	}

	log.Printf("Running on " + env.APP_ENV + " mode")
	if env.PORT == 3000 {
		log.Printf("Running on default port 3000")
	}
	if env.GIN_MODE == "debug" && env.APP_ENV == "production" {
		log.Printf("Gin is running in debug mode. Recommend setting release mode for production.")
	}

	return &env
}
