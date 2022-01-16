package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"os"
	"street/ent"
	"street/pkg/controller"
	"street/pkg/data"
	"street/web/account"
	"street/web/episode"
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

func storeSetup() controller.Controller {

	client, err := ent.Open("postgres", os.Getenv("database"))
	if err != nil {
		panic(err)
	}

	client.Schema.Create(context.Background())
	store := data.New(client)
	return controller.New(store)
}

func setup() *gin.Engine {
	r := gin.Default()
	ctrl := storeSetup()
	r.Use(ctrl.Original(account.TryAccessToken), ctrl.Original(profile.TryProfile), middleware.TryUriUUID)

	g := r.Group("/account")
	account.Routers(g, ctrl)

	g = r.Group("/profile")
	profile.Routers(g, ctrl)

	g = r.Group("/series")
	series.Routers(g, ctrl)

	g = r.Group("/episode")
	episode.Routers(g, ctrl)

	return r
}

func main() {

	setup().Run(os.Getenv("address"))
}
