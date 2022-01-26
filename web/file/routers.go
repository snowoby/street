package file

import (
	"github.com/gin-gonic/gin"
	"street/pkg/controller"
	"street/web/middleware"
)

func Routers(group *gin.RouterGroup, ctrl controller.Controller) {
	//group.Use(account.MustLogin, profile.MustProfile)
	group.POST("/", ctrl.General(create))
	group.PUT("/", ctrl.Bare(upload))
	group.POST("/:id", middleware.MustUriUUID, ctrl.Owned(owned), ctrl.Bare(done))

}
