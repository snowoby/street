package main

import (
	"context"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"street/account"
	"street/db"
	"street/ent"
	"street/handler"
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

func main() {
	r := gin.Default()
	h := storeSetup()
	g := r.Group("/account")

	account.Routers(g, h)

	//r.Use(account.AccessTokenMiddleware(store))

	r.Run("127.0.0.1:8088")
}
