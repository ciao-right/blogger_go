package utils

import (
	"blogger/global"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ErrToString(err error) string {
	errString := fmt.Sprintf("%s", err)
	return errString
}

func FailResCom(c *gin.Context, message error, data interface{}) {
	global.Logger.WithError(message).Info(message)
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": ErrToString(message),
		"data":    data,
	})
}
