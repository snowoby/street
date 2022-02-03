package data

import (
	"github.com/go-redis/redis/v8"
	"github.com/hibiken/asynq"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"golang.org/x/net/context"
	"os"
	"street/ent"
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
	Account        *account
	Token          *token
	SiteConfig     *siteConfig
	MultiPartRedis *multiPartRedis
	Storage        *storage.Storage
	Series         *series
	Episode        *episode
	Profile        *profile
	File           *file
	Task           *task.Task
}

func NewDefaultEnv() *Store {
	err := godotenv.Load()
	client, err := ent.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err)
	}
	err = client.Schema.Create(context.Background())
	if err != nil {
		panic(err)
	}

	filePartRedis := redis.NewClient(&redis.Options{
		Addr: os.Getenv("redis"),
		DB:   0, // use default DB
	})
	taskClient := asynq.NewClient(asynq.RedisClientOpt{Addr: os.Getenv("redis"), DB: 1})

	store := New(client, storage.New(storage.NewDefaultConfig()), filePartRedis, task.New(taskClient))
	return store
}

func New(client *ent.Client, s3 *storage.Storage, fileRedis *redis.Client, taskClient *task.Task) *Store {

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
		Task:           taskClient,
	}
}
