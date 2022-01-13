package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"street/errs"
	"street/pkg/data/value"
)

func TryUriUUID(ctx *gin.Context) {
	type ID struct {
		ID uuid.UUID `uri:"id" binding:"uuid"`
	}

	var id ID
	err := ctx.ShouldBindUri(id)
	if err != nil {
		ctx.Set(value.StringObjectUUID, nil)
	} else {
		ctx.Set(value.StringObjectUUID, &id)
	}
	ctx.Next()
}

func MustUriUUID(ctx *gin.Context) {
	_, ok := ctx.Get(value.StringObjectUUID)
	if !ok {
		ctx.AbortWithStatusJSON(errs.NotFoundError.Code(), errs.NotFoundError)
	}
	ctx.Next()
}
