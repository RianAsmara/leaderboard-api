package common

import (
	"context"

	"github.com/go-redis/redis"
)

func InitRedis() (*redis.Client, context.Context) {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379", // Redis server address
	})

	ctx := context.Background()

	return rdb, ctx
}
