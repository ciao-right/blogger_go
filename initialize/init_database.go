package initialize

import (
	"blogger/global"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDatabase() *gorm.DB {
	if db, err := gorm.Open(mysql.Open(getDatabaseConfig()), &gorm.Config{}); err == nil {
		return db
	} else {
		panic("连接失败")
	}
}

func getDatabaseConfig() string {
	username := global.Viper.GetString("mysql.username")
	password := global.Viper.GetString("mysql.password")
	address := global.Viper.GetString("mysql.address")
	port := global.Viper.GetString("mysql.port")
	databaseName := global.Viper.GetString("mysql.database_name")
	fmt.Println(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, address, port, databaseName))
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, address, port, databaseName)
}
