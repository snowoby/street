package data

import (
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"street/pkg/data/config"
	"street/pkg/data/db"
	"street/pkg/data/redis"
	"street/pkg/data/storage"
	"street/pkg/data/task"
)

type Store struct {
	DB             *db.DB
	Storage        *storage.Storage
	SiteConfig     *config.Site
	MultiPartRedis *redis.MultiPartRedis
	Task           *task.Task
}

func NewDefaultEnv() *Store {
	_ = godotenv.Load()

	store := New(config.NewDefault(),
		db.New(db.NewDefaultConfig()),
		storage.New(storage.NewDefaultConfig()),
		redis.New(redis.NewDefaultConfig()),
		task.New(task.NewDefaultConfig()))
	return store
}

func New(site *config.Site, dbClient *db.DB, s3 *storage.Storage, fileRedis *redis.MultiPartRedis, taskClient *task.Task) *Store {

	return &Store{
		SiteConfig:     site,
		DB:             dbClient,
		MultiPartRedis: fileRedis,
		Storage:        s3,
		Task:           taskClient,
	}
}
