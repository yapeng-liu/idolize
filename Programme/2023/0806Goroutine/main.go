package main

import (
	"context"
	"fmt"
	"strings"
	"time"
)

func main() {
	ab(3, 5)
	fmt.Println("退出来了")
	fmt.Println()
	//var lastLoginTime int64 = 1691477925
	//needTime := time.Unix(lastLoginTime, 0)
	//每日0点刷新用户任务进度表中所有的日常任务
	var nums []int
	nums = make([]int, 0)
	nums = nil
	if len(nums) == 0 {
		fmt.Println(len(nums))
	}
	event1 := "点赞"
	event2 := "发帖"
	event3 := "评论"
	parts := strings.Split("点赞_发帖_评论", "_")
	fmt.Println(parts)
	for _, part := range parts {
		if part == event1 {
			fmt.Println(event1)
		}
		if part == event2 {
			fmt.Println(event2)
		}
		if part == event3 {
			fmt.Println(event3)
		}
	}

	//now := time.Now()
	//startToday := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, &time.Location{})
	//endToday := startToday.AddDate(0, 0, 1).Add(-time.Nanosecond)
	//fmt.Println(time.Now().Unix(), startToday.Unix(), endToday.Unix())
	//if needTime.After(startToday) && needTime.Before(endToday) {
	//	fmt.Println("就在今日")
	//} else {
	//	fmt.Println("不知道是哪一天")
	//}
	time.Sleep(5 * time.Second)
	fmt.Println("主函数退出")
}
func ab(a, b int) int {
	ctx := context.Background()
	go func(ctx context.Context) {
		time.Sleep(3 * time.Second)
		fmt.Println(" 子协程 三秒执行完了")
	}(ctx)
	fmt.Println("相加 执行完毕")
	return a + b
}
