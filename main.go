package main

import (
	"context"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"street/account"
	"street/db"
	"street/ent"
)

func main() {
	r := gin.Default()
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	client.Schema.Create(context.Background())
	store := db.New(client)
	if err != nil {
		panic(err)
	}
	g := r.Group("/account")

	account.Routers(g, store)

	//r.Use(account.AccessTokenMiddleware(store))

	r.Run("127.0.0.1:8088")
}
