package system

import (
	"blogger/model/system"
	"github.com/gin-gonic/gin"
	"net/http"
)

type BaseApi struct{}

func (b *BaseApi) Login(ctx *gin.Context) {
	var user system.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"msg":  "something has error",
			"code": 0,
		})
	}

}
