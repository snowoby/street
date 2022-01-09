package main

import (
	"context"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"street/account"
	"street/db"
	"street/ent"
	"street/handler"
	"street/middleware"
	"street/profile"
)

func storeSetup() handler.Handler {
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		panic(err)
	}

	client.Schema.Create(context.Background())
	store := db.New(client)
	h := handler.New(store)
	return h
}

func setup() *gin.Engine {
	r := gin.Default()
	h := storeSetup()
	g := r.Group("/account")
	account.Routers(g, h)

	g = r.Group("/profile")
	profile.PublicRouters(g, h)
	g.Use(h.P(middleware.AccessToken))
	profile.Routers(g, h)
	return r
}

func main() {

	setup().Run("127.0.0.1:8088")
}
