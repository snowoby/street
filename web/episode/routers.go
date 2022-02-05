package episode

import (
	"github.com/gin-gonic/gin"
	"street/pkg/controller"
	"street/web/middleware"
	"street/web/profile"
)

func Routers(group *gin.RouterGroup, ctrl controller.Controller) {
	group.GET("/", ctrl.Bare(getAll))
	group.GET("/:id", middleware.MustUriUUID, ctrl.Bare(get))
	group.Use(profile.MustUseProfile)
	group.POST("/:pid", ctrl.General(create))
	group.Use(middleware.MustUriUUID, ctrl.Owned(owned))
	group.PUT("/:pid/:id", ctrl.Bare(update))
	group.DELETE("/:pid/:id", ctrl.Bare(del))
}
