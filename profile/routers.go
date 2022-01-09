package profile

import (
	"github.com/gin-gonic/gin"
	"street/handler"
)

func Routers(group *gin.RouterGroup, h handler.Handler) {
	group.GET("/", h.P(accountProfiles))
	group.POST("/", h.P(createProfile))
	group.PUT("/:id", h.P(Middleware), h.P(updateProfile))

}

func PublicRouters(group *gin.RouterGroup, h handler.Handler) {
	group.GET("/:id", h.P(getProfile))
}
