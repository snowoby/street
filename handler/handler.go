package handler

import (
	"github.com/gin-gonic/gin"
	"street/data"
)

type F func(ctx *gin.Context, store *data.Store)

type handler struct {
	store *data.Store
}

type Handler interface {
	// P is for Process
	P(f F) gin.HandlerFunc
}

func New(store *data.Store) *handler {
	return &handler{store: store}
}

func (handler *handler) P(f F) gin.HandlerFunc {
	return func(context *gin.Context) {
		f(context, handler.store)
	}
}
