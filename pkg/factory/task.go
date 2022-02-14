package factory

import (
	"github.com/gin-gonic/gin"
	"github.com/hibiken/asynq"
	"gopkg.in/gographics/imagick.v3/imagick"
	"log"
	"os"
	"street/ent"
	"street/pkg/auth"
	"street/pkg/d"
)

func init() {
	imagick.Initialize()
}

type service struct {
	db             *ent.Client
	auth           auth.Service
	router         *gin.RouterGroup
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

func New(server *asynq.Server, db *ent.Client,
	auth auth.Service,
	router *gin.RouterGroup,
	storageService *storageService) *service {
	return &service{
		db:             db,
		auth:           auth,
		router:         router,
		storageService: storageService,
		server:         server,
	}
}

func (s *service) Run() {
	mux := asynq.NewServeMux()
	mux.HandleFunc(d.StringTaskImageCompress, s.HandleImageCompressTask)
	mux.HandleFunc(d.StringTaskAvatar, s.HandleAvatarCompressTask)

	if err := s.server.Run(mux); err != nil {
		log.Fatal(err)
	}
}
