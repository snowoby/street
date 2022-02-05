package series

import (
	"github.com/gin-gonic/gin"
	"street/pkg/controller"
	"street/web/middleware"
	"street/web/profile"
)

func Routers(group *gin.RouterGroup, ctrl controller.Controller) {
	group.GET("/:id", middleware.MustUriUUID, ctrl.Bare(get))

	group.Use(profile.TryProfile)
	group.POST("/", ctrl.General(create))
	group.Use(middleware.MustUriUUID)
	group.Use(ctrl.Owned(owned))

	group.PUT("/:id", ctrl.Bare(update))
	group.DELETE("/:id", ctrl.Bare(del))
}
