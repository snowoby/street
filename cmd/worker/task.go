package main

import (
	"github.com/hibiken/asynq"
	"gopkg.in/gographics/imagick.v3/imagick"
	"log"
	"os"
	"street/pkg/controller"
	"street/pkg/data"
	"street/pkg/data/value"
)

func init() {
	imagick.Initialize()
}

func worker(controller controller.Controller) {
	srv := asynq.NewServer(
		asynq.RedisClientOpt{Addr: os.Getenv("redis"), DB: 1},
		asynq.Config{
			// Specify how many concurrent workers to use
			Concurrency: 10,
			// Optionally specify multiple queues with different priority.
			Queues: map[string]int{
				"critical": 2,
				"default":  2,
				"low":      2,
			},
			// See the godoc for other configuration options
		},
	)

	mux := asynq.NewServeMux()
	//mux.HandleFunc(task.TypeWelcomeEmail, task.HandleWelcomeEmailTask)
	mux.HandleFunc(value.StringTaskImageCompress, controller.Task(HandleImageCompressTask))
	mux.HandleFunc(value.StringTaskAvatar, controller.Task(HandleAvatarCompressTask))

	if err := srv.Run(mux); err != nil {
		log.Fatal(err)
	}
}

func main() {
	ctrl := controller.New(data.NewDefaultEnv())
	worker(ctrl)
}
