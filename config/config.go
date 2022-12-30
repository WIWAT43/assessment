package config

import (
	"github.com/spf13/viper"
	"log"
	"strings"
)

type Configs struct {
	DbDriver string
	DbSource string
}

func GetConfig() Configs {
	return Configs{
		viper.GetString("postgres.dbDriver"),
		viper.GetString("postgres.database_url"),
	}
}

func InitViper() {
	viper.AddConfigPath("../../")
	viper.SetConfigName("config")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("cannot read in viper config:%s", err)
	}

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
}
