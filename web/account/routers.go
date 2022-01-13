package account

import (
	"github.com/gin-gonic/gin"
	"street/pkg/controller"
)

func Routers(group *gin.RouterGroup, ctrl controller.Controller) {
	group.POST("/refresh", ctrl.Bare(MustRefresh))
	group.POST("/login", ctrl.Bare(login))
	group.POST("/register", ctrl.General(register))
	group.GET("/", MustLogin, ctrl.General(info))
}
