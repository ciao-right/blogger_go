package system

import (
	"blogger/model/system"
	"blogger/service"
	"blogger/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type BaseApi struct{}

func (b *BaseApi) Login(ctx *gin.Context) {
	var user system.BaseLogin
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		fmt.Println(err)
		utils.FailResCom(ctx, err, nil)
		return
	}
	validate := validator.New()
	vErr := validate.Struct(&user)
	if vErr != nil {
		fmt.Println(vErr)
		utils.FailResCom(ctx, vErr, nil)
		return
	}
	logic := service.Service{}.BaseService
	option := system.BaseLogin{
		Account:  user.Account,
		Password: user.Password,
	}
	userInfo, sErr := logic.Login(option)
	token, _ := logic.GetToken(option.Account, option.Password)
	if sErr != nil {
		fmt.Println(sErr)
		utils.FailResCom(ctx, sErr, nil)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "登陆成功",
		"data":    userInfo,
		"token":   token,
	})
}

func (b *BaseApi) Register(ctx *gin.Context) {
	var user system.User
	v := validator.New()
	jErr := ctx.ShouldBindJSON(&user)
	if jErr != nil {
		utils.FailResCom(ctx, jErr, nil)
		return
	}
	vErr := v.Struct(&user)
	if vErr != nil {
		utils.FailResCom(ctx, vErr, nil)
		return
	}
	logic := service.Service{}.BaseService
	err := logic.Register(user)
	if err != nil {
		utils.FailResCom(ctx, err, nil)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    1,
	})
}
