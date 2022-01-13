package profile

import (
	"github.com/gin-gonic/gin"
	"street/pkg/handler"
	"street/web/account"
)

func Routers(group *gin.RouterGroup, h handler.Handler) {
	group.GET("/:id", h.P(getProfile))
	group.Use(h.P(account.MustLogin))
	group.GET("/", h.P(accountProfiles))
	group.POST("/", h.P(createProfile))
	group.PUT("/:id", h.P(MustProfile), h.P(updateProfile))

}
