package site

import (
	"github.com/gin-gonic/gin"
	"street/pkg/data"
)

func siteConfig(_ *gin.Context, data *data.Store) (int, interface{}, error) {
	return 200, data.SiteConfig, nil
}
