package db

import (
	"golang.org/x/net/context"
	"os"
	"street/ent"
)

type DB struct {
	Account *account
	Token   *token
	Series  *series
	Episode *episode
	Profile *profile
	File    *file
}

func NewDefaultConfig() *ent.Client {
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

func New(client *ent.Client) *DB {
	return &DB{
		Account: &account{client.Account},
		Token:   &token{client.Token},
		Series:  &series{client.Series},
		Episode: &episode{client.Episode},
		Profile: &profile{client.Profile},
		File:    &file{client.File},
	}
}
