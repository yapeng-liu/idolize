package main

import (
	"fmt"
	"time"
)

// Observer 观察者（服务器）接口（包含主题发布、取消发布和任务完成通知）
type Observer interface {
	Publish()
	CancelPublish()
	Told(taskProgressId int64) bool
}

// ConcreteObserver 具体观察者 数据库
type ConcreteObserver struct {
}

// Told 任务事件触发 查询任务是否完成
func (cor *ConcreteObserver) Told(taskProgressId int64) bool {
	fmt.Print("任务 触发次数总数+1")
	time.Sleep(1 * time.Second)
	fmt.Print("查询数据库")
	time.Sleep(1 * time.Second)
	fmt.Print("数据库查询完毕", taskProgressId, "任务触发次数已达到 任务完成")
	time.Sleep(1 * time.Second)
	return true
}

// Publish 将任务进行发布
func (cor *ConcreteObserver) Publish() {
	fmt.Print("发布任务....")
	time.Sleep(1 * time.Second)
	fmt.Print("将任务注册到主题当中....")

}

// CancelPublish 将任务
func (cor *ConcreteObserver) CancelPublish() {
	fmt.Print("取消任务")
}
