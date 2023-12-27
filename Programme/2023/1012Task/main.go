package main

import "fmt"

const (
	TaskA = 1 << iota // 第0个bit位代表机器人签到事件
	TaskB             // 第1个bit位代表动态评论事件
	TaskC             // 第2个bit位代表动态点赞事件
	TaskD             //第2个bit位代表发表动态事件
	// 可以继续定义更多的任务
)

func main() {
	var (
		taskA    uint32 // 签到任务
		taskB    uint32 // 评论任务
		taskC    uint32 // 点赞任务
		taskD    uint32 // 发布动态动态
		taskABCD uint32 //任务触发
	)
	// 设置任务
	taskA |= TaskA
	taskB |= TaskB
	taskC |= TaskC
	taskD |= TaskD
	taskABCD |= TaskA
	taskABCD |= TaskB
	taskABCD |= TaskC
	taskABCD |= TaskD
	fmt.Println(taskA, " ", taskB, " ", taskC, " ", taskD, " ", taskABCD)

	// 检查任务是否触发
	if taskA&taskB == 1 {
		fmt.Println("任务A已经触发")
	} else {
		fmt.Println("任务A未触发")
	}
	if taskABCD&TaskA != 0 {
		fmt.Println("任务ABCD已经触发")
	}
	if taskA&TaskA == 1 {
		fmt.Println("任务A已经触发")
	}
}

//docker exec docoi-redis-1 redis-cli -a baiG6boo.z SAVE
//docker cp docoi-redis-1:/data/dump.rdb /data/app/docoi/data/redis/data
//docker exec docoi-mongo-3-1 mongodump -u admin -p ab96.5bc2d -o /data/a.file
//docker exec docoi-mongo-2-1 mongodump -u admin -p ab96.5bc2d -o /data/a.file
//docker exec docoi-mongo-1-1 mongodump -u admin -p ab96.5bc2d -o /data/a.file
//
//delfile=`ls -l -crt  $backup_dir/mongo1-* | awk '{print $9 }' | head -1`
//
//count=`ls -l -crt  $backup_dir/mongo1-* | awk '{print $9 }' | wc -l`
//
//coscli cp ./gamping-2023-10-15-03-00-01.sql cos://SleepWalkingClub/backUp/
//
//echo "y" | coscli rm cos://SleepWalkingClub/backUp/gamping-2023-10-15-03-00-01.sql
//
//ls -l -crt  $backup_dir/*.rdb | awk '{print $9 }' | head -1
//ls -l -crt /data/dbbackup/mysql/*.sql | awk -F "/" '{print $5 }' | head -1
//
//ls -l -crt $backup_dir/*.sql | awk -F "/" '{print $5 }' | head -1
//ls -d $backup_dir/mongo1-*/ | awk -F "/" '{print $5}' | head -1
//
//coscli cp $backup_dir/$dd.rdb cos://SleepWalkingClub/backUp/redis/$dd.rdb
//echo "y" | coscli rm cos://SleepWalkingClub/backUp/redis/$delfile
