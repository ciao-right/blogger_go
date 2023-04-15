package sys_base

import (
	"blogger/global"
	"blogger/model/system"
	"blogger/utils"
	"errors"
)

type BaseService struct{}

func (b *BaseService) Login(loginInfo system.BaseLogin) (user system.User, err error) {
	db := global.Db
	if hasAccount(loginInfo.Account) {
		res := db.Model(&system.User{}).Where("account = ?", loginInfo.Account).Find(&user)
		if res.Error != nil {
			return user, res.Error
		} else {
			if utils.Check(loginInfo.Password, user.Password) {
				return user, nil
			} else {
				return user, errors.New("该账户不存在")
			}
		}
	} else {
		return user, errors.New("该账户不存在")
	}
}

func (b *BaseService) GetToken(account, password string) (token string, err error) {
	return utils.GenerateToken(account, password)
}

func (b *BaseService) Register(user system.User) error {
	db := global.Db
	user.Password = utils.Encode(user.Password)
	res := db.Create(&user)
	return res.Error
}

func hasAccount(account string) bool {
	var user system.User
	res := global.Db.Model(&system.User{}).Where("account = ?", account).First(&user)
	if res.Error != nil {
		global.Logger.Error(res.Error)
		return false
	} else {
		return res.RowsAffected > 0
	}
}
