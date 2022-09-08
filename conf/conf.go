package conf

import (
	"os"
	"strings"

	"github.com/spf13/viper"
)

type config struct {
	App         *app
	Api         *api
	DB          *db
	Credentials *credentials
}

type app struct {
	Environment string
	Debug       bool
}

type api struct {
	Port string
	User string
	Pass string
}

type db struct {
	DSN string
}

type credentials struct {
	Firebase *firebase
}

type firebase struct {
	SecretKey string `mapstructure:"secret_key"`
}

var C config

func init() {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	viper.SetConfigName("conf")
	viper.SetConfigType("yml")
	viper.AddConfigPath(dir + "/conf")

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := viper.Unmarshal(&C); err != nil {
		panic(err)
	}
}
