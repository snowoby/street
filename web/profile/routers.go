package profile

import (
	"github.com/gin-gonic/gin"
	"street/pkg/controller"
	"street/web/account"
	"street/web/middleware"
)

func Routers(group *gin.RouterGroup, ctrl controller.Controller) {
	group.GET("/:id", middleware.MustUriUUID, ctrl.Bare(get))
	group.Use(account.MustLogin)
	//group.GET("/", ctrl.General(accountProfiles))
	group.POST("/", ctrl.General(create))
	group.PUT("/:id", ctrl.Owned(owned), ctrl.General(update))

}
