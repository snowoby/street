package composer

import (
	"github.com/gin-gonic/gin"
	"street/ent"
	"street/errs"
	"street/pkg/d"
	"street/pkg/operator"
)

func extractOperator(ctx *gin.Context) *operator.Identity {
	acc := ctx.MustGet(d.StringAccount).(*ent.Account)
	return operator.New(acc)
}

func Bare(f func(ctx *gin.Context) (int, interface{}, error)) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		code, responseValue, err := f(ctx)
		resultProcess(ctx, code, responseValue, err)
	}
}

func Authed(f func(ctx *gin.Context, operator *operator.Identity) (int, interface{}, error)) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		code, responseValue, err := f(ctx, extractOperator(ctx))
		resultProcess(ctx, code, responseValue, err)
	}
}

type id struct {
	ID string `uri:"id" binding:"uuid,required"`
}

func AuthedID(f func(ctx *gin.Context, operator *operator.Identity, id string) (int, interface{}, error)) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var eid id
		err := ctx.ShouldBindUri(&eid)
		if err != nil {
			resultProcess(ctx, 0, nil, err)
			ctx.Abort()
			return
		}

		code, responseValue, err := f(ctx, extractOperator(ctx), eid.ID)
		resultProcess(ctx, code, responseValue, err)
		ctx.Next()

	}
}

func AuthedIDCheck(f func(ctx *gin.Context, operator *operator.Identity, id string) error) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var eid id
		if err := ctx.ShouldBindUri(&eid); err != nil {
			resultProcess(ctx, errs.NotFoundError.Code, errs.NotFoundError, errs.NotFoundError)
			ctx.Abort()
			return
		}

		err := f(ctx, extractOperator(ctx), eid.ID)
		if err != nil {
			resultProcess(ctx, errs.NotBelongsToOperator.Code, errs.NotBelongsToOperator, errs.NotBelongsToOperator)
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}

func ID(f func(ctx *gin.Context, id string) (int, interface{}, error)) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var eid id
		err := ctx.ShouldBindUri(&eid)
		if err != nil {
			resultProcess(ctx, errs.NotFoundError.Code, errs.NotFoundError, errs.NotFoundError)
			return
		}
		code, responseValue, err := f(ctx, eid.ID)
		resultProcess(ctx, code, responseValue, err)
		ctx.Next()
		return
	}
}

func resultProcess(ctx *gin.Context, code int, responseValue interface{}, err error) {
	if err != nil {
		responseError := errs.Detect(err)
		ctx.AbortWithStatusJSON(responseError.Code, struct {
			Message string `json:"message"`
		}{responseError.Message})
		return
	}

	ctx.JSON(code, responseValue)
}
