package main

import (
	"context"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"os"
	"street/ent"
	"street/pkg/controller"
	"street/pkg/data"
	"street/pkg/data/storage"
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

func storeSetup() controller.Controller {

	client, err := ent.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err)
	}

	err = client.Schema.Create(context.Background())
	if err != nil {
		panic(err)
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	store := data.New(client, storage.New(), rdb)
	return controller.New(store)
}

func setup() *gin.Engine {
	r := gin.Default()
	ctrl := storeSetup()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{os.Getenv("site")}
	config.AddAllowMethods("OPTIONS")
	r.Use(cors.New(config))

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
