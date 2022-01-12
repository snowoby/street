package handler

import (
	"github.com/gin-gonic/gin"
	"street/data"
	"street/data/value"
	"street/ent"
	"street/errs"
)

type identity struct {
	Account *ent.Account
	Profile *ent.Profile
}

type F func(ctx *gin.Context, store *data.Store)
type NF func(ctx *gin.Context, store *data.Store, visitor *identity) (int, interface{}, errs.ResponseError)
type OwnerFunc func(ctx *gin.Context, store data.Owner)

type handler struct {
	store *data.Store
}

type Handler interface {
	// P is for Process
	P(f F) gin.HandlerFunc
	NP(nf NF) gin.HandlerFunc
	Store() *data.Store
	OwnerFunc(f OwnerFunc, store data.Owner) gin.HandlerFunc
}

func New(store *data.Store) *handler {
	return &handler{store: store}
}

func (handler *handler) Store() *data.Store {
	return handler.store
}

func (handler *handler) NP(nf NF) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		a, _ := ctx.Get(value.StringAccount)
		p, _ := ctx.Get(value.StringProfile)
		id := &identity{
			Account: a.(*ent.Account),
			Profile: p.(*ent.Profile),
		}

		code, responseValue, err := nf(ctx, handler.store, id)

		if err != nil {
			ctx.AbortWithStatusJSON(err.Code(), err.Message())
			return
		}

		ctx.JSON(code, responseValue)
	}
}

func (handler *handler) P(f F) gin.HandlerFunc {
	return func(context *gin.Context) {
		f(context, handler.store)
	}
}

func (handler *handler) OwnerFunc(f OwnerFunc, store data.Owner) gin.HandlerFunc {
	return func(context *gin.Context) {
		f(context, store)
	}
}
