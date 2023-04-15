package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func InitApi[T any](ctx *gin.Context, target T) error {
	err := ctx.ShouldBindJSON(&target)
	if err != nil {
		FailResCom(ctx, err, nil)
		return err
	}
	validate := validator.New()
	vErr := validate.Struct(&target)
	if vErr != nil {
		fmt.Println(vErr)
		FailResCom(ctx, vErr, nil)
		return vErr
	}
	return nil
}

func InitPage(page, pageSize *int) {
	if *page == 0 {
		*page = 1
	}
	if *pageSize == 0 {
		*pageSize = 10
	}
}
