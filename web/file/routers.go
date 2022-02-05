package file

import (
	"github.com/gin-gonic/gin"
	"street/pkg/controller"
	"street/web/account"
	"street/web/middleware"
	"street/web/profile"
)

func Routers(group *gin.RouterGroup, ctrl controller.Controller) {
	group.Use(account.MustLogin, profile.MustHaveProfile)

	single := group.Group("single")
	single.POST("", ctrl.General(createSingle))
	single.PUT("/:id", middleware.MustUriUUID, ctrl.Owned(owned), ctrl.General(putSingle))

	large := group.Group("large")
	large.POST("/", ctrl.General(createMulti))
	large.PUT("/:id/:part_id", middleware.MustUriUUID, ctrl.Bare(uploadMulti))
	large.POST("/:id", middleware.MustUriUUID, ctrl.Owned(owned), ctrl.Bare(doneMulti))

}
