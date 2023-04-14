package initialize

import (
	"github.com/spf13/viper"
	"log"
	"os"
)

func InitViper() *viper.Viper {
	path, _ := os.Getwd()
	v := viper.New()
	v.AddConfigPath(path)
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	if err := v.ReadInConfig(); err != nil {
		log.Fatal(err)
	}
	return v
}
