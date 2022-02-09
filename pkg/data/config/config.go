package config

import (
	"os"
	"time"
)

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
