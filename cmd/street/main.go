package main

import (
	"context"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"street/ent"
	"street/pkg/controller"
	"street/pkg/data"
	"street/web/account"
	"street/web/profile"
)

func storeSetup() controller.Controller {
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
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
	r.Use(ctrl.Bare(account.TryAccessToken), ctrl.Bare(profile.TryProfile))

	g := r.Group("/account")
	account.Routers(g, ctrl)

	g = r.Group("/profile")
	profile.Routers(g, ctrl)

	return r
}

func main() {

	setup().Run("127.0.0.1:8088")
}
