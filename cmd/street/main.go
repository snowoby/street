package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/hibiken/asynq"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/mvrilo/go-redoc"
	"github.com/mvrilo/go-redoc/gin"
	"golang.org/x/net/context"
	"os"
	"street/ent"
	"street/pkg/account"
	"street/pkg/auth"
	"street/pkg/comment"
	"street/pkg/episode"
	"street/pkg/profile"
	"street/pkg/series"
	"street/pkg/site"
	"street/pkg/storage"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func setup() *gin.Engine {

	doc := redoc.Redoc{
		Title:       "Example API",
		Description: "Example API Description",
		SpecFile:    "./docs/swagger.json",
		SpecPath:    "/openapi.json",
		DocsPath:    "/docs",
	}

	r := gin.Default()
	r.RedirectTrailingSlash = false

	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})
	r.Use(ginredoc.New(doc))

	entClient := NewDefaultEnt()
	redisClient := NewDefaultRedis()
	asynqClient := NewDefaultAsynq()
	s3 := NewDefaultS3()
	authSrv := auth.New(entClient)
	// gin executes middleware after route is matched
	// https://github.com/gin-gonic/gin/issues/2413#issuecomment-645768561
	site.New(site.NewDefault(), r.Group("/site"))

	account.New(entClient, authSrv, r.Group("/account"))
	profile.New(entClient, authSrv, r.Group("/profile"))
	episode.New(entClient, authSrv, r.Group("/episode"))
	comment.New(entClient, authSrv, r.Group("/comment"))
	series.New(entClient, authSrv, r.Group("/series"))
	storage.New(entClient, authSrv, redisClient, r.Group("/file"), asynqClient, s3)

	return r
}

func main() {

	err := setup().Run(fmt.Sprintf("%s:%s", os.Getenv("address"), os.Getenv("port")))
	if err != nil {
		panic(err)
	}
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

func NewDefaultRedis() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: os.Getenv("redis"),
		DB:   0,
	})
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
