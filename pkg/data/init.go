package data

import (
	"github.com/go-redis/redis/v8"
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
	Account        *account
	Token          *token
	SiteConfig     *siteConfig
	MultiPartRedis *multiPartRedis
	Storage        *storage.Storage
	Series         *series
	Episode        *episode
	Profile        *profile
	File           *file
}

func New(client *ent.Client, s3 *storage.Storage, fileRedis *redis.Client) *Store {

	return &Store{
		Account: &account{client.Account},
		Token:   &token{client.Token},
		SiteConfig: &siteConfig{
			//TODO config
			RefreshTokenExpireTime: time.Hour * 24 * 7 * 4,
			AccessTokenExpireTime:  time.Hour,
		},
		MultiPartRedis: &multiPartRedis{fileRedis},
		Storage:        s3,
		Series:         &series{client.Series},
		Episode:        &episode{client.Episode},
		Profile:        &profile{client.Profile},
		File:           &file{client.File},
	}
}

//func (s *Store) DB() *db {
//	return s.db
//}
//
//func (s *Store) Config() *siteConfig {
//	return s.siteConfig
//}
//
//func (s *Store) Series() *series {
//	return s.series
//}
//
//func (s *Store) Episode() *episode {
//	return s.episode
//}
//
//func (s *Store) Profile() *profile {
//	return s.profile
//}
//
//func (s *Store) File() *file {
//	return s.file
//}
//func (s *Store) Storage() *storage.Storage {
//	return s.storage
//}
//
//func (s *Store) MultiPartRedis() *file {
//	return s.file
//}
