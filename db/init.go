package db

import (
	"street/ent"
	"time"
)

type config struct {
	RefreshTokenExpireTime time.Duration `json:"refresh_token_expire_time"`
	AccessTokenExpireTime  time.Duration `json:"access_token_expire_time"`
	Domain                 string        `json:"domain"`
}
type store struct {
	*ent.Client
	config
}

func New(client *ent.Client) *store {
	return &store{
		client,
		config{
			//TODO config
			RefreshTokenExpireTime: time.Hour * 24 * 7 * 4,
			AccessTokenExpireTime:  time.Hour,
		},
	}
}

func (s *store) DB() *ent.Client {
	return s.Client
}

func (s *store) Config() config {
	return s.config
}

type Store interface {
	DB() *ent.Client
	Config() config
}
