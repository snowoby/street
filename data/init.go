package data

import (
	"street/ent"
	"time"
)

type config struct {
	RefreshTokenExpireTime time.Duration `json:"refresh_token_expire_time"`
	AccessTokenExpireTime  time.Duration `json:"access_token_expire_time"`
	Domain                 string        `json:"domain"`
}
type Store struct {
	*db
	config
	*series
}

type db struct {
	client *ent.Client
}

func New(client *ent.Client) *Store {
	return &Store{
		&db{client},
		config{
			//TODO config
			RefreshTokenExpireTime: time.Hour * 24 * 7 * 4,
			AccessTokenExpireTime:  time.Hour,
		},
		&series{client.Series},
	}
}

func (s *Store) DB() *db {
	return s.db
}

func (s *Store) Config() config {
	return s.config
}

func (s *Store) Series() *series {
	return s.series
}
