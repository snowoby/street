package profile

import (
	"github.com/gin-gonic/gin"
	"street/handler"
	"street/middleware"
)

func Routers(group *gin.RouterGroup, h handler.Handler) {
	group.GET("/:id", h.P(getEpisode))
	group.Use(h.P(middleware.MustProfile))
	group.POST("/", h.P(createEpisode))
	group.PUT("/:id", h.P(updateEpisode))
	group.DELETE("/:id", h.P(deleteEpisode))
}
