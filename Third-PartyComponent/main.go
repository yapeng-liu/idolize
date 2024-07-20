package main

import (
	"fmt"
	"time"
)

func main() {
	//var (
	//	ctx = context.Background()
	//)
	//mysql
	//mysqlConnect, err := mysqlInit(ctx)
	//if err != nil {
	//	return
	//}
	//err = mysqlOptions(ctx, mysqlConnect)
	//if err != nil {
	//	return
	//}
	////redis
	//redisConnect, err := redisInit(ctx)
	//if err != nil {
	//	return
	//}
	//err = redisOptions(ctx, redisConnect)
	//if err != nil {
	//	return
	//}
	////mongo
	//mongoConnect, err := mongoInit(ctx)
	//if err != nil {
	//	return
	//}
	//err = mongoOptions(ctx, mongoConnect)
	//if err != nil {
	//	fmt.Println("mongoOptions err:", err)
	//	return
	//}

	type f func() time.Time
	m, ok := map[int32]f{
		1: GetNextDay,
		2: GetNextFirstWeekDay,
		3: GetNextFirstMonthDay,
	}[2]
	if !ok {
		fmt.Println("4")
		return
	}
	fmt.Println(":")
	fmt.Println("result:", m())
	//endTime, ok := map[int32]string{
	//	1: "1",
	//	2: "2",
	//	3: "3",
	//}[3]
	//if !ok {
	//	fmt.Println("4")
	//	return
	//}
	//fmt.Println(endTime)
}

// GetNextDay 获取明天凌晨的时间
func GetNextDay() time.Time {
	fmt.Println("1")
	deadline := time.Now().AddDate(0, 0, 1)
	return time.Date(deadline.Year(), deadline.Month(), deadline.Day(), 0, 0, 0, 0, deadline.Location())
}

// GetNextFirstWeekDay 获取下周一凌晨的时间
func GetNextFirstWeekDay() time.Time {
	fmt.Println("2")
	now := time.Now()
	weekday := int(now.Weekday())
	if weekday == 0 {
		weekday = 7
	}

	deadline := now.AddDate(0, 0, 8-weekday)
	return time.Date(deadline.Year(), deadline.Month(), deadline.Day(), 0, 0, 0, 0, deadline.Location())
}

// GetNextFirstMonthDay 获取下月一号凌晨的时间
func GetNextFirstMonthDay() time.Time {
	fmt.Println("3")
	now := time.Now()
	d := now.Day()
	deadline := now.AddDate(0, 1, -d+1)
	return time.Date(deadline.Year(), deadline.Month(), deadline.Day(), 0, 0, 0, 0, deadline.Location())
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
