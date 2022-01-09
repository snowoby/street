package profile

import (
	"github.com/gin-gonic/gin"
	"street/handler"
)

func Routers(group *gin.RouterGroup, h handler.Handler) {
	group.POST("/", h.P(createEpisode))
	group.PUT("/:id", h.P(updateEpisode))
	group.DELETE("/:id", h.P(deleteEpisode))
}

func PublicRouters(group *gin.RouterGroup, h handler.Handler) {
	group.GET("/:id", h.P(getEpisode))
}
