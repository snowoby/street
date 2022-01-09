package handler

import (
	"github.com/gin-gonic/gin"
	"street/db"
)

type F func(ctx *gin.Context, store *db.Store)

type handler struct {
	store *db.Store
}

type Handler interface {
	P(f F) gin.HandlerFunc
}

func New(store *db.Store) *handler {
	return &handler{store: store}
}

// P is for Process
func (handler *handler) P(f F) gin.HandlerFunc {
	return func(context *gin.Context) {
		f(context, handler.store)
	}
}
