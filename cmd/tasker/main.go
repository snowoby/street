package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/hibiken/asynq"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"golang.org/x/net/context"
	"os"
	"street/ent"
	"street/pkg/factory"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
		return
	}
}
func main() {

	server := asynq.NewServer(
		asynq.RedisClientOpt{Addr: os.Getenv("redis"), DB: 1},
		asynq.Config{
			// Specify how many concurrent workers to use
			Concurrency: 10,
			// Optionally specify multiple queues with different priority.
			Queues: map[string]int{
				"avatar":  8,
				"image":   7,
				"default": 2,
				"low":     1,
			},
			// See the godoc for other configuration options
		},
	)

	factory.New(server, NewDefaultEnt(), NewDefaultS3()).Run()
}

func NewDefaultEnt() *ent.Client {
	client, err := ent.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err)
	}
	err = client.Schema.Create(context.Background())
	if err != nil {
		panic(err)
	}
	return client
}

func NewDefaultS3() *aws.Config {
	return &aws.Config{
		Credentials:      credentials.NewStaticCredentials(os.Getenv("s3_accesskey"), os.Getenv("s3_secretkey"), ""),
		Endpoint:         aws.String(os.Getenv("storage_access_endpoint")),
		Region:           aws.String(os.Getenv("s3_region")),
		DisableSSL:       aws.Bool(true),
		S3ForcePathStyle: aws.Bool(true),
	}
}

func NewDefaultAsynq() *asynq.Client {
	return asynq.NewClient(asynq.RedisClientOpt{Addr: os.Getenv("redis"), DB: 1})
}
