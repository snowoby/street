package controller

import (
	"github.com/gin-gonic/gin"
	"street/ent"
	"street/errs"
	"street/pkg/data"
	"street/pkg/data/value"
)

type F func(ctx *gin.Context, store *data.Store)
type Func func(ctx *gin.Context, store *data.Store, visitor *Identity) (int, interface{}, error)
type OwnerFunc func(ctx *gin.Context, store data.Owner)

type controller struct {
	store *data.Store
}

type Controller interface {
	Bare(f F) gin.HandlerFunc
	General(f Func) gin.HandlerFunc
	Store() *data.Store
	Owned(f OwnerFunc, store data.Owner) gin.HandlerFunc
}

func New(store *data.Store) *controller {
	return &controller{store: store}
}

func (controller *controller) Store() *data.Store {
	return controller.store
}

func (controller *controller) General(nf Func) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		a, _ := ctx.Get(value.StringAccount)
		p, _ := ctx.Get(value.StringProfile)
		id := &Identity{
			account: a.(*ent.Account),
			profile: p.(*ent.Profile),
		}
		code, responseValue, err := nf(ctx, controller.store, id)

		if err != nil {
			responseError := errs.Detect(err)
			ctx.AbortWithStatusJSON(responseError.Code(), responseError.Message())
			return
		}

		ctx.JSON(code, responseValue)
	}
}

func (controller *controller) Bare(f F) gin.HandlerFunc {
	return func(context *gin.Context) {
		f(context, controller.store)
	}
}

func (controller *controller) Owned(f OwnerFunc, store data.Owner) gin.HandlerFunc {
	return func(context *gin.Context) {
		f(context, store)
	}
}
