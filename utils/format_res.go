package utils

import (
	"blogger/global"
	"blogger/model/system"
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

func SuccessResCom(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    data,
	})
}

func FormatResForTime(list []system.ModelTime) {
	var resDeptList []struct {
		data       system.ModelTime
		CreatedOn  string `json:"created_on" `
		ModifiedOn string `json:"modified_on" `
	}

	for _, v := range list {
		resDeptList = append(resDeptList, struct {
			data       system.ModelTime
			CreatedOn  string `json:"created_on" `
			ModifiedOn string `json:"modified_on" `
		}{
			data:       v,
			CreatedOn:  v.GetCreatedOn().Format("2006-01-02 15:04:05"),
			ModifiedOn: v.GetModifiedOn().Format("2006-01-02 15:04:05"),
		})
	}

}
