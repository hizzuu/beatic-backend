package conf

import (
	"os"
	"strings"

	"github.com/davecgh/go-spew/spew"
	"github.com/spf13/viper"
)

type config struct {
	DB          *db
	Credentials *credentials
}

type db struct {
	Dbms                 string
	Name                 string
	User                 string
	Pass                 string
	Net                  string
	Host                 string
	Port                 string
	Parsetime            bool
	AllowNativePasswords bool
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

	spew.Dump(C)
}
