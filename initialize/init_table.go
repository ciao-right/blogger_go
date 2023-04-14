package initialize

import (
	"blogger/global"
	"blogger/model/system"
	"errors"
)

func NormalCreateTable(table interface{}) error {
	if HasTable(table) {
		return errors.New("已经存在该表")
	}
	err := global.Db.AutoMigrate(table)
	if err != nil {
		return err
	}
	return nil
}

func HasTable(table interface{}) bool {
	db := global.Db.Migrator()
	return db.HasTable(table)

}

func InitTable() error {
	var err error
	err = NormalCreateTable(&system.User{})
	err = NormalCreateTable(&system.SysDept{})
	return err
}
