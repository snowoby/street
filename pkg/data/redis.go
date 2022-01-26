package data

import "github.com/go-redis/redis/v8"

type rds struct {
	client *redis.Client
}
