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
	s.router.GET("/", func(ctx *gin.Context) {
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
		Domain:                 os.Getenv("DOMAIN"),
		storageAccessEndpoint:  os.Getenv("STORAGE_ACCESS_ENDPOINT"),
		StorageEndpoint:        os.Getenv("STORAGE_PUBLIC_ENDPOINT"),
		StorageBucket:          os.Getenv("STORAGE_BUCKET"),
		SiteName:               os.Getenv("SITE_NAME"),
		RefreshTokenExpireTime: time.Hour * 24 * 7 * 4,
		AccessTokenExpireTime:  time.Hour,
	}
}

func (s *Site) StorageAccessEndpoint() string {
	return s.storageAccessEndpoint
}
