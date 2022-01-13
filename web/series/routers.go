package profile

import (
	"github.com/gin-gonic/gin"
	"street/pkg/handler"
	middleware2 "street/web/middleware"
)

func Routers(group *gin.RouterGroup, h handler.Handler) {
	group.GET("/:id", h.P(get))
	group.Use(h.P(middleware2.MustProfile))
	group.POST("/", h.P(create))
	group.Use(h.OwnerFunc(middleware2.MustBeOwner, h.Store().Series()))
	group.PUT("/:id", h.P(update))
	group.DELETE("/:id", h.P(del))
}
