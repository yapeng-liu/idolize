package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/sync/errgroup"
	"gorm.io/gorm"
)

func main() {
	var (
		group, ctx   = errgroup.WithContext(context.Background())
		mysqlConnect *gorm.DB
		redisConnect *redis.Client
		mongoConnect *mongo.Client
		err          error
	)
	//mysql
	group.Go(func() error {
		mysqlConnect, err = mysqlInit(ctx)
		if err != nil {
			return err
		}
		return nil
	})

	//redis
	group.Go(func() error {
		redisConnect, err = redisInit(ctx)
		if err != nil {
			return err
		}
		return nil
	})

	//mongo
	group.Go(func() error {
		mongoConnect, err = mongoInit(ctx)
		if err != nil {
			fmt.Println("mongoInit err:", err)
			return err
		}
		return nil
	})
	err = group.Wait()
	if err != nil {
		fmt.Println("group.Wait() err", err)
		panic(err)
	}
	ctx = context.Background()
	err = redisOptions(ctx, redisConnect)
	if err != nil {
		fmt.Println("redisOptions err", err)
		return
	}
	err = mysqlOptions(ctx, mysqlConnect)
	if err != nil {
		fmt.Println("mysqlOptions err", err)
		return
	}
	err = mongoOptions(context.Background(), mongoConnect)
	if err != nil {
		fmt.Println("mongoOptions err:", err)
		return
	}
}

//ls -d /data/dbbackup/mongo/mongo1-*/ | awk -F "/" ls'{print $5}' | wc -l

//redis-cli -p 30015 -h 42.194.237.166 -a baiG6boo.z --scan --pattern 'user:user:*' | xargs -L 1 redis-cli -p 30015 -h 42.194.237.166 -a baiG6boo.z del
//
//redis-cli -p 6379 -h 192.168.1.61 -a LJFDSLjl89234 --scan --pattern 'user:user:*' | xargs -L 1 redis-cli -p 6379 -h 192.168.1.61 -a LJFDSLjl89234 del
//
//update users set avatar="https://cdn.docoi.cc/avatars/mys_avatar_boy.png?TS=20" where avatar="https://cdn.docoi.cc/avatars/mys_avatar_boy.png?TS=10";
//
//update users set avatar="https://cdn.docoi.cc/avatars/mys_avatar_girl.png?TS=20" where avatar="https://cdn.docoi.cc/avatars/mys_avatar_girl.png?TS=10";
//
//SELECT utp.*
//FROM gamping.user_task_progress utp
//INNER JOIN (
//SELECT user_id, MAX(completed_count) AS max_completed
//FROM gamping.user_task_progress
//WHERE server_task_id IN (139,140,141,142,143,144,145)
//GROUP BY user_id
//) max_utp ON utp.user_id = max_utp.user_id AND utp.completed_count = max_utp.max_completed
//GROUP BY utp.user_id
//ORDER BY utp.user_id
//LIMIT 10;
