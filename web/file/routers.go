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

	single := group.Group("single")
	single.POST("/:pid", ctrl.General(createSingle))
	single.PUT("/:pid/:id", middleware.MustUriUUID, ctrl.General(putSingle))

	large := group.Group("large")
	large.POST("/:pid", ctrl.General(createMulti))
	large.PUT("/:pid/:id/:part_id", middleware.MustUriUUID, ctrl.Bare(uploadMulti))
	large.POST("/:pid/:id", middleware.MustUriUUID, ctrl.Owned(owned), ctrl.Bare(doneMulti))

}
