package series

import (
	"github.com/gin-gonic/gin"
	"street/pkg/controller"
	"street/web/middleware"
	"street/web/profile"
)

func Routers(group *gin.RouterGroup, ctrl controller.Controller) {
	group.GET("/:id", middleware.MustUriUUID, ctrl.Bare(get))
	group.Use(profile.MustProfile)
	group.POST("/", ctrl.Bare(create))
	group.PUT("/:id", middleware.MustUriUUID, ctrl.Owned(owned), ctrl.Bare(update))
	group.DELETE("/:id", middleware.MustUriUUID, ctrl.Owned(owned), ctrl.Bare(del))
}
