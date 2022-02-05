package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/mvrilo/go-redoc"
	"github.com/mvrilo/go-redoc/gin"
	"os"
	"street/pkg/controller"
	"street/pkg/data"
	"street/web/account"
	"street/web/episode"
	"street/web/file"
	"street/web/middleware"
	"street/web/profile"
	"street/web/series"
	"street/web/site"
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
	r.Use(ginredoc.New(doc))
	ctrl := controller.New(data.NewDefaultEnv())

	r.Use(cors.Default())

	r.Use(ctrl.Original(account.TryAccessToken), middleware.TryUriUUID, profile.TryProfile)

	g := r.Group("/account")
	account.Routers(g, ctrl)

	g = r.Group("/profile")
	profile.Routers(g, ctrl)

	g = r.Group("/series")
	series.Routers(g, ctrl)

	g = r.Group("/episode")
	episode.Routers(g, ctrl)

	g = r.Group("/file")
	file.Routers(g, ctrl)

	g = r.Group("/site")
	site.Routers(g, ctrl)

	return r
}

func main() {

	err := setup().Run(fmt.Sprintf("%s:%s", os.Getenv("address"), os.Getenv("port")))
	if err != nil {
		panic(err)
	}
}
