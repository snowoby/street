package config

import (
	"os"
	"street/ent"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/go-redis/redis/v8"
	"github.com/hibiken/asynq"
)

func NewDefaultEnt() *ent.Client {
	client, err := ent.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err)
	}
	// err = client.Schema.Create(context.Background())
	// if err != nil {
	// 	panic(err)
	// }
	return client
}

func NewDefaultRedis() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS"),
		DB:   0,
	})
}

func NewDefaultS3() *aws.Config {
	return &aws.Config{
		Credentials:      credentials.NewStaticCredentials(os.Getenv("STORAGE_ACCESS_KEY"), os.Getenv("STORAGE_SECRET_KEY"), ""),
		Endpoint:         aws.String(os.Getenv("STORAGE_ACCESS_ENDPOINT")),
		Region:           aws.String(os.Getenv("STORAGE_REGION")),
		DisableSSL:       aws.Bool(true),
		S3ForcePathStyle: aws.Bool(true),
	}
}

func NewDefaultAsynq() *asynq.Client {
	return asynq.NewClient(asynq.RedisClientOpt{Addr: os.Getenv("REDIS"), DB: 1})
}
