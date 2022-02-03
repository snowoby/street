package data

import (
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"street/pkg/data/db"
	"street/pkg/data/redis"
	"street/pkg/data/storage"
	"street/pkg/data/task"
	"time"
)

type siteConfig struct {
	RefreshTokenExpireTime time.Duration `json:"refresh_token_expire_time"`
	AccessTokenExpireTime  time.Duration `json:"access_token_expire_time"`
	Domain                 string        `json:"domain"`
}
type Store struct {
	DB             *db.DB
	Storage        *storage.Storage
	SiteConfig     *siteConfig
	MultiPartRedis *redis.MultiPartRedis
	Task           *task.Task
}

func NewDefaultEnv() *Store {
	_ = godotenv.Load()

	store := New(db.New(db.NewDefaultConfig()),
		storage.New(storage.NewDefaultConfig()),
		redis.New(redis.NewDefaultConfig()),
		task.New(task.NewDefaultConfig()))
	return store
}

func New(dbClient *db.DB, s3 *storage.Storage, fileRedis *redis.MultiPartRedis, taskClient *task.Task) *Store {

	return &Store{
		SiteConfig: &siteConfig{
			//TODO config
			RefreshTokenExpireTime: time.Hour * 24 * 7 * 4,
			AccessTokenExpireTime:  time.Hour,
		},
		DB:             dbClient,
		MultiPartRedis: fileRedis,
		Storage:        s3,
		Task:           taskClient,
	}
}
