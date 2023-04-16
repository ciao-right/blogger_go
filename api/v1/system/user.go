package system

import (
	"blogger/model/system"
	"blogger/service"
	"blogger/utils"
	"github.com/gin-gonic/gin"
	"strconv"
)

type UserApi struct{}

func (u *UserApi) Update(c *gin.Context) {
	var user system.User
	err := utils.InitPost(c, &user)
	if err != nil {
		utils.FailResCom(c, err, nil)
		return
	}
	logic := service.Service{}.UserService
	sErr := logic.Update(user)
	if sErr != nil {
		utils.FailResCom(c, sErr, nil)
		return
	}
	utils.SuccessResCom(c, 1)
}

func (u *UserApi) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	logic := service.Service{}.UserService
	sErr := logic.Delete(id)
	if sErr != nil {
		utils.FailResCom(c, sErr, nil)
		return
	}
	utils.SuccessResCom(c, 1)
}

func (u *UserApi) GetList(c *gin.Context) {
	var userSearch system.SearchPage
	userSearch.Condition = struct {
		name string
		dept int
		sex  int
	}{}
	err := utils.InitPost(c, &userSearch)
	utils.InitPage(&userSearch.Page, &userSearch.PageSize)
	if err != nil {
		utils.FailResCom(c, err, nil)
		return
	}
	logic := service.Service{}.UserService
	list, total, sErr := logic.GetList(userSearch.Page, userSearch.PageSize, userSearch.Condition)
	if sErr != nil {
		utils.FailResCom(c, sErr, nil)
		return
	}

	type _userList struct {
		system.User
		CreatedOn  string `json:"created_on" `
		ModifiedOn string `json:"modified_on" `
	}

	var resDeptList []_userList
	for _, user := range list {
		resDeptList = append(resDeptList, _userList{
			User:       user,
			CreatedOn:  user.GetCreatedOn().Format("2006-01-02 15:04:05"),
			ModifiedOn: user.GetModifiedOn().Format("2006-01-02 15:04:05"),
		})
	}

	var data = make(map[string]interface{})
	data["list"] = resDeptList
	data["total"] = total
	utils.SuccessResCom(c, data)

}
