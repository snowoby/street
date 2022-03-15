package factory

import (
	"log"
	"os"
	"street/ent"
	"street/pkg/base3"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/hibiken/asynq"
	"gopkg.in/gographics/imagick.v3/imagick"
)

func init() {
	imagick.Initialize()
}

type service struct {
	db             *ent.Client
	storageService *storageService
	server         *asynq.Server
}

func DefaultServer() *asynq.Server {
	return asynq.NewServer(
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
}

func New(server *asynq.Server, db *ent.Client, s3Config *aws.Config) *service {
	return &service{
		db:             db,
		storageService: &storageService{base3.New(s3Config)},
		server:         server,
	}
}

func (s *service) Run() {
	mux := asynq.NewServeMux()
	// mux.HandleFunc(d.StringTaskImageCompress, s.HandleImageCompressTask)
	// mux.HandleFunc(d.StringTaskAvatar, s.HandleAvatarCompressTask)

	if err := s.server.Run(mux); err != nil {
		log.Fatal(err)
	}
}
