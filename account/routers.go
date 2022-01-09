package account

import (
	"github.com/gin-gonic/gin"
	"street/handler"
)

func Routers(group *gin.RouterGroup, h handler.Handler) {

	group.POST("/login", h.P(login))
	group.POST("/register", h.P(register))
	group.POST("/refresh", h.P(refreshToken))
	group.GET("/", h.P(AccessTokenMiddleware), h.P(info))

}
