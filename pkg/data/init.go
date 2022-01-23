package data

import (
	"street/ent"
	"street/pkg/data/storage"
	"time"
)

type siteConfig struct {
	RefreshTokenExpireTime time.Duration `json:"refresh_token_expire_time"`
	AccessTokenExpireTime  time.Duration `json:"access_token_expire_time"`
	Domain                 string        `json:"domain"`
}
type Store struct {
	*db
	*siteConfig
	*storage.Storage
	*series
	*episode
	*profile
	*file
}

type db struct {
	client *ent.Client
}

func New(client *ent.Client, s3 *storage.Storage) *Store {

	return &Store{
		&db{client},
		&siteConfig{
			//TODO config
			RefreshTokenExpireTime: time.Hour * 24 * 7 * 4,
			AccessTokenExpireTime:  time.Hour,
		},
		s3,
		&series{client.Series},
		&episode{client.Episode},
		&profile{client.Profile},
		&file{client.File},
	}
}

func (s *Store) DB() *db {
	return s.db
}

func (s *Store) Config() *siteConfig {
	return s.siteConfig
}

func (s *Store) Series() *series {
	return s.series
}

func (s *Store) Episode() *episode {
	return s.episode
}

func (s *Store) Profile() *profile {
	return s.profile
}

func (s *Store) File() *file {
	return s.file
}
