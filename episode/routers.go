package profile

import (
	"github.com/gin-gonic/gin"
	"street/handler"
	"street/middleware"
)

func Routers(group *gin.RouterGroup, h handler.Handler) {
	group.GET("/:id", h.P(get))
	group.Use(h.P(middleware.MustProfile))
	group.POST("/", h.P(create))
	group.PUT("/:id", h.P(update))
	group.DELETE("/:id", h.P(del))
}
