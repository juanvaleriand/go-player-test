package utils

import (
	"context"
	"fmt"
	"go-player-test/internal/config"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func InitRedis(cfg config.Config) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.RedisAddr,
		Password: cfg.RedisPass,
		DB:       0,
	})
	pong, err := rdb.Ping(ctx).Result()
	if err != nil {
		fmt.Println(pong, err)
	}
	return rdb
}
