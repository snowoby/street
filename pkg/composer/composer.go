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
		if err := ctx.ShouldBindUri(&eid); err == nil {
			code, responseValue, err := f(ctx, extractOperator(ctx), eid.ID)
			resultProcess(ctx, code, responseValue, err)
		}
		resultProcess(ctx, errs.NotFoundError.Code, errs.NotFoundError, errs.NotFoundError)
		return

	}
}

func AuthedIDCheck(f func(ctx *gin.Context, operator *operator.Identity, id string) error) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var eid id
		if err := ctx.ShouldBindUri(&eid); err != nil {
			resultProcess(ctx, errs.NotFoundError.Code, errs.NotFoundError, errs.NotFoundError)
			return
		}

		err := f(ctx, nil, eid.ID)
		if err != nil {
			resultProcess(ctx, errs.NotBelongsToOperator.Code, errs.NotBelongsToOperator, errs.NotBelongsToOperator)
			return
		}
		ctx.Next()
	}
}

func ID(f func(ctx *gin.Context, id string) (int, interface{}, error)) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var eid id
		if err := ctx.ShouldBindUri(&eid); err == nil {
			code, responseValue, err := f(ctx, eid.ID)
			resultProcess(ctx, code, responseValue, err)
		}
		resultProcess(ctx, 0, errs.NotFoundError, errs.NotFoundError)
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
