package main

import (
	"fmt"
	"time"
)

// Task 任务接口
type Task interface {
	TaskInit()
	TaskUnInit()
	TaskExecute()
}

// ConcreteTask 具体任务
type ConcreteTask struct {
	TaskManageId   int64      //任务Id
	TaskProgressId int64      //任务进度唯一Id
	subject        Subject    //对应一个主题(任务)
	TriggerEvert   []int32    //多个任务触发事件
	TaskStatus     int32      //任务用户完成状态
	subscriber     Subscriber //订阅者
	observer       Observer   //观察者
}

// TaskInit 将用户注册到主题中
func (ctk *ConcreteTask) TaskInit() {
	fmt.Println(ctk.subscriber.GetUserId(), "用户接受任务,初始化")
	ctk.subject.Register(ctk.subscriber)
}

// TaskUnInit 在主题中注销用户
func (ctk *ConcreteTask) TaskUnInit() {
	fmt.Println(ctk.subscriber.GetUserId(), "用户已经完成任务,取消初始化")
	ctk.subject.CancelRegister(ctk.subscriber)
}

// TaskExecute 用户执行任务
func (ctk *ConcreteTask) TaskExecute() {
	fmt.Print(ctk.subscriber.GetUserId(), "用户执行任务当中，任务id=", ctk.TaskProgressId)
	time.Sleep(1 * time.Second)
	fmt.Print(ctk.subscriber.GetUserId(), "用户执行任务当中....")
	time.Sleep(2 * time.Second)
	fmt.Print(ctk.subscriber.GetUserId(), "用户执行任务完成")
	time.Sleep(1 * time.Second)
	// 通知观察者任务完成
	status := ctk.observer.Told(ctk.TaskProgressId)
	if status == true {
		ctk.TaskUnInit()
	}
}
