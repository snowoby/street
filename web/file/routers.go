package file

import (
	"github.com/gin-gonic/gin"
	"street/pkg/controller"
	"street/web/account"
	"street/web/middleware"
	"street/web/profile"
)

func Routers(group *gin.RouterGroup, ctrl controller.Controller) {
	group.Use(account.MustLogin, profile.MustProfile)
	group.POST("/:pid", ctrl.General(create))
	group.PUT("/:pid/:id/:part_id", ctrl.Bare(upload))
	group.POST("/:pid/:id", middleware.MustUriUUID, ctrl.Owned(owned), ctrl.Bare(done))

}
