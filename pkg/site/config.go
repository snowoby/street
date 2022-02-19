package site

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"time"
)

type service struct {
	router *gin.RouterGroup
	site   *Site
}

func New(site *Site, router *gin.RouterGroup) *service {
	s := &service{
		site:   site,
		router: router,
	}
	s.registerRouters()
	return s
}

func (s *service) registerRouters() {
	s.router.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, s.site)
	})

}

type Site struct {
	RefreshTokenExpireTime time.Duration `json:"refreshTokenExpireTime"`
	AccessTokenExpireTime  time.Duration `json:"accessTokenExpireTime"`
	SiteName               string        `json:"siteName"`
	Domain                 string        `json:"domain"`
	storageAccessEndpoint  string
	StorageEndpoint        string `json:"storageEndpoint"`
	StorageBucket          string `json:"storageBucket"`
}

func NewDefault() *Site {
	return &Site{
		Domain:                 os.Getenv("site"),
		storageAccessEndpoint:  os.Getenv("storage_access_endpoint"),
		StorageEndpoint:        os.Getenv("storage_endpoint"),
		StorageBucket:          os.Getenv("storage_bucket"),
		SiteName:               os.Getenv("site_name"),
		RefreshTokenExpireTime: time.Hour * 24 * 7 * 4,
		AccessTokenExpireTime:  time.Hour,
	}
}

func (s *Site) StorageAccessEndpoint() string {
	return s.storageAccessEndpoint
}
