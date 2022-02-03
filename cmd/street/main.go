package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"os"
	"street/pkg/controller"
	"street/pkg/data"
	"street/web/account"
	"street/web/episode"
	"street/web/file"
	"street/web/middleware"
	"street/web/profile"
	"street/web/series"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func setup() *gin.Engine {
	r := gin.Default()
	ctrl := controller.New(data.NewDefaultEnv())

	r.Use(cors.Default())

	r.Use(ctrl.Original(account.TryAccessToken), ctrl.Original(profile.TryProfile), middleware.TryUriUUID)

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

	return r
}

func main() {

	err := setup().Run(fmt.Sprintf("%s:%s", os.Getenv("address"), os.Getenv("port")))
	if err != nil {
		panic(err)
	}
}
