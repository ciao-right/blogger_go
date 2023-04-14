package main

import (
	"blogger/global"
	"blogger/initialize"
)

func main() {
	global.Viper = initialize.InitViper()
	global.Db = initialize.InitDatabase()
	global.Logger = initialize.InitLog()
	err := initialize.InitTable()
	if err != nil {
		global.Logger.Info(err)
	}
	r := initialize.InitEngine()
	global.Logger.Info("运行成功")
	r.Run()

}
