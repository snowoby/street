package utils

import (
	"github.com/gin-gonic/gin"
	"street/errs"
)

func MustBindUri(ctx *gin.Context, bind interface{}) bool {
	err := ctx.ShouldBindUri(bind)
	if err != nil {
		e := errs.BindingError(err)
		ctx.JSON(e.Code, e)
		return false
	}
	return true
}

func MustBindJSON(ctx *gin.Context, bind interface{}) bool {
	err := ctx.ShouldBindJSON(bind)
	if err != nil {
		e := errs.BindingError(err)
		ctx.JSON(e.Code, e)
		return false
	}
	return true
}