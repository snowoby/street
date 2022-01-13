package controller

import (
	"github.com/gin-gonic/gin"
	"street/ent"
	"street/errs"
	"street/pkg/data"
	"street/pkg/data/value"
)

type F func(ctx *gin.Context, store *data.Store) (int, interface{}, error)
type Func func(ctx *gin.Context, store *data.Store, visitor *Identity) (int, interface{}, error)

//type UUIDFunc func(ctx *gin.Context, store *data.Store, visitor *Identity, uuid *uuid.UUID) (int, interface{}, error)
type OwnerFunc func(ctx *gin.Context, store data.Owner)

type controller struct {
	store *data.Store
}

type Controller interface {
	Bare(f F) gin.HandlerFunc
	// General will try to bind visitor identity as param.
	General(f Func) gin.HandlerFunc
	//GeneralUUID(f UUIDFunc) gin.HandlerFunc
	Store() *data.Store
	Owned(f OwnerFunc, store data.Owner) gin.HandlerFunc
}

func New(store *data.Store) *controller {
	return &controller{store: store}
}

func resultProcess(ctx *gin.Context, code int, responseValue interface{}, err error) {
	if err != nil {
		responseError := errs.Detect(err)
		ctx.AbortWithStatusJSON(responseError.Code(), responseError.Message())
		return
	}

	ctx.JSON(code, responseValue)
}

func (controller *controller) Store() *data.Store {
	return controller.store
}

func extractOperator(ctx *gin.Context) *Identity {
	a, _ := ctx.Get(value.StringAccount)
	p, _ := ctx.Get(value.StringProfile)
	return &Identity{
		account: a.(*ent.Account),
		profile: p.(*ent.Profile),
	}
}

func (controller *controller) Bare(f F) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		code, responseValue, err := f(ctx, controller.store)
		resultProcess(ctx, code, responseValue, err)

	}
}

func (controller *controller) General(nf Func) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := extractOperator(ctx)
		code, responseValue, err := nf(ctx, controller.store, id)
		resultProcess(ctx, code, responseValue, err)
	}
}

//
//func (controller *controller) GeneralUUID(f UUIDFunc) gin.HandlerFunc {
//	return func(ctx *gin.Context) {
//		identity := extractOperator(ctx)
//		id, _ := extractUUID(ctx)
//
//		code, responseValue, err := f(ctx, controller.store, identity, id)
//
//		if err != nil {
//			responseError := errs.Detect(err)
//			ctx.AbortWithStatusJSON(responseError.Code(), responseError.Message())
//			return
//		}
//
//		ctx.JSON(code, responseValue)
//	}
//
//}

func (controller *controller) Owned(f OwnerFunc, store data.Owner) gin.HandlerFunc {
	return func(context *gin.Context) {
		f(context, store)
	}
}
