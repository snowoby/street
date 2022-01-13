package middleware

import (
	"github.com/gin-gonic/gin"
	"street/ent"
	"street/errs"
	"street/pkg/data"
	"street/pkg/data/value"
	"street/pkg/utils"
)

func MustBeOwner(ctx *gin.Context, owner data.Owner) {
	operator := ctx.MustGet(value.StringProfile).(*ent.Profile)
	var object ID
	ok := utils.MustBindUri(ctx, &object)
	if !ok {
		ctx.AbortWithStatusJSON(errs.NotFoundError.Code, errs.NotFoundError)
		return
	}
	ok, err := owner.IsOwner(ctx, operator.ID, object.ID)
	if err != nil {
		e := errs.DatabaseError(err)
		ctx.AbortWithStatusJSON(e.Code, e)
		return
	}
	if !ok {
		ctx.AbortWithStatusJSON(errs.NotBelongsToOperator.Code, errs.NotBelongsToOperator)
		return
	}

	ctx.Next()
	return
}
