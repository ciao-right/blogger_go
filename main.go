package main

import (
	"blogger/global"
	"blogger/initialize"
	log "github.com/sirupsen/logrus"
)

func main() {
	global.Viper = initialize.InitViper()
	global.Db = initialize.InitDatabase()
	global.Logger = initialize.InitLog()
	err := initialize.InitTable()
	if err != nil {
		log.Info(err)
	}
	r := initialize.InitEngine()
	global.Logger.WithFields(log.Fields{
		"system": "blogger",
	}).Info("server run at " + global.Viper.GetString("system.port"))
	global.Logger.WithError(err).Info("main")
	r.Run(global.Viper.GetString("system.port"))
}
