package profile

import (
	"github.com/gin-gonic/gin"
	"street/handler"
	"street/middleware"
)

func Routers(group *gin.RouterGroup, h handler.Handler) {
	group.GET("/:id", h.P(getProfile))
	group.Use(h.P(middleware.MustLogin))
	group.GET("/", h.P(accountProfiles))
	group.POST("/", h.P(createProfile))
	group.PUT("/:id", h.P(middleware.MustProfile), h.P(updateProfile))

}
