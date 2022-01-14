package profile

import (
	"github.com/gin-gonic/gin"
	"street/pkg/controller"
	"street/web/profile"
)

func Routers(group *gin.RouterGroup, ctrl controller.Controller) {
	group.GET("/:id", ctrl.Bare(get))
	group.Use(profile.MustProfile)
	group.POST("/", ctrl.General(create))
	group.PUT("/:id", ctrl.Bare(update))
	group.DELETE("/:id", ctrl.Bare(del))
}
