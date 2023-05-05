package sys_base

import (
	"blogger/global"
	"blogger/model/system"
	"errors"
)

type UserService struct{}

func (u *UserService) Update(user system.User) error {
	db := global.Db
	if user.ID == 0 {
		return errors.New("id 不能为空")
	}
	res := db.Model(&system.SysDept{}).Where("id = ?", user.ID).Updates(user)
	return res.Error
}

func (u *UserService) Delete(id int) error {
	if id == 0 {
		return errors.New("id 不能为空")
	}
	db := global.Db
	res := db.Where("id = ?", id).Delete(&system.SysDept{})
	return res.Error
}

func (u *UserService) GetList(page, pageSize int, user interface{}) ([]system.User, int, error) {
	db := global.Db
	var users []system.User
	res := db.Where(user).Preload("Dept").Offset((page - 1) * pageSize).Limit(pageSize).Find(&users)
	return users, getTotal(user), res.Error
}
func getTotal(user interface{}) int {
	db := global.Db
	var total int64
	db.Model(&system.User{}).Where(user).Count(&total)
	return int(total)
}
