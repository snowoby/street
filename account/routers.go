package account

import (
	"github.com/gin-gonic/gin"
	"street/handler"
	"street/middleware"
)

func Routers(group *gin.RouterGroup, h handler.Handler) {

	group.POST("/login", h.P(login))
	group.POST("/register", h.P(register))
	group.POST("/refresh", h.P(middleware.RefreshToken))
	group.GET("/", h.P(middleware.MustLogin), h.P(info))

}
