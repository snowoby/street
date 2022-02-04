package site

import (
	"github.com/gin-gonic/gin"
	"street/pkg/controller"
)

func Routers(group *gin.RouterGroup, ctrl controller.Controller) {
	group.GET("", ctrl.Bare(siteConfig))
}
