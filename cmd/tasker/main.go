package main

import (
	"fmt"
	"github.com/hibiken/asynq"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"os"
	"street/cmd/config"
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

	factory.New(server, config.NewDefaultEnt(), config.NewDefaultS3()).Run()
}
