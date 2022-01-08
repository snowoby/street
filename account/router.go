package account

import (
	"github.com/gin-gonic/gin"
	"street/db"
)

type Flatten func(ctx *gin.Context, store *store)

func Handle(s db.Store, handle Flatten) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		handle(ctx, &store{s})
	}
}

func Routers(group *gin.RouterGroup, store db.Store) {
	group.POST("/login", Handle(store, login))
	group.POST("/register", Handle(store, register))
}
