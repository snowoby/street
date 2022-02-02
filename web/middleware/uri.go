package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"street/ent/schema"
	"street/errs"
	"street/pkg/data/value"
)

func TryUriUUID(ctx *gin.Context) {
	type ID struct {
		ID string `uri:"id" binding:"uuid,required"`
	}

	var id ID
	err := ctx.ShouldBindUri(&id)
	if err == nil {
		ctx.Set(value.StringObjectUUID, uuid.MustParse(id.ID))
	}
	ctx.Next()
}

func TryUriSID(ctx *gin.Context) {
	type ID struct {
		ID schema.ID `uri:"id" binding:"sid"`
	}

	var id ID
	err := ctx.ShouldBindUri(id)
	if err == nil {
		ctx.Set(value.StringSID, id.ID)
	}
	ctx.Next()
}

func MustUriUUID(ctx *gin.Context) {
	_, ok := ctx.Get(value.StringObjectUUID)
	if !ok {
		ctx.AbortWithStatusJSON(errs.NotFoundError.Code, errs.NotFoundError)
	}
	ctx.Next()
}

func MustUriSID(ctx *gin.Context) {
	_, ok := ctx.Get(value.StringSID)
	if !ok {
		ctx.AbortWithStatusJSON(errs.NotFoundError.Code, errs.NotFoundError)
	}
	ctx.Next()
}
