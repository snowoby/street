package street

import (
	"fmt"
	"os"
	"street/cmd/config"
	"street/pkg/account"
	"street/pkg/auth"
	"street/pkg/comment"
	"street/pkg/episode"
	"street/pkg/profile"
	"street/pkg/series"
	"street/pkg/site"
	"street/pkg/storage"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/mvrilo/go-redoc"
	ginredoc "github.com/mvrilo/go-redoc/gin"
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
		c.Writer.Header().Set("Access-Control-Allow-Origin", os.Getenv("DOMAIN"))
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

	entClient := config.NewDefaultEnt()
	redisClient := config.NewDefaultRedis()
	asynqClient := config.NewDefaultAsynq()
	s3 := config.NewDefaultS3()
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
	address := "127.0.0.1:28089"
	if os.Getenv("MODE") != "debug" {
		address = "0.0.0.0:8088"
	}
	err := setup().Run(address)
	if err != nil {
		panic(err)
	}
}

func Main() {
	main()
}
