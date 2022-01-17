package episode

import (
	"github.com/gin-gonic/gin"
	"street/pkg/controller"
	"street/web/middleware"
	"street/web/profile"
)

func Routers(group *gin.RouterGroup, ctrl controller.Controller) {
	group.GET("/:id", middleware.MustUriUUID, ctrl.Bare(get))
	group.Use(profile.MustProfile)
	group.POST("/", ctrl.General(create))
	group.Use(middleware.MustUriUUID, ctrl.Owned(owned))
	group.PUT("/:id", ctrl.Bare(update))
	group.DELETE("/:id", ctrl.Bare(del))
}