package profile

import (
	"github.com/gin-gonic/gin"
	"street/pkg/controller"
	"street/web/account"
)

func Routers(group *gin.RouterGroup, ctrl controller.Controller) {
	group.GET("/:id", ctrl.Bare(get))
	group.Use(account.MustLogin)
	group.GET("/", ctrl.General(accountProfiles))
	group.POST("/", ctrl.General(create))
	group.PUT("/:id", MustProfile, ctrl.General(update))

}
