package main

import (
	"context"
	"github.com/redis/go-redis/v9"
)

func redisInit(ctx context.Context) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "192.168.1.61:6379",
		Password: "LJFDSLjl89234",
		DB:       0, // use default DB
	})

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}

	return rdb, nil
}

func redisOptions(ctx context.Context, connect *redis.Client) error {
	return nil
}
