package series

import (
	"github.com/gin-gonic/gin"
	"street/pkg/controller"
	"street/web/middleware"
	"street/web/profile"
)

func Routers(group *gin.RouterGroup, ctrl controller.Controller) {
	group.GET("/:id", middleware.MustUriUUID, ctrl.Bare(get))
	group.Use(profile.MustUseProfile)
	group.POST("/:pid", ctrl.General(create))
	group.PUT("/:pid/:id", middleware.MustUriUUID, ctrl.Owned(owned), ctrl.Bare(update))
	group.DELETE("/:pid/:id", middleware.MustUriUUID, ctrl.Owned(owned), ctrl.Bare(del))
}
