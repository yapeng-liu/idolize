package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

const (
	AppOnlineKey = "app_online_test:set"
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
	var (
		err    error
		result int64
	)
	result, err = connect.SAdd(ctx, AppOnlineKey, 10).Result()
	if err != nil {
		fmt.Printf("SAdd err: %v", err)
		return err
	}
	fmt.Printf("SAdd result: %v\n", result)

	result, err = connect.SRem(ctx, AppOnlineKey, 10, 22).Result()
	if err != nil {
		fmt.Printf("SRem err: %v", err)
		return err
	}
	fmt.Printf("SRem result: %v\n", result)

	result, err = connect.SAdd(ctx, AppOnlineKey, 15, 18, 20, 26).Result()
	if err != nil {
		fmt.Printf("SAdd err: %v", err)
		return err
	}
	fmt.Printf("SAdd result: %v\n", result)

	result, err = connect.SCard(ctx, AppOnlineKey).Result()
	if err != nil {
		fmt.Printf("SCard err: %v", err)
		return err
	}
	fmt.Printf("SCard result: %v\n", result)

	result, err = connect.Del(ctx, AppOnlineKey).Result()
	if err != nil {
		fmt.Printf("SCard err: %v", err)
		return err
	}
	fmt.Printf("SCard result: %v\n", result)

	result, err = connect.SCard(ctx, AppOnlineKey).Result()
	if err != nil {
		fmt.Printf("SCard err: %v", err)
		return err
	}
	fmt.Printf("SCard result: %v\n", result)

	return nil
}
